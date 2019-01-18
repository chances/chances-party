﻿using System;
using JetBrains.Annotations;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata;

namespace Models
{
    [UsedImplicitly]
    public partial class PartyModelContainer : DbContext
    {
        public PartyModelContainer(DbContextOptions options)
            : base(options)
        {
        }

        public virtual DbSet<GuestList> GuestList { get; set; }
        public virtual DbSet<Party> Party { get; set; }
        public virtual DbSet<TrackList> TrackList { get; set; }
        public virtual DbSet<User> User { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Party>(entity =>
            {
                entity.HasIndex(e => e.RoomCode)
                    .HasName("party_room_code_uindex")
                    .IsUnique();

                entity.HasOne(d => d.Guests)
                    .WithMany(p => p.Party)
                    .HasForeignKey(d => d.GuestsId)
                    .HasConstraintName("party_guests_fk");

                entity.HasOne(d => d.History)
                    .WithMany(p => p.PartyHistory)
                    .HasForeignKey(d => d.HistoryId)
                    .HasConstraintName("party_history_fk");

                entity.HasOne(d => d.Queue)
                    .WithMany(p => p.PartyQueue)
                    .HasForeignKey(d => d.QueueId)
                    .HasConstraintName("party_queue_fk");
            });

            modelBuilder.Entity<TrackList>(entity =>
            {
                entity.HasIndex(e => e.SpotifyPlaylistId)
                    .HasName("track_list_spotify_playlist_id_uindex")
                    .IsUnique();

                entity.Property(e => e.CreatedAt).HasDefaultValueSql("now()");

                entity.Property(e => e.UpdatedAt).HasDefaultValueSql("now()");
            });

            modelBuilder.Entity<User>(entity =>
            {
                entity.HasIndex(e => e.PartyId)
                    .HasName("user_party_id_uindex")
                    .IsUnique();

                entity.HasIndex(e => e.Username)
                    .HasName("unique_username")
                    .IsUnique();

                entity.Property(e => e.CreatedAt).HasDefaultValueSql("now()");

                entity.Property(e => e.UpdatedAt).HasDefaultValueSql("now()");

                entity.HasOne(d => d.Party)
                    .WithOne(p => p.User)
                    .HasForeignKey<User>(d => d.PartyId)
                    .HasConstraintName("user_party_fk");
            });
        }
    }
}