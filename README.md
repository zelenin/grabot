# grabot

Go wrapper for [Telegram Bot API](https://core.telegram.org/bots/api) with full support of Bot API 4.1.

## API client

```go
import (
    "github.com/zelenin/grabot/client"
)

...

var chatId int64 = 12345678
var channelName = "@channelname"

token := "<bot_token>"

apiClient, _ := client.New(token/*, client.WithStdLogger*/)

apiClient.SendMessage(&client.SendMessageRequest{
    ChatId: client.IntChatId(chatId),
    Text:   "Hello. I'm bot.",
})

file, _ := client.NewFileInputFile("/path/to/image.jpg")

// or
// file := client.NewUrlInputFile("https://example.com/path/to/image.jpg")

apiClient.SendPhoto(&client.SendPhotoRequest{
    ChatId:  client.StringChatId(channelName),
    Photo:   file,
    Caption: client.OptionalString("Your answer?"),
    ReplyMarkup: &client.InlineKeyboardMarkup{
        InlineKeyboard: [][]client.InlineKeyboardButton{
            {
                {Text: "Variant 1", Url: client.OptionalString("https://example.com/variant/1")},
                {Text: "Variant 2", Url: client.OptionalString("https://example.com/variant/2")},
            },
            {
                {Text: "Variant 3", Url: client.OptionalString("https://example.com/variant/3")},
                {Text: "Variant 4", Url: client.OptionalString("https://example.com/variant/4")},
            },
        },
    },
})
```

## Updates

### Webhook

```go
import (
    "net/http"
    "log"
    "github.com/zelenin/grabot/updates"
    "github.com/zelenin/grabot/client"
    "context"
)

...

token := "<bot_token>"
apiClient, _ := client.New(token/*, client.WithStdLogger*/)
setWebhook(apiClient)

webhookHandler := updates.NewWebhookHandler(func(ctx context.Context, update *client.Update) {
    log.Printf("%#v", update)
})

mux := http.NewServeMux()
mux.HandleFunc("/webhook", webhookHandler.ServeHTTP)
srv := &http.Server{
    Addr:    ":8443",
    Handler: mux,
}
log.Fatal(srv.ListenAndServeTLS("./server.crt", "./server.key"))
```

### Long polling

```go
import (
    "log"
    "context"
    "time"
    "github.com/zelenin/grabot/updates"
    "github.com/zelenin/grabot/client"
)

... 

token := "<bot_token>"

apiClient, _ := client.New(token/*, client.WithStdLogger*/)

ctx, _ := context.WithCancel(context.Background())

longPoller := updates.NewLongPoller(apiClient)
updatesChan, errsChan := longPoller.LongPoll(ctx, &client.GetUpdatesRequest{
    Offset: client.OptionalInt(0),
}, 1*time.Second)

for {
    select {
    case update := <-updatesChan:
        log.Printf("%#v", update)
    
    case err := <-errsChan:
        log.Printf("error: %s", err)
    }
}
```

## Bot

```go

import (
    "github.com/zelenin/grabot/client"
    "github.com/zelenin/grabot/bot"
    "github.com/zelenin/grabot/updates"
    "log"
    "time"
    "context"
)

token := "<bot_token>"
apiClient, _ := client.New(token/*, client.WithStdLogger*/)

router := bot.NewRouter()

router.AddRoute(bot.NewRoute(bot.BotCommandMatcher("/start"), func(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
    log.Printf("Handle update #%d [bot command]", update.UpdateId)
    
    updateHandler(ctx, update)
}))

router.AddRoute(bot.NewRoute(bot.HashtagMatcher("#hashtag"), func(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
    log.Printf("Handle update #%d [hashtag]", update.UpdateId)
    
    updateHandler(ctx, update)
}))

router.AddRoute(bot.NewRoute(bot.MentionMatcher("@NameOfTheBot"), func(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
    log.Printf("Handle update #%d [mention]", update.UpdateId)
    
    updateHandler(ctx, update)
}))

router.AddRoute(bot.NewRoute(
    func(update *client.Update) bool {
        return update.Message != nil && update.Message.From != nil && *update.Message.From.Username == "username"
    },
    func(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
        log.Printf("Handle update #%d [from %s]", update.UpdateId, *update.Message.From.Username)
        
        updateHandler(ctx, update)
    }, 
))

router.AddRoute(bot.NewRoute(bot.MessageMatcher(), func(ctx context.Context, update *client.Update, updateHandler updates.UpdateHandler) {
    log.Printf("Handle update #%d [message #%d]", update.UpdateId, update.Message.MessageId)
    
    updateHandler(ctx, update)
}))

grabot := bot.NewBot(apiClient)

grabot.Add(bot.NewRouteMiddleware(router))

ctx, _ := context.WithCancel(context.Background())

longPoller := updates.NewLongPoller(apiClient)
updatesChan, errsChan := longPoller.LongPoll(ctx, &client.GetUpdatesRequest{
    Offset: client.OptionalInt(0),
}, 1*time.Second)

for {
    select {
    case update := <-updatesChan:
        ctx := context.Background()
        go grabot.Handle(ctx, update)
        
    case err := <-errsChan:
        log.Printf("error: %s", err)
    }
}
```

## Rate limiter

```go
import (
    "github.com/zelenin/grabot/client"
    "github.com/zelenin/grabot/ratelimiter"
)

...

var chatId int64 = 12345678

token := "<bot_token>"
apiClient, _ := client.New(token, client.WithStdLogger)

rlimiter := ratelimiter.New()

for i := 0; i < 100; i++ {
    rlimiter.AddTask(ratelimiter.NewTask(client.IntChatId(chatId), func() {
        apiClient.SendMessage(&client.SendMessageRequest{
            ChatId: client.IntChatId(chatId),
            Text:   "Hello. I'm bot.",
        })
    }))
}
```

## Author

[Aleksandr Zelenin](https://github.com/zelenin/), e-mail: [aleksandr@zelenin.me](mailto:aleksandr@zelenin.me)
