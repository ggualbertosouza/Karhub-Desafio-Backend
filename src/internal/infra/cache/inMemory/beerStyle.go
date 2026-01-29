package InMemoryCache

import (
	"context"
	"errors"
	BsRepository "github/ggualbertosouza/Karhub-Desafio-Backend/src/internal/infra/repositories"
	"github/ggualbertosouza/Karhub-Desafio-Backend/src/pkg/postgres"
	"log"
	"sync"
)

type BeerStyleCache struct {
	mu     sync.RWMutex
	data   []BeerStyleList
	closed bool
}

var BsCache *BeerStyleCache

func InitBeerStyleCache(ctx context.Context) error {
	log.Println("[BeerStyleCache] global init started")

	cache := &BeerStyleCache{}

	if err := cache.Init(ctx); err != nil {
		log.Printf("[BeerStyleCache] global init failed: %v", err)
		return err
	}

	BsCache = cache

	log.Println("[BeerStyleCache] global init completed")
	return nil
}

func (c *BeerStyleCache) Init(ctx context.Context) error {
	log.Println("[BeerStyleCache] init started")

	if err := c.Populate(ctx); err != nil {
		return err
	}

	log.Println("[BeerStyleCache] init completed successfully")
	return nil
}

func (c *BeerStyleCache) Populate(ctx context.Context) error {
	log.Println("[BeerStyleCache] populate started")

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return errors.New("beer style cache is closed")
	}

	repo := BsRepository.NewQuery(postgres.PostgresDb)

	styles, err := repo.ListAll(ctx)
	if err != nil {
		log.Printf("[BeerStyleCache] failed to load beer styles: %v", err)
		return err
	}

	mapped, err := mapToCache(styles)
	if err != nil {
		return err
	}

	c.data = mapped

	log.Printf("[BeerStyleCache] populate completed: %d items", len(mapped))
	return nil
}

func (c *BeerStyleCache) Get(ctx context.Context) ([]BeerStyleList, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.closed {
		return nil, errors.New("beer style cache is closed")
	}

	if c.data == nil {
		return nil, errors.New("beer style cache not initialized")
	}

	return c.data, nil
}

func CloseBeerStyleCache() {
	if BsCache == nil {
		return
	}

	log.Println("[BeerStyleCache] closing global cache")

	BsCache.mu.Lock()
	defer BsCache.mu.Unlock()

	BsCache.data = nil
	BsCache.closed = true

	log.Println("[BeerStyleCache] global cache closed")
}
