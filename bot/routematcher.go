package bot

import (
    "strings"
    "github.com/zelenin/grabot/client"
)

func BotCommandMatcher(botCommand string) RouteMatcher {
    botCommand = normalizeBotCommand(botCommand)

    return func(update *client.Update) bool {
        if update.Message == nil {
            return false
        }

        if update.Message.Entities == nil {
            return false
        }

        for _, entity := range *update.Message.Entities {
            if entity.Type == client.MessageEntityBotCommand {
                text := normalizeBotCommand(substring(*update.Message.Text, entity.Offset, entity.Length))
                if text == botCommand {
                    return true
                }
            }
        }

        return false
    }
}

func HashtagMatcher(hashtag string) RouteMatcher {
    hashtag = normalizeHashtag(hashtag)

    return func(update *client.Update) bool {
        if update.Message == nil {
            return false
        }

        if update.Message.Entities == nil {
            return false
        }

        for _, entity := range *update.Message.Entities {
            if entity.Type == client.MessageEntityHashtag {
                text := normalizeHashtag(substring(*update.Message.Text, entity.Offset, entity.Length))
                if text == hashtag {
                    return true
                }
            }
        }

        return false
    }
}

func MentionMatcher(mention string) RouteMatcher {
    mention = normalizeMention(mention)

    return func(update *client.Update) bool {
        if update.Message == nil {
            return false
        }

        if update.Message.Entities == nil {
            return false
        }

        for _, entity := range *update.Message.Entities {
            if entity.Type == client.MessageEntityHashtag {
                text := normalizeMention(substring(*update.Message.Text, entity.Offset, entity.Length))
                if text == mention {
                    return true
                }
            }
        }

        return false
    }
}

func ChosenInlineResultMatcher() RouteMatcher {
    return func(update *client.Update) bool {
        return update.ChosenInlineResult != nil
    }
}

func CallbackQueryMatcher() RouteMatcher {
    return func(update *client.Update) bool {
        return update.CallbackQuery != nil
    }
}

func InlineQueryMatcher() RouteMatcher {
    return func(update *client.Update) bool {
        return update.InlineQuery != nil
    }
}

func MessageMatcher() RouteMatcher {
    return func(update *client.Update) bool {
        return update.Message != nil
    }
}

func PreCheckoutQueryMatcher() RouteMatcher {
    return func(update *client.Update) bool {
        return update.PreCheckoutQuery != nil
    }
}

func ShippingQueryMatcher() RouteMatcher {
    return func(update *client.Update) bool {
        return update.ShippingQuery != nil
    }
}

func substring(s string, offset int64, length int64) string {
    end := offset + length

    var start int64
    var i int64
    for index, _ := range s {
        if i == offset {
            start = int64(index)
        }
        if i == end {
            return s[start:int64(index)]
        }

        i++
    }
    return s[start:]
}

func normalizeBotCommand(botCommand string) string {
    botCommand = strings.TrimPrefix(botCommand, "/")

    botCommandParts := strings.Split(botCommand, "@")

    return botCommandParts[0]
}

func normalizeHashtag(hashtag string) string {
    return strings.TrimPrefix(hashtag, "#")
}

func normalizeMention(mention string) string {
    return strings.TrimPrefix(mention, "@")
}
