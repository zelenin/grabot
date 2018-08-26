package client

import (
    "github.com/fatih/structs"
)

// Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.
type GetUpdatesRequest struct {
    // Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
    Offset *int64 `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
    // Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
    Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
    // Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
    Timeout *int64 `json:"timeout,omitempty" structs:"timeout,omitempty,omitnested"`
    // List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
    AllowedUpdates *[]UpdateType `json:"allowed_updates,omitempty" structs:"allowed_updates,omitempty,omitnested"`
}

// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
type SetWebhookRequest struct {
    // HTTPS url to send updates to. Use an empty string to remove webhook integration
    Url string `json:"url" structs:"url,omitnested"`
    // Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
    Certificate InputFile `json:"certificate,omitempty" structs:"certificate,omitempty,omitnested"`
    // Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput.
    MaxConnections *int64 `json:"max_connections,omitempty" structs:"max_connections,omitempty,omitnested"`
    // List the types of updates you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
    AllowedUpdates *[]UpdateType `json:"allowed_updates,omitempty" structs:"allowed_updates,omitempty,omitnested"`
}

// Use this method to send text messages. On success, the sent Message is returned.
type SendMessageRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Text of the message to be sent
    Text string `json:"text" structs:"text,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Disables link previews for links in this message
    DisableWebPagePreview *bool `json:"disable_web_page_preview,omitempty" structs:"disable_web_page_preview,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to forward messages of any kind. On success, the sent Message is returned.
type ForwardMessageRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
    FromChatId ChatId `json:"from_chat_id" structs:"from_chat_id,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // Message identifier in the chat specified in from_chat_id
    MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
}

// Use this method to send photos. On success, the sent Message is returned.
type SendPhotoRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. More info on Sending Files »
    Photo InputFile `json:"photo" structs:"photo,omitnested"`
    // Photo caption (may also be used when resending photos by file_id), 0-200 characters
    Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
type SendAudioRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
    Audio InputFile `json:"audio" structs:"audio,omitnested"`
    // Audio caption, 0-200 characters
    Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Duration of the audio in seconds
    Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
    // Performer
    Performer *string `json:"performer,omitempty" structs:"performer,omitempty,omitnested"`
    // Track name
    Title *string `json:"title,omitempty" structs:"title,omitempty,omitnested"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty" structs:"thumb,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocumentRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
    Document InputFile `json:"document" structs:"document,omitnested"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty" structs:"thumb,omitempty,omitnested"`
    // Document caption (may also be used when resending documents by file_id), 0-200 characters
    Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideoRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
    Video InputFile `json:"video" structs:"video,omitnested"`
    // Duration of sent video in seconds
    Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
    // Video width
    Width *int64 `json:"width,omitempty" structs:"width,omitempty,omitnested"`
    // Video height
    Height *int64 `json:"height,omitempty" structs:"height,omitempty,omitnested"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty" structs:"thumb,omitempty,omitnested"`
    // Video caption (may also be used when resending videos by file_id), 0-200 characters
    Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Pass True, if the uploaded video is suitable for streaming
    SupportsStreaming *bool `json:"supports_streaming,omitempty" structs:"supports_streaming,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimationRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
    Animation InputFile `json:"animation" structs:"animation,omitnested"`
    // Duration of sent animation in seconds
    Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
    // Animation width
    Width *int64 `json:"width,omitempty" structs:"width,omitempty,omitnested"`
    // Animation height
    Height *int64 `json:"height,omitempty" structs:"height,omitempty,omitnested"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty" structs:"thumb,omitempty,omitnested"`
    // Animation caption (may also be used when resending animation by file_id), 0-200 characters
    Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoiceRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
    Voice InputFile `json:"voice" structs:"voice,omitnested"`
    // Voice message caption, 0-200 characters
    Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Duration of the voice message in seconds
    Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNoteRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
    VideoNote InputFile `json:"video_note" structs:"video_note,omitnested"`
    // Duration of sent video in seconds
    Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
    // Video width and height
    Length *int64 `json:"length,omitempty" structs:"length,omitempty,omitnested"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty" structs:"thumb,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
type SendMediaGroupRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
    Media []InputMediaGroup `json:"media" structs:"media,omitnested"`
    // Sends the messages silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the messages are a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
}

// Use this method to send point on the map. On success, the sent Message is returned.
type SendLocationRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Latitude of the location
    Latitude float64 `json:"latitude" structs:"latitude,omitnested"`
    // Longitude of the location
    Longitude float64 `json:"longitude" structs:"longitude,omitnested"`
    // Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
    LivePeriod *int64 `json:"live_period,omitempty" structs:"live_period,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit live location messages sent by the bot or via the bot (for inline bots). A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocationRequest struct {
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
    // Latitude of new location
    Latitude float64 `json:"latitude" structs:"latitude,omitnested"`
    // Longitude of new location
    Longitude float64 `json:"longitude" structs:"longitude,omitnested"`
    // A JSON-serialized object for a new inline keyboard.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
type StopMessageLiveLocationRequest struct {
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
    // A JSON-serialized object for a new inline keyboard.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send information about a venue. On success, the sent Message is returned.
type SendVenueRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Latitude of the venue
    Latitude float64 `json:"latitude" structs:"latitude,omitnested"`
    // Longitude of the venue
    Longitude float64 `json:"longitude" structs:"longitude,omitnested"`
    // Name of the venue
    Title string `json:"title" structs:"title,omitnested"`
    // Address of the venue
    Address string `json:"address" structs:"address,omitnested"`
    // Foursquare identifier of the venue
    FoursquareId *string `json:"foursquare_id,omitempty" structs:"foursquare_id,omitempty,omitnested"`
    // Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
    FoursquareType *string `json:"foursquare_type,omitempty" structs:"foursquare_type,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send phone contacts. On success, the sent Message is returned.
type SendContactRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Contact's phone number
    PhoneNumber string `json:"phone_number" structs:"phone_number,omitnested"`
    // Contact's first name
    FirstName string `json:"first_name" structs:"first_name,omitnested"`
    // Contact's last name
    LastName *string `json:"last_name,omitempty" structs:"last_name,omitempty,omitnested"`
    // Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard *string `json:"vcard,omitempty" structs:"vcard,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

type ChatAction string

func (chatAction ChatAction) String() string {
    return string(chatAction)
}

const (
    ChatActionTyping          ChatAction = "typing"
    ChatActionUploadPhoto     ChatAction = "upload_photo"
    ChatActionRecordVideo     ChatAction = "record_video"
    ChatActionUploadVideo     ChatAction = "upload_video"
    ChatActionRecordAudio     ChatAction = "record_audio"
    ChatActionUploadAudio     ChatAction = "upload_audio"
    ChatActionUploadDocument  ChatAction = "upload_document"
    ChatActionFindLocation    ChatAction = "find_location"
    ChatActionRecordVideoNote ChatAction = "record_video_note"
    ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
type SendChatActionRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
    Action ChatAction `json:"action" structs:"action,omitnested"`
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotosRequest struct {
    // Unique identifier of the target user
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Sequential number of the first photo to be returned. By default, all photos are returned.
    Offset *int64 `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
    // Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
    Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
}

// Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
type GetFileRequest struct {
    // File identifier to get info about
    FileId string `json:"file_id" structs:"file_id,omitnested"`
}

// Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type KickChatMemberRequest struct {
    // Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Unique identifier of the target user
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
    UntilDate *int64 `json:"until_date,omitempty" structs:"until_date,omitempty,omitnested"`
}

// Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.
type UnbanChatMemberRequest struct {
    // Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Unique identifier of the target user
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.
type RestrictChatMemberRequest struct {
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Unique identifier of the target user
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
    UntilDate *int64 `json:"until_date,omitempty" structs:"until_date,omitempty,omitnested"`
    // Pass True, if the user can send text messages, contacts, locations and venues
    CanSendMessages *bool `json:"can_send_messages,omitempty" structs:"can_send_messages,omitempty,omitnested"`
    // Pass True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
    CanSendMediaMessages *bool `json:"can_send_media_messages,omitempty" structs:"can_send_media_messages,omitempty,omitnested"`
    // Pass True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
    CanSendOtherMessages *bool `json:"can_send_other_messages,omitempty" structs:"can_send_other_messages,omitempty,omitnested"`
    // Pass True, if the user may add web page previews to their messages, implies can_send_media_messages
    CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty" structs:"can_add_web_page_previews,omitempty,omitnested"`
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMemberRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Unique identifier of the target user
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Pass True, if the administrator can change chat title, photo and other settings
    CanChangeInfo *bool `json:"can_change_info,omitempty" structs:"can_change_info,omitempty,omitnested"`
    // Pass True, if the administrator can create channel posts, channels only
    CanPostMessages *bool `json:"can_post_messages,omitempty" structs:"can_post_messages,omitempty,omitnested"`
    // Pass True, if the administrator can edit messages of other users and can pin messages, channels only
    CanEditMessages *bool `json:"can_edit_messages,omitempty" structs:"can_edit_messages,omitempty,omitnested"`
    // Pass True, if the administrator can delete messages of other users
    CanDeleteMessages *bool `json:"can_delete_messages,omitempty" structs:"can_delete_messages,omitempty,omitnested"`
    // Pass True, if the administrator can invite new users to the chat
    CanInviteUsers *bool `json:"can_invite_users,omitempty" structs:"can_invite_users,omitempty,omitnested"`
    // Pass True, if the administrator can restrict, ban or unban chat members
    CanRestrictMembers *bool `json:"can_restrict_members,omitempty" structs:"can_restrict_members,omitempty,omitnested"`
    // Pass True, if the administrator can pin messages, supergroups only
    CanPinMessages *bool `json:"can_pin_messages,omitempty" structs:"can_pin_messages,omitempty,omitnested"`
    // Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
    CanPromoteMembers *bool `json:"can_promote_members,omitempty" structs:"can_promote_members,omitempty,omitnested"`
}

// Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.
type ExportChatInviteLinkRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatPhotoRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // New chat photo, uploaded using multipart/form-data
    Photo InputFile `json:"photo" structs:"photo,omitnested"`
}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type DeleteChatPhotoRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatTitleRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // New chat title, 1-255 characters
    Title string `json:"title" structs:"title,omitnested"`
}

// Use this method to change the description of a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatDescriptionRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // New chat description, 0-255 characters
    Description *string `json:"description,omitempty" structs:"description,omitempty,omitnested"`
}

// Use this method to pin a message in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
type PinChatMessageRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Identifier of a message to pin
    MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
    // Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
}

// Use this method to unpin a message in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
type UnpinChatMessageRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChatRequest struct {
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChatRequest struct {
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
type GetChatAdministratorsRequest struct {
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get the number of members in a chat. Returns Int on success.
type GetChatMembersCountRequest struct {
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get information about a member of a chat. Returns a ChatMember object on success.
type GetChatMemberRequest struct {
    // Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Unique identifier of the target user
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSetRequest struct {
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Name of the sticker set to be set as the group sticker set
    StickerSetName string `json:"sticker_set_name" structs:"sticker_set_name,omitnested"`
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSetRequest struct {
    // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
type AnswerCallbackQueryRequest struct {
    // Unique identifier for the query to be answered
    CallbackQueryId string `json:"callback_query_id" structs:"callback_query_id,omitnested"`
    // Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
    Text *string `json:"text,omitempty" structs:"text,omitempty,omitnested"`
    // If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
    ShowAlert *bool `json:"show_alert,omitempty" structs:"show_alert,omitempty,omitnested"`
    // URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
    Url *string `json:"url,omitempty" structs:"url,omitempty,omitnested"`
    // The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
    CacheTime *int64 `json:"cache_time,omitempty" structs:"cache_time,omitempty,omitnested"`
}

// Use this method to edit text and game messages sent by the bot or via the bot (for inline bots). On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageTextRequest struct {
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
    // New text of the message
    Text string `json:"text" structs:"text,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // Disables link previews for links in this message
    DisableWebPagePreview *bool `json:"disable_web_page_preview,omitempty" structs:"disable_web_page_preview,omitempty,omitnested"`
    // A JSON-serialized object for an inline keyboard.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit captions of messages sent by the bot or via the bot (for inline bots). On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageCaptionRequest struct {
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
    // New caption of the message
    Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
    // A JSON-serialized object for an inline keyboard.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit audio, document, photo, or video messages. If a message is a part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageMediaRequest struct {
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
    // A JSON-serialized object for a new media content of the message
    Media InputMedia `json:"media,omitempty" structs:"media,omitempty,omitnested"`
    // A JSON-serialized object for a new inline keyboard.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit only the reply markup of messages sent by the bot or via the bot (for inline bots).  On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageReplyMarkupRequest struct {
    // Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
    // A JSON-serialized object for an inline keyboard.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Bots can delete outgoing messages in groups and supergroups.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
type DeleteMessageRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Identifier of the message to delete
    MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
}

// Use this method to send .webp stickers. On success, the sent Message is returned.
type SendStickerRequest struct {
    // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
    ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
    // Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
    Sticker InputFile `json:"sticker" structs:"sticker,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
    ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
type GetStickerSetRequest struct {
    // Name of the sticker set
    Name string `json:"name" structs:"name,omitnested"`
}

// Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
type UploadStickerFileRequest struct {
    // User identifier of sticker file owner
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »
    PngSticker InputFile `json:"png_sticker" structs:"png_sticker,omitnested"`
}

// Use this method to create new sticker set owned by a user. The bot will be able to edit the created sticker set. Returns True on success.
type CreateNewStickerSetRequest struct {
    // User identifier of created sticker set owner
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
    Name string `json:"name" structs:"name,omitnested"`
    // Sticker set title, 1-64 characters
    Title string `json:"title" structs:"title,omitnested"`
    // Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
    PngSticker InputFile `json:"png_sticker" structs:"png_sticker,omitnested"`
    // One or more emoji corresponding to the sticker
    Emojis string `json:"emojis" structs:"emojis,omitnested"`
    // Pass True, if a set of mask stickers should be created
    ContainsMasks *bool `json:"contains_masks,omitempty" structs:"contains_masks,omitempty,omitnested"`
    // A JSON-serialized object for position where the mask should be placed on faces
    MaskPosition *MaskPosition `json:"mask_position,omitempty" structs:"mask_position,omitempty,omitnested"`
}

// Use this method to add a new sticker to a set created by the bot. Returns True on success.
type AddStickerToSetRequest struct {
    // User identifier of sticker set owner
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Sticker set name
    Name string `json:"name" structs:"name,omitnested"`
    // Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
    PngSticker InputFile `json:"png_sticker" structs:"png_sticker,omitnested"`
    // One or more emoji corresponding to the sticker
    Emojis string `json:"emojis" structs:"emojis,omitnested"`
    // A JSON-serialized object for position where the mask should be placed on faces
    MaskPosition *MaskPosition `json:"mask_position,omitempty" structs:"mask_position,omitempty,omitnested"`
}

// Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.
type SetStickerPositionInSetRequest struct {
    // File identifier of the sticker
    Sticker string `json:"sticker" structs:"sticker,omitnested"`
    // New sticker position in the set, zero-based
    Position int64 `json:"position" structs:"position,omitnested"`
}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
type DeleteStickerFromSetRequest struct {
    // File identifier of the sticker
    Sticker string `json:"sticker" structs:"sticker,omitnested"`
}

// Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
type AnswerInlineQueryRequest struct {
    // Unique identifier for the answered query
    InlineQueryId string `json:"inline_query_id" structs:"inline_query_id,omitnested"`
    // A JSON-serialized array of results for the inline query
    Results []InlineQueryResult `json:"results" structs:"results,omitnested"`
    // The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
    CacheTime *int64 `json:"cache_time,omitempty" structs:"cache_time,omitempty,omitnested"`
    // Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
    IsPersonal *bool `json:"is_personal,omitempty" structs:"is_personal,omitempty,omitnested"`
    // Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
    NextOffset *string `json:"next_offset,omitempty" structs:"next_offset,omitempty,omitnested"`
    // If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
    SwitchPmText *string `json:"switch_pm_text,omitempty" structs:"switch_pm_text,omitempty,omitnested"`
    // Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
    SwitchPmParameter *string `json:"switch_pm_parameter,omitempty" structs:"switch_pm_parameter,omitempty,omitnested"`
}

// Use this method to send invoices. On success, the sent Message is returned.
type SendInvoiceRequest struct {
    // Unique identifier for the target private chat
    ChatId int64 `json:"chat_id" structs:"chat_id,omitnested"`
    // Product name, 1-32 characters
    Title string `json:"title" structs:"title,omitnested"`
    // Product description, 1-255 characters
    Description string `json:"description" structs:"description,omitnested"`
    // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
    Payload string `json:"payload" structs:"payload,omitnested"`
    // Payments provider token, obtained via Botfather
    ProviderToken string `json:"provider_token" structs:"provider_token,omitnested"`
    // Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
    StartParameter string `json:"start_parameter" structs:"start_parameter,omitnested"`
    // Three-letter ISO 4217 currency code, see more on currencies
    Currency string `json:"currency" structs:"currency,omitnested"`
    // Price breakdown, a list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
    Prices []LabeledPrice `json:"prices" structs:"prices,omitnested"`
    // JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
    ProviderData *string `json:"provider_data,omitempty" structs:"provider_data,omitempty,omitnested"`
    // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
    PhotoUrl *string `json:"photo_url,omitempty" structs:"photo_url,omitempty,omitnested"`
    // Photo size
    PhotoSize *int64 `json:"photo_size,omitempty" structs:"photo_size,omitempty,omitnested"`
    // Photo width
    PhotoWidth *int64 `json:"photo_width,omitempty" structs:"photo_width,omitempty,omitnested"`
    // Photo height
    PhotoHeight *int64 `json:"photo_height,omitempty" structs:"photo_height,omitempty,omitnested"`
    // Pass True, if you require the user's full name to complete the order
    NeedName *bool `json:"need_name,omitempty" structs:"need_name,omitempty,omitnested"`
    // Pass True, if you require the user's phone number to complete the order
    NeedPhoneNumber *bool `json:"need_phone_number,omitempty" structs:"need_phone_number,omitempty,omitnested"`
    // Pass True, if you require the user's email address to complete the order
    NeedEmail *bool `json:"need_email,omitempty" structs:"need_email,omitempty,omitnested"`
    // Pass True, if you require the user's shipping address to complete the order
    NeedShippingAddress *bool `json:"need_shipping_address,omitempty" structs:"need_shipping_address,omitempty,omitnested"`
    // Pass True, if user's phone number should be sent to provider
    SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty" structs:"send_phone_number_to_provider,omitempty,omitnested"`
    // Pass True, if user's email address should be sent to provider
    SendEmailToProvider *bool `json:"send_email_to_provider,omitempty" structs:"send_email_to_provider,omitempty,omitnested"`
    // Pass True, if the final price depends on the shipping method
    IsFlexible *bool `json:"is_flexible,omitempty" structs:"is_flexible,omitempty,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
type AnswerShippingQueryRequest struct {
    // Unique identifier for the query to be answered
    ShippingQueryId string `json:"shipping_query_id" structs:"shipping_query_id,omitnested"`
    // Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
    Ok bool `json:"ok" structs:"ok,omitnested"`
    // Required if ok is True. A JSON-serialized array of available shipping options.
    ShippingOptions *[]ShippingOption `json:"shipping_options,omitempty" structs:"shipping_options,omitempty,omitnested"`
    // Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
    ErrorMessage *string `json:"error_message,omitempty" structs:"error_message,omitempty,omitnested"`
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQueryRequest struct {
    // Unique identifier for the query to be answered
    PreCheckoutQueryId string `json:"pre_checkout_query_id" structs:"pre_checkout_query_id,omitnested"`
    // Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
    Ok bool `json:"ok" structs:"ok,omitnested"`
    // Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
    ErrorMessage *string `json:"error_message,omitempty" structs:"error_message,omitempty,omitnested"`
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
//Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
type SetPassportDataErrorsRequest struct {
    // User identifier
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // A JSON-serialized array describing the errors
    Errors []PassportElementError `json:"errors" structs:"errors,omitnested"`
}

// Use this method to send a game. On success, the sent Message is returned.
type SendGameRequest struct {
    // Unique identifier for the target chat
    ChatId int64 `json:"chat_id" structs:"chat_id,omitnested"`
    // Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
    GameShortName string `json:"game_short_name" structs:"game_short_name,omitnested"`
    // Sends the message silently. Users will receive a notification with no sound.
    DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
    // If the message is a reply, ID of the original message
    ReplyToMessageId *int64 `json:"reply_to_message_id,omitempty" structs:"reply_to_message_id,omitempty,omitnested"`
    // A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScoreRequest struct {
    // User identifier
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // New score, must be non-negative
    Score int64 `json:"score" structs:"score,omitnested"`
    // Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
    Force *bool `json:"force,omitempty" structs:"force,omitempty,omitnested"`
    // Pass True, if the game message should not be automatically edited to include the current scoreboard
    DisableEditMessage *bool `json:"disable_edit_message,omitempty" structs:"disable_edit_message,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Unique identifier for the target chat
    ChatId *int64 `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
type GetGameHighScoresRequest struct {
    // Target user id
    UserId int64 `json:"user_id" structs:"user_id,omitnested"`
    // Required if inline_message_id is not specified. Unique identifier for the target chat
    ChatId *int64 `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
    // Required if inline_message_id is not specified. Identifier of the sent message
    MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
    // Required if chat_id and message_id are not specified. Identifier of the inline message
    InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
}

func requestToMap(request interface{}) map[string]interface{} {
    return structs.Map(request)
}
