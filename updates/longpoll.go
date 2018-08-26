package updates

import (
    "time"
    "github.com/zelenin/grabot/client"
    "context"
)

type LongPoller interface {
    LongPoll(ctx context.Context, initReq *client.GetUpdatesRequest, interval time.Duration) (chan *client.Update, chan error)
}

type BasicLongPoller struct {
    client *client.Client
}

func (longPoller *BasicLongPoller) LongPoll(ctx context.Context, initReq *client.GetUpdatesRequest, interval time.Duration) (chan *client.Update, chan error) {
    updates := make(chan *client.Update, 1000)
    errs := make(chan error, 1000)

    go longPoller.longPoll(ctx, initReq, interval, updates, errs)

    return updates, errs
}

func (longPoller *BasicLongPoller) longPoll(ctx context.Context, initReq *client.GetUpdatesRequest, interval time.Duration, updatesChan chan *client.Update, errs chan error) {
    ticker := time.NewTicker(interval)

    for {
        select {
        case <-ticker.C:
            updates, err := longPoller.client.GetUpdates(initReq)
            if err != nil {
                errs <- err
                continue
            }

            for _, update := range updates {
                updatesChan <- update
                if *initReq.Offset <= update.UpdateId {
                    initReq.Offset = client.OptionalInt(update.UpdateId + 1)
                }
            }

        case <-ctx.Done():
            errs <- ctx.Err()
            return
        }
    }
}

func NewLongPoller(client *client.Client) LongPoller {
    return &BasicLongPoller{
        client: client,
    }
}
