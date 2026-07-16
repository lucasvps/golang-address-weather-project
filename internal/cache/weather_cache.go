package cache

import (
	"time"

	"example.com/address-weather-project/internal/domain"
)

type cacheItem struct {
	Data      *domain.WeatherResponse
	ExpiresAt time.Time
}

type WeatherCache struct {
	items map[string]cacheItem
	ttl   time.Duration
	// mutex sync.RWMutex
}

func NewWeatherCache(ttl time.Duration) *WeatherCache {
	return &WeatherCache{
		ttl:   ttl,
		items: make(map[string]cacheItem),
	}
}

func (wc *WeatherCache) Get(postalCode string) *domain.WeatherResponse {
	cacheItem, ok := wc.items[postalCode]

	if !ok {
		return nil
	}

	if cacheItem.ExpiresAt.Before(time.Now()) {
		delete(wc.items, postalCode)
		return nil
	}

	return cacheItem.Data
}

func (wc *WeatherCache) Set(postalCode string, data *domain.WeatherResponse) {
	var ct cacheItem

	ct.ExpiresAt = time.Now().Add(wc.ttl)
	ct.Data = data

	wc.items[postalCode] = ct
}
