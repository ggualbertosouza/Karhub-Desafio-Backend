package InMemoryCache

import (
	"encoding/json"
	"errors"
	SpotifyEntity "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/domain/spotify"
	"log"
	"os"
	"sync"
)

type PlaylistCache struct {
	mu     sync.RWMutex
	data   map[string]SpotifyEntity.Playlist
	closed bool
}

var PlaylistMockCache *PlaylistCache

func InitPlaylistMockCache(filePath string) error {
	log.Println("[PlaylistCache] global init started")

	cache := &PlaylistCache{}

	if err := cache.Init(filePath); err != nil {
		log.Printf("[PlaylistCache] global init failed: %v", err)
		return err
	}

	PlaylistMockCache = cache

	log.Println("[PlaylistCache] global init completed")
	return nil
}

func (c *PlaylistCache) Init(filePath string) error {
	log.Println("[PlaylistCache] init started")

	if err := c.loadFromFile(filePath); err != nil {
		return err
	}

	log.Println("[PlaylistCache] init completed successfully")
	return nil
}

func (c *PlaylistCache) loadFromFile(filePath string) error {
	log.Printf("[PlaylistCache] loading playlists from file: %s", filePath)

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return errors.New("playlist cache is closed")
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var playlists map[string]SpotifyEntity.Playlist

	if err := json.Unmarshal(file, &playlists); err != nil {
		return err
	}

	c.data = playlists

	log.Printf("[PlaylistCache] loaded %d playlists", len(playlists))
	return nil
}

func (c *PlaylistCache) Get(style string) (*SpotifyEntity.Playlist, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.closed {
		return nil, errors.New("playlist cache is closed")
	}

	if c.data == nil {
		return nil, errors.New("playlist cache not initialized")
	}

	playlist, ok := c.data[style]
	if !ok {
		return nil, errors.New("playlist not found for beer style")
	}

	return &playlist, nil
}

func ClosePlaylistMockCache() {
	if PlaylistMockCache == nil {
		return
	}

	log.Println("[PlaylistCache] closing global cache")

	PlaylistMockCache.mu.Lock()
	defer PlaylistMockCache.mu.Unlock()

	PlaylistMockCache.data = nil
	PlaylistMockCache.closed = true

	log.Println("[PlaylistCache] global cache closed")
}
