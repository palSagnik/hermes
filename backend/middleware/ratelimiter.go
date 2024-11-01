package middleware

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type Visitor struct {
	Limiter *rate.Limiter
	LastSeen time.Time
}

// A map to hold the rate limiters for each visitor and a mut
var (
	visitors = make(map[string]*Visitor)
	mu sync.Mutex
)

// Creating a new limiter for a new ip address
// If an ip address already exists then retrieve it
func GetVisitor(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := rate.NewLimiter(1, 4)
		visitors[ip] = &Visitor{
			Limiter: limiter,
			LastSeen: time.Now(),
		}
	}

	v.LastSeen = time.Now()
	return v.Limiter
}

func CleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.LastSeen) > 3*(time.Minute) {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}
