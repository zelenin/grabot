package ratelimiter

import (
    "time"
    "sync"
    "github.com/zelenin/grabot/client"
)

type tickerStorage struct {
    tickers             map[string]*time.Ticker
    chatTickerDuration  time.Duration
    groupTickerDuration time.Duration
    mu                  sync.Mutex
}

func (storage *tickerStorage) Get(id client.ChatId) *time.Ticker {
    storage.mu.Lock()
    defer storage.mu.Unlock()

    return storage.get(id)
}

func (storage *tickerStorage) get(id client.ChatId) *time.Ticker {
    ticker, ok := storage.tickers[id.String()]
    if !ok {
        dur := storage.chatTickerDuration
        if isGroup(id) {
            dur = storage.groupTickerDuration
        }

        ticker = time.NewTicker(dur)

        storage.tickers[id.String()] = ticker

        go func(storage *tickerStorage, id client.ChatId) {
            <-time.After(10 * time.Minute)
            storage.remove(id)
        }(storage, id)
    }

    return ticker
}

func (storage *tickerStorage) remove(id client.ChatId) {
    storage.mu.Lock()
    defer storage.mu.Unlock()

    delete(storage.tickers, id.String())
}
