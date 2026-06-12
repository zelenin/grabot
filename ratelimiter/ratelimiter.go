package ratelimiter

import (
	"context"
	"sync"
	"time"

	"github.com/zelenin/grabot/client"
	"golang.org/x/time/rate"
)

// rate limits:
// 1 message/sec to chat
// 30 message/sec
// 20 message/min to group

type RateLimiter struct {
	commonLimiter *rate.Limiter
	mu            sync.Mutex
	limiters      map[string]*limiter
}

func New() *RateLimiter {
	rateLimiter := &RateLimiter{
		commonLimiter: rate.NewLimiter(rate.Every(1*time.Second/30), 1),
	}

	go rateLimiter.gc()

	return rateLimiter
}

func (rl *RateLimiter) gc() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		rl.mu.Lock()
		for k, v := range rl.limiters {
			if v.accessedAt < (time.Now().Unix() - 5*60) {
				delete(rl.limiters, k)
			}
		}
		defer rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Wait(ctx context.Context, chatId client.ChatId) error {
	err := rl.commonLimiter.Wait(ctx)
	if err != nil {
		return err
	}

	return rl.getLimiter(chatId).Wait(ctx)
}

func (rl *RateLimiter) getLimiter(chatId client.ChatId) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	_, ok := rl.limiters[chatId.String()]
	if !ok {
		if isGroup(chatId) {
			rl.limiters[chatId.String()] = &limiter{
				limiter:    rate.NewLimiter(rate.Every(1*time.Minute/20), 1),
				accessedAt: 0,
			}
		} else {
			rl.limiters[chatId.String()] = &limiter{
				limiter:    rate.NewLimiter(rate.Every(1*time.Second), 1),
				accessedAt: 0,
			}
		}
	}

	rl.limiters[chatId.String()].accessedAt = time.Now().Unix()

	return rl.limiters[chatId.String()].limiter
}

func isGroup(id client.ChatId) bool {
	strId := id.String()

	return strId[0] == '-' || strId[0] == '@'
}

type limiter struct {
	limiter    *rate.Limiter
	accessedAt int64
}
