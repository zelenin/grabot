package ratelimiter

import (
    "time"
    "github.com/zelenin/grabot/client"
)

// rate limits:
//
// 1 message/sec to chat
// 30 message/sec
// 20 message/min to group

func New() *RateLimiter {
    rateLimiter := &RateLimiter{
        commonTicker: newCommonTicker(),
        tasks:        make(chan Task, 1000),
        tickerStorage: &tickerStorage{
            tickers:             make(map[string]*time.Ticker),
            chatTickerDuration:  time.Second,
            groupTickerDuration: time.Minute / 20,
        },
    }

    go rateLimiter.run()

    return rateLimiter
}

type RateLimiter struct {
    commonTicker  *time.Ticker
    tasks         chan Task
    tickerStorage *tickerStorage
}

func (limiter *RateLimiter) AddTask(task Task) {
    go func(tasks chan Task, task Task) {
        tasks <- task
    }(limiter.tasks, task)
}

func (limiter *RateLimiter) run() {
    for {
        task := <-limiter.tasks

        <-limiter.commonTicker.C
        <-limiter.tickerStorage.Get(task.Id).C

        go task.Job()
    }
}

type Task struct {
    Id  client.ChatId
    Job Job
}

func NewTask(id client.ChatId, job Job) Task {
    return Task{
        Id:  id,
        Job: job,
    }
}

type Job func()

func newCommonTicker() *time.Ticker {
    return time.NewTicker(time.Second / 30)
}

func newChatTicker() *time.Ticker {
    return time.NewTicker(time.Second)
}

func newGroupTicker() *time.Ticker {
    return time.NewTicker(time.Minute / 20)
}

func isGroup(id client.ChatId) bool {
    strId := id.String()

    return strId[0] == '-' || strId[0] == '@'
}
