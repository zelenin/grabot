package bot

import (
    "context"
    "github.com/zelenin/grabot/updates"
    "github.com/zelenin/grabot/client"
)

type routeMiddleware struct {
    router *Router
}

func (middleware *routeMiddleware) Process(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
    route := middleware.router.Match(update)

    if route != nil {
        route.Handle(ctx, update, updateHandler)
        return
    }

    updateHandler(ctx, update)
}

func NewRouteMiddleware(router *Router) Middleware {
    middleware := &routeMiddleware{
        router: router,
    }

    return middleware.Process
}

type RouteMatcher func(update *client.Update) bool

type Route struct {
    matcher RouteMatcher
    handler Middleware
}

func (route *Route) Handle(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
    route.handler(ctx, update, updateHandler)
}

func NewRoute(matcher RouteMatcher, handler Middleware) *Route {
    return &Route{
        matcher: matcher,
        handler: handler,
    }
}

type Router struct {
    routes []*Route
}

func NewRouter() *Router {
    return &Router{
        routes: []*Route{},
    }
}

func (router *Router) AddRoute(route *Route) {
    router.routes = append(router.routes, route)
}

func (router *Router) Match(update *client.Update) *Route {
    for _, route := range router.routes {
        if route.matcher(update) {
            return route
        }
    }

    return nil
}
