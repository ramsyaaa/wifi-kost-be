package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type cacheMiddleware struct {
	rds *redis.Client
}

func NewCacheMiddleware(rds *redis.Client) *cacheMiddleware {
	return &cacheMiddleware{rds: rds}
}

func (m *cacheMiddleware) CacheUser(c *fiber.Ctx) error {
	cacheKey := getCacheKey(c)

	// return cache if exists
	cached, err := m.rds.Get(c.Context(), cacheKey).Result()
	if err == nil && cached != "" {
		c.Response().Header.Set("X-Cache", "HIT")
		c.Response().Header.Set("Content-Type", "application/json")
		return c.Status(200).SendString(cached)
	}

	// run handler if cached not exists
	result := c.Next()

	// then cache the response for 12 hours if status OK (200)
	status := c.Response().StatusCode()
	if status == 200 {
		response := c.Response().Body()
		m.rds.Set(c.Context(), cacheKey, response, 12*time.Hour)
	}

	return result
}

func (m *cacheMiddleware) InvalidateUserCache(c *fiber.Ctx) error {
	result := c.Next()

	// invalidate cache if status OK (200)
	status := c.Response().StatusCode()
	if status == 200 {
		cacheKey := getCacheKey(c)
		m.rds.Del(c.Context(), cacheKey)
	}

	return result
}

func getCacheKey(c *fiber.Ctx) string {
	return "user:" + c.Params("msisdn")
}
