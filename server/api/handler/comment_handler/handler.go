// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handlers

import (
	"douyin-user/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendResponse pack response
func SendResponse(c *app.RequestContext, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}

type NoteParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
type CommentRequestParam struct {
	UserId      int64  `thrift:"UserId,1" frugal:"1,default,i64" json:"UserId"`
	VideoId     int64  `thrift:"VideoId,2" frugal:"2,default,i64" json:"VideoId"`
	ActionType  int64  `thrift:"actionType,3" frugal:"3,default,i64" json:"actionType"`
	CommentText string `thrift:"commentText,4" frugal:"4,default,string" json:"commentText"`
	CommentId   int64  `thrift:"commentId,5" frugal:"5,default,i64" json:"commentId"`
}
