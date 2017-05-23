package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"gopkg.in/gin-contrib/cors.v1"
)

func main() {
	dotenvErr := godotenv.Load()
	if dotenvErr != nil {
		log.Println("Warning: .env file is not present. Using system provided environment variables")
	}

	setupAuth()

	// === Data Stores ===
	// Postgres
	db = initDatabase()
	defer db.Close()
	// Redis
	pool = newRedisPool()
	defer pool.Close()

	g := gin.New()
	g.Use(gin.Logger())

	// === Middleware ===
	// CORS
	corsOrigins :=
		strings.Split(getenv("CORS_ORIGINS", "https://chancesnow.me"), ",")
	g.Use(cors.New(cors.Config{
		AllowOrigins:     corsOrigins,
		AllowMethods:     []string{"GET", "PUT", "POST", "PATCH", "DELETE"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Session
	g.Use(partySession())

	g.Use(handleErrors())
	g.Use(gin.Recovery())

	// Require SSL in production
	g.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			if gin.IsDebugging() == false && c.Request.TLS == nil {
				c.Error(errBadRequest.WithDetail("Party requires SSL"))
				c.Abort()
				return
			}

			c.Next()
		}
	}())

	// Static files
	g.Static("/css/", "./public")

	g.LoadHTMLGlob("views/*")

	// Application routes
	g.GET("/", func(c *gin.Context) {
		session := DefaultSession(c)
		flashedError, err := session.Error()
		if err != nil {
			c.Error(errInternal.CausedBy(err))
		}
		if IsLoggedIn(c) {
			currentUser := CurrentUser(c)
			var spotifyUser spotify.PrivateUser
			err := currentUser.SpotifyUser.Unmarshal(&spotifyUser)
			if err != nil {
				c.Error(errInternal.CausedBy(err))
				c.Abort()
				return
			}

			spotifyClient, err := ClientFromSession(c)
			if err != nil {
				c.Error(errInternal.CausedBy(err))
				c.Abort()
				return
			}

			var currentPlaylist *spotify.SimplePlaylist
			currentPlaylist = nil
			playlists := Playlists(*spotifyClient) // TODO: Cache these?
			for _, playlist := range playlists {
				if currentUser.SpotifyPlaylistID.String == playlist.ID.String() {
					currentPlaylist = &playlist
					break
				}
			}
			c.HTML(http.StatusOK, "index.html", gin.H{
				"user":            spotifyUser,
				"currentPlaylist": currentPlaylist,
				"playlists":       playlists,
				"error":           flashedError,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": flashedError,
			})
		}
	})

	playlist := g.Group("/playlist")
	playlist.Use(AuthRequired())
	{
		playlist.PATCH("/", patchPlaylist)
	}

	g.GET("/auth/login", login)
	g.GET("/auth/callback", spotifyCallback)
	g.GET("/auth/logout", func(c *gin.Context) {
		if IsLoggedIn(c) {
			session := DefaultSession(c)
			session.Delete("USER")
		}

		c.Redirect(http.StatusSeeOther, "/")
	})

	g.Run()
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getenvOrFatal(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatalf("error: Missing environment variable: %v\n", key)
	}
	return value
}
