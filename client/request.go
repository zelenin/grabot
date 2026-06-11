package client

// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
type GetUpdatesRequest struct {
	// Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will be forgotten.
	Offset *int64 `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
	// Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
	// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	Timeout *int64 `json:"timeout,omitempty" structs:"timeout,omitempty,omitnested"`
	// A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty" structs:"allowed_updates,omitempty,omitnested"`
}

// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request (a request with response HTTP status code different from 2XY), we will repeat the request and give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header "X-Telegram-Bot-Api-Secret-Token" with the secret token as content.
type SetWebhookRequest struct {
	// HTTPS URL to send updates to. Use an empty string to remove webhook integration.
	Url string `json:"url" structs:"url,omitnested"`
	// Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	Certificate InputFile `json:"certificate,omitempty" structs:"certificate,omitempty,omitnested"`
	// The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	IpAddress *string `json:"ip_address,omitempty" structs:"ip_address,omitempty,omitnested"`
	// The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	MaxConnections *int64 `json:"max_connections,omitempty" structs:"max_connections,omitempty,omitnested"`
	// A JSON-serialized list of the update types you want your bot to receive. For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default). If not specified, the previous setting will be used. Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty" structs:"allowed_updates,omitempty,omitnested"`
	// Pass True to drop all pending updates
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty" structs:"drop_pending_updates,omitempty,omitnested"`
	// A secret token to be sent in a header "X-Telegram-Bot-Api-Secret-Token" in every webhook request, 1-256 characters. Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
	SecretToken *string `json:"secret_token,omitempty" structs:"secret_token,omitempty,omitnested"`
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
type DeleteWebhookRequest struct {
	// Pass True to drop all pending updates
	DropPendingUpdates *bool `json:"drop_pending_updates,omitempty" structs:"drop_pending_updates,omitempty,omitnested"`
}

// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
type GetWebhookInfoRequest struct{}

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
type GetMeRequest struct{}

// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
type LogOutRequest struct{}

// Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
type CloseRequest struct{}

// Use this method to send text messages. On success, the sent Message is returned.
type SendMessageRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text" structs:"text,omitnested"`
	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty" structs:"entities,omitempty,omitnested"`
	// Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty" structs:"link_preview_options,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to forward messages of any kind. Service messages and messages with protected content can't be forwarded. On success, the sent Message is returned.
type ForwardMessageRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be forwarded; required if the message is forwarded to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Unique identifier for the chat where the original message was sent (or username of the target bot, supergroup or channel in the format @username)
	FromChatId ChatId `json:"from_chat_id" structs:"from_chat_id,omitnested"`
	// New start timestamp for the forwarded video in the message
	VideoStartTimestamp *int64 `json:"video_start_timestamp,omitempty" structs:"video_start_timestamp,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the forwarded message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; only available when forwarding to private chats
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Message identifier in the chat specified in from_chat_id
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
}

// Use this method to forward multiple messages of any kind. If some of the specified messages can't be found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
type ForwardMessagesRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the messages will be forwarded; required if the messages are forwarded to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Unique identifier for the chat where the original messages were sent (or username of the target bot, supergroup or channel in the format @username)
	FromChatId ChatId `json:"from_chat_id" structs:"from_chat_id,omitnested"`
	// A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward. The identifiers must be specified in a strictly increasing order.
	MessageIds []int64 `json:"message_ids" structs:"message_ids,omitnested"`
	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
}

// Use this method to copy messages of any kind. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
type CopyMessageRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Unique identifier for the chat where the original message was sent (or username of the target bot, supergroup or channel in the format @username)
	FromChatId ChatId `json:"from_chat_id" structs:"from_chat_id,omitnested"`
	// Message identifier in the chat specified in from_chat_id
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// New start timestamp for the copied video in the message
	VideoStartTimestamp *int64 `json:"video_start_timestamp,omitempty" structs:"video_start_timestamp,omitempty,omitnested"`
	// New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept.
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the new caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Pass True, if the caption must be shown above the message media. Ignored if a new caption isn't specified.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty" structs:"show_caption_above_media,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; only available when copying to private chats
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to copy messages of any kind. If some of the specified messages can't be found or copied, they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
type CopyMessagesRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the messages will be sent; required if the messages are sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Unique identifier for the chat where the original messages were sent (or username of the target bot, supergroup or channel in the format @username)
	FromChatId ChatId `json:"from_chat_id" structs:"from_chat_id,omitnested"`
	// A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy. The identifiers must be specified in a strictly increasing order.
	MessageIds []int64 `json:"message_ids" structs:"message_ids,omitnested"`
	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to copy the messages without their captions
	RemoveCaption *bool `json:"remove_caption,omitempty" structs:"remove_caption,omitempty,omitnested"`
}

// Use this method to send photos. On success, the sent Message is returned.
type SendPhotoRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total. Width and height ratio must be at most 20. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Photo InputFile `json:"photo" structs:"photo,omitnested"`
	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the photo caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty" structs:"show_caption_above_media,omitempty,omitnested"`
	// Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty" structs:"has_spoiler,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send live photos. On success, the sent Message is returned.
type SendLivePhotoRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Live photo video to send. The video must be no longer than 10 seconds and must not exceed 10 MB in size. Pass a file_id as String to send a video that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Sending live photos by a URL is currently unsupported.
	LivePhoto InputFile `json:"live_photo" structs:"live_photo,omitnested"`
	// The static photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Sending live photos by a URL is currently unsupported.
	Photo InputFile `json:"photo" structs:"photo,omitnested"`
	// Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the video caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty" structs:"show_caption_above_media,omitempty,omitnested"`
	// Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty" structs:"has_spoiler,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
type SendAudioRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Audio InputFile `json:"audio" structs:"audio,omitnested"`
	// Audio caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the audio caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Duration of the audio in seconds
	Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
	// Performer
	Performer *string `json:"performer,omitempty" structs:"performer,omitempty,omitnested"`
	// Track name
	Title *string `json:"title,omitempty" structs:"title,omitempty,omitnested"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail InputFile `json:"thumbnail,omitempty" structs:"thumbnail,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocumentRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Document InputFile `json:"document" structs:"document,omitnested"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail InputFile `json:"thumbnail,omitempty" structs:"thumbnail,omitempty,omitnested"`
	// Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the document caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty" structs:"disable_content_type_detection,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideoRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Video InputFile `json:"video" structs:"video,omitnested"`
	// Duration of sent video in seconds
	Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
	// Video width
	Width *int64 `json:"width,omitempty" structs:"width,omitempty,omitnested"`
	// Video height
	Height *int64 `json:"height,omitempty" structs:"height,omitempty,omitnested"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail InputFile `json:"thumbnail,omitempty" structs:"thumbnail,omitempty,omitnested"`
	// Cover for the video in the message. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Cover InputFile `json:"cover,omitempty" structs:"cover,omitempty,omitnested"`
	// Start timestamp for the video in the message
	StartTimestamp *int64 `json:"start_timestamp,omitempty" structs:"start_timestamp,omitempty,omitnested"`
	// Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the video caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty" structs:"show_caption_above_media,omitempty,omitnested"`
	// Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty" structs:"has_spoiler,omitempty,omitnested"`
	// Pass True if the uploaded video is suitable for streaming
	SupportsStreaming *bool `json:"supports_streaming,omitempty" structs:"supports_streaming,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimationRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Animation InputFile `json:"animation" structs:"animation,omitnested"`
	// Duration of sent animation in seconds
	Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
	// Animation width
	Width *int64 `json:"width,omitempty" structs:"width,omitempty,omitnested"`
	// Animation height
	Height *int64 `json:"height,omitempty" structs:"height,omitempty,omitnested"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail InputFile `json:"thumbnail,omitempty" structs:"thumbnail,omitempty,omitnested"`
	// Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the animation caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty" structs:"show_caption_above_media,omitempty,omitnested"`
	// Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty" structs:"has_spoiler,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoiceRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Voice InputFile `json:"voice" structs:"voice,omitnested"`
	// Voice message caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the voice message caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Duration of the voice message in seconds
	Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNoteRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Sending video notes by a URL is currently unsupported.
	VideoNote InputFile `json:"video_note" structs:"video_note,omitnested"`
	// Duration of sent video in seconds
	Duration *int64 `json:"duration,omitempty" structs:"duration,omitempty,omitnested"`
	// Video width and height, i.e. diameter of the video message
	Length *int64 `json:"length,omitempty" structs:"length,omitempty,omitnested"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file, so you can pass "attach://<file_attach_name>" if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail InputFile `json:"thumbnail,omitempty" structs:"thumbnail,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send paid media. On success, the sent Message is returned.
type SendPaidMediaRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username. If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance. Otherwise, they will be credited to the bot's balance.
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// The number of Telegram Stars that must be paid to buy access to the media; 1-25000
	StarCount int64 `json:"star_count" structs:"star_count,omitnested"`
	// A JSON-serialized array describing the media to be sent; up to 10 items
	Media []InputPaidMedia `json:"media" structs:"media,omitnested"`
	// Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload *string `json:"payload,omitempty" structs:"payload,omitempty,omitnested"`
	// Media caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the media caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty" structs:"show_caption_above_media,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send a group of photos, live photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Message objects that were sent is returned.
type SendMediaGroupRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the messages will be sent; required if the messages are sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// A JSON-serialized array describing messages to be sent, must include 2-10 items
	Media any `json:"media" structs:"media,omitnested"`
	// Sends messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
}

// Use this method to send point on the map. On success, the sent Message is returned.
type SendLocationRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Latitude of the location
	Latitude float64 `json:"latitude" structs:"latitude,omitnested"`
	// Longitude of the location
	Longitude float64 `json:"longitude" structs:"longitude,omitnested"`
	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty" structs:"horizontal_accuracy,omitempty,omitnested"`
	// Period in seconds during which the location will be updated (see Live Locations, should be between 60 and 86400, or 0x7FFFFFFF for live locations that can be edited indefinitely
	LivePeriod *int64 `json:"live_period,omitempty" structs:"live_period,omitempty,omitnested"`
	// For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int64 `json:"heading,omitempty" structs:"heading,omitempty,omitnested"`
	// For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int64 `json:"proximity_alert_radius,omitempty" structs:"proximity_alert_radius,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send information about a venue. On success, the sent Message is returned.
type SendVenueRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
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
	// Foursquare type of the venue, if known. (For example, "arts_entertainment/default", "arts_entertainment/aquarium" or "food/icecream".)
	FoursquareType *string `json:"foursquare_type,omitempty" structs:"foursquare_type,omitempty,omitnested"`
	// Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty" structs:"google_place_id,omitempty,omitnested"`
	// Google Places type of the venue. (See supported types.)
	GooglePlaceType *string `json:"google_place_type,omitempty" structs:"google_place_type,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send phone contacts. On success, the sent Message is returned.
type SendContactRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
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
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send a native poll. On success, the sent Message is returned.
type SendPollRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username. Polls can't be sent to channel direct messages chats.
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Poll question, 1-300 characters
	Question string `json:"question" structs:"question,omitnested"`
	// Mode for parsing entities in the question. See formatting options for more details. Currently, only custom emoji entities are allowed.
	QuestionParseMode *string `json:"question_parse_mode,omitempty" structs:"question_parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the poll question. It can be specified instead of question_parse_mode.
	QuestionEntities []MessageEntity `json:"question_entities,omitempty" structs:"question_entities,omitempty,omitnested"`
	// A JSON-serialized list of 1-12 answer options
	Options []InputPollOption `json:"options" structs:"options,omitnested"`
	// True, if the poll needs to be anonymous, defaults to True
	IsAnonymous *bool `json:"is_anonymous,omitempty" structs:"is_anonymous,omitempty,omitnested"`
	// Poll type, "quiz" or "regular", defaults to "regular"
	Type *string `json:"type,omitempty" structs:"type,omitempty,omitnested"`
	// Pass True, if the poll allows multiple answers, defaults to False
	AllowsMultipleAnswers *bool `json:"allows_multiple_answers,omitempty" structs:"allows_multiple_answers,omitempty,omitnested"`
	// Pass True, if the poll allows to change chosen answer options, defaults to False for quizzes and to True for regular polls
	AllowsRevoting *bool `json:"allows_revoting,omitempty" structs:"allows_revoting,omitempty,omitnested"`
	// Pass True, if the poll options must be shown in random order
	ShuffleOptions *bool `json:"shuffle_options,omitempty" structs:"shuffle_options,omitempty,omitnested"`
	// Pass True, if answer options can be added to the poll after creation; not supported for anonymous polls and quizzes
	AllowAddingOptions *bool `json:"allow_adding_options,omitempty" structs:"allow_adding_options,omitempty,omitnested"`
	// Pass True, if poll results must be shown only after the poll closes
	HideResultsUntilCloses *bool `json:"hide_results_until_closes,omitempty" structs:"hide_results_until_closes,omitempty,omitnested"`
	// Pass True, if voting is limited to users who have been members of the chat where the poll is being sent for more than 24 hours; for channel chats only
	MembersOnly *bool `json:"members_only,omitempty" structs:"members_only,omitempty,omitnested"`
	// A JSON-serialized list of 0-12 two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which users can vote in the poll; for channel chats only. Use "FT" as a country code to allow users with anonymous numbers to vote. If omitted or empty, then users from any country can participate in the poll.
	CountryCodes []string `json:"country_codes,omitempty" structs:"country_codes,omitempty,omitnested"`
	// A JSON-serialized list of monotonically increasing 0-based identifiers of the correct answer options, required for polls in quiz mode
	CorrectOptionIds []int64 `json:"correct_option_ids,omitempty" structs:"correct_option_ids,omitempty,omitnested"`
	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters with at most 2 line feeds after entities parsing
	Explanation *string `json:"explanation,omitempty" structs:"explanation,omitempty,omitnested"`
	// Mode for parsing entities in the explanation. See formatting options for more details.
	ExplanationParseMode *string `json:"explanation_parse_mode,omitempty" structs:"explanation_parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the poll explanation. It can be specified instead of explanation_parse_mode.
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty" structs:"explanation_entities,omitempty,omitnested"`
	// Media added to the quiz explanation
	ExplanationMedia *InputPollMedia `json:"explanation_media,omitempty" structs:"explanation_media,omitempty,omitnested"`
	// Amount of time in seconds the poll will be active after creation, 5-2628000. Can't be used together with close_date.
	OpenPeriod *int64 `json:"open_period,omitempty" structs:"open_period,omitempty,omitnested"`
	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 2628000 seconds in the future. Can't be used together with open_period.
	CloseDate *int64 `json:"close_date,omitempty" structs:"close_date,omitempty,omitnested"`
	// Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed *bool `json:"is_closed,omitempty" structs:"is_closed,omitempty,omitnested"`
	// Description of the poll to be sent, 0-1024 characters after entities parsing
	Description *string `json:"description,omitempty" structs:"description,omitempty,omitnested"`
	// Mode for parsing entities in the poll description. See formatting options for more details.
	DescriptionParseMode *string `json:"description_parse_mode,omitempty" structs:"description_parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the poll description, which can be specified instead of description_parse_mode
	DescriptionEntities []MessageEntity `json:"description_entities,omitempty" structs:"description_entities,omitempty,omitnested"`
	// Media added to the poll description
	Media *InputPollMedia `json:"media,omitempty" structs:"media,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send a checklist on behalf of a connected business account. On success, the sent Message is returned.
type SendChecklistRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier for the target chat or username of the target bot in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// A JSON-serialized object for the checklist to send
	Checklist InputChecklist `json:"checklist" structs:"checklist,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object for description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// A JSON-serialized object for an inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendDiceRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Emoji on which the dice throw animation is based. Currently, must be one of "🎲", "🎯", "🏀", "⚽", "🎳", or "🎰". Dice can have values 1-6 for "🎲", "🎯" and "🎳", values 1-5 for "🏀" and "⚽", and values 1-64 for "🎰". Defaults to "🎲".
	Emoji *string `json:"emoji,omitempty" structs:"emoji,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to stream a partial message to a user while the message is being generated. Note that the streamed draft is ephemeral and acts as a temporary 30-second preview - once the output is finalized, you must call sendMessage with the complete message to persist it in the user's chat. Returns True on success.
type SendMessageDraftRequest struct {
	// Unique identifier for the target private chat
	ChatId int64 `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Unique identifier of the message draft; must be non-zero. Changes of drafts with the same identifier are animated.
	DraftId int64 `json:"draft_id" structs:"draft_id,omitnested"`
	// Text of the message to be sent, 0-4096 characters after entities parsing. Pass an empty text to show a "Thinking..." placeholder.
	Text *string `json:"text,omitempty" structs:"text,omitempty,omitnested"`
	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty" structs:"entities,omitempty,omitnested"`
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
type SendChatActionRequest struct {
	// Unique identifier of the business connection on behalf of which the action will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot or supergroup in the format @username. Channel chats and channel direct messages chats aren't supported.
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread or topic of a forum; for supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers, find_location for location data, record_video_note or upload_video_note for video notes.
	Action string `json:"action" structs:"action,omitnested"`
}

// Use this method to change the chosen reactions on a message. Service messages of some types can't be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Bots can't use paid reactions. Returns True on success.
type SetMessageReactionRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// A JSON-serialized list of reaction types to set on the message. Currently, as non-premium users, bots can set up to one reaction per message. A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators. Paid reactions can't be used by bots.
	Reaction []ReactionType `json:"reaction,omitempty" structs:"reaction,omitempty,omitnested"`
	// Pass True to set the reaction with a big animation
	IsBig *bool `json:"is_big,omitempty" structs:"is_big,omitempty,omitnested"`
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotosRequest struct {
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset *int64 `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
	// Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
}

// Use this method to get a list of profile audios for a user. Returns a UserProfileAudios object.
type GetUserProfileAudiosRequest struct {
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Sequential number of the first audio to be returned. By default, all audios are returned.
	Offset *int64 `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
	// Limits the number of audios to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess. Returns True on success.
type SetUserEmojiStatusRequest struct {
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status.
	EmojiStatusCustomEmojiId *string `json:"emoji_status_custom_emoji_id,omitempty" structs:"emoji_status_custom_emoji_id,omitempty,omitnested"`
	// Expiration date of the emoji status, if any
	EmojiStatusExpirationDate *int64 `json:"emoji_status_expiration_date,omitempty" structs:"emoji_status_expiration_date,omitempty,omitnested"`
}

// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
type GetFileRequest struct {
	// File identifier to get information about
	FileId string `json:"file_id" structs:"file_id,omitnested"`
}

// Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Date when the user will be unbanned; Unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever. Applied for supergroups and channels only.
	UntilDate *int64 `json:"until_date,omitempty" structs:"until_date,omitempty,omitnested"`
	// Pass True to delete all messages from the chat for the user that is being removed. If False, the user will be able to see messages in the group that were sent before the user was removed. Always True for supergroups and channels.
	RevokeMessages *bool `json:"revoke_messages,omitempty" structs:"revoke_messages,omitempty,omitnested"`
}

// Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
type UnbanChatMemberRequest struct {
	// Unique identifier for the target group or username of the target supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Do nothing if the user is not banned
	OnlyIfBanned *bool `json:"only_if_banned,omitempty" structs:"only_if_banned,omitempty,omitnested"`
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// A JSON-serialized object for new user permissions
	Permissions ChatPermissions `json:"permissions" structs:"permissions,omitnested"`
	// Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UseIndependentChatPermissions *bool `json:"use_independent_chat_permissions,omitempty" structs:"use_independent_chat_permissions,omitempty,omitnested"`
	// Date when restrictions will be lifted for the user; Unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever.
	UntilDate *int64 `json:"until_date,omitempty" structs:"until_date,omitempty,omitnested"`
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Pass True if the administrator's presence in the chat is hidden
	IsAnonymous *bool `json:"is_anonymous,omitempty" structs:"is_anonymous,omitempty,omitnested"`
	// Pass True if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages, ignore slow mode, and send messages to the chat without paying Telegram Stars. Implied by any other administrator privilege.
	CanManageChat *bool `json:"can_manage_chat,omitempty" structs:"can_manage_chat,omitempty,omitnested"`
	// Pass True if the administrator can delete messages of other users
	CanDeleteMessages *bool `json:"can_delete_messages,omitempty" structs:"can_delete_messages,omitempty,omitnested"`
	// Pass True if the administrator can manage video chats
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty" structs:"can_manage_video_chats,omitempty,omitnested"`
	// Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics. For backward compatibility, defaults to True for promotions of channel administrators.
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty" structs:"can_restrict_members,omitempty,omitnested"`
	// Pass True if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanPromoteMembers *bool `json:"can_promote_members,omitempty" structs:"can_promote_members,omitempty,omitnested"`
	// Pass True if the administrator can change chat title, photo and other settings
	CanChangeInfo *bool `json:"can_change_info,omitempty" structs:"can_change_info,omitempty,omitnested"`
	// Pass True if the administrator can invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty" structs:"can_invite_users,omitempty,omitnested"`
	// Pass True if the administrator can post stories to the chat
	CanPostStories *bool `json:"can_post_stories,omitempty" structs:"can_post_stories,omitempty,omitnested"`
	// Pass True if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories *bool `json:"can_edit_stories,omitempty" structs:"can_edit_stories,omitempty,omitnested"`
	// Pass True if the administrator can delete stories posted by other users
	CanDeleteStories *bool `json:"can_delete_stories,omitempty" structs:"can_delete_stories,omitempty,omitnested"`
	// Pass True if the administrator can post messages in the channel, approve suggested posts, or access channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty" structs:"can_post_messages,omitempty,omitnested"`
	// Pass True if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty" structs:"can_edit_messages,omitempty,omitnested"`
	// Pass True if the administrator can pin messages; for supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty" structs:"can_pin_messages,omitempty,omitnested"`
	// Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty" structs:"can_manage_topics,omitempty,omitnested"`
	// Pass True if the administrator can manage direct messages within the channel and decline suggested posts; for channels only
	CanManageDirectMessages *bool `json:"can_manage_direct_messages,omitempty" structs:"can_manage_direct_messages,omitempty,omitnested"`
	// Pass True if the administrator can edit the tags of regular members; for groups and supergroups only
	CanManageTags *bool `json:"can_manage_tags,omitempty" structs:"can_manage_tags,omitempty,omitnested"`
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
type SetChatAdministratorCustomTitleRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// New custom title for the administrator; 0-16 characters, emoji are not allowed
	CustomTitle string `json:"custom_title" structs:"custom_title,omitnested"`
}

// Use this method to set a tag for a regular member in a group or a supergroup. The bot must be an administrator in the chat for this to work and must have the can_manage_tags administrator right. Returns True on success.
type SetChatMemberTagRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// New tag for the member; 0-16 characters, emoji are not allowed
	Tag *string `json:"tag,omitempty" structs:"tag,omitempty,omitnested"`
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
type BanChatSenderChatRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target sender chat
	SenderChatId int64 `json:"sender_chat_id" structs:"sender_chat_id,omitnested"`
}

// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
type UnbanChatSenderChatRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target sender chat
	SenderChatId int64 `json:"sender_chat_id" structs:"sender_chat_id,omitnested"`
}

// Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
type SetChatPermissionsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// A JSON-serialized object for new default chat permissions
	Permissions ChatPermissions `json:"permissions" structs:"permissions,omitnested"`
	// Pass True if chat permissions are set independently. Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages, can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions; the can_send_polls permission will imply the can_send_messages permission.
	UseIndependentChatPermissions *bool `json:"use_independent_chat_permissions,omitempty" structs:"use_independent_chat_permissions,omitempty,omitnested"`
}

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
type ExportChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
type CreateChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty" structs:"name,omitempty,omitnested"`
	// Point in time (Unix timestamp) when the link will expire
	ExpireDate *int64 `json:"expire_date,omitempty" structs:"expire_date,omitempty,omitnested"`
	// The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int64 `json:"member_limit,omitempty" structs:"member_limit,omitempty,omitnested"`
	// True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified.
	CreatesJoinRequest *bool `json:"creates_join_request,omitempty" structs:"creates_join_request,omitempty,omitnested"`
}

// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
type EditChatInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// The invite link to edit
	InviteLink string `json:"invite_link" structs:"invite_link,omitnested"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty" structs:"name,omitempty,omitnested"`
	// Point in time (Unix timestamp) when the link will expire
	ExpireDate *int64 `json:"expire_date,omitempty" structs:"expire_date,omitempty,omitnested"`
	// The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int64 `json:"member_limit,omitempty" structs:"member_limit,omitempty,omitnested"`
	// True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified.
	CreatesJoinRequest *bool `json:"creates_join_request,omitempty" structs:"creates_join_request,omitempty,omitnested"`
}

// Use this method to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink. Returns the new invite link as a ChatInviteLink object.
type CreateChatSubscriptionInviteLinkRequest struct {
	// Unique identifier for the target channel chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty" structs:"name,omitempty,omitnested"`
	// The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPeriod int64 `json:"subscription_period" structs:"subscription_period,omitnested"`
	// The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-10000
	SubscriptionPrice int64 `json:"subscription_price" structs:"subscription_price,omitnested"`
}

// Use this method to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a ChatInviteLink object.
type EditChatSubscriptionInviteLinkRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// The invite link to edit
	InviteLink string `json:"invite_link" structs:"invite_link,omitnested"`
	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty" structs:"name,omitempty,omitnested"`
}

// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
type RevokeChatInviteLinkRequest struct {
	// Unique identifier of the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// The invite link to revoke
	InviteLink string `json:"invite_link" structs:"invite_link,omitnested"`
}

// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type ApproveChatJoinRequestRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
type DeclineChatJoinRequestRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// New chat photo, uploaded using multipart/form-data
	Photo InputFile `json:"photo" structs:"photo,omitnested"`
}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type DeleteChatPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatTitleRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// New chat title, 1-128 characters
	Title string `json:"title" structs:"title,omitnested"`
}

// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
type SetChatDescriptionRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// New chat description, 0-255 characters
	Description *string `json:"description,omitempty" structs:"description,omitempty,omitnested"`
}

// Use this method to add a message to the list of pinned messages in a chat. In private chats and channel direct messages chats, all non-service messages can be pinned. Conversely, the bot must be an administrator with the 'can_pin_messages' right or the 'can_edit_messages' right to pin messages in groups and channels respectively. Returns True on success.
type PinChatMessageRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be pinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of a message to pin
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// Pass True if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels and private chats.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
}

// Use this method to remove a message from the list of pinned messages in a chat. In private chats and channel direct messages chats, all messages can be unpinned. Conversely, the bot must be an administrator with the 'can_pin_messages' right or the 'can_edit_messages' right to unpin messages in groups and channels respectively. Returns True on success.
type UnpinChatMessageRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be unpinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of the message to unpin. Required if business_connection_id is specified. If not specified, the most recent pinned message (by sending date) will be unpinned.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
}

// Use this method to clear the list of pinned messages in a chat. In private chats and channel direct messages chats, no additional rights are required to unpin all pinned messages. Conversely, the bot must be an administrator with the 'can_pin_messages' right or the 'can_edit_messages' right to unpin all pinned messages in groups and channels respectively. Returns True on success.
type UnpinAllChatMessagesRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel in the format @username. Channel direct messages chats aren't supported; leave the corresponding channel instead.
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
type GetChatRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get a list of administrators in a chat. Returns an Array of ChatMember objects.
type GetChatAdministratorsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Pass True to additionally receive all bots that are administrators of the chat. By default, bots other than the current bot are omitted.
	ReturnBots *bool `json:"return_bots,omitempty" structs:"return_bots,omitempty,omitnested"`
}

// Use this method to get the number of members in a chat. Returns Int on success.
type GetChatMemberCountRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
type GetChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to get the last messages from the personal chat (i.e., the chat currently added to their profile) of a given user. On success, an array of Message objects is returned.
type GetUserPersonalChatMessagesRequest struct {
	// Unique identifier for the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// The maximum number of messages to return; 1-20
	Limit int64 `json:"limit" structs:"limit,omitnested"`
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name" structs:"sticker_set_name,omitnested"`
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSetRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
type GetForumTopicIconStickersRequest struct{}

// Use this method to create a topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator right. Returns information about the created topic as a ForumTopic object.
type CreateForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Topic name, 1-128 characters
	Name string `json:"name" structs:"name,omitnested"`
	// Color of the topic icon in RGB format. Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F).
	IconColor *int64 `json:"icon_color,omitempty" structs:"icon_color,omitempty,omitnested"`
	// Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty" structs:"icon_custom_emoji_id,omitempty,omitnested"`
}

// Use this method to edit name and icon of a topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type EditForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId int64 `json:"message_thread_id" structs:"message_thread_id,omitnested"`
	// New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept.
	Name *string `json:"name,omitempty" structs:"name,omitempty,omitnested"`
	// New unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers. Pass an empty string to remove the icon. If not specified, the current icon will be kept.
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty" structs:"icon_custom_emoji_id,omitempty,omitnested"`
}

// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type CloseForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId int64 `json:"message_thread_id" structs:"message_thread_id,omitnested"`
}

// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
type ReopenForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId int64 `json:"message_thread_id" structs:"message_thread_id,omitnested"`
}

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
type DeleteForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId int64 `json:"message_thread_id" structs:"message_thread_id,omitnested"`
}

// Use this method to clear the list of pinned messages in a forum topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
type UnpinAllForumTopicMessagesRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId int64 `json:"message_thread_id" structs:"message_thread_id,omitnested"`
}

// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type EditGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// New topic name, 1-128 characters
	Name string `json:"name" structs:"name,omitnested"`
}

// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type CloseGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
type ReopenGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
type HideGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
type UnhideGeneralForumTopicRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
type UnpinAllGeneralForumTopicMessagesRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
type AnswerCallbackQueryRequest struct {
	// Unique identifier for the query to be answered
	CallbackQueryId string `json:"callback_query_id" structs:"callback_query_id,omitnested"`
	// Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters.
	Text *string `json:"text,omitempty" structs:"text,omitempty,omitnested"`
	// If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert *bool `json:"show_alert,omitempty" structs:"show_alert,omitempty,omitnested"`
	// URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @BotFather, specify the URL that opens your game - note that this will only work if the query comes from a callback_game button. Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	Url *string `json:"url,omitempty" structs:"url,omitempty,omitnested"`
	// The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime *int64 `json:"cache_time,omitempty" structs:"cache_time,omitempty,omitnested"`
}

// Use this method to reply to a received guest message. On success, a SentGuestMessage object is returned.
type AnswerGuestQueryRequest struct {
	// Unique identifier for the query to be answered
	GuestQueryId string `json:"guest_query_id" structs:"guest_query_id,omitnested"`
	// A JSON-serialized object describing the message to be sent
	Result InlineQueryResult `json:"result" structs:"result,omitnested"`
}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
type GetUserChatBoostsRequest struct {
	// Unique identifier for the chat or username of the channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
type GetBusinessConnectionRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
}

// Use this method to get the token of a managed bot. Returns the token as String on success.
type GetManagedBotTokenRequest struct {
	// User identifier of the managed bot whose token will be returned
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to revoke the current token of a managed bot and generate a new one. Returns the new token as String on success.
type ReplaceManagedBotTokenRequest struct {
	// User identifier of the managed bot whose token will be replaced
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to get the access settings of a managed bot. Returns a BotAccessSettings object on success.
type GetManagedBotAccessSettingsRequest struct {
	// User identifier of the managed bot whose access settings will be returned
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Use this method to change the access settings of a managed bot. Returns True on success.
type SetManagedBotAccessSettingsRequest struct {
	// User identifier of the managed bot whose access settings will be changed
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Pass True, if only selected users can access the bot. The bot's owner can always access it.
	IsAccessRestricted bool `json:"is_access_restricted" structs:"is_access_restricted,omitnested"`
	// A JSON-serialized list of up to 10 identifiers of users who will have access to the bot in addition to its owner. Ignored if is_access_restricted is false.
	AddedUserIds []int64 `json:"added_user_ids,omitempty" structs:"added_user_ids,omitempty,omitnested"`
}

// Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
type SetMyCommandsRequest struct {
	// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Commands []BotCommand `json:"commands" structs:"commands,omitnested"`
	// A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope,omitempty" structs:"scope,omitempty,omitnested"`
	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands.
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
type DeleteMyCommandsRequest struct {
	// A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope,omitempty" structs:"scope,omitempty,omitnested"`
	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands.
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
type GetMyCommandsRequest struct {
	// A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope,omitempty" structs:"scope,omitempty,omitnested"`
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to change the bot's name. Returns True on success.
type SetMyNameRequest struct {
	// New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	Name *string `json:"name,omitempty" structs:"name,omitempty,omitnested"`
	// A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to get the current bot name for the given user language. Returns BotName on success.
type GetMyNameRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
type SetMyDescriptionRequest struct {
	// New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	Description *string `json:"description,omitempty" structs:"description,omitempty,omitnested"`
	// A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
type GetMyDescriptionRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
type SetMyShortDescriptionRequest struct {
	// New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	ShortDescription *string `json:"short_description,omitempty" structs:"short_description,omitempty,omitnested"`
	// A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
type GetMyShortDescriptionRequest struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty" structs:"language_code,omitempty,omitnested"`
}

// Changes the profile photo of the bot. Returns True on success.
type SetMyProfilePhotoRequest struct {
	// The new profile photo to set
	Photo InputProfilePhoto `json:"photo" structs:"photo,omitnested"`
}

// Removes the profile photo of the bot. Requires no parameters. Returns True on success.
type RemoveMyProfilePhotoRequest struct{}

// Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
type SetChatMenuButtonRequest struct {
	// Unique identifier for the target private chat. If not specified, the bot's default menu button will be changed.
	ChatId *int64 `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault.
	MenuButton *MenuButton `json:"menu_button,omitempty" structs:"menu_button,omitempty,omitnested"`
}

// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
type GetChatMenuButtonRequest struct {
	// Unique identifier for the target private chat. If not specified, the bot's default menu button will be returned.
	ChatId *int64 `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
}

// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
type SetMyDefaultAdministratorRightsRequest struct {
	// A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	Rights *ChatAdministratorRights `json:"rights,omitempty" structs:"rights,omitempty,omitnested"`
	// Pass True to change the default administrator rights of the bot in channels. Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
	ForChannels *bool `json:"for_channels,omitempty" structs:"for_channels,omitempty,omitnested"`
}

// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
type GetMyDefaultAdministratorRightsRequest struct {
	// Pass True to get default administrator rights of the bot in channels. Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
	ForChannels *bool `json:"for_channels,omitempty" structs:"for_channels,omitempty,omitnested"`
}

// Returns the list of gifts that can be sent by the bot to users and channel chats. Requires no parameters. Returns a Gifts object.
type GetAvailableGiftsRequest struct{}

// Sends a gift to the given user or channel chat. The gift can't be converted to Telegram Stars by the receiver. Returns True on success.
type SendGiftRequest struct {
	// Required if chat_id is not specified. Unique identifier of the target user who will receive the gift.
	UserId *int64 `json:"user_id,omitempty" structs:"user_id,omitempty,omitnested"`
	// Required if user_id is not specified. Unique identifier for the chat or username of the channel (in the format @username) that will receive the gift.
	ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Identifier of the gift; limited gifts can't be sent to channel chats
	GiftId string `json:"gift_id" structs:"gift_id,omitnested"`
	// Pass True to pay for the gift upgrade from the bot's balance, thereby making the upgrade free for the receiver
	PayForUpgrade *bool `json:"pay_for_upgrade,omitempty" structs:"pay_for_upgrade,omitempty,omitnested"`
	// Text that will be shown along with the gift; 0-128 characters
	Text *string `json:"text,omitempty" structs:"text,omitempty,omitnested"`
	// Mode for parsing entities in the text. See formatting options for more details. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", "custom_emoji", and "date_time" are ignored.
	TextParseMode *string `json:"text_parse_mode,omitempty" structs:"text_parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of text_parse_mode. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", "custom_emoji", and "date_time" are ignored.
	TextEntities []MessageEntity `json:"text_entities,omitempty" structs:"text_entities,omitempty,omitnested"`
}

// Gifts a Telegram Premium subscription to the given user. Returns True on success.
type GiftPremiumSubscriptionRequest struct {
	// Unique identifier of the target user who will receive a Telegram Premium subscription
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Number of months the Telegram Premium subscription will be active for the user; must be one of 3, 6, or 12
	MonthCount int64 `json:"month_count" structs:"month_count,omitnested"`
	// Number of Telegram Stars to pay for the Telegram Premium subscription; must be 1000 for 3 months, 1500 for 6 months, and 2500 for 12 months
	StarCount int64 `json:"star_count" structs:"star_count,omitnested"`
	// Text that will be shown along with the service message about the subscription; 0-128 characters
	Text *string `json:"text,omitempty" structs:"text,omitempty,omitnested"`
	// Mode for parsing entities in the text. See formatting options for more details. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", "custom_emoji", and "date_time" are ignored.
	TextParseMode *string `json:"text_parse_mode,omitempty" structs:"text_parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the gift text. It can be specified instead of text_parse_mode. Entities other than "bold", "italic", "underline", "strikethrough", "spoiler", "custom_emoji", and "date_time" are ignored.
	TextEntities []MessageEntity `json:"text_entities,omitempty" structs:"text_entities,omitempty,omitnested"`
}

// Verifies a user on behalf of the organization which is represented by the bot. Returns True on success.
type VerifyUserRequest struct {
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Custom description for the verification; 0-70 characters. Must be empty if the organization isn't allowed to provide a custom verification description.
	CustomDescription *string `json:"custom_description,omitempty" structs:"custom_description,omitempty,omitnested"`
}

// Verifies a chat on behalf of the organization which is represented by the bot. Returns True on success.
type VerifyChatRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username. Channel direct messages chats can't be verified.
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Custom description for the verification; 0-70 characters. Must be empty if the organization isn't allowed to provide a custom verification description.
	CustomDescription *string `json:"custom_description,omitempty" structs:"custom_description,omitempty,omitnested"`
}

// Removes verification from a user who is currently verified on behalf of the organization represented by the bot. Returns True on success.
type RemoveUserVerificationRequest struct {
	// Unique identifier of the target user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
}

// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot. Returns True on success.
type RemoveChatVerificationRequest struct {
	// Unique identifier for the target chat or username of the target bot or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
}

// Marks incoming message as read on behalf of a business account. Requires the can_read_messages business bot right. Returns True on success.
type ReadBusinessMessageRequest struct {
	// Unique identifier of the business connection on behalf of which to read the message
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier of the chat in which the message was received. The chat must have been active in the last 24 hours.
	ChatId int64 `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier of the message to mark as read
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
}

// Delete messages on behalf of a business account. Requires the can_delete_sent_messages business bot right to delete messages sent by the bot itself, or the can_delete_all_messages business bot right to delete any message. Returns True on success.
type DeleteBusinessMessagesRequest struct {
	// Unique identifier of the business connection on behalf of which to delete the messages
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// A JSON-serialized list of 1-100 identifiers of messages to delete. All messages must be from the same chat. See deleteMessage for limitations on which messages can be deleted.
	MessageIds []int64 `json:"message_ids" structs:"message_ids,omitnested"`
}

// Changes the first and last name of a managed business account. Requires the can_change_name business bot right. Returns True on success.
type SetBusinessAccountNameRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// The new value of the first name for the business account; 1-64 characters
	FirstName string `json:"first_name" structs:"first_name,omitnested"`
	// The new value of the last name for the business account; 0-64 characters
	LastName *string `json:"last_name,omitempty" structs:"last_name,omitempty,omitnested"`
}

// Changes the username of a managed business account. Requires the can_change_username business bot right. Returns True on success.
type SetBusinessAccountUsernameRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// The new value of the username for the business account; 0-32 characters
	Username *string `json:"username,omitempty" structs:"username,omitempty,omitnested"`
}

// Changes the bio of a managed business account. Requires the can_change_bio business bot right. Returns True on success.
type SetBusinessAccountBioRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// The new value of the bio for the business account; 0-140 characters
	Bio *string `json:"bio,omitempty" structs:"bio,omitempty,omitnested"`
}

// Changes the profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
type SetBusinessAccountProfilePhotoRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// The new profile photo to set
	Photo InputProfilePhoto `json:"photo" structs:"photo,omitnested"`
	// Pass True to set the public photo, which will be visible even if the main photo is hidden by the business account's privacy settings. An account can have only one public photo.
	IsPublic *bool `json:"is_public,omitempty" structs:"is_public,omitempty,omitnested"`
}

// Removes the current profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
type RemoveBusinessAccountProfilePhotoRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Pass True to remove the public photo, which is visible even if the main photo is hidden by the business account's privacy settings. After the main photo is removed, the previous profile photo (if present) becomes the main photo.
	IsPublic *bool `json:"is_public,omitempty" structs:"is_public,omitempty,omitnested"`
}

// Changes the privacy settings pertaining to incoming gifts in a managed business account. Requires the can_change_gift_settings business bot right. Returns True on success.
type SetBusinessAccountGiftSettingsRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Pass True, if a button for sending a gift to the user or by the business account must always be shown in the input field
	ShowGiftButton bool `json:"show_gift_button" structs:"show_gift_button,omitnested"`
	// Types of gifts accepted by the business account
	AcceptedGiftTypes AcceptedGiftTypes `json:"accepted_gift_types" structs:"accepted_gift_types,omitnested"`
}

// Returns the amount of Telegram Stars owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns StarAmount on success.
type GetBusinessAccountStarBalanceRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
}

// Transfers Telegram Stars from the business account balance to the bot's balance. Requires the can_transfer_stars business bot right. Returns True on success.
type TransferBusinessAccountStarsRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Number of Telegram Stars to transfer; 1-10000
	StarCount int64 `json:"star_count" structs:"star_count,omitnested"`
}

// Returns the gifts received and owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns OwnedGifts on success.
type GetBusinessAccountGiftsRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Pass True to exclude gifts that aren't saved to the account's profile page
	ExcludeUnsaved *bool `json:"exclude_unsaved,omitempty" structs:"exclude_unsaved,omitempty,omitnested"`
	// Pass True to exclude gifts that are saved to the account's profile page
	ExcludeSaved *bool `json:"exclude_saved,omitempty" structs:"exclude_saved,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased an unlimited number of times
	ExcludeUnlimited *bool `json:"exclude_unlimited,omitempty" structs:"exclude_unlimited,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased a limited number of times and can be upgraded to unique
	ExcludeLimitedUpgradable *bool `json:"exclude_limited_upgradable,omitempty" structs:"exclude_limited_upgradable,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased a limited number of times and can't be upgraded to unique
	ExcludeLimitedNonUpgradable *bool `json:"exclude_limited_non_upgradable,omitempty" structs:"exclude_limited_non_upgradable,omitempty,omitnested"`
	// Pass True to exclude unique gifts
	ExcludeUnique *bool `json:"exclude_unique,omitempty" structs:"exclude_unique,omitempty,omitnested"`
	// Pass True to exclude gifts that were assigned from the TON blockchain and can't be resold or transferred in Telegram
	ExcludeFromBlockchain *bool `json:"exclude_from_blockchain,omitempty" structs:"exclude_from_blockchain,omitempty,omitnested"`
	// Pass True to sort results by gift price instead of send date. Sorting is applied before pagination.
	SortByPrice *bool `json:"sort_by_price,omitempty" structs:"sort_by_price,omitempty,omitnested"`
	// Offset of the first entry to return as received from the previous request; use empty string to get the first chunk of results
	Offset *string `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
	// The maximum number of gifts to be returned; 1-100. Defaults to 100.
	Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
}

// Returns the gifts owned and hosted by a user. Returns OwnedGifts on success.
type GetUserGiftsRequest struct {
	// Unique identifier of the user
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Pass True to exclude gifts that can be purchased an unlimited number of times
	ExcludeUnlimited *bool `json:"exclude_unlimited,omitempty" structs:"exclude_unlimited,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased a limited number of times and can be upgraded to unique
	ExcludeLimitedUpgradable *bool `json:"exclude_limited_upgradable,omitempty" structs:"exclude_limited_upgradable,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased a limited number of times and can't be upgraded to unique
	ExcludeLimitedNonUpgradable *bool `json:"exclude_limited_non_upgradable,omitempty" structs:"exclude_limited_non_upgradable,omitempty,omitnested"`
	// Pass True to exclude gifts that were assigned from the TON blockchain and can't be resold or transferred in Telegram
	ExcludeFromBlockchain *bool `json:"exclude_from_blockchain,omitempty" structs:"exclude_from_blockchain,omitempty,omitnested"`
	// Pass True to exclude unique gifts
	ExcludeUnique *bool `json:"exclude_unique,omitempty" structs:"exclude_unique,omitempty,omitnested"`
	// Pass True to sort results by gift price instead of send date. Sorting is applied before pagination.
	SortByPrice *bool `json:"sort_by_price,omitempty" structs:"sort_by_price,omitempty,omitnested"`
	// Offset of the first entry to return as received from the previous request; use an empty string to get the first chunk of results
	Offset *string `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
	// The maximum number of gifts to be returned; 1-100. Defaults to 100.
	Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
}

// Returns the gifts owned by a chat. Returns OwnedGifts on success.
type GetChatGiftsRequest struct {
	// Unique identifier for the target chat or username of the target channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Pass True to exclude gifts that aren't saved to the chat's profile page. Always True, unless the bot has the can_post_messages administrator right in the channel.
	ExcludeUnsaved *bool `json:"exclude_unsaved,omitempty" structs:"exclude_unsaved,omitempty,omitnested"`
	// Pass True to exclude gifts that are saved to the chat's profile page. Always False, unless the bot has the can_post_messages administrator right in the channel.
	ExcludeSaved *bool `json:"exclude_saved,omitempty" structs:"exclude_saved,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased an unlimited number of times
	ExcludeUnlimited *bool `json:"exclude_unlimited,omitempty" structs:"exclude_unlimited,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased a limited number of times and can be upgraded to unique
	ExcludeLimitedUpgradable *bool `json:"exclude_limited_upgradable,omitempty" structs:"exclude_limited_upgradable,omitempty,omitnested"`
	// Pass True to exclude gifts that can be purchased a limited number of times and can't be upgraded to unique
	ExcludeLimitedNonUpgradable *bool `json:"exclude_limited_non_upgradable,omitempty" structs:"exclude_limited_non_upgradable,omitempty,omitnested"`
	// Pass True to exclude gifts that were assigned from the TON blockchain and can't be resold or transferred in Telegram
	ExcludeFromBlockchain *bool `json:"exclude_from_blockchain,omitempty" structs:"exclude_from_blockchain,omitempty,omitnested"`
	// Pass True to exclude unique gifts
	ExcludeUnique *bool `json:"exclude_unique,omitempty" structs:"exclude_unique,omitempty,omitnested"`
	// Pass True to sort results by gift price instead of send date. Sorting is applied before pagination.
	SortByPrice *bool `json:"sort_by_price,omitempty" structs:"sort_by_price,omitempty,omitnested"`
	// Offset of the first entry to return as received from the previous request; use an empty string to get the first chunk of results
	Offset *string `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
	// The maximum number of gifts to be returned; 1-100. Defaults to 100.
	Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
}

// Converts a given regular gift to Telegram Stars. Requires the can_convert_gifts_to_stars business bot right. Returns True on success.
type ConvertGiftToStarsRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier of the regular gift that should be converted to Telegram Stars
	OwnedGiftId string `json:"owned_gift_id" structs:"owned_gift_id,omitnested"`
}

// Upgrades a given regular gift to a unique gift. Requires the can_transfer_and_upgrade_gifts business bot right. Additionally requires the can_transfer_stars business bot right if the upgrade is paid. Returns True on success.
type UpgradeGiftRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier of the regular gift that should be upgraded to a unique one
	OwnedGiftId string `json:"owned_gift_id" structs:"owned_gift_id,omitnested"`
	// Pass True to keep the original gift text, sender and receiver in the upgraded gift
	KeepOriginalDetails *bool `json:"keep_original_details,omitempty" structs:"keep_original_details,omitempty,omitnested"`
	// The amount of Telegram Stars that will be paid for the upgrade from the business account balance. If gift.prepaid_upgrade_star_count > 0, then pass 0, otherwise, the can_transfer_stars business bot right is required and gift.upgrade_star_count must be passed.
	StarCount *int64 `json:"star_count,omitempty" structs:"star_count,omitempty,omitnested"`
}

// Transfers an owned unique gift to another user. Requires the can_transfer_and_upgrade_gifts business bot right. Requires can_transfer_stars business bot right if the transfer is paid. Returns True on success.
type TransferGiftRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier of the regular gift that should be transferred
	OwnedGiftId string `json:"owned_gift_id" structs:"owned_gift_id,omitnested"`
	// Unique identifier of the chat which will own the gift. The chat must be active in the last 24 hours.
	NewOwnerChatId int64 `json:"new_owner_chat_id" structs:"new_owner_chat_id,omitnested"`
	// The amount of Telegram Stars that will be paid for the transfer from the business account balance. If positive, then the can_transfer_stars business bot right is required.
	StarCount *int64 `json:"star_count,omitempty" structs:"star_count,omitempty,omitnested"`
}

// Posts a story on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
type PostStoryRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Content of the story
	Content InputStoryContent `json:"content" structs:"content,omitnested"`
	// Period after which the story is moved to the archive, in seconds; must be one of 6 * 3600, 12 * 3600, 86400, or 2 * 86400
	ActivePeriod int64 `json:"active_period" structs:"active_period,omitnested"`
	// Caption of the story, 0-2048 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the story caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// A JSON-serialized list of clickable areas to be shown on the story
	Areas []StoryArea `json:"areas,omitempty" structs:"areas,omitempty,omitnested"`
	// Pass True to keep the story accessible after it expires
	PostToChatPage *bool `json:"post_to_chat_page,omitempty" structs:"post_to_chat_page,omitempty,omitnested"`
	// Pass True if the content of the story must be protected from forwarding and screenshotting
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
}

// Reposts a story on behalf of a business account from another business account. Both business accounts must be managed by the same bot, and the story on the source account must have been posted (or reposted) by the bot. Requires the can_manage_stories business bot right for both business accounts. Returns Story on success.
type RepostStoryRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier of the chat which posted the story that should be reposted
	FromChatId int64 `json:"from_chat_id" structs:"from_chat_id,omitnested"`
	// Unique identifier of the story that should be reposted
	FromStoryId int64 `json:"from_story_id" structs:"from_story_id,omitnested"`
	// Period after which the story is moved to the archive, in seconds; must be one of 6 * 3600, 12 * 3600, 86400, or 2 * 86400
	ActivePeriod int64 `json:"active_period" structs:"active_period,omitnested"`
	// Pass True to keep the story accessible after it expires
	PostToChatPage *bool `json:"post_to_chat_page,omitempty" structs:"post_to_chat_page,omitempty,omitnested"`
	// Pass True if the content of the story must be protected from forwarding and screenshotting
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
}

// Edits a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
type EditStoryRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier of the story to edit
	StoryId int64 `json:"story_id" structs:"story_id,omitnested"`
	// Content of the story
	Content InputStoryContent `json:"content" structs:"content,omitnested"`
	// Caption of the story, 0-2048 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the story caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// A JSON-serialized list of clickable areas to be shown on the story
	Areas []StoryArea `json:"areas,omitempty" structs:"areas,omitempty,omitnested"`
}

// Deletes a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns True on success.
type DeleteStoryRequest struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier of the story to delete
	StoryId int64 `json:"story_id" structs:"story_id,omitnested"`
}

// Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
type AnswerWebAppQueryRequest struct {
	// Unique identifier for the query to be answered
	WebAppQueryId string `json:"web_app_query_id" structs:"web_app_query_id,omitnested"`
	// A JSON-serialized object describing the message to be sent
	Result InlineQueryResult `json:"result" structs:"result,omitnested"`
}

// Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.
type SavePreparedInlineMessageRequest struct {
	// Unique identifier of the target user that can use the prepared message
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// A JSON-serialized object describing the message to be sent
	Result InlineQueryResult `json:"result" structs:"result,omitnested"`
	// Pass True if the message can be sent to private chats with users
	AllowUserChats *bool `json:"allow_user_chats,omitempty" structs:"allow_user_chats,omitempty,omitnested"`
	// Pass True if the message can be sent to private chats with bots
	AllowBotChats *bool `json:"allow_bot_chats,omitempty" structs:"allow_bot_chats,omitempty,omitnested"`
	// Pass True if the message can be sent to group and supergroup chats
	AllowGroupChats *bool `json:"allow_group_chats,omitempty" structs:"allow_group_chats,omitempty,omitnested"`
	// Pass True if the message can be sent to channel chats
	AllowChannelChats *bool `json:"allow_channel_chats,omitempty" structs:"allow_channel_chats,omitempty,omitnested"`
}

// Stores a keyboard button that can be used by a user within a Mini App. Returns a PreparedKeyboardButton object.
type SavePreparedKeyboardButtonRequest struct {
	// Unique identifier of the target user that can use the button
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// A JSON-serialized object describing the button to be saved. The button must be of the type request_users, request_chat, or request_managed_bot.
	Button KeyboardButton `json:"button" structs:"button,omitnested"`
}

// Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageTextRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username.
	ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text" structs:"text,omitnested"`
	// Mode for parsing entities in the message text. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities []MessageEntity `json:"entities,omitempty" structs:"entities,omitempty,omitnested"`
	// Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty" structs:"link_preview_options,omitempty,omitnested"`
	// A JSON-serialized object for an inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageCaptionRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username.
	ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
	// New caption of the message, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty" structs:"caption,omitempty,omitnested"`
	// Mode for parsing entities in the message caption. See formatting options for more details.
	ParseMode *string `json:"parse_mode,omitempty" structs:"parse_mode,omitempty,omitnested"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty" structs:"caption_entities,omitempty,omitnested"`
	// Pass True, if the caption must be shown above the message media. Supported only for animation, photo and video messages.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty" structs:"show_caption_above_media,omitempty,omitnested"`
	// A JSON-serialized object for an inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit animation, audio, document, live photo, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo, a live photo, or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageMediaRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username.
	ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
	// A JSON-serialized object for a new media content of the message
	Media InputMedia `json:"media" structs:"media,omitnested"`
	// A JSON-serialized object for a new inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocationRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username.
	ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
	// Latitude of new location
	Latitude float64 `json:"latitude" structs:"latitude,omitnested"`
	// Longitude of new location
	Longitude float64 `json:"longitude" structs:"longitude,omitnested"`
	// New period in seconds during which the location can be updated, starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever. Otherwise, the new value must not exceed the current live_period by more than a day, and the live location expiration date must remain within the next 90 days. If not specified, then live_period remains unchanged.
	LivePeriod *int64 `json:"live_period,omitempty" structs:"live_period,omitempty,omitnested"`
	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty" structs:"horizontal_accuracy,omitempty,omitnested"`
	// Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int64 `json:"heading,omitempty" structs:"heading,omitempty,omitnested"`
	// The maximum distance for proximity alerts about approaching another chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int64 `json:"proximity_alert_radius,omitempty" structs:"proximity_alert_radius,omitempty,omitnested"`
	// A JSON-serialized object for a new inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
type StopMessageLiveLocationRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username.
	ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the message with live location to stop.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
	// A JSON-serialized object for a new inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit a checklist on behalf of a connected business account. On success, the edited Message is returned.
type EditMessageChecklistRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId string `json:"business_connection_id" structs:"business_connection_id,omitnested"`
	// Unique identifier for the target chat or username of the target bot in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// A JSON-serialized object for the new checklist
	Checklist InputChecklist `json:"checklist" structs:"checklist,omitnested"`
	// A JSON-serialized object for the new inline keyboard for the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageReplyMarkupRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username.
	ChatId ChatId `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the message to edit.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
	// A JSON-serialized object for an inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
type StopPollRequest struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of the original message with the poll
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// A JSON-serialized object for a new message inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to approve a suggested post in a direct messages chat. The bot must have the 'can_post_messages' administrator right in the corresponding channel chat. Returns True on success.
type ApproveSuggestedPostRequest struct {
	// Unique identifier for the target direct messages chat
	ChatId int64 `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of a suggested post message to approve
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// Point in time (Unix timestamp) when the post is expected to be published; omit if the date has already been specified when the suggested post was created. If specified, then the date must be not more than 2678400 seconds (30 days) in the future.
	SendDate *int64 `json:"send_date,omitempty" structs:"send_date,omitempty,omitnested"`
}

// Use this method to decline a suggested post in a direct messages chat. The bot must have the 'can_manage_direct_messages' administrator right in the corresponding channel chat. Returns True on success.
type DeclineSuggestedPostRequest struct {
	// Unique identifier for the target direct messages chat
	ChatId int64 `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of a suggested post message to decline
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// Comment for the creator of the suggested post; 0-128 characters
	Comment *string `json:"comment,omitempty" structs:"comment,omitempty,omitnested"`
}

// Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages administrator right in a supergroup or a channel, it can delete any message there.
// - If the bot has can_manage_direct_messages administrator right in a channel, it can delete any message in the corresponding direct messages chat.
// Returns True on success.
type DeleteMessageRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of the message to delete
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
}

// Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
type DeleteMessagesRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// A JSON-serialized list of 1-100 identifiers of messages to delete. See deleteMessage for limitations on which messages can be deleted.
	MessageIds []int64 `json:"message_ids" structs:"message_ids,omitnested"`
}

// Use this method to remove a reaction from a message in a group or a supergroup chat. The bot must have the 'can_delete_messages' administrator right in the chat. Returns True on success.
type DeleteMessageReactionRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of the target message
	MessageId int64 `json:"message_id" structs:"message_id,omitnested"`
	// Identifier of the user whose reaction will be removed, if the reaction was added by a user
	UserId *int64 `json:"user_id,omitempty" structs:"user_id,omitempty,omitnested"`
	// Identifier of the chat whose reaction will be removed, if the reaction was added by a chat
	ActorChatId *int64 `json:"actor_chat_id,omitempty" structs:"actor_chat_id,omitempty,omitnested"`
}

// Use this method to remove up to 10000 recent reactions in a group or a supergroup chat added by a given user or chat. The bot must have the 'can_delete_messages' administrator right in the chat. Returns True on success.
type DeleteAllMessageReactionsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Identifier of the user whose reactions will be removed, if the reactions were added by a user
	UserId *int64 `json:"user_id,omitempty" structs:"user_id,omitempty,omitnested"`
	// Identifier of the chat whose reactions will be removed, if the reactions were added by a chat
	ActorChatId *int64 `json:"actor_chat_id,omitempty" structs:"actor_chat_id,omitempty,omitnested"`
}

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
type SendStickerRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Video and animated stickers can't be sent via an HTTP URL.
	Sticker InputFile `json:"sticker" structs:"sticker,omitnested"`
	// Emoji associated with the sticker; only for just uploaded stickers
	Emoji *string `json:"emoji,omitempty" structs:"emoji,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
type GetStickerSetRequest struct {
	// Name of the sticker set
	Name string `json:"name" structs:"name,omitnested"`
}

// Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
type GetCustomEmojiStickersRequest struct {
	// A JSON-serialized list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
	CustomEmojiIds []string `json:"custom_emoji_ids" structs:"custom_emoji_ids,omitnested"`
}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
type UploadStickerFileRequest struct {
	// User identifier of sticker file owner
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See https://core.telegram.org/stickers for technical requirements. More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Sticker InputFile `json:"sticker" structs:"sticker,omitnested"`
	// Format of the sticker, must be one of "static", "animated", "video"
	StickerFormat string `json:"sticker_format" structs:"sticker_format,omitnested"`
}

// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
type CreateNewStickerSetRequest struct {
	// User identifier of created sticker set owner
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name" structs:"name,omitnested"`
	// Sticker set title, 1-64 characters
	Title string `json:"title" structs:"title,omitnested"`
	// A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
	Stickers []InputSticker `json:"stickers" structs:"stickers,omitnested"`
	// Type of stickers in the set, pass "regular", "mask", or "custom_emoji". By default, a regular sticker set is created.
	StickerType *string `json:"sticker_type,omitempty" structs:"sticker_type,omitempty,omitnested"`
	// Pass True if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
	NeedsRepainting *bool `json:"needs_repainting,omitempty" structs:"needs_repainting,omitempty,omitnested"`
}

// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
type AddStickerToSetRequest struct {
	// User identifier of sticker set owner
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Sticker set name
	Name string `json:"name" structs:"name,omitnested"`
	// A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
	Sticker InputSticker `json:"sticker" structs:"sticker,omitnested"`
}

// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
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

// Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
type ReplaceStickerInSetRequest struct {
	// User identifier of the sticker set owner
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Sticker set name
	Name string `json:"name" structs:"name,omitnested"`
	// File identifier of the replaced sticker
	OldSticker string `json:"old_sticker" structs:"old_sticker,omitnested"`
	// A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
	Sticker InputSticker `json:"sticker" structs:"sticker,omitnested"`
}

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
type SetStickerEmojiListRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker" structs:"sticker,omitnested"`
	// A JSON-serialized list of 1-20 emoji associated with the sticker
	EmojiList []string `json:"emoji_list" structs:"emoji_list,omitnested"`
}

// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
type SetStickerKeywordsRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker" structs:"sticker,omitnested"`
	// A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters
	Keywords []string `json:"keywords,omitempty" structs:"keywords,omitempty,omitnested"`
}

// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
type SetStickerMaskPositionRequest struct {
	// File identifier of the sticker
	Sticker string `json:"sticker" structs:"sticker,omitnested"`
	// A JSON-serialized object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
	MaskPosition *MaskPosition `json:"mask_position,omitempty" structs:"mask_position,omitempty,omitnested"`
}

// Use this method to set the title of a created sticker set. Returns True on success.
type SetStickerSetTitleRequest struct {
	// Sticker set name
	Name string `json:"name" structs:"name,omitnested"`
	// Sticker set title, 1-64 characters
	Title string `json:"title" structs:"title,omitnested"`
}

// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
type SetStickerSetThumbnailRequest struct {
	// Sticker set name
	Name string `json:"name" structs:"name,omitnested"`
	// User identifier of the sticker set owner
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animation-requirements for animated sticker technical requirements), or a .WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files: https://core.telegram.org/bots/api#sending-files. Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
	Thumbnail InputFile `json:"thumbnail,omitempty" structs:"thumbnail,omitempty,omitnested"`
	// Format of the thumbnail, must be one of "static" for a .WEBP or .PNG image, "animated" for a .TGS animation, or "video" for a .WEBM video
	Format string `json:"format" structs:"format,omitnested"`
}

// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
type SetCustomEmojiStickerSetThumbnailRequest struct {
	// Sticker set name
	Name string `json:"name" structs:"name,omitnested"`
	// Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail
	CustomEmojiId *string `json:"custom_emoji_id,omitempty" structs:"custom_emoji_id,omitempty,omitnested"`
}

// Use this method to delete a sticker set that was created by the bot. Returns True on success.
type DeleteStickerSetRequest struct {
	// Sticker set name
	Name string `json:"name" structs:"name,omitnested"`
}

// Use this method to send answers to an inline query. On success, True is returned.
// No more than 50 results per query are allowed.
type AnswerInlineQueryRequest struct {
	// Unique identifier for the answered query
	InlineQueryId string `json:"inline_query_id" structs:"inline_query_id,omitnested"`
	// A JSON-serialized array of results for the inline query
	Results []InlineQueryResult `json:"results" structs:"results,omitnested"`
	// The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	CacheTime *int64 `json:"cache_time,omitempty" structs:"cache_time,omitempty,omitnested"`
	// Pass True if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query.
	IsPersonal *bool `json:"is_personal,omitempty" structs:"is_personal,omitempty,omitnested"`
	// Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don't support pagination. Offset length can't exceed 64 bytes.
	NextOffset *string `json:"next_offset,omitempty" structs:"next_offset,omitempty,omitnested"`
	// A JSON-serialized object describing a button to be shown above inline query results
	Button *InlineQueryResultsButton `json:"button,omitempty" structs:"button,omitempty,omitnested"`
}

// Use this method to send invoices. On success, the sent Message is returned.
type SendInvoiceRequest struct {
	// Unique identifier for the target chat or username of the target bot, supergroup or channel in the format @username
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Identifier of the direct messages topic to which the message will be sent; required if the message is sent to a direct messages chat
	DirectMessagesTopicId *int64 `json:"direct_messages_topic_id,omitempty" structs:"direct_messages_topic_id,omitempty,omitnested"`
	// Product name, 1-32 characters
	Title string `json:"title" structs:"title,omitnested"`
	// Product description, 1-255 characters
	Description string `json:"description" structs:"description,omitnested"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload" structs:"payload,omitnested"`
	// Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken *string `json:"provider_token,omitempty" structs:"provider_token,omitempty,omitnested"`
	// Three-letter ISO 4217 currency code, see more on currencies. Pass "XTR" for payments in Telegram Stars.
	Currency string `json:"currency" structs:"currency,omitnested"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	Prices []LabeledPrice `json:"prices" structs:"prices,omitnested"`
	// The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	MaxTipAmount *int64 `json:"max_tip_amount,omitempty" structs:"max_tip_amount,omitempty,omitnested"`
	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty" structs:"suggested_tip_amounts,omitempty,omitnested"`
	// Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter.
	StartParameter *string `json:"start_parameter,omitempty" structs:"start_parameter,omitempty,omitnested"`
	// JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty" structs:"provider_data,omitempty,omitnested"`
	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoUrl *string `json:"photo_url,omitempty" structs:"photo_url,omitempty,omitnested"`
	// Photo size in bytes
	PhotoSize *int64 `json:"photo_size,omitempty" structs:"photo_size,omitempty,omitnested"`
	// Photo width
	PhotoWidth *int64 `json:"photo_width,omitempty" structs:"photo_width,omitempty,omitnested"`
	// Photo height
	PhotoHeight *int64 `json:"photo_height,omitempty" structs:"photo_height,omitempty,omitnested"`
	// Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedName *bool `json:"need_name,omitempty" structs:"need_name,omitempty,omitnested"`
	// Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber *bool `json:"need_phone_number,omitempty" structs:"need_phone_number,omitempty,omitnested"`
	// Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail *bool `json:"need_email,omitempty" structs:"need_email,omitempty,omitnested"`
	// Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress *bool `json:"need_shipping_address,omitempty" structs:"need_shipping_address,omitempty,omitnested"`
	// Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty" structs:"send_phone_number_to_provider,omitempty,omitnested"`
	// Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider *bool `json:"send_email_to_provider,omitempty" structs:"send_email_to_provider,omitempty,omitnested"`
	// Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
	IsFlexible *bool `json:"is_flexible,omitempty" structs:"is_flexible,omitempty,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// A JSON-serialized object containing the parameters of the suggested post to send; for direct messages chats only. If the message is sent as a reply to another suggested post, then that suggested post is automatically declined.
	SuggestedPostParameters *SuggestedPostParameters `json:"suggested_post_parameters,omitempty" structs:"suggested_post_parameters,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to create a link for an invoice. Returns the created invoice link as String on success.
type CreateInvoiceLinkRequest struct {
	// Unique identifier of the business connection on behalf of which the link will be created. For payments in Telegram Stars only.
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Product name, 1-32 characters
	Title string `json:"title" structs:"title,omitnested"`
	// Product description, 1-255 characters
	Description string `json:"description" structs:"description,omitnested"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload string `json:"payload" structs:"payload,omitnested"`
	// Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	ProviderToken *string `json:"provider_token,omitempty" structs:"provider_token,omitempty,omitnested"`
	// Three-letter ISO 4217 currency code, see more on currencies. Pass "XTR" for payments in Telegram Stars.
	Currency string `json:"currency" structs:"currency,omitnested"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	Prices []LabeledPrice `json:"prices" structs:"prices,omitnested"`
	// The number of seconds the subscription will be active for before the next payment. The currency must be set to "XTR" (Telegram Stars) if the parameter is used. Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user. Subscription price must no exceed 10000 Telegram Stars.
	SubscriptionPeriod *int64 `json:"subscription_period,omitempty" structs:"subscription_period,omitempty,omitnested"`
	// The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	MaxTipAmount *int64 `json:"max_tip_amount,omitempty" structs:"max_tip_amount,omitempty,omitnested"`
	// A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty" structs:"suggested_tip_amounts,omitempty,omitnested"`
	// JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	ProviderData *string `json:"provider_data,omitempty" structs:"provider_data,omitempty,omitnested"`
	// URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
	PhotoUrl *string `json:"photo_url,omitempty" structs:"photo_url,omitempty,omitnested"`
	// Photo size in bytes
	PhotoSize *int64 `json:"photo_size,omitempty" structs:"photo_size,omitempty,omitnested"`
	// Photo width
	PhotoWidth *int64 `json:"photo_width,omitempty" structs:"photo_width,omitempty,omitnested"`
	// Photo height
	PhotoHeight *int64 `json:"photo_height,omitempty" structs:"photo_height,omitempty,omitnested"`
	// Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedName *bool `json:"need_name,omitempty" structs:"need_name,omitempty,omitnested"`
	// Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber *bool `json:"need_phone_number,omitempty" structs:"need_phone_number,omitempty,omitnested"`
	// Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail *bool `json:"need_email,omitempty" structs:"need_email,omitempty,omitnested"`
	// Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress *bool `json:"need_shipping_address,omitempty" structs:"need_shipping_address,omitempty,omitnested"`
	// Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendPhoneNumberToProvider *bool `json:"send_phone_number_to_provider,omitempty" structs:"send_phone_number_to_provider,omitempty,omitnested"`
	// Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider *bool `json:"send_email_to_provider,omitempty" structs:"send_email_to_provider,omitempty,omitnested"`
	// Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
	IsFlexible *bool `json:"is_flexible,omitempty" structs:"is_flexible,omitempty,omitnested"`
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
type AnswerShippingQueryRequest struct {
	// Unique identifier for the query to be answered
	ShippingQueryId string `json:"shipping_query_id" structs:"shipping_query_id,omitnested"`
	// Pass True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	Ok bool `json:"ok" structs:"ok,omitnested"`
	// Required if ok is True. A JSON-serialized array of available shipping options.
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty" structs:"shipping_options,omitempty,omitnested"`
	// Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable"). Telegram will display this message to the user.
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

// A method to get the current Telegram Stars balance of the bot. Requires no parameters. On success, returns a StarAmount object.
type GetMyStarBalanceRequest struct{}

// Returns the bot's Telegram Star transactions in chronological order. On success, returns a StarTransactions object.
type GetStarTransactionsRequest struct {
	// Number of transactions to skip in the response
	Offset *int64 `json:"offset,omitempty" structs:"offset,omitempty,omitnested"`
	// The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int64 `json:"limit,omitempty" structs:"limit,omitempty,omitnested"`
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
type RefundStarPaymentRequest struct {
	// Identifier of the user whose payment will be refunded
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id" structs:"telegram_payment_charge_id,omitnested"`
}

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
type EditUserStarSubscriptionRequest struct {
	// Identifier of the user whose subscription will be edited
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Telegram payment identifier for the subscription
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id" structs:"telegram_payment_charge_id,omitnested"`
	// Pass True to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.
	IsCanceled bool `json:"is_canceled" structs:"is_canceled,omitnested"`
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
type SetPassportDataErrorsRequest struct {
	// User identifier
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// A JSON-serialized array describing the errors
	Errors []PassportElementError `json:"errors" structs:"errors,omitnested"`
}

// Use this method to send a game. On success, the sent Message is returned.
type SendGameRequest struct {
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty" structs:"business_connection_id,omitempty,omitnested"`
	// Unique identifier for the target chat or username of the target bot in the format @username. Games can't be sent to channel direct messages chats and channel chats.
	ChatId ChatId `json:"chat_id" structs:"chat_id,omitnested"`
	// Unique identifier for the target message thread (topic) of a forum; for forum supergroups and private chats of bots with forum topic mode enabled only
	MessageThreadId *int64 `json:"message_thread_id,omitempty" structs:"message_thread_id,omitempty,omitnested"`
	// Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
	GameShortName string `json:"game_short_name" structs:"game_short_name,omitnested"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty" structs:"disable_notification,omitempty,omitnested"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty" structs:"protect_content,omitempty,omitnested"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance.
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty" structs:"allow_paid_broadcast,omitempty,omitnested"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty" structs:"message_effect_id,omitempty,omitnested"`
	// Description of the message to reply to
	ReplyParameters *ReplyParameters `json:"reply_parameters,omitempty" structs:"reply_parameters,omitempty,omitnested"`
	// A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty" structs:"reply_markup,omitempty,omitnested"`
}

// Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScoreRequest struct {
	// User identifier
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// New score, must be non-negative
	Score int64 `json:"score" structs:"score,omitnested"`
	// Pass True if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters.
	Force *bool `json:"force,omitempty" structs:"force,omitempty,omitnested"`
	// Pass True if the game message should not be automatically edited to include the current scoreboard
	DisableEditMessage *bool `json:"disable_edit_message,omitempty" structs:"disable_edit_message,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat.
	ChatId *int64 `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the sent message.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
type GetGameHighScoresRequest struct {
	// Target user id
	UserId int64 `json:"user_id" structs:"user_id,omitnested"`
	// Required if inline_message_id is not specified. Unique identifier for the target chat.
	ChatId *int64 `json:"chat_id,omitempty" structs:"chat_id,omitempty,omitnested"`
	// Required if inline_message_id is not specified. Identifier of the sent message.
	MessageId *int64 `json:"message_id,omitempty" structs:"message_id,omitempty,omitnested"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message.
	InlineMessageId *string `json:"inline_message_id,omitempty" structs:"inline_message_id,omitempty,omitnested"`
}
