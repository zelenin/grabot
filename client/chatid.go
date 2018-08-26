package client

import (
    "strconv"
    "strings"
)

const (
    minSecretId  = -2002147483648
    maxSecretId  = -1997852516353
    zeroSecretId = -2000000000000

    minChannelId = -1002147483647
    maxChannelId = -1000000000000

    minChatId = -2147483647

    maxUserId = 2147483647
)

// bot api chat types:
//
// - private (incl. bot)
// - group
// - supergroup
// - channel

// tdlib api chat types:
//
// - private (incl. bot)
// - basic group
// - supergroup (incl. channel)
// - secret

// It's chat ids. Not entity ids.
//
// secret chat id in [-2.002.147.483.648, -1.997.852.516.353)
// supergroup chat id in [-1.002.147.483.647, -1.000.000.000.000)
// basic group chat id in [-2.147.483.647, 0)
// private chat id in (0, 2.147.483.647]

type ChatId interface {
    String() string
}

func IntChatId(chatId int64) ChatId {
    return &intChatId{
        chatId: chatId,
    }
}

type intChatId struct {
    chatId int64
}

func (chatId *intChatId) String() string {
    return strconv.FormatInt(chatId.chatId, 10)
}

func (intChatId intChatId) IsPrivate() bool {
    return intChatId.chatId > 0 && intChatId.chatId <= maxUserId
}

func (intChatId intChatId) ToUserId() int64 {
    return intChatId.chatId
}

func (intChatId intChatId) IsBasicGroup() bool {
    return intChatId.chatId < 0 && minChatId <= intChatId.chatId
}

func (intChatId intChatId) ToBasicGroupId() int64 {
    return -intChatId.chatId
}

func (intChatId intChatId) IsSupergroup() bool {
    return intChatId.chatId < 0 && minChannelId <= intChatId.chatId && intChatId.chatId < maxChannelId
}

func (intChatId intChatId) ToSupergroupId() int64 {
    return maxChannelId - intChatId.chatId
}

func (intChatId intChatId) IsSecretChat() bool {
    return intChatId.chatId < 0 && minSecretId <= intChatId.chatId && intChatId.chatId < maxSecretId
}

func (intChatId intChatId) ToSecretChatId() int64 {
    return intChatId.chatId - zeroSecretId
}

// @channelname
func StringChatId(chatId string) ChatId {
    if !strings.HasPrefix(chatId, "@") {
        chatId = "@" + chatId
    }
    return &stringChatId{
        chatId: chatId,
    }
}

type stringChatId struct {
    chatId string
}

func (chatId *stringChatId) String() string {
    return chatId.chatId
}
