// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/pkg/errno"
)

// ErrResponse defines the return messages when an error occurred.
// Reference will be omitted if it does not exist.
// swagger:model
type ErrResponse struct {
	// Code defines the business error code.
	Code int `json:"code"`

	// Message contains the detail of this message.
	// This message is suitable to be exposed to external
	Message string `json:"message"`
}

// WriteResponse write an error or the response data into http response body.
// It use errors.ParseCoder to parse any error into errors.Coder
// errors.Coder contains error code, user-safe error message and http status code.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		code, message := errno.DecodeErr(err)
		c.JSON(http.StatusOK, ErrResponse{
			Code:    code,
			Message: message,
		})

		return
	}

	c.JSON(http.StatusOK, data)
}
