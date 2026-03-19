package middleware

import (
	"backend/pkg/response"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// in-memory store for all clients
var clients = make(map[string]*client)

// protect concurrency access to map
var mu sync.Mutex

// remove clients unused to prevent memory leak
func cleanupClients() {
	// run every 1 minute
	for {
		time.Sleep(1 * time.Minute)
		mu.Lock() // prevent multiple goroutine access at the same time (allowed 1 request only)

		for key, cl := range clients {
			// if not used in 3 min it's should be delete
			if time.Since(cl.lastSeen) > 3*time.Minute {
				delete(clients, key)
			}
		}
		mu.Unlock() // unlock for next reqeusts can access to this
	}
}

func RateLimiter(r rate.Limit, brust int) gin.HandlerFunc {
	go cleanupClients()

	return func(c *gin.Context) {
		// use IP + path as key to separate limiter per endpoint
		key := c.ClientIP() + ":" + c.FullPath()

		mu.Lock()

		cl, exits := clients[key]
		// check key not exits then create new map client
		if !exits {
			clients[key] = &client{
				limiter:  rate.NewLimiter(r, brust),
				lastSeen: time.Now(),
			}

			cl = clients[key]
		}

		// update last seen
		cl.lastSeen = time.Now()

		mu.Unlock()

		// check if request is allowed
		if cl.limiter.Allow() {
			c.Next()
			return
		}

		response.Error(c, http.StatusTooManyRequests, response.TooManyReq)
		c.Abort()
	}
}
