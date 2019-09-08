using System.Collections.Generic;
using System.Linq;
using Models;
using Server.Models;
using Spotify.API.NetCore.Models;

namespace Server.ViewModels
{
  public class Administrator
  {
    public Models.Spotify.PrivateProfile User { get; }
    public string UserUrl => User?.ExternalUrls["spotify"] ?? null;
    public bool LoggedIn => User != null;

    public Playlist CurrentPlaylist { get; }
    public bool HasCurrentPlaylist => CurrentPlaylist != null;

    public Party CurrentParty { get; }
    public bool HasCurrentParty => CurrentParty != null;

    public IEnumerable<Playlist> Playlists { get; }
    public bool HasPlaylists => Playlists?.Any() ?? false;

    public string Error { get; }
    public bool HasError => Error?.Any() ?? false;

    public Models.Spotify.Image LargestUserImage
    {
      get
      {
        if (User.Images == null || User.Images.Count == 0)
        {
          return null;
        }

        if (!User.Images[0].Width.HasValue)
        {
          return User.Images[0];
        }

        var largestImage = User.Images[0];
        var size = User.Images[0].Width * User.Images[0].Height;
        foreach (var image in User.Images)
        {
          if (size < image.Width * image.Height)
          {
            largestImage = image;
          }
        }

        return largestImage;
      }
    }
    public bool HasUserImage => User.Images != null && User.Images.Count > 0;

    public Administrator(string error = null)
      : this(null, null, null, null, error)
    {
    }

    public Administrator(
      Models.Spotify.PrivateProfile user,
      IEnumerable<Playlist> playlists,
      Playlist playlist,
      Party party
    ) : this(user, playlists, playlist, party, null)
    {
    }

    public Administrator(
      Models.Spotify.PrivateProfile user,
      IEnumerable<Playlist> playlists,
      Playlist playlist,
      Party party,
      string error
    )
    {
      User = user;
      Playlists = playlists;
      CurrentPlaylist = playlist;
      CurrentParty = party;
      Error = error;
    }
  }
}
