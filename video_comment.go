package douyingo

import (
	"context"
	"github.com/CriarBrand/douyin-go/conf"
	"net/url"
)

// VideoCommentListReq 评论列表请求(企业号)
type VideoCommentListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	ItemId      string // 视频id
}

// VideoComment 评论列表(企业号)
type VideoComment struct {
	ReplyCommentTotal int32  `json:"reply_comment_total"` // 回复评论数
	Top               bool   `json:"top"`                 // 是否置顶评论
	CommentId         string `json:"comment_id"`          // 评论id
	CommentUserId     string `json:"comment_user_id"`     // 评论用户id
	Content           string `json:"content"`             // 评论内容
	CreateTime        int64  `json:"create_time"`         // 时间戳
	DiggCount         int32  `json:"digg_count"`          // 点赞数
}

// VideoCommentListData 评论列表(企业号)
type VideoCommentListData struct {
	List    []VideoComment `json:"list"`     // 评论列表
	Cursor  int64          `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool           `json:"has_more"` // 更多数据
	DYError
	Extra DYExtra `json:"extra"`
}

// VideoCommentListRes 评论列表(企业号)
type VideoCommentListRes struct {
	Data ItemCommentListData `json:"data"`
}

// VideoCommentList 获取评论列表(企业号)
func (m *Manager) VideoCommentList(req VideoCommentListReq) (res VideoCommentListRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d&item_id=%s", conf.API_VIDEO_COMMENT_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count, itemId), nil, nil)
	return res, err
}

// VideoCommentReplyListReq 评论回复列表请求(企业号)
type VideoCommentReplyListReq struct {
	OpenId      string // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Cursor      int64  // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
	Count       int64  // 每页数量
	ItemId      string // 视频id
	CommentId   string // 评论id
}

// VideoCommentReplyListData 评论回复列表(企业号)
type VideoCommentReplyListData struct {
	List    []VideoComment `json:"list"`     // 评论回复列表
	Cursor  int64          `json:"cursor"`   // 用于下一页请求的cursor
	HasMore bool           `json:"has_more"` // 更多数据
	DYError
	Extra DYExtra `json:"extra"`
}

// VideoCommentReplyListRes 评论回复列表(企业号)
type VideoCommentReplyListRes struct {
	Data ItemCommentReplyListData `json:"data"`
}

// VideoCommentReplyList 获取评论回复列表(企业号)
func (m *Manager) VideoCommentReplyList(req VideoCommentReplyListReq) (res VideoCommentReplyListRes, err error) {
	itemId := url.QueryEscape(req.ItemId)
	commentId := url.QueryEscape(req.CommentId)
	err = m.client.CallWithJson(context.Background(), &res, "GET", m.url("%s?access_token=%s&open_id=%s&cursor=%d&count=%d&item_id=%s&comment_id=%s", conf.API_VIDEO_COMMENT_REPLY_LIST, req.AccessToken, req.OpenId, req.Cursor, req.Count, itemId, commentId), nil, nil)
	return res, err
}

// VideoCommentReplyReq 回复视频评论请求(企业号)
type VideoCommentReplyReq struct {
	OpenId      string                // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string                // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoCommentReplyBody // 回复视频评论body
}

// VideoCommentReplyBody 回复视频评论(企业号)
type VideoCommentReplyBody struct {
	CommentId string `json:"comment_id,omitempty"` // 需要回复的评论id（如果需要回复的是视频不传此字段）
	Content   string `json:"content"`              // 评论内容
	ItemId    string `json:"item_id"`              // 视频id
}

// VideoCommentReplyData 回复视频评论(企业号)
type VideoCommentReplyData struct {
	CommentId string `json:"comment_id"` // 评论id
}

// VideoCommentReplyRes 回复视频评论(企业号)
type VideoCommentReplyRes struct {
	Data  VideoCommentReplyData `json:"data"`
	Extra DYExtra               `json:"extra"`
}

// VideoCommentReply 回复视频评论(企业号)
func (m *Manager) VideoCommentReply(req VideoCommentReplyReq) (res VideoCommentReplyRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_COMMENT_REPLY, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}

// VideoCommentTopReq 置顶视频评论请求(企业号)
type VideoCommentTopReq struct {
	OpenId      string              // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string              // 调用/oauth/access_token/生成的token，此token需要用户授权。
	Body        VideoCommentTopBody // 回复视频评论body
}

// VideoCommentTopBody 置顶视频评论请求body(企业号)
type VideoCommentTopBody struct {
	CommentId string `json:"comment_id,omitempty"` // 需要回复的评论id（如果需要回复的是视频不传此字段）
	Top       bool   `json:"top"`                  // true: 置顶, false: 取消置顶
	ItemId    string `json:"item_id"`              // 视频id
}

// VideoCommentTopRes 置顶视频评论返回(企业号)
type VideoCommentTopRes struct {
	Data  DYError `json:"data"`
	Extra DYExtra `json:"extra"`
}

// VideoCommentTop 置顶视频评论(企业号)
func (m *Manager) VideoCommentTop(req VideoCommentTopReq) (res VideoCommentTopRes, err error) {
	err = m.client.CallWithJson(context.Background(), &res, "POST", m.url("%s?access_token=%s&open_id=%s", conf.API_VIDEO_COMMENT_TOP, req.AccessToken, req.OpenId), nil, req.Body)
	return res, err
}
