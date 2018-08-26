package client

import (
    "io"
    "strings"
    "os"
)

type UpdateType string

const (
    UpdateTypeMessage            UpdateType = "message"
    UpdateTypeEditedMessage      UpdateType = "edited_message"
    UpdateTypeChannelPost        UpdateType = "channel_post"
    UpdateTypeEditedChannelPost  UpdateType = "edited_channel_post"
    UpdateTypeInlineQuery        UpdateType = "inline_query"
    UpdateTypeChosenInlineResult UpdateType = "chosen_inline_result"
    UpdateTypeCallbackQuery      UpdateType = "callback_query"
    UpdateTypeShippingQuery      UpdateType = "shipping_query"
    UpdateTypePreCheckoutQuery   UpdateType = "pre_checkout_query"
)

func (updateType UpdateType) String() string {
    return string(updateType)
}

// This object represents an incoming update.At most one of the optional parameters can be present in any given update.
type Update struct {
    // The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
    UpdateId int64 `json:"update_id"`
    // New incoming message of any kind — text, photo, sticker, etc.
    Message *Message `json:"message,omitempty"`
    // New version of a message that is known to the bot and was edited
    EditedMessage *Message `json:"edited_message,omitempty"`
    // New incoming channel post of any kind — text, photo, sticker, etc.
    ChannelPost *Message `json:"channel_post,omitempty"`
    // New version of a channel post that is known to the bot and was edited
    EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
    // New incoming inline query
    InlineQuery *InlineQuery `json:"inline_query,omitempty"`
    // The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
    ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
    // New incoming callback query
    CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
    // New incoming shipping query. Only for invoices with flexible price
    ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
    // New incoming pre-checkout query. Contains full information about checkout
    PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
}

// Contains information about the current status of a webhook.
type WebhookInfo struct {
    // Webhook URL, may be empty if webhook is not set up
    Url string `json:"url"`
    // True, if a custom certificate was provided for webhook certificate checks
    HasCustomCertificate bool `json:"has_custom_certificate"`
    // Number of updates awaiting delivery
    PendingUpdateCount int64 `json:"pending_update_count"`
    // Unix time for the most recent error that happened when trying to deliver an update via webhook
    LastErrorDate *int64 `json:"last_error_date,omitempty"`
    // Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
    LastErrorMessage *string `json:"last_error_message,omitempty"`
    // Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
    MaxConnections *int64 `json:"max_connections,omitempty"`
    // A list of update types the bot is subscribed to. Defaults to all update types
    AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

// This object represents a Telegram user or bot.
type User struct {
    // Unique identifier for this user or bot
    Id int64 `json:"id"`
    // True, if this user is a bot
    IsBot bool `json:"is_bot"`
    // User‘s or bot’s first name
    FirstName string `json:"first_name"`
    // User‘s or bot’s last name
    LastName *string `json:"last_name,omitempty"`
    // User‘s or bot’s username
    Username *string `json:"username,omitempty"`
    // IETF language tag of the user's language
    LanguageCode *string `json:"language_code,omitempty"`
}

// This object represents a chat.
type Chat struct {
    // Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
    Id int64 `json:"id"`
    // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
    Type string `json:"type"`
    // Title, for supergroups, channels and group chats
    Title *string `json:"title,omitempty"`
    // Username, for private chats, supergroups and channels if available
    Username *string `json:"username,omitempty"`
    // First name of the other party in a private chat
    FirstName *string `json:"first_name,omitempty"`
    // Last name of the other party in a private chat
    LastName *string `json:"last_name,omitempty"`
    // True if a group has ‘All Members Are Admins’ enabled.
    AllMembersAreAdministrators *bool `json:"all_members_are_administrators,omitempty"`
    // Chat photo. Returned only in getChat.
    Photo *ChatPhoto `json:"photo,omitempty"`
    // Description, for supergroups and channel chats. Returned only in getChat.
    Description *string `json:"description,omitempty"`
    // Chat invite link, for supergroups and channel chats. Returned only in getChat.
    InviteLink *string `json:"invite_link,omitempty"`
    // Pinned message, for supergroups and channel chats. Returned only in getChat.
    PinnedMessage *Message `json:"pinned_message,omitempty"`
    // For supergroups, name of group sticker set. Returned only in getChat.
    StickerSetName *string `json:"sticker_set_name,omitempty"`
    // True, if the bot can change the group sticker set. Returned only in getChat.
    CanSetStickerSet *bool `json:"can_set_sticker_set,omitempty"`
}

// This object represents a message.
type Message struct {
    // Unique message identifier inside this chat
    MessageId int64 `json:"message_id"`
    // Sender, empty for messages sent to channels
    From *User `json:"from,omitempty"`
    // Date the message was sent in Unix time
    Date int64 `json:"date"`
    // Conversation the message belongs to
    Chat Chat `json:"chat"`
    // For forwarded messages, sender of the original message
    ForwardFrom *User `json:"forward_from,omitempty"`
    // For messages forwarded from channels, information about the original channel
    ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
    // For messages forwarded from channels, identifier of the original message in the channel
    ForwardFromMessageId *int64 `json:"forward_from_message_id,omitempty"`
    // For messages forwarded from channels, signature of the post author if present
    ForwardSignature *string `json:"forward_signature,omitempty"`
    // For forwarded messages, date the original message was sent in Unix time
    ForwardDate *int64 `json:"forward_date,omitempty"`
    // For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
    ReplyToMessage *Message `json:"reply_to_message,omitempty"`
    // Date the message was last edited in Unix time
    EditDate *int64 `json:"edit_date,omitempty"`
    // The unique identifier of a media message group this message belongs to
    MediaGroupId *string `json:"media_group_id,omitempty"`
    // Signature of the post author for messages in channels
    AuthorSignature *string `json:"author_signature,omitempty"`
    // For text messages, the actual UTF-8 text of the message, 0-4096 characters.
    Text *string `json:"text,omitempty"`
    // For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
    Entities *[]MessageEntity `json:"entities,omitempty"`
    // For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
    CaptionEntities *[]MessageEntity `json:"caption_entities,omitempty"`
    // Message is an audio file, information about the file
    Audio *Audio `json:"audio,omitempty"`
    // Message is a general file, information about the file
    Document *Document `json:"document,omitempty"`
    // Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
    Animation *Animation `json:"animation,omitempty"`
    // Message is a game, information about the game. More about games »
    Game *Game `json:"game,omitempty"`
    // Message is a photo, available sizes of the photo
    Photo *[]PhotoSize `json:"photo,omitempty"`
    // Message is a sticker, information about the sticker
    Sticker *Sticker `json:"sticker,omitempty"`
    // Message is a video, information about the video
    Video *Video `json:"video,omitempty"`
    // Message is a voice message, information about the file
    Voice *Voice `json:"voice,omitempty"`
    // Message is a video note, information about the video message
    VideoNote *VideoNote `json:"video_note,omitempty"`
    // Caption for the audio, document, photo, video or voice, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Message is a shared contact, information about the contact
    Contact *Contact `json:"contact,omitempty"`
    // Message is a shared location, information about the location
    Location *Location `json:"location,omitempty"`
    // Message is a venue, information about the venue
    Venue *Venue `json:"venue,omitempty"`
    // New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
    NewChatMembers *[]User `json:"new_chat_members,omitempty"`
    // A member was removed from the group, information about them (this member may be the bot itself)
    LeftChatMember *User `json:"left_chat_member,omitempty"`
    // A chat title was changed to this value
    NewChatTitle *string `json:"new_chat_title,omitempty"`
    // A chat photo was change to this value
    NewChatPhoto *[]PhotoSize `json:"new_chat_photo,omitempty"`
    // Service message: the chat photo was deleted
    DeleteChatPhoto *bool `json:"delete_chat_photo,omitempty"`
    // Service message: the group has been created
    GroupChatCreated *bool `json:"group_chat_created,omitempty"`
    // Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
    SupergroupChatCreated *bool `json:"supergroup_chat_created,omitempty"`
    // Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
    ChannelChatCreated *bool `json:"channel_chat_created,omitempty"`
    // The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
    MigrateToChatId *int64 `json:"migrate_to_chat_id,omitempty"`
    // The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
    MigrateFromChatId *int64 `json:"migrate_from_chat_id,omitempty"`
    // Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
    PinnedMessage *Message `json:"pinned_message,omitempty"`
    // Message is an invoice for a payment, information about the invoice. More about payments »
    Invoice *Invoice `json:"invoice,omitempty"`
    // Message is a service message about a successful payment, information about the payment. More about payments »
    SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
    // The domain name of the website on which the user has logged in. More about Telegram Login »
    ConnectedWebsite *string `json:"connected_website,omitempty"`
    // Telegram Passport data
    PassportData *PassportData `json:"passport_data,omitempty"`
}

type MessageEntityType string

func (messageEntityType MessageEntityType) String() string {
    return string(messageEntityType)
}

const (
    MessageEntityMention     MessageEntityType = "mention"
    MessageEntityHashtag     MessageEntityType = "hashtag"
    MessageEntityCashtag     MessageEntityType = "cashtag"
    MessageEntityBotCommand  MessageEntityType = "bot_command"
    MessageEntityUrl         MessageEntityType = "url"
    MessageEntityEmail       MessageEntityType = "email"
    MessageEntityPhoneNumber MessageEntityType = "phone_number"
    MessageEntityBold        MessageEntityType = "bold"
    MessageEntityItalic      MessageEntityType = "italic"
    MessageEntityCode        MessageEntityType = "code"
    MessageEntityPre         MessageEntityType = "pre"
    MessageEntityTextLink    MessageEntityType = "text_link"
    MessageEntityTextMention MessageEntityType = "text_mention"
)

// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
    // Type of the entity. Can be mention (@username), hashtag, cashtag, bot_command, url, email, phone_number, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
    Type MessageEntityType `json:"type"`
    // Offset in UTF-16 code units to the start of the entity
    Offset int64 `json:"offset"`
    // Length of the entity in UTF-16 code units
    Length int64 `json:"length"`
    // For “text_link” only, url that will be opened after user taps on the text
    Url *string `json:"url,omitempty"`
    // For “text_mention” only, the mentioned user
    User *User `json:"user,omitempty"`
}

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // Photo width
    Width int64 `json:"width"`
    // Photo height
    Height int64 `json:"height"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // Duration of the audio in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Performer of the audio as defined by sender or by audio tags
    Performer *string `json:"performer,omitempty"`
    // Title of the audio as defined by sender or by audio tags
    Title *string `json:"title,omitempty"`
    // MIME type of the file as defined by sender
    MimeType *string `json:"mime_type,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
    // Thumbnail of the album cover to which the music file belongs
    Thumb *PhotoSize `json:"thumb,omitempty"`
}

// This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
    // Unique file identifier
    FileId string `json:"file_id"`
    // Document thumbnail as defined by sender
    Thumb *PhotoSize `json:"thumb,omitempty"`
    // Original filename as defined by sender
    FileName *string `json:"file_name,omitempty"`
    // MIME type of the file as defined by sender
    MimeType *string `json:"mime_type,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents a video file.
type Video struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // Video width as defined by sender
    Width int64 `json:"width"`
    // Video height as defined by sender
    Height int64 `json:"height"`
    // Duration of the video in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Video thumbnail
    Thumb *PhotoSize `json:"thumb,omitempty"`
    // Mime type of a file as defined by sender
    MimeType *string `json:"mime_type,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
    // Unique file identifier
    FileId string `json:"file_id"`
    // Video width as defined by sender
    Width int64 `json:"width"`
    // Video height as defined by sender
    Height int64 `json:"height"`
    // Duration of the video in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Animation thumbnail as defined by sender
    Thumb *PhotoSize `json:"thumb,omitempty"`
    // Original animation filename as defined by sender
    FileName *string `json:"file_name,omitempty"`
    // MIME type of the file as defined by sender
    MimeType *string `json:"mime_type,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents a voice note.
type Voice struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // Duration of the audio in seconds as defined by sender
    Duration int64 `json:"duration"`
    // MIME type of the file as defined by sender
    MimeType *string `json:"mime_type,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // Video width and height as defined by sender
    Length int64 `json:"length"`
    // Duration of the video in seconds as defined by sender
    Duration int64 `json:"duration"`
    // Video thumbnail
    Thumb *PhotoSize `json:"thumb,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents a phone contact.
type Contact struct {
    // Contact's phone number
    PhoneNumber string `json:"phone_number"`
    // Contact's first name
    FirstName string `json:"first_name"`
    // Contact's last name
    LastName *string `json:"last_name,omitempty"`
    // Contact's user identifier in Telegram
    UserId *int64 `json:"user_id,omitempty"`
    // Additional data about the contact in the form of a vCard
    Vcard *string `json:"vcard,omitempty"`
}

// This object represents a point on the map.
type Location struct {
    // Longitude as defined by sender
    Longitude float64 `json:"longitude"`
    // Latitude as defined by sender
    Latitude float64 `json:"latitude"`
}

// This object represents a venue.
type Venue struct {
    // Venue location
    Location Location `json:"location"`
    // Name of the venue
    Title string `json:"title"`
    // Address of the venue
    Address string `json:"address"`
    // Foursquare identifier of the venue
    FoursquareId *string `json:"foursquare_id,omitempty"`
    // Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
    FoursquareType *string `json:"foursquare_type,omitempty"`
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
    // Total number of profile pictures the target user has
    TotalCount int64 `json:"total_count"`
    // Requested profile pictures (in up to 4 sizes each)
    Photos [][]PhotoSize `json:"photos"`
}

// This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
type File struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // File size, if known
    FileSize *int64 `json:"file_size,omitempty"`
    // File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
    FilePath *string `json:"file_path,omitempty"`
}

type ReplyMarkup interface{}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
    // Array of button rows, each represented by an Array of KeyboardButton objects
    Keyboard [][]KeyboardButton `json:"keyboard"`
    // Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
    ResizeKeyboard *bool `json:"resize_keyboard,omitempty"`
    // Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
    OneTimeKeyboard *bool `json:"one_time_keyboard,omitempty"`
    // Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
    Selective *bool `json:"selective,omitempty"`
}

// This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields are mutually exclusive.
type KeyboardButton struct {
    // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
    Text string `json:"text"`
    // If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
    RequestContact *bool `json:"request_contact,omitempty"`
    // If True, the user's current location will be sent when the button is pressed. Available in private chats only
    RequestLocation *bool `json:"request_location,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
    // Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
    RemoveKeyboard bool `json:"remove_keyboard"`
    // Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
    Selective *bool `json:"selective,omitempty"`
}

// This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
    // Array of button rows, each represented by an Array of InlineKeyboardButton objects
    InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
    // Label text on the button
    Text string `json:"text"`
    // HTTP or tg:// url to be opened when button is pressed
    Url *string `json:"url,omitempty"`
    // Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
    CallbackData *string `json:"callback_data,omitempty"`
    // If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
    SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`
    // If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
    SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`
    // Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row.
    CallbackGame CallbackGame `json:"callback_game,omitempty"`
    // Specify True, to send a Pay button.NOTE: This type of button must always be the first button in the first row.
    Pay *bool `json:"pay,omitempty"`
}

// This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
    // Unique identifier for this query
    Id string `json:"id"`
    // Sender
    From User `json:"from"`
    // Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
    Message *Message `json:"message,omitempty"`
    // Identifier of the message sent via the bot in inline mode, that originated the query.
    InlineMessageId *string `json:"inline_message_id,omitempty"`
    // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
    ChatInstance string `json:"chat_instance"`
    // Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
    Data *string `json:"data,omitempty"`
    // Short name of a Game to be returned, serves as the unique identifier for the game
    GameShortName *string `json:"game_short_name,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
    // Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
    ForceReply bool `json:"force_reply"`
    // Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
    Selective *bool `json:"selective,omitempty"`
}

// This object represents a chat photo.
type ChatPhoto struct {
    // Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
    SmallFileId string `json:"small_file_id"`
    // Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
    BigFileId string `json:"big_file_id"`
}

// This object contains information about one member of a chat.
type ChatMember struct {
    // Information about the user
    User User `json:"user"`
    // The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
    Status string `json:"status"`
    // Restricted and kicked only. Date when restrictions will be lifted for this user, unix time
    UntilDate *int64 `json:"until_date,omitempty"`
    // Administrators only. True, if the bot is allowed to edit administrator privileges of that user
    CanBeEdited *bool `json:"can_be_edited,omitempty"`
    // Administrators only. True, if the administrator can change the chat title, photo and other settings
    CanChangeInfo *bool `json:"can_change_info,omitempty"`
    // Administrators only. True, if the administrator can post in the channel, channels only
    CanPostMessages *bool `json:"can_post_messages,omitempty"`
    // Administrators only. True, if the administrator can edit messages of other users and can pin messages, channels only
    CanEditMessages *bool `json:"can_edit_messages,omitempty"`
    // Administrators only. True, if the administrator can delete messages of other users
    CanDeleteMessages *bool `json:"can_delete_messages,omitempty"`
    // Administrators only. True, if the administrator can invite new users to the chat
    CanInviteUsers *bool `json:"can_invite_users,omitempty"`
    // Administrators only. True, if the administrator can restrict, ban or unban chat members
    CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`
    // Administrators only. True, if the administrator can pin messages, supergroups only
    CanPinMessages *bool `json:"can_pin_messages,omitempty"`
    // Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
    CanPromoteMembers *bool `json:"can_promote_members,omitempty"`
    // Restricted only. True, if the user can send text messages, contacts, locations and venues
    CanSendMessages *bool `json:"can_send_messages,omitempty"`
    // Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
    CanSendMediaMessages *bool `json:"can_send_media_messages,omitempty"`
    // Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
    CanSendOtherMessages *bool `json:"can_send_other_messages,omitempty"`
    // Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
    CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
}

// Contains information about why a request was unsuccessful.
type ResponseParameters struct {
    // The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
    MigrateToChatId *int64 `json:"migrate_to_chat_id,omitempty"`
    // In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
    RetryAfter *int64 `json:"retry_after,omitempty"`
}

// This object represents the content of a media message to be sent. It should be one of InputMediaAnimation, InputMediaDocument, InputMediaAudio, InputMediaPhoto, InputMediaVideo
type InputMedia interface{}

// InputMediaPhoto and InputMediaVideo
type InputMediaGroup interface{}

// Represents a photo to be sent.
type InputMediaPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
    Media string `json:"media"`
    // Caption of the photo to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
}

// Represents a video to be sent.
type InputMediaVideo struct {
    // Type of the result, must be video
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using multipart/form-data under <file_attach_name> name
    Media string `json:"media"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty"`
    // Caption of the video to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Video width
    Width *int64 `json:"width,omitempty"`
    // Video height
    Height *int64 `json:"height,omitempty"`
    // Video duration
    Duration *int64 `json:"duration,omitempty"`
    // Pass True, if the uploaded video is suitable for streaming
    SupportsStreaming *bool `json:"supports_streaming,omitempty"`
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
    // Type of the result, must be animation
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name name. More info on Sending Files »
    Media string `json:"media"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty"`
    // Caption of the animation to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Animation width
    Width *int64 `json:"width,omitempty"`
    // Animation height
    Height *int64 `json:"height,omitempty"`
    // Animation duration
    Duration *int64 `json:"duration,omitempty"`
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
    // Type of the result, must be audio
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
    Media string `json:"media"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty"`
    // Caption of the audio to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Duration of the audio in seconds
    Duration *int64 `json:"duration,omitempty"`
    // Performer of the audio
    Performer *int64 `json:"performer,omitempty"`
    // Title of the audio
    Title *int64 `json:"title,omitempty"`
}

// Represents a general file to be sent.
type InputMediaDocument struct {
    // Type of the result, must be document
    Type string `json:"type"`
    // File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
    Media string `json:"media"`
    // Thumbnail of the file sent. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 90. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
    Thumb InputFile `json:"thumb,omitempty"`
    // Caption of the document to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
}

// This object represents the contents of a file to be uploaded. Must be posted using multipart/form-data in the usual way that files are uploaded via the browser.
type InputFile interface {
    Close()
    GetReader() io.Reader
    Name() string
    IsStream() bool
}

// If the file is already stored somewhere on the Telegram servers, you don't need to reupload it: each file object has a file_id field, simply pass this file_id as a parameter instead of uploading. There are no limits for files sent this way.
type FileIdInputFile struct {
    FileId string
}

func NewFileIdInputFile(fileId string) *FileIdInputFile {
    return &FileIdInputFile{
        FileId: fileId,
    }
}

func (FileIdInputFile) Close() {}

func (inputFile *FileIdInputFile) GetReader() io.Reader {
    return strings.NewReader(inputFile.FileId)
}

func (inputFile *FileIdInputFile) Name() string {
    return inputFile.FileId
}

func (inputFile *FileIdInputFile) IsStream() bool {
    return false
}

// Provide Telegram with an HTTP URL for the file to be sent. Telegram will download and send the file. 5 MB max size for photos and 20 MB max for other types of content.
type UrlInputFile struct {
    Url string
}

func NewUrlInputFile(url string) *UrlInputFile {
    return &UrlInputFile{
        Url: url,
    }
}

func (UrlInputFile) Close() {}

func (inputFile *UrlInputFile) GetReader() io.Reader {
    return strings.NewReader(inputFile.Url)
}

func (inputFile *UrlInputFile) Name() string {
    return inputFile.Url
}

func (inputFile *UrlInputFile) IsStream() bool {
    return false
}

// Post the file using multipart/form-data in the usual way that files are uploaded via the browser. 10 MB max size for photos, 50 MB for other files.
type FileInputFile struct {
    Reader   io.Reader
    FileName string
}

func NewFileInputFile(path string) (*FileInputFile, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    stat, err := file.Stat()
    if err != nil {
        return nil, err
    }

    return &FileInputFile{
        Reader:   file,
        FileName: stat.Name(),
    }, nil
}

func (inputFile *FileInputFile) Close() {
    rc, ok := inputFile.Reader.(io.ReadCloser)
    if ok {
        rc.Close()
    }
}

func (inputFile *FileInputFile) GetReader() io.Reader {
    return inputFile.Reader
}

func (inputFile *FileInputFile) Name() string {
    return inputFile.FileName
}

func (inputFile *FileInputFile) IsStream() bool {
    return true
}

// This object represents a sticker.
type Sticker struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // Sticker width
    Width int64 `json:"width"`
    // Sticker height
    Height int64 `json:"height"`
    // Sticker thumbnail in the .webp or .jpg format
    Thumb *PhotoSize `json:"thumb,omitempty"`
    // Emoji associated with the sticker
    Emoji *string `json:"emoji,omitempty"`
    // Name of the sticker set to which the sticker belongs
    SetName *string `json:"set_name,omitempty"`
    // For mask stickers, the position where the mask should be placed
    MaskPosition *MaskPosition `json:"mask_position,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

// This object represents a sticker set.
type StickerSet struct {
    // Sticker set name
    Name string `json:"name"`
    // Sticker set title
    Title string `json:"title"`
    // True, if the sticker set contains masks
    ContainsMasks bool `json:"contains_masks"`
    // List of all set stickers
    Stickers []Sticker `json:"stickers"`
}

// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
    // The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
    Point string `json:"point"`
    // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
    XShift float64 `json:"x_shift"`
    // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
    YShift float64 `json:"y_shift"`
    // Mask scaling coefficient. For example, 2.0 means double size.
    Scale float64 `json:"scale"`
}

// This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
    // Unique identifier for this query
    Id string `json:"id"`
    // Sender
    From User `json:"from"`
    // Sender location, only for bots that request user location
    Location *Location `json:"location,omitempty"`
    // Text of the query (up to 512 characters)
    Query string `json:"query"`
    // Offset of the results to be returned, can be controlled by the bot
    Offset string `json:"offset"`
}

// This object represents one result of an inline query. Telegram clients currently support results of the following 20 types
type InlineQueryResult interface{}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
    // Type of the result, must be article
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Title of the result
    Title string `json:"title"`
    // Content of the message to be sent
    InputMessageContent InputMessageContent `json:"input_message_content"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // URL of the result
    Url *string `json:"url,omitempty"`
    // Pass True, if you don't want the URL to be shown in the message
    HideUrl *bool `json:"hide_url,omitempty"`
    // Short description of the result
    Description *string `json:"description,omitempty"`
    // Url of the thumbnail for the result
    ThumbUrl *string `json:"thumb_url,omitempty"`
    // Thumbnail width
    ThumbWidth *int64 `json:"thumb_width,omitempty"`
    // Thumbnail height
    ThumbHeight *int64 `json:"thumb_height,omitempty"`
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
    PhotoUrl string `json:"photo_url"`
    // URL of the thumbnail for the photo
    ThumbUrl string `json:"thumb_url"`
    // Width of the photo
    PhotoWidth *int64 `json:"photo_width,omitempty"`
    // Height of the photo
    PhotoHeight *int64 `json:"photo_height,omitempty"`
    // Title for the result
    Title *string `json:"title,omitempty"`
    // Short description of the result
    Description *string `json:"description,omitempty"`
    // Caption of the photo to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the photo
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
    // Type of the result, must be gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the GIF file. File size must not exceed 1MB
    GifUrl string `json:"gif_url"`
    // Width of the GIF
    GifWidth *int64 `json:"gif_width,omitempty"`
    // Height of the GIF
    GifHeight *int64 `json:"gif_height,omitempty"`
    // Duration of the GIF
    GifDuration *int64 `json:"gif_duration,omitempty"`
    // URL of the static thumbnail for the result (jpeg or gif)
    ThumbUrl string `json:"thumb_url"`
    // Title for the result
    Title *string `json:"title,omitempty"`
    // Caption of the GIF file to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the GIF animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
    // Type of the result, must be mpeg4_gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the MP4 file. File size must not exceed 1MB
    Mpeg4Url string `json:"mpeg4_url"`
    // Video width
    Mpeg4Width *int64 `json:"mpeg4_width,omitempty"`
    // Video height
    Mpeg4Height *int64 `json:"mpeg4_height,omitempty"`
    // Video duration
    Mpeg4Duration *int64 `json:"mpeg4_duration,omitempty"`
    // URL of the static thumbnail (jpeg or gif) for the result
    ThumbUrl string `json:"thumb_url"`
    // Title for the result
    Title *string `json:"title,omitempty"`
    // Caption of the MPEG-4 file to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the video animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
    // Type of the result, must be video
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the embedded video player or video file
    VideoUrl string `json:"video_url"`
    // Mime type of the content of video url, “text/html” or “video/mp4”
    MimeType string `json:"mime_type"`
    // URL of the thumbnail (jpeg only) for the video
    ThumbUrl string `json:"thumb_url"`
    // Title for the result
    Title string `json:"title"`
    // Caption of the video to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Video width
    VideoWidth *int64 `json:"video_width,omitempty"`
    // Video height
    VideoHeight *int64 `json:"video_height,omitempty"`
    // Video duration in seconds
    VideoDuration *int64 `json:"video_duration,omitempty"`
    // Short description of the result
    Description *string `json:"description,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an mp3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
    // Type of the result, must be audio
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the audio file
    AudioUrl string `json:"audio_url"`
    // Title
    Title string `json:"title"`
    // Caption, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Performer
    Performer *string `json:"performer,omitempty"`
    // Audio duration in seconds
    AudioDuration *int64 `json:"audio_duration,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the audio
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
    // Type of the result, must be voice
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid URL for the voice recording
    VoiceUrl string `json:"voice_url"`
    // Recording title
    Title string `json:"title"`
    // Caption, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Recording duration in seconds
    VoiceDuration *int64 `json:"voice_duration,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the voice recording
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
    // Type of the result, must be document
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // Title for the result
    Title string `json:"title"`
    // Caption of the document to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // A valid URL for the file
    DocumentUrl string `json:"document_url"`
    // Mime type of the content of the file, either “application/pdf” or “application/zip”
    MimeType string `json:"mime_type"`
    // Short description of the result
    Description *string `json:"description,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the file
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // URL of the thumbnail (jpeg only) for the file
    ThumbUrl *string `json:"thumb_url,omitempty"`
    // Thumbnail width
    ThumbWidth *int64 `json:"thumb_width,omitempty"`
    // Thumbnail height
    ThumbHeight *int64 `json:"thumb_height,omitempty"`
}

// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
    // Type of the result, must be location
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Location latitude in degrees
    Latitude float64 `json:"latitude"`
    // Location longitude in degrees
    Longitude float64 `json:"longitude"`
    // Location title
    Title string `json:"title"`
    // Period in seconds for which the location can be updated, should be between 60 and 86400.
    LivePeriod *int64 `json:"live_period,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the location
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Url of the thumbnail for the result
    ThumbUrl *string `json:"thumb_url,omitempty"`
    // Thumbnail width
    ThumbWidth *int64 `json:"thumb_width,omitempty"`
    // Thumbnail height
    ThumbHeight *int64 `json:"thumb_height,omitempty"`
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
    // Type of the result, must be venue
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Latitude of the venue location in degrees
    Latitude float64 `json:"latitude"`
    // Longitude of the venue location in degrees
    Longitude float64 `json:"longitude"`
    // Title of the venue
    Title string `json:"title"`
    // Address of the venue
    Address string `json:"address"`
    // Foursquare identifier of the venue if known
    FoursquareId *string `json:"foursquare_id,omitempty"`
    // Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
    FoursquareType *string `json:"foursquare_type,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the venue
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Url of the thumbnail for the result
    ThumbUrl *string `json:"thumb_url,omitempty"`
    // Thumbnail width
    ThumbWidth *int64 `json:"thumb_width,omitempty"`
    // Thumbnail height
    ThumbHeight *int64 `json:"thumb_height,omitempty"`
}

// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
    // Type of the result, must be contact
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 Bytes
    Id string `json:"id"`
    // Contact's phone number
    PhoneNumber string `json:"phone_number"`
    // Contact's first name
    FirstName string `json:"first_name"`
    // Contact's last name
    LastName *string `json:"last_name,omitempty"`
    // Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard *string `json:"vcard,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the contact
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
    // Url of the thumbnail for the result
    ThumbUrl *string `json:"thumb_url,omitempty"`
    // Thumbnail width
    ThumbWidth *int64 `json:"thumb_width,omitempty"`
    // Thumbnail height
    ThumbHeight *int64 `json:"thumb_height,omitempty"`
}

// Represents a Game.
type InlineQueryResultGame struct {
    // Type of the result, must be game
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // Short name of the game
    GameShortName string `json:"game_short_name"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
    // Type of the result, must be photo
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier of the photo
    PhotoFileId string `json:"photo_file_id"`
    // Title for the result
    Title *string `json:"title,omitempty"`
    // Short description of the result
    Description *string `json:"description,omitempty"`
    // Caption of the photo to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the photo
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
    // Type of the result, must be gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the GIF file
    GifFileId string `json:"gif_file_id"`
    // Title for the result
    Title *string `json:"title,omitempty"`
    // Caption of the GIF file to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the GIF animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
    // Type of the result, must be mpeg4_gif
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the MP4 file
    Mpeg4FileId string `json:"mpeg4_file_id"`
    // Title for the result
    Title *string `json:"title,omitempty"`
    // Caption of the MPEG-4 file to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the video animation
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
    // Type of the result, must be sticker
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier of the sticker
    StickerFileId string `json:"sticker_file_id"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the sticker
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
    // Type of the result, must be document
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // Title for the result
    Title string `json:"title"`
    // A valid file identifier for the file
    DocumentFileId string `json:"document_file_id"`
    // Short description of the result
    Description *string `json:"description,omitempty"`
    // Caption of the document to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the file
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
    // Type of the result, must be video
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the video file
    VideoFileId string `json:"video_file_id"`
    // Title for the result
    Title string `json:"title"`
    // Short description of the result
    Description *string `json:"description,omitempty"`
    // Caption of the video to be sent, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the video
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
    // Type of the result, must be voice
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the voice message
    VoiceFileId string `json:"voice_file_id"`
    // Voice message title
    Title string `json:"title"`
    // Caption, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the voice message
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// Represents a link to an mp3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
    // Type of the result, must be audio
    Type string `json:"type"`
    // Unique identifier for this result, 1-64 bytes
    Id string `json:"id"`
    // A valid file identifier for the audio file
    AudioFileId string `json:"audio_file_id"`
    // Caption, 0-200 characters
    Caption *string `json:"caption,omitempty"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Inline keyboard attached to the message
    ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
    // Content of the message to be sent instead of the audio
    InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

// This object represents the content of a message to be sent as a result of an inline query. Telegram clients currently support the following 4 types
type InputMessageContent interface{}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
    // Text of the message to be sent, 1-4096 characters
    MessageText string `json:"message_text"`
    // Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
    ParseMode *string `json:"parse_mode,omitempty"`
    // Disables link previews for links in the sent message
    DisableWebPagePreview *bool `json:"disable_web_page_preview,omitempty"`
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
    // Latitude of the location in degrees
    Latitude float64 `json:"latitude"`
    // Longitude of the location in degrees
    Longitude float64 `json:"longitude"`
    // Period in seconds for which the location can be updated, should be between 60 and 86400.
    LivePeriod *int64 `json:"live_period,omitempty"`
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
    // Latitude of the venue in degrees
    Latitude float64 `json:"latitude"`
    // Longitude of the venue in degrees
    Longitude float64 `json:"longitude"`
    // Name of the venue
    Title string `json:"title"`
    // Address of the venue
    Address string `json:"address"`
    // Foursquare identifier of the venue, if known
    FoursquareId *string `json:"foursquare_id,omitempty"`
    // Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
    FoursquareType *string `json:"foursquare_type,omitempty"`
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
    // Contact's phone number
    PhoneNumber string `json:"phone_number"`
    // Contact's first name
    FirstName string `json:"first_name"`
    // Contact's last name
    LastName *string `json:"last_name,omitempty"`
    // Additional data about the contact in the form of a vCard, 0-2048 bytes
    Vcard *string `json:"vcard,omitempty"`
}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
    // The unique identifier for the result that was chosen
    ResultId string `json:"result_id"`
    // The user that chose the result
    From User `json:"from"`
    // Sender location, only for bots that require user location
    Location *Location `json:"location,omitempty"`
    // Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
    InlineMessageId *string `json:"inline_message_id,omitempty"`
    // The query that was used to obtain the result
    Query string `json:"query"`
}

// This object represents a portion of the price for goods or services.
type LabeledPrice struct {
    // Portion label
    Label string `json:"label"`
    // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    Amount int64 `json:"amount"`
}

// This object contains basic information about an invoice.
type Invoice struct {
    // Product name
    Title string `json:"title"`
    // Product description
    Description string `json:"description"`
    // Unique bot deep-linking parameter that can be used to generate this invoice
    StartParameter string `json:"start_parameter"`
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int64 `json:"total_amount"`
}

// This object represents a shipping address.
type ShippingAddress struct {
    // ISO 3166-1 alpha-2 country code
    CountryCode string `json:"country_code"`
    // State, if applicable
    State string `json:"state"`
    // City
    City string `json:"city"`
    // First line for the address
    StreetLine1 string `json:"street_line1"`
    // Second line for the address
    StreetLine2 string `json:"street_line2"`
    // Address post code
    PostCode string `json:"post_code"`
}

// This object represents information about an order.
type OrderInfo struct {
    // User name
    Name *string `json:"name,omitempty"`
    // User's phone number
    PhoneNumber *string `json:"phone_number,omitempty"`
    // User email
    Email *string `json:"email,omitempty"`
    // User shipping address
    ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// This object represents one shipping option.
type ShippingOption struct {
    // Shipping option identifier
    Id string `json:"id"`
    // Option title
    Title string `json:"title"`
    // List of price portions
    Prices []LabeledPrice `json:"prices"`
}

// This object contains basic information about a successful payment.
type SuccessfulPayment struct {
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int64 `json:"total_amount"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload"`
    // Identifier of the shipping option chosen by the user
    ShippingOptionId *string `json:"shipping_option_id,omitempty"`
    // Order info provided by the user
    OrderInfo *OrderInfo `json:"order_info,omitempty"`
    // Telegram payment identifier
    TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
    // Provider payment identifier
    ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}

// This object contains information about an incoming shipping query.
type ShippingQuery struct {
    // Unique query identifier
    Id string `json:"id"`
    // User who sent the query
    From User `json:"from"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload"`
    // User specified shipping address
    ShippingAddress ShippingAddress `json:"shipping_address"`
}

// This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
    // Unique query identifier
    Id string `json:"id"`
    // User who sent the query
    From User `json:"from"`
    // Three-letter ISO 4217 currency code
    Currency string `json:"currency"`
    // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
    TotalAmount int64 `json:"total_amount"`
    // Bot specified invoice payload
    InvoicePayload string `json:"invoice_payload"`
    // Identifier of the shipping option chosen by the user
    ShippingOptionId *string `json:"shipping_option_id,omitempty"`
    // Order info provided by the user
    OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

// Contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
    // Array with information about documents and other Telegram Passport elements that was shared with the bot
    Data []EncryptedPassportElement `json:"data"`
    // Encrypted credentials required to decrypt the data
    Credentials EncryptedCredentials `json:"credentials"`
}

// This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
    // Unique identifier for this file
    FileId string `json:"file_id"`
    // File size
    FileSize int64 `json:"file_size"`
    // Unix time when the file was uploaded
    FileDate int64 `json:"file_date"`
}

// Contains information about documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
    // Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
    Type string `json:"type"`
    // Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “identity_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
    Data *string `json:"data,omitempty"`
    // User's verified phone number, available only for “phone_number” type
    PhoneNumber *string `json:"phone_number,omitempty"`
    // User's verified email address, available only for “email” type
    Email *string `json:"email,omitempty"`
    // Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
    Files *[]PassportFile `json:"files,omitempty"`
    // Encrypted file with the front side of the document, provided by the user. Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
    FrontSide *PassportFile `json:"front_side,omitempty"`
    // Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
    ReverseSide *PassportFile `json:"reverse_side,omitempty"`
    // Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
    Selfie *PassportFile `json:"selfie,omitempty"`
}

// Contains data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
    // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
    Data string `json:"data"`
    // Base64-encoded data hash for data authentication
    Hash string `json:"hash"`
    // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
    Secret string `json:"secret"`
}

// This object represents an error in the Telegram Passport element which was submitted that should be resolved by the user. It should be one of: PassportElementErrorDataField, PassportElementErrorFrontSide, PassportElementErrorReverseSide, PassportElementErrorSelfie, PassportElementErrorFile, PassportElementErrorFiles
type PassportElementError interface{}

// Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
    // Error source, must be data
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
    Type string `json:"type"`
    // Name of the data field which has the error
    FieldName string `json:"field_name"`
    // Base64-encoded data hash
    DataHash string `json:"data_hash"`
    // Error message
    Message string `json:"message"`
}

// Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
    // Error source, must be front_side
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
    Type string `json:"type"`
    // Base64-encoded hash of the file with the front side of the document
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

// Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
    // Error source, must be reverse_side
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
    Type string `json:"type"`
    // Base64-encoded hash of the file with the reverse side of the document
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

// Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
    // Error source, must be selfie
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
    Type string `json:"type"`
    // Base64-encoded hash of the file with the selfie
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

// Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
    // Error source, must be file
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
    Type string `json:"type"`
    // Base64-encoded file hash
    FileHash string `json:"file_hash"`
    // Error message
    Message string `json:"message"`
}

// Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
    // Error source, must be files
    Source string `json:"source"`
    // The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
    Type string `json:"type"`
    // List of base64-encoded file hashes
    FileHashes []string `json:"file_hashes"`
    // Error message
    Message string `json:"message"`
}

// This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
    // Title of the game
    Title string `json:"title"`
    // Description of the game
    Description string `json:"description"`
    // Photo that will be displayed in the game message in chats.
    Photo []PhotoSize `json:"photo"`
    // Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
    Text *string `json:"text,omitempty"`
    // Special entities that appear in text, such as usernames, URLs, bot commands, etc.
    TextEntities *[]MessageEntity `json:"text_entities,omitempty"`
    // Animation that will be displayed in the game message in chats. Upload via BotFather
    Animation *AnimationGames `json:"animation,omitempty"`
}

// You can provide an animation for your game so that it looks stylish in chats (check out Lumberjack for an example). This object represents an animation file to be displayed in the message containing a game.
type AnimationGames struct {
    // Unique file identifier
    FileId string `json:"file_id"`
    // Animation thumbnail as defined by sender
    Thumb *PhotoSize `json:"thumb,omitempty"`
    // Original animation filename as defined by sender
    FileName *string `json:"file_name,omitempty"`
    // MIME type of the file as defined by sender
    MimeType *string `json:"mime_type,omitempty"`
    // File size
    FileSize *int64 `json:"file_size,omitempty"`
}

type CallbackGame interface{}

// This object represents one row of the high scores table for a game.
type GameHighScore struct {
    // Position in high score table for the game
    Position int64 `json:"position"`
    // User
    User User `json:"user"`
    // Score
    Score int64 `json:"score"`
}

func OptionalBool(b bool) *bool {
    return &b
}

func OptionalFloat(f float64) *float64 {
    return &f
}

func OptionalInt(i int64) *int64 {
    return &i
}

func OptionalString(s string) *string {
    return &s
}
