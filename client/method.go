package client

import "encoding/json"

// Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.
func (client *Client) GetUpdates(req *GetUpdatesRequest) ([]*Update, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getUpdates", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp []*Update

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
func (client *Client) SetWebhook(req *SetWebhookRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setWebhook", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.
func (client *Client) DeleteWebhook() (bool, error) {
    apiResp, err := client.Request("deleteWebhook", nil)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (client *Client) GetWebhookInfo() (*WebhookInfo, error) {
    apiResp, err := client.Request("getWebhookInfo", nil)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *WebhookInfo

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in form of a User object.
func (client *Client) GetMe() (*User, error) {
    apiResp, err := client.Request("getMe", nil)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *User

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send text messages. On success, the sent Message is returned.
func (client *Client) SendMessage(req *SendMessageRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendMessage", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to forward messages of any kind. On success, the sent Message is returned.
func (client *Client) ForwardMessage(req *ForwardMessageRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("forwardMessage", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send photos. On success, the sent Message is returned.
func (client *Client) SendPhoto(req *SendPhotoRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendPhoto", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (client *Client) SendAudio(req *SendAudioRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendAudio", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (client *Client) SendDocument(req *SendDocumentRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendDocument", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
func (client *Client) SendVideo(req *SendVideoRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendVideo", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
func (client *Client) SendAnimation(req *SendAnimationRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendAnimation", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (client *Client) SendVoice(req *SendVoiceRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendVoice", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
func (client *Client) SendVideoNote(req *SendVideoNoteRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendVideoNote", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
func (client *Client) SendMediaGroup(req *SendMediaGroupRequest) ([]*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendMediaGroup", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp []*Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send point on the map. On success, the sent Message is returned.
func (client *Client) SendLocation(req *SendLocationRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendLocation", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to edit live location messages sent by the bot or via the bot (for inline bots). A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
func (client *Client) EditMessageLiveLocation(req *EditMessageLiveLocationRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("editMessageLiveLocation", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to stop updating a live location message sent by the bot or via the bot (for inline bots) before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
func (client *Client) StopMessageLiveLocation(req *StopMessageLiveLocationRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("stopMessageLiveLocation", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send information about a venue. On success, the sent Message is returned.
func (client *Client) SendVenue(req *SendVenueRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendVenue", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to send phone contacts. On success, the sent Message is returned.
func (client *Client) SendContact(req *SendContactRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendContact", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
func (client *Client) SendChatAction(req *SendChatActionRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendChatAction", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (client *Client) GetUserProfilePhotos(req *GetUserProfilePhotosRequest) (*UserProfilePhotos, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getUserProfilePhotos", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *UserProfilePhotos

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
func (client *Client) GetFile(req *GetFileRequest) (*File, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getFile", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *File

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (client *Client) KickChatMember(req *KickChatMemberRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("kickChatMember", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.
func (client *Client) UnbanChatMember(req *UnbanChatMemberRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("unbanChatMember", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all boolean parameters to lift restrictions from a user. Returns True on success.
func (client *Client) RestrictChatMember(req *RestrictChatMemberRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("restrictChatMember", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success.
func (client *Client) PromoteChatMember(req *PromoteChatMemberRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("promoteChatMember", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.
func (client *Client) ExportChatInviteLink(req *ExportChatInviteLinkRequest) (string, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("exportChatInviteLink", params)
    if err != nil {
        return "", err
    }

    if !apiResp.Ok {
        return "", newError(apiResp)
    }

    var resp string

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return "", err
    }

    return resp, nil
}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (client *Client) SetChatPhoto(req *SetChatPhotoRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setChatPhoto", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (client *Client) DeleteChatPhoto(req *DeleteChatPhotoRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("deleteChatPhoto", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (client *Client) SetChatTitle(req *SetChatTitleRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setChatTitle", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to change the description of a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (client *Client) SetChatDescription(req *SetChatDescriptionRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setChatDescription", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to pin a message in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
func (client *Client) PinChatMessage(req *PinChatMessageRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("pinChatMessage", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to unpin a message in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
func (client *Client) UnpinChatMessage(req *UnpinChatMessageRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("unpinChatMessage", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (client *Client) LeaveChat(req *LeaveChatRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("leaveChat", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
func (client *Client) GetChat(req *GetChatRequest) (*Chat, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getChat", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Chat

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
func (client *Client) GetChatAdministrators(req *GetChatAdministratorsRequest) ([]*ChatMember, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getChatAdministrators", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp []*ChatMember

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to get the number of members in a chat. Returns Int on success.
func (client *Client) GetChatMembersCount(req *GetChatMembersCountRequest) (int64, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getChatMembersCount", params)
    if err != nil {
        return 0, err
    }

    if !apiResp.Ok {
        return 0, newError(apiResp)
    }

    var resp int64

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return 0, err
    }

    return resp, nil
}

// Use this method to get information about a member of a chat. Returns a ChatMember object on success.
func (client *Client) GetChatMember(req *GetChatMemberRequest) (*ChatMember, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getChatMember", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *ChatMember

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (client *Client) SetChatStickerSet(req *SetChatStickerSetRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setChatStickerSet", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
func (client *Client) DeleteChatStickerSet(req *DeleteChatStickerSetRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("deleteChatStickerSet", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
func (client *Client) AnswerCallbackQuery(req *AnswerCallbackQueryRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("answerCallbackQuery", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to edit text and game messages sent by the bot or via the bot (for inline bots). On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
func (client *Client) EditMessageText(req *EditMessageTextRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("editMessageText", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to edit captions of messages sent by the bot or via the bot (for inline bots). On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
func (client *Client) EditMessageCaption(req *EditMessageCaptionRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("editMessageCaption", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to edit audio, document, photo, or video messages. If a message is a part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
func (client *Client) EditMessageMedia(req *EditMessageMediaRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("editMessageMedia", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to edit only the reply markup of messages sent by the bot or via the bot (for inline bots).  On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
func (client *Client) EditMessageReplyMarkup(req *EditMessageReplyMarkupRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("editMessageReplyMarkup", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Bots can delete outgoing messages in groups and supergroups.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
func (client *Client) DeleteMessage(req *DeleteMessageRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("deleteMessage", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to send .webp stickers. On success, the sent Message is returned.
func (client *Client) SendSticker(req *SendStickerRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendSticker", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
func (client *Client) GetStickerSet(req *GetStickerSetRequest) (*StickerSet, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getStickerSet", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *StickerSet

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
func (client *Client) UploadStickerFile(req *UploadStickerFileRequest) (*File, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("uploadStickerFile", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *File

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to create new sticker set owned by a user. The bot will be able to edit the created sticker set. Returns True on success.
func (client *Client) CreateNewStickerSet(req *CreateNewStickerSetRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("createNewStickerSet", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to add a new sticker to a set created by the bot. Returns True on success.
func (client *Client) AddStickerToSet(req *AddStickerToSetRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("addStickerToSet", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.
func (client *Client) SetStickerPositionInSet(req *SetStickerPositionInSetRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setStickerPositionInSet", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (client *Client) DeleteStickerFromSet(req *DeleteStickerFromSetRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("deleteStickerFromSet", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
func (client *Client) AnswerInlineQuery(req *AnswerInlineQueryRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("answerInlineQuery", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to send invoices. On success, the sent Message is returned.
func (client *Client) SendInvoice(req *SendInvoiceRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendInvoice", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
func (client *Client) AnswerShippingQuery(req *AnswerShippingQueryRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("answerShippingQuery", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (client *Client) AnswerPreCheckoutQuery(req *AnswerPreCheckoutQueryRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("answerPreCheckoutQuery", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
//Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
func (client *Client) SetPassportDataErrors(req *SetPassportDataErrorsRequest) (bool, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setPassportDataErrors", params)
    if err != nil {
        return false, err
    }

    if !apiResp.Ok {
        return false, newError(apiResp)
    }

    var resp bool

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return false, err
    }

    return resp, nil
}

// Use this method to send a game. On success, the sent Message is returned.
func (client *Client) SendGame(req *SendGameRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("sendGame", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (client *Client) SetGameScore(req *SetGameScoreRequest) (*Message, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("setGameScore", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp *Message

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
func (client *Client) GetGameHighScores(req *GetGameHighScoresRequest) ([]*GameHighScore, error) {
    params := requestToMap(req)

    apiResp, err := client.Request("getGameHighScores", params)
    if err != nil {
        return nil, err
    }

    if !apiResp.Ok {
        return nil, newError(apiResp)
    }

    var resp []*GameHighScore

    err = json.Unmarshal(apiResp.Result, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}
