package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type RateLimiter struct {
	ips    map[string]*rate.Limiter
	mu     *sync.RWMutex
	rate   rate.Limit
	burst  int
	expiry time.Duration
}

func NewRateLimiter(r rate.Limit, b int, expiry time.Duration) *RateLimiter {
	return &RateLimiter{
		ips:    make(map[string]*rate.Limiter),
		mu:     &sync.RWMutex{},
		rate:   r,
		burst:  b,
		expiry: expiry,
	}
}

func RateLimit(limit float64, burst int) gin.HandlerFunc {
	limiter := NewRateLimiter(rate.Limit(limit), burst, 1*time.Hour)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !limiter.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (r *RateLimiter) Allow(ip string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	limiter, exists := r.ips[ip]
	if !exists {
		limiter = rate.NewLimiter(r.rate, r.burst)
		r.ips[ip] = limiter
	}

	return limiter.Allow()
}
