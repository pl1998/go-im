/**
  @author:panliang
  @data:2021/6/26
  @note
**/
package auth

import "C"
import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	messageModel "im_app/im/http/models/msg"
	userModel "im_app/im/http/models/user"
	"im_app/pkg/helpler"
	"im_app/pkg/model"
	"im_app/pkg/response"
)

type (
	UsersController struct{}
	UsersList       struct {
		ID            uint64    `json:"id"`
		Email         string    `json:"email"`
		Avatar        string    `json:"avatar"`
		Name          string    `json:"name"`
		Msg           string    `json:"msg"`
		Status        int       `json:"status"`
		IsRead        int       `json:"is_read"`
		SendTime      string    `json:"send_time"`
		SendMsg       string    `json:"send_msg"`
		MsgTotal      int       `json:"msg_total"`
		ClientType    int       `json:"client_type"`
		Bio           int       `json:"bio"`
		Sex           int       `json:"sex"`
		LastLoginTime time.Time `gorm:"type:time" json:"last_login_time"`
	}
	NotFriendList struct {
		Name   string `json:"name"`
		ID     uint64 `json:"id"`
		Avatar string `json:"avatar"`
		Status string `json:"status"`
	}
)

// @BasePath /api

// @Summary 获取非好友用户列表
// @Description 获取非好友用户列表
// @Tags 获取非好友用户列表
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param name query string false "账号"
// @Produce json
// @Success 200
// @Router /UsersList [get]
func (*UsersController) GetUsersList(c *gin.Context) {
	name := c.Query("name")
	user := userModel.AuthUser

	var Users []userModel.Users
	var NotFriendList []NotFriendList

	subQuery := model.DB.Select("f_id").
		Where("m_id=?", user.ID).
		Table("im_friends")

	if len(name) > 0 {

		model.DB.Model(&Users).
			Select("im_users.id,im_users.name,im_users.avatar,im_friend_records.status").
			Joins("left join im_friend_records on im_friend_records.f_id=im_users.id"+
				" and im_users.id not in(?) and im_users.id!=? and im_users.name like '%?%' limit 10 ", subQuery, user.ID, name).
			Scan(&NotFriendList)
	} else {

		model.DB.Model(&Users).
			Select("im_users.id,im_users.name,im_users.avatar,im_friend_records.status").
			Joins("left join im_friend_records on im_friend_records.f_id=im_users.id"+
				" and im_users.id not in(?) and im_users.id!=? limit 10 ", subQuery, user.ID).
			Scan(&NotFriendList)
	}

	response.SuccessResponse(map[string]interface{}{
		"list": NotFriendList,
	}, 200).ToJson(c)
}

// @Summary 历史消息读取[废弃]
// @Description 历史消息读取
// @Tags 历史消息读取
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization
// @Param Authorization	header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param voice formData file true "图片上传"
// @Produce json
// @Success 200
// @Router /ReadMessage [get]
func (*UsersController) ReadMessage(c *gin.Context) {
	user := userModel.AuthUser
	channel_a, channel_b := helpler.ProduceChannelName(strconv.Itoa(int(user.ID)), c.Query("to_id"))
	messageModel.ReadMsg(channel_a, channel_b)
	response.SuccessResponse(gin.H{}, 200).ToJson(c)
}
