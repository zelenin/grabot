package updates

import (
    "context"
    "github.com/zelenin/grabot/client"
)

type UpdateHandler func(ctx context.Context, update *client.Update)
