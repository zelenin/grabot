package bot

import (
    "log"
    "context"
    "github.com/zelenin/grabot/updates"
    "github.com/zelenin/grabot/client"
)

type Middleware func(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler)

func LoggingMiddleware(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
    log.Printf("%#v", update)

    updateHandler(ctx, update)
}

func NoOpMiddleware(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {}

type middlewarePipe struct {
    middlewares        []Middleware
    fallbackMiddleware Middleware
    current            int
}

func (pipe *middlewarePipe) Handle(ctx context.Context, update *client.Update) {
    if pipe.current == len(pipe.middlewares) {
        pipe.fallbackMiddleware(ctx, update, updateHandler(pipe))
        return
    }

    middleware := pipe.middlewares[pipe.current]

    pipe.current++

    middleware(ctx, update, updateHandler(pipe))
}

func newMiddlewarePipe(middlewares []Middleware) *middlewarePipe {
    return &middlewarePipe{
        middlewares:        middlewares,
        fallbackMiddleware: NoOpMiddleware,
        current:            0,
    }
}

func updateHandler(pipe *middlewarePipe) updates.UpdateHandler {
    return func(ctx context.Context, update *client.Update) {
        pipe.Handle(ctx, update)
    }
}
