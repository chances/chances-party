using System.Collections.Generic;
using Newtonsoft.Json;
using SpotifyAPI = Spotify.API.NetCore.Models;

namespace Server.Models.Spotify
{
  public class PrivateProfile : SpotifyAPI.BasicModel
  {
    public PrivateProfile()
    {
    }

    [JsonProperty("birthdate")]
    public string Birthdate { get; set; }
    [JsonProperty("country")]
    public string Country { get; set; }
    [JsonProperty("display_name")]
    public string DisplayName { get; set; }
    [JsonProperty("email")]
    public string Email { get; set; }
    [JsonProperty("external_urls")]
    public Dictionary<string, string> ExternalUrls { get; set; }
    [JsonProperty("followers")]
    public SpotifyAPI.Followers Followers { get; set; }
    [JsonProperty("href")]
    public string Href { get; set; }
    [JsonProperty("id")]
    public string Id { get; set; }
    [JsonProperty("images")]
    public List<Image> Images { get; set; }
    [JsonProperty("product")]
    public string Product { get; set; }
    [JsonProperty("type")]
    public string Type { get; set; }
    [JsonProperty("uri")]
    public string Uri { get; set; }
  }
}
