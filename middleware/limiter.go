package middleware

// import (
// 	"sync"
// 	"time"
// )

// type RateLimiter struct {
// 	requests map[string][]time.Time
// 	mu       sync.Mutex
// }

// func (rl *RateLimiter) isAllowed(ip string) bool {
// 	rl.mu.Lock()
// 	defer rl.mu.Unlock()

// 	now := time.Now()
// 	// Clean old requests
// 	rl.requests[ip] = filterOldRequests(rl.requests[ip], now)

// 	// Check rate limit
// 	if len(rl.requests[ip]) >= 100 {
// 		return false
// 	}

// 	rl.requests[ip] = append(rl.requests[ip], now)
// 	return true
// }
