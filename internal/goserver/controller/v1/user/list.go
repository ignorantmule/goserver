// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/internal/pkg/log"
	"github.com/marmotedu/goserver/pkg/core"
	"github.com/marmotedu/goserver/pkg/errno"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// List list the users in the storage.
// Only administrator can call this function.
func (u *UserController) List(c *gin.Context) {
	log.L(c).Info("list user function called.")

	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	users, err := u.srv.Users().List(c, r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, users)
}
