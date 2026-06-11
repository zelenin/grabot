package client

import "encoding/json"

// Use this method to receive incoming updates using long polling (wiki). Returns an Array of Update objects.
func (client *Client) GetUpdates(req *GetUpdatesRequest) ([]Update, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getUpdates", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []Update

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to specify a URL and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified URL, containing a JSON-serialized Update. In case of an unsuccessful request (a request with response HTTP status code different from 2XY), we will repeat the request and give up after a reasonable amount of attempts. Returns True on success.
// If you'd like to make sure that the webhook was set by you, you can specify secret data in the parameter secret_token. If specified, the request will contain a header "X-Telegram-Bot-Api-Secret-Token" with the secret token as content.
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

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
func (client *Client) DeleteWebhook(req *DeleteWebhookRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteWebhook", params)
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

// A simple method for testing your bot's authentication token. Requires no parameters. Returns basic information about the bot in form of a User object.
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

// Use this method to log out from the cloud Bot API server before launching the bot locally. You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates. After a successful call, you can immediately log in on a local server, but will not be able to log in back to the cloud Bot API server for 10 minutes. Returns True on success. Requires no parameters.
func (client *Client) LogOut() (bool, error) {

	apiResp, err := client.Request("logOut", nil)
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

// Use this method to close the bot instance before moving it from one local server to another. You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart. The method will return error 429 in the first 10 minutes after the bot is launched. Returns True on success. Requires no parameters.
func (client *Client) Close() (bool, error) {

	apiResp, err := client.Request("close", nil)
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

// Use this method to forward messages of any kind. Service messages and messages with protected content can't be forwarded. On success, the sent Message is returned.
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

// Use this method to forward multiple messages of any kind. If some of the specified messages can't be found or forwarded, they are skipped. Service messages and messages with protected content can't be forwarded. Album grouping is kept for forwarded messages. On success, an array of MessageId of the sent messages is returned.
func (client *Client) ForwardMessages(req *ForwardMessagesRequest) ([]MessageId, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("forwardMessages", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []MessageId

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to copy messages of any kind. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message. Returns the MessageId of the sent message on success.
func (client *Client) CopyMessage(req *CopyMessageRequest) (*MessageId, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("copyMessage", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *MessageId

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to copy messages of any kind. If some of the specified messages can't be found or copied, they are skipped. Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied. A quiz poll can be copied only if the value of the field correct_option_id is known to the bot. The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message. Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
func (client *Client) CopyMessages(req *CopyMessagesRequest) ([]MessageId, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("copyMessages", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []MessageId

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

// Use this method to send live photos. On success, the sent Message is returned.
func (client *Client) SendLivePhoto(req *SendLivePhotoRequest) (*Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendLivePhoto", params)
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

// Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
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

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
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

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
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

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
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

// Use this method to send paid media. On success, the sent Message is returned.
func (client *Client) SendPaidMedia(req *SendPaidMediaRequest) (*Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendPaidMedia", params)
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

// Use this method to send a group of photos, live photos, videos, documents or audios as an album. Documents and audio files can be only grouped in an album with messages of the same type. On success, an array of Message objects that were sent is returned.
func (client *Client) SendMediaGroup(req *SendMediaGroupRequest) ([]Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendMediaGroup", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []Message

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

// Use this method to send a native poll. On success, the sent Message is returned.
func (client *Client) SendPoll(req *SendPollRequest) (*Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendPoll", params)
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

// Use this method to send a checklist on behalf of a connected business account. On success, the sent Message is returned.
func (client *Client) SendChecklist(req *SendChecklistRequest) (*Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendChecklist", params)
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

// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
func (client *Client) SendDice(req *SendDiceRequest) (*Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendDice", params)
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

// Use this method to stream a partial message to a user while the message is being generated. Note that the streamed draft is ephemeral and acts as a temporary 30-second preview - once the output is finalized, you must call sendMessage with the complete message to persist it in the user's chat. Returns True on success.
func (client *Client) SendMessageDraft(req *SendMessageDraftRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendMessageDraft", params)
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

// Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
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

// Use this method to change the chosen reactions on a message. Service messages of some types can't be reacted to. Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel. Bots can't use paid reactions. Returns True on success.
func (client *Client) SetMessageReaction(req *SetMessageReactionRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setMessageReaction", params)
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

// Use this method to get a list of profile audios for a user. Returns a UserProfileAudios object.
func (client *Client) GetUserProfileAudios(req *GetUserProfileAudiosRequest) (*UserProfileAudios, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getUserProfileAudios", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *UserProfileAudios

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App method requestEmojiStatusAccess. Returns True on success.
func (client *Client) SetUserEmojiStatus(req *SetUserEmojiStatusRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setUserEmojiStatus", params)
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

// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type and name (if available) when the File object is received.
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

// Use this method to ban a user in a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
func (client *Client) BanChatMember(req *BanChatMemberRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("banChatMember", params)
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

// Use this method to unban a previously banned user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it. So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned. Returns True on success.
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

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
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

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Pass False for all boolean parameters to demote a user. Returns True on success.
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

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
func (client *Client) SetChatAdministratorCustomTitle(req *SetChatAdministratorCustomTitleRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setChatAdministratorCustomTitle", params)
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

// Use this method to set a tag for a regular member in a group or a supergroup. The bot must be an administrator in the chat for this to work and must have the can_manage_tags administrator right. Returns True on success.
func (client *Client) SetChatMemberTag(req *SetChatMemberTagRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setChatMemberTag", params)
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

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels. The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights. Returns True on success.
func (client *Client) BanChatSenderChat(req *BanChatSenderChatRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("banChatSenderChat", params)
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

// Use this method to unban a previously banned channel chat in a supergroup or channel. The bot must be an administrator for this to work and must have the appropriate administrator rights. Returns True on success.
func (client *Client) UnbanChatSenderChat(req *UnbanChatSenderChatRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("unbanChatSenderChat", params)
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

// Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights. Returns True on success.
func (client *Client) SetChatPermissions(req *SetChatPermissionsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setChatPermissions", params)
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

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the new invite link as String on success.
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

// Use this method to create an additional invite link for a chat. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
func (client *Client) CreateChatInviteLink(req *CreateChatInviteLinkRequest) (*ChatInviteLink, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("createChatInviteLink", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ChatInviteLink

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit a non-primary invite link created by the bot. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the edited invite link as a ChatInviteLink object.
func (client *Client) EditChatInviteLink(req *EditChatInviteLinkRequest) (*ChatInviteLink, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editChatInviteLink", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ChatInviteLink

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to create a subscription invite link for a channel chat. The bot must have the can_invite_users administrator rights. The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink. Returns the new invite link as a ChatInviteLink object.
func (client *Client) CreateChatSubscriptionInviteLink(req *CreateChatSubscriptionInviteLinkRequest) (*ChatInviteLink, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("createChatSubscriptionInviteLink", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ChatInviteLink

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit a subscription invite link created by the bot. The bot must have the can_invite_users administrator rights. Returns the edited invite link as a ChatInviteLink object.
func (client *Client) EditChatSubscriptionInviteLink(req *EditChatSubscriptionInviteLinkRequest) (*ChatInviteLink, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editChatSubscriptionInviteLink", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ChatInviteLink

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to revoke an invite link created by the bot. If the primary link is revoked, a new link is automatically generated. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns the revoked invite link as ChatInviteLink object.
func (client *Client) RevokeChatInviteLink(req *RevokeChatInviteLinkRequest) (*ChatInviteLink, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("revokeChatInviteLink", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ChatInviteLink

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to approve a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (client *Client) ApproveChatJoinRequest(req *ApproveChatJoinRequestRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("approveChatJoinRequest", params)
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

// Use this method to decline a chat join request. The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right. Returns True on success.
func (client *Client) DeclineChatJoinRequest(req *DeclineChatJoinRequestRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("declineChatJoinRequest", params)
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

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
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

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
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

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
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

// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Returns True on success.
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

// Use this method to add a message to the list of pinned messages in a chat. In private chats and channel direct messages chats, all non-service messages can be pinned. Conversely, the bot must be an administrator with the 'can_pin_messages' right or the 'can_edit_messages' right to pin messages in groups and channels respectively. Returns True on success.
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

// Use this method to remove a message from the list of pinned messages in a chat. In private chats and channel direct messages chats, all messages can be unpinned. Conversely, the bot must be an administrator with the 'can_pin_messages' right or the 'can_edit_messages' right to unpin messages in groups and channels respectively. Returns True on success.
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

// Use this method to clear the list of pinned messages in a chat. In private chats and channel direct messages chats, no additional rights are required to unpin all pinned messages. Conversely, the bot must be an administrator with the 'can_pin_messages' right or the 'can_edit_messages' right to unpin all pinned messages in groups and channels respectively. Returns True on success.
func (client *Client) UnpinAllChatMessages(req *UnpinAllChatMessagesRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("unpinAllChatMessages", params)
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

// Use this method to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
func (client *Client) GetChat(req *GetChatRequest) (*ChatFullInfo, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getChat", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ChatFullInfo

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to get a list of administrators in a chat. Returns an Array of ChatMember objects.
func (client *Client) GetChatAdministrators(req *GetChatAdministratorsRequest) ([]ChatMember, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getChatAdministrators", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []ChatMember

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to get the number of members in a chat. Returns Int on success.
func (client *Client) GetChatMemberCount(req *GetChatMemberCountRequest) (int64, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getChatMemberCount", params)
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

// Use this method to get information about a member of a chat. The method is only guaranteed to work for other users if the bot is an administrator in the chat. Returns a ChatMember object on success.
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

// Use this method to get the last messages from the personal chat (i.e., the chat currently added to their profile) of a given user. On success, an array of Message objects is returned.
func (client *Client) GetUserPersonalChatMessages(req *GetUserPersonalChatMessagesRequest) ([]Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getUserPersonalChatMessages", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []Message

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
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

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
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

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user. Requires no parameters. Returns an Array of Sticker objects.
func (client *Client) GetForumTopicIconStickers() ([]Sticker, error) {

	apiResp, err := client.Request("getForumTopicIconStickers", nil)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []Sticker

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to create a topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator right. Returns information about the created topic as a ForumTopic object.
func (client *Client) CreateForumTopic(req *CreateForumTopicRequest) (*ForumTopic, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("createForumTopic", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ForumTopic

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit name and icon of a topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (client *Client) EditForumTopic(req *EditForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editForumTopic", params)
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

// Use this method to close an open topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (client *Client) CloseForumTopic(req *CloseForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("closeForumTopic", params)
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

// Use this method to reopen a closed topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic. Returns True on success.
func (client *Client) ReopenForumTopic(req *ReopenForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("reopenForumTopic", params)
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

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights. Returns True on success.
func (client *Client) DeleteForumTopic(req *DeleteForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteForumTopic", params)
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

// Use this method to clear the list of pinned messages in a forum topic in a forum supergroup chat or a private chat with a user. In the case of a supergroup chat the bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (client *Client) UnpinAllForumTopicMessages(req *UnpinAllForumTopicMessagesRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("unpinAllForumTopicMessages", params)
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

// Use this method to edit the name of the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (client *Client) EditGeneralForumTopic(req *EditGeneralForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editGeneralForumTopic", params)
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

// Use this method to close an open 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (client *Client) CloseGeneralForumTopic(req *CloseGeneralForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("closeGeneralForumTopic", params)
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

// Use this method to reopen a closed 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically unhidden if it was hidden. Returns True on success.
func (client *Client) ReopenGeneralForumTopic(req *ReopenGeneralForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("reopenGeneralForumTopic", params)
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

// Use this method to hide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. The topic will be automatically closed if it was open. Returns True on success.
func (client *Client) HideGeneralForumTopic(req *HideGeneralForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("hideGeneralForumTopic", params)
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

// Use this method to unhide the 'General' topic in a forum supergroup chat. The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights. Returns True on success.
func (client *Client) UnhideGeneralForumTopic(req *UnhideGeneralForumTopicRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("unhideGeneralForumTopic", params)
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

// Use this method to clear the list of pinned messages in a General forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup. Returns True on success.
func (client *Client) UnpinAllGeneralForumTopicMessages(req *UnpinAllGeneralForumTopicMessagesRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("unpinAllGeneralForumTopicMessages", params)
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

// Use this method to reply to a received guest message. On success, a SentGuestMessage object is returned.
func (client *Client) AnswerGuestQuery(req *AnswerGuestQueryRequest) (*SentGuestMessage, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("answerGuestQuery", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *SentGuestMessage

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
func (client *Client) GetUserChatBoosts(req *GetUserChatBoostsRequest) (*UserChatBoosts, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getUserChatBoosts", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *UserChatBoosts

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to get information about the connection of the bot with a business account. Returns a BusinessConnection object on success.
func (client *Client) GetBusinessConnection(req *GetBusinessConnectionRequest) (*BusinessConnection, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getBusinessConnection", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *BusinessConnection

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to get the token of a managed bot. Returns the token as String on success.
func (client *Client) GetManagedBotToken(req *GetManagedBotTokenRequest) (string, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getManagedBotToken", params)
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

// Use this method to revoke the current token of a managed bot and generate a new one. Returns the new token as String on success.
func (client *Client) ReplaceManagedBotToken(req *ReplaceManagedBotTokenRequest) (string, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("replaceManagedBotToken", params)
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

// Use this method to get the access settings of a managed bot. Returns a BotAccessSettings object on success.
func (client *Client) GetManagedBotAccessSettings(req *GetManagedBotAccessSettingsRequest) (*BotAccessSettings, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getManagedBotAccessSettings", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *BotAccessSettings

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to change the access settings of a managed bot. Returns True on success.
func (client *Client) SetManagedBotAccessSettings(req *SetManagedBotAccessSettingsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setManagedBotAccessSettings", params)
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

// Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
func (client *Client) SetMyCommands(req *SetMyCommandsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setMyCommands", params)
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

// Use this method to delete the list of the bot's commands for the given scope and user language. After deletion, higher level commands will be shown to affected users. Returns True on success.
func (client *Client) DeleteMyCommands(req *DeleteMyCommandsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteMyCommands", params)
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

// Use this method to get the current list of the bot's commands for the given scope and user language. Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
func (client *Client) GetMyCommands(req *GetMyCommandsRequest) ([]BotCommand, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getMyCommands", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []BotCommand

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to change the bot's name. Returns True on success.
func (client *Client) SetMyName(req *SetMyNameRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setMyName", params)
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

// Use this method to get the current bot name for the given user language. Returns BotName on success.
func (client *Client) GetMyName(req *GetMyNameRequest) (*BotName, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getMyName", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *BotName

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
func (client *Client) SetMyDescription(req *SetMyDescriptionRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setMyDescription", params)
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

// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
func (client *Client) GetMyDescription(req *GetMyDescriptionRequest) (*BotDescription, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getMyDescription", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *BotDescription

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot. Returns True on success.
func (client *Client) SetMyShortDescription(req *SetMyShortDescriptionRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setMyShortDescription", params)
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

// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
func (client *Client) GetMyShortDescription(req *GetMyShortDescriptionRequest) (*BotShortDescription, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getMyShortDescription", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *BotShortDescription

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Changes the profile photo of the bot. Returns True on success.
func (client *Client) SetMyProfilePhoto(req *SetMyProfilePhotoRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setMyProfilePhoto", params)
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

// Removes the profile photo of the bot. Requires no parameters. Returns True on success.
func (client *Client) RemoveMyProfilePhoto() (bool, error) {

	apiResp, err := client.Request("removeMyProfilePhoto", nil)
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

// Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
func (client *Client) SetChatMenuButton(req *SetChatMenuButtonRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setChatMenuButton", params)
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

// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button. Returns MenuButton on success.
func (client *Client) GetChatMenuButton(req *GetChatMenuButtonRequest) (*MenuButton, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getChatMenuButton", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *MenuButton

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels. These rights will be suggested to users, but they are free to modify the list before adding the bot. Returns True on success.
func (client *Client) SetMyDefaultAdministratorRights(req *SetMyDefaultAdministratorRightsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setMyDefaultAdministratorRights", params)
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

// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
func (client *Client) GetMyDefaultAdministratorRights(req *GetMyDefaultAdministratorRightsRequest) (*ChatAdministratorRights, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getMyDefaultAdministratorRights", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *ChatAdministratorRights

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Returns the list of gifts that can be sent by the bot to users and channel chats. Requires no parameters. Returns a Gifts object.
func (client *Client) GetAvailableGifts() (*Gifts, error) {

	apiResp, err := client.Request("getAvailableGifts", nil)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *Gifts

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Sends a gift to the given user or channel chat. The gift can't be converted to Telegram Stars by the receiver. Returns True on success.
func (client *Client) SendGift(req *SendGiftRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("sendGift", params)
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

// Gifts a Telegram Premium subscription to the given user. Returns True on success.
func (client *Client) GiftPremiumSubscription(req *GiftPremiumSubscriptionRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("giftPremiumSubscription", params)
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

// Verifies a user on behalf of the organization which is represented by the bot. Returns True on success.
func (client *Client) VerifyUser(req *VerifyUserRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("verifyUser", params)
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

// Verifies a chat on behalf of the organization which is represented by the bot. Returns True on success.
func (client *Client) VerifyChat(req *VerifyChatRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("verifyChat", params)
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

// Removes verification from a user who is currently verified on behalf of the organization represented by the bot. Returns True on success.
func (client *Client) RemoveUserVerification(req *RemoveUserVerificationRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("removeUserVerification", params)
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

// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot. Returns True on success.
func (client *Client) RemoveChatVerification(req *RemoveChatVerificationRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("removeChatVerification", params)
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

// Marks incoming message as read on behalf of a business account. Requires the can_read_messages business bot right. Returns True on success.
func (client *Client) ReadBusinessMessage(req *ReadBusinessMessageRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("readBusinessMessage", params)
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

// Delete messages on behalf of a business account. Requires the can_delete_sent_messages business bot right to delete messages sent by the bot itself, or the can_delete_all_messages business bot right to delete any message. Returns True on success.
func (client *Client) DeleteBusinessMessages(req *DeleteBusinessMessagesRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteBusinessMessages", params)
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

// Changes the first and last name of a managed business account. Requires the can_change_name business bot right. Returns True on success.
func (client *Client) SetBusinessAccountName(req *SetBusinessAccountNameRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setBusinessAccountName", params)
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

// Changes the username of a managed business account. Requires the can_change_username business bot right. Returns True on success.
func (client *Client) SetBusinessAccountUsername(req *SetBusinessAccountUsernameRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setBusinessAccountUsername", params)
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

// Changes the bio of a managed business account. Requires the can_change_bio business bot right. Returns True on success.
func (client *Client) SetBusinessAccountBio(req *SetBusinessAccountBioRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setBusinessAccountBio", params)
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

// Changes the profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
func (client *Client) SetBusinessAccountProfilePhoto(req *SetBusinessAccountProfilePhotoRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setBusinessAccountProfilePhoto", params)
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

// Removes the current profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.
func (client *Client) RemoveBusinessAccountProfilePhoto(req *RemoveBusinessAccountProfilePhotoRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("removeBusinessAccountProfilePhoto", params)
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

// Changes the privacy settings pertaining to incoming gifts in a managed business account. Requires the can_change_gift_settings business bot right. Returns True on success.
func (client *Client) SetBusinessAccountGiftSettings(req *SetBusinessAccountGiftSettingsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setBusinessAccountGiftSettings", params)
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

// Returns the amount of Telegram Stars owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns StarAmount on success.
func (client *Client) GetBusinessAccountStarBalance(req *GetBusinessAccountStarBalanceRequest) (*StarAmount, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getBusinessAccountStarBalance", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *StarAmount

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Transfers Telegram Stars from the business account balance to the bot's balance. Requires the can_transfer_stars business bot right. Returns True on success.
func (client *Client) TransferBusinessAccountStars(req *TransferBusinessAccountStarsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("transferBusinessAccountStars", params)
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

// Returns the gifts received and owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns OwnedGifts on success.
func (client *Client) GetBusinessAccountGifts(req *GetBusinessAccountGiftsRequest) (*OwnedGifts, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getBusinessAccountGifts", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *OwnedGifts

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Returns the gifts owned and hosted by a user. Returns OwnedGifts on success.
func (client *Client) GetUserGifts(req *GetUserGiftsRequest) (*OwnedGifts, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getUserGifts", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *OwnedGifts

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Returns the gifts owned by a chat. Returns OwnedGifts on success.
func (client *Client) GetChatGifts(req *GetChatGiftsRequest) (*OwnedGifts, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getChatGifts", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *OwnedGifts

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Converts a given regular gift to Telegram Stars. Requires the can_convert_gifts_to_stars business bot right. Returns True on success.
func (client *Client) ConvertGiftToStars(req *ConvertGiftToStarsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("convertGiftToStars", params)
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

// Upgrades a given regular gift to a unique gift. Requires the can_transfer_and_upgrade_gifts business bot right. Additionally requires the can_transfer_stars business bot right if the upgrade is paid. Returns True on success.
func (client *Client) UpgradeGift(req *UpgradeGiftRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("upgradeGift", params)
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

// Transfers an owned unique gift to another user. Requires the can_transfer_and_upgrade_gifts business bot right. Requires can_transfer_stars business bot right if the transfer is paid. Returns True on success.
func (client *Client) TransferGift(req *TransferGiftRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("transferGift", params)
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

// Posts a story on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
func (client *Client) PostStory(req *PostStoryRequest) (*Story, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("postStory", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *Story

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Reposts a story on behalf of a business account from another business account. Both business accounts must be managed by the same bot, and the story on the source account must have been posted (or reposted) by the bot. Requires the can_manage_stories business bot right for both business accounts. Returns Story on success.
func (client *Client) RepostStory(req *RepostStoryRequest) (*Story, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("repostStory", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *Story

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Edits a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns Story on success.
func (client *Client) EditStory(req *EditStoryRequest) (*Story, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editStory", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *Story

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Deletes a story previously posted by the bot on behalf of a managed business account. Requires the can_manage_stories business bot right. Returns True on success.
func (client *Client) DeleteStory(req *DeleteStoryRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteStory", params)
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

// Use this method to set the result of an interaction with a Web App and send a corresponding message on behalf of the user to the chat from which the query originated. On success, a SentWebAppMessage object is returned.
func (client *Client) AnswerWebAppQuery(req *AnswerWebAppQueryRequest) (*SentWebAppMessage, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("answerWebAppQuery", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *SentWebAppMessage

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.
func (client *Client) SavePreparedInlineMessage(req *SavePreparedInlineMessageRequest) (*PreparedInlineMessage, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("savePreparedInlineMessage", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *PreparedInlineMessage

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Stores a keyboard button that can be used by a user within a Mini App. Returns a PreparedKeyboardButton object.
func (client *Client) SavePreparedKeyboardButton(req *SavePreparedKeyboardButtonRequest) (*PreparedKeyboardButton, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("savePreparedKeyboardButton", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *PreparedKeyboardButton

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit text and game messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (client *Client) EditMessageText(req *EditMessageTextRequest) (any, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editMessageText", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp any

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit captions of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (client *Client) EditMessageCaption(req *EditMessageCaptionRequest) (any, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editMessageCaption", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp any

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit animation, audio, document, live photo, photo, or video messages, or to add media to text messages. If a message is part of a message album, then it can be edited only to an audio for audio albums, only to a document for document albums and to a photo, a live photo, or a video otherwise. When an inline message is edited, a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (client *Client) EditMessageMedia(req *EditMessageMediaRequest) (any, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editMessageMedia", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp any

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
func (client *Client) EditMessageLiveLocation(req *EditMessageLiveLocationRequest) (any, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editMessageLiveLocation", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp any

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to stop updating a live location message before live_period expires. On success, if the message is not an inline message, the edited Message is returned, otherwise True is returned.
func (client *Client) StopMessageLiveLocation(req *StopMessageLiveLocationRequest) (any, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("stopMessageLiveLocation", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp any

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to edit a checklist on behalf of a connected business account. On success, the edited Message is returned.
func (client *Client) EditMessageChecklist(req *EditMessageChecklistRequest) (*Message, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editMessageChecklist", params)
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

// Use this method to edit only the reply markup of messages. On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned. Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
func (client *Client) EditMessageReplyMarkup(req *EditMessageReplyMarkupRequest) (any, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editMessageReplyMarkup", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp any

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll is returned.
func (client *Client) StopPoll(req *StopPollRequest) (*Poll, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("stopPoll", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *Poll

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to approve a suggested post in a direct messages chat. The bot must have the 'can_post_messages' administrator right in the corresponding channel chat. Returns True on success.
func (client *Client) ApproveSuggestedPost(req *ApproveSuggestedPostRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("approveSuggestedPost", params)
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

// Use this method to decline a suggested post in a direct messages chat. The bot must have the 'can_manage_direct_messages' administrator right in the corresponding channel chat. Returns True on success.
func (client *Client) DeclineSuggestedPost(req *DeclineSuggestedPostRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("declineSuggestedPost", params)
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

// Use this method to delete multiple messages simultaneously. If some of the specified messages can't be found, they are skipped. Returns True on success.
func (client *Client) DeleteMessages(req *DeleteMessagesRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteMessages", params)
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

// Use this method to remove a reaction from a message in a group or a supergroup chat. The bot must have the 'can_delete_messages' administrator right in the chat. Returns True on success.
func (client *Client) DeleteMessageReaction(req *DeleteMessageReactionRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteMessageReaction", params)
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

// Use this method to remove up to 10000 recent reactions in a group or a supergroup chat added by a given user or chat. The bot must have the 'can_delete_messages' administrator right in the chat. Returns True on success.
func (client *Client) DeleteAllMessageReactions(req *DeleteAllMessageReactionsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteAllMessageReactions", params)
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

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.
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

// Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.
func (client *Client) GetCustomEmojiStickers(req *GetCustomEmojiStickersRequest) ([]Sticker, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getCustomEmojiStickers", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []Sticker

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.
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

// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.
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

// Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.
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

// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
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

// Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.
func (client *Client) ReplaceStickerInSet(req *ReplaceStickerInSetRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("replaceStickerInSet", params)
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

// Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (client *Client) SetStickerEmojiList(req *SetStickerEmojiListRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setStickerEmojiList", params)
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

// Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.
func (client *Client) SetStickerKeywords(req *SetStickerKeywordsRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setStickerKeywords", params)
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

// Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.
func (client *Client) SetStickerMaskPosition(req *SetStickerMaskPositionRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setStickerMaskPosition", params)
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

// Use this method to set the title of a created sticker set. Returns True on success.
func (client *Client) SetStickerSetTitle(req *SetStickerSetTitleRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setStickerSetTitle", params)
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

// Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.
func (client *Client) SetStickerSetThumbnail(req *SetStickerSetThumbnailRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setStickerSetThumbnail", params)
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

// Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.
func (client *Client) SetCustomEmojiStickerSetThumbnail(req *SetCustomEmojiStickerSetThumbnailRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setCustomEmojiStickerSetThumbnail", params)
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

// Use this method to delete a sticker set that was created by the bot. Returns True on success.
func (client *Client) DeleteStickerSet(req *DeleteStickerSetRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("deleteStickerSet", params)
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

// Use this method to send answers to an inline query. On success, True is returned.
// No more than 50 results per query are allowed.
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

// Use this method to create a link for an invoice. Returns the created invoice link as String on success.
func (client *Client) CreateInvoiceLink(req *CreateInvoiceLinkRequest) (string, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("createInvoiceLink", params)
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

// A method to get the current Telegram Stars balance of the bot. Requires no parameters. On success, returns a StarAmount object.
func (client *Client) GetMyStarBalance() (*StarAmount, error) {

	apiResp, err := client.Request("getMyStarBalance", nil)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *StarAmount

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Returns the bot's Telegram Star transactions in chronological order. On success, returns a StarTransactions object.
func (client *Client) GetStarTransactions(req *GetStarTransactionsRequest) (*StarTransactions, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getStarTransactions", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp *StarTransactions

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Refunds a successful payment in Telegram Stars. Returns True on success.
func (client *Client) RefundStarPayment(req *RefundStarPaymentRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("refundStarPayment", params)
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

// Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.
func (client *Client) EditUserStarSubscription(req *EditUserStarSubscriptionRequest) (bool, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("editUserStarSubscription", params)
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
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.
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

// Use this method to set the score of the specified user in a game message. On success, if the message is not an inline message, the Message is returned, otherwise True is returned. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
func (client *Client) SetGameScore(req *SetGameScoreRequest) (any, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("setGameScore", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp any

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Use this method to get data for high score tables. Will return the score of the specified user and several of their neighbors in a game. Returns an Array of GameHighScore objects.
func (client *Client) GetGameHighScores(req *GetGameHighScoresRequest) ([]GameHighScore, error) {
	params := requestToMap(req)

	apiResp, err := client.Request("getGameHighScores", params)
	if err != nil {
		return nil, err
	}

	if !apiResp.Ok {
		return nil, newError(apiResp)
	}

	var resp []GameHighScore

	err = json.Unmarshal(apiResp.Result, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
