using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.Claims;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Http;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.DependencyInjection;
using Models;
using Newtonsoft.Json;
using Server.Configuration;
using Server.Models;
using Server.Services.Background;

namespace Server.Services.Authentication
{
  public static class CookiesAuthenticationScheme
  {
    public const string Name = CookieAuthenticationDefaults.AuthenticationScheme;

    private const string CookieName = "cpSESSION";
    private const string ProductionCookieDomain = ".chancesnow.me";

    public const string GuestUserJsonClaim = "urn:party:guest:userJson";

    public static void Configure(CookieAuthenticationOptions options, ITicketStore sessionStore, AppConfiguration config)
    {
      var isProduction = config.Mode.IsProduction();

      options.SessionStore = sessionStore;
      options.LoginPath = "/auth/login";
      options.LogoutPath = "/auth/logout";
      options.ReturnUrlParameter = "return_to";
      options.ExpireTimeSpan = TimeSpan.FromHours(12);
      options.SlidingExpiration = true;
      options.Cookie.Name = CookieName;
      options.Cookie.SecurePolicy = isProduction ? CookieSecurePolicy.Always : CookieSecurePolicy.None;
      options.Cookie.Domain = isProduction ? ProductionCookieDomain : "";

      options.Events = new CookieAuthenticationEvents
      {
        OnRedirectToLogin = context =>
        {
          // https://github.com/aspnet/Security/issues/1394
          // https://github.com/aspnet/AspNetCore/blob/62351067ff4c1401556725b401478e648b66acdc/src/Security/Authentication/Cookies/src/CookieAuthenticationEvents.cs#L42
          context.Response.Headers["Location"] = context.RedirectUri;
          context.Response.StatusCode = 401;
          return Task.CompletedTask;
        },

        OnSignedIn = async context =>
        {
          if (context.Principal.IsInRole(Roles.Guest)) return;

          var dbContext = context.HttpContext.RequestServices.GetRequiredService<PartyModelContainer>();
          await SpotifyAuthenticationScheme.UpsertPartyUser(dbContext, context.Principal.Claims);
        },

        OnValidatePrincipal = async context =>
        {
          var claims = context.Principal.Claims.ToList();

          if (context.Principal.IsInRole(Roles.Guest))
          {
            ValidateGuestPrincipal(context);
            if (context.ShouldRenew == false)
            {
              await context.HttpContext.SignOutAsync(Name);
            }
            return;
          }

          // Try to replace Spotify access token if it's expired
          if (SpotifyAuthenticationScheme.IsAccessTokenExpired(claims))
          {
            var updatedPrincipal = await SpotifyAuthenticationScheme
              .RefreshAccessToken(claims, config.Spotify);

            // Replace principal if Spotify access token was refreshed
            if (updatedPrincipal != null)
            {
              context.ReplacePrincipal(updatedPrincipal);
              context.ShouldRenew = true;
              return;
            }

            // Reject principal otherwise
            context.RejectPrincipal();
            context.ShouldRenew = false;

            await context.HttpContext.SignOutAsync(Name);
          }
        }
      };
    }

    public static ClaimsPrincipal CreateGuestPrincipal(Guest guest)
    {
      var claims = new List<Claim>
      {
        new Claim(ClaimsIdentity.DefaultRoleClaimType, Roles.Guest),
        new Claim(ClaimTypes.NameIdentifier, guest.Token),
        new Claim(GuestUserJsonClaim, JsonConvert.SerializeObject(guest)),
      };
      return new ClaimsPrincipal(new ClaimsIdentity(claims, "Guest"));
    }

    public static Guest GetGuest(IEnumerable<Claim> principalClaims)
    {
      if (principalClaims == null) return null;

      var guestJson = principalClaims.ToList()
        .Where(c => c.Type == GuestUserJsonClaim)
        .Select(c => c.Value)
        .FirstOrDefault();
      return string.IsNullOrWhiteSpace(guestJson)
        ? null
        : JsonConvert.DeserializeObject<Guest>(guestJson);
    }

    private static void ValidateGuestPrincipal(CookieValidatePrincipalContext context)
    {
      var guest = GetGuest(context.Principal.Claims);

      // Reject guest principal if it:
      //  - Couldn't be loaded,
      //  - Is expired,
      //  - Or Origin header is mismatched
      if (guest == null || guest.IsExpired || !guest.OriginMatches(context.Request.Headers))
      {
        context.RejectPrincipal();
        context.ShouldRenew = false;
        return;
      }

      // Sliding expiration rule, update expiry if time left on expiry is more than half through
      if (guest.Expiry.Subtract(DateTime.UtcNow).TotalMinutes >= 15) return;

      guest.UpdatedAt = DateTime.UtcNow;
      guest.Expiry = guest.UpdatedAt.AddMinutes(30);

      // Update the guest in its related party's guest list
      var bg = context.HttpContext.RequestServices.GetRequiredService<IBackgroundTaskQueue>();
      bg.QueueTask(async token =>
      {
        var db = context.HttpContext.RequestServices.GetRequiredService<PartyModelContainer>();
        var party = await db.Party
          .Include(p => p.Guests)
          .FirstOrDefaultAsync(p => p.Id == guest.PartyId, token);
        if (party == null) return;

        var guests = party.GuestList();
        var existingGuest = guests.FirstOrDefault(g => g.Token == guest.Token);
        if (existingGuest == null) return;

        var existingGuestIndex = guests.IndexOf(existingGuest);
        guests[existingGuestIndex] = guest;
        party.UpdateGuestList(guests);

        await db.SaveChangesAsync(token);
      });

      context.ReplacePrincipal(CreateGuestPrincipal(guest));
      context.ShouldRenew = true;
    }
  }
}
