package douyingo

import (
	"context"
	"errors"
	"github.com/CriarBrand/douyin-go/conf"
)

// ImMessageReq 发私信给用户请求(企业号)
type ImMessageReq struct {
	OpenId          string        // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken     string        // 调用/oauth/access_token/生成的token，此token需要用户授权。
	ImMessageSchema string        // 调给body里Content
	Body            ImMessageBody // 回复视频评论body
}

type ImMessageContent struct {
	Text    string `json:"text,omitempty"`
	MediaId string `json:"media_id,omitempty"`
	ItemId  string `json:"item_id,omitempty"`
	CardId  string `json:"card_id,omitempty"`
}

// ImMessageBody 发私信给用户请求body(企业号)
type ImMessageBody struct {
	Content     ImMessageContent `json:"content"`              // 消息体见下方schema，需进行json转义
	MessageType string           `json:"message_type"`         // 消息内容格式:"text"(文本消息) "image"(图片消息) "video"(视频消息) "card"(卡片消息)
	PersonaId   string           `json:"persona_id,omitempty"` // 客服id，传则走客服会话，否则为普通会话
	ToUserId    string           `json:"to_user_id"`           // 消息的接收方openid
}

// ImMessageData 发私信给用户data(企业号)
type ImMessageData struct {
	DYError
	ServerMsgId string `json:"server_msg_id,omitempty"` // 内部使用（暂时不懂是什么意思）
}

// ImMessageRes 发私信给用户返回(企业号)
type ImMessageRes struct {
	Data  ImMessageData `json:"data"`
	Extra DYExtra       `json:"extra"`
}

// ImMessage 发私信给用户(企业号)
func (m *Manager) ImMessage(req ImMessageReq) (res ImMessageRes, err error) {
	switch req.Body.MessageType {
	case "text":
		req.Body.Content.Text = req.ImMessageSchema
	case "image":
		req.Body.Content.MediaId = req.ImMessageSchema
	case "video":
		req.Body.Content.ItemId = req.ImMessageSchema
	case "card":
		req.Body.Content.CardId = req.ImMessageSchema
	default:
		return res, errors.New("MessageType必须是“text”，“image”，“video“，”card“其中一个")
	}
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_IM_MESSAGE_SEND, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}
