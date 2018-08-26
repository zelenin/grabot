package updates

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "github.com/zelenin/grabot/client"
)

type WebhookHandler struct {
    updateHandler UpdateHandler
}

func (handler WebhookHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    defer req.Body.Close()

    if !isValidWebhookRequest(req) {
        res.WriteHeader(http.StatusBadRequest)
        return
    }

    data, err := ioutil.ReadAll(req.Body)
    if err != nil {
        res.WriteHeader(http.StatusBadRequest)
        return
    }

    var update client.Update

    err = json.Unmarshal(data, &update)
    if err != nil {
        res.WriteHeader(http.StatusBadRequest)
        return
    }

    handler.updateHandler(req.Context(), &update)
}

func isValidWebhookRequest(req *http.Request) bool {
    if req.Method != http.MethodPost {
        return false
    }

    if req.Header.Get("Content-Type") != "application/json" {
        return false
    }

    if req.Body == nil {
        return false
    }

    return true
}

func NewWebhookHandler(updateHandler UpdateHandler) *WebhookHandler {
    return &WebhookHandler{
        updateHandler: updateHandler,
    }
}
