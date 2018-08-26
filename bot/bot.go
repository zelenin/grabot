package bot

import (
    "context"
    "github.com/zelenin/grabot/client"
)

type Bot struct {
    client      *client.Client
    middlewares []Middleware
}

func NewBot(client *client.Client) *Bot {
    return &Bot{
        client:      client,
        middlewares: []Middleware{},
    }
}

func (bot *Bot) Add(middleware Middleware) {
    bot.middlewares = append(bot.middlewares, middleware)
}

func (bot *Bot) Handle(ctx context.Context, update *client.Update) {
    if ctx == nil {
        ctx = context.Background()
    }

    updateHandler := newMiddlewarePipe(bot.middlewares)

    updateHandler.Handle(ctx, update)
}
