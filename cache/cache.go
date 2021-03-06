package cache

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// Store Party cache entries
type Store struct {
	pool *redis.Pool
}

// Entry stores a cache line given a value and expiry time
type Entry struct {
	Value   interface{}
	Expiry  time.Time
	Rolling bool
}

// IsForever returns true if entry is cached indefinitely
func (e *Entry) IsForever() bool {
	return e.Expiry.Equal(time.Unix(0, 0).UTC())
}

// IsExpired returns true if entry is expired
func (e *Entry) IsExpired() bool {
	isForever := e.IsForever()
	return !isForever && e.Expiry.Before(time.Now().UTC())
}

// Value creates a plain cache entry given a value
func Value(value interface{}) Entry {
	return Entry{
		Value:   value,
		Rolling: false,
	}
}

// Expires creates a cache entry that expires at the given time
func Expires(expiry time.Time, value interface{}) Entry {
	return Entry{
		Expiry:  expiry.UTC(),
		Value:   value,
		Rolling: false,
	}
}

// ExpiresRolling creates a cache entry that expires at the given time, though expiry will roll forward when accessed
func ExpiresRolling(expiry time.Time, value interface{}) Entry {
	return Entry{
		Expiry:  expiry.UTC(),
		Value:   value,
		Rolling: true,
	}
}

// Forever creates a cache entry that shall never expire
func Forever(value interface{}) Entry {
	return Entry{
		Expiry:  time.Unix(0, 0).UTC(),
		Value:   value,
		Rolling: false,
	}
}

// NewStore instantiates a cache store given a Redis pool
func NewStore(p *redis.Pool) Store {
	gob.Register(map[string]string{})
	gob.Register(Entry{})
	gob.Register(gin.H{})
	return Store{
		pool: p,
	}
}

// Exists checks for existence of a cache entry key
func (s *Store) Exists(key string) (bool, error) {
	return s.redisExists(key)
}

// Get a cache entry
func (s *Store) Get(key string) (*Entry, error) {
	valueBytes, err := s.redisGetBytes(key)
	if err != nil {
		return nil, err
	}

	var value Entry
	valueBuffer := bytes.NewBuffer(valueBytes)
	dec := gob.NewDecoder(valueBuffer)
	err = dec.Decode(&value)
	if err != nil {
		return nil, err
	}

	// Refresh the entry in Redis if it's expiry rolls with access
	if !value.IsForever() && value.Rolling {
		err = s.redisSet(key, valueBuffer)
		if err != nil {
			return nil, err
		}
	}

	// Delete the entry from the cache if it's expired
	if value.IsExpired() {
		s.Delete(key)
	}

	return &value, nil
}

// GetOrDefer a cache entry, deferring to supplier func if entry doesn't exist
// or is expired. If key exists and is expired, entry is deleted
func (s *Store) GetOrDefer(key string, deferFn func() (*Entry, error)) (*Entry, error) {
	exists, err := s.redisExists(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		entry, err := deferFn()
		if err != nil {
			return nil, err
		}
		err = s.Set(key, *entry)
		if err != nil {
			return nil, err
		}
		return entry, nil
	}

	entry, err := s.Get(key)
	empty := err != nil && err.Error() == "EOF"

	if empty || entry.IsExpired() {
		entry, err = deferFn()
		if err != nil {
			return nil, err
		}
		err = s.Set(key, *entry)
		if err != nil {
			return nil, err
		}
		return entry, nil
	}

	if err != nil {
		return nil, err
	}

	return entry, nil
}

// Set a cache entry
func (s *Store) Set(key string, e Entry) error {
	valueBuffer := new(bytes.Buffer)
	enc := gob.NewEncoder(valueBuffer)
	err := enc.Encode(e)
	if err != nil {
		return err
	}

	err = s.redisSet(key, valueBuffer)
	if err != nil {
		return err
	}
	// If key is not forever then tell redis to expire it automagically
	if !e.IsForever() {
		lifetime := e.Expiry.Sub(time.Now())
		return s.redisExpire(key, lifetime)
	}

	return nil
}

// Delete a cache entry
func (s *Store) Delete(key string) error {
	return s.redisDelete(key)
}

// DeleteKeys deletes all given keys
func (s *Store) DeleteKeys(keys []string) error {
	return s.redisDeleteKeys(keys)
}
