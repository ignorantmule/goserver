// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/internal/pkg/log"
	v1 "github.com/marmotedu/goserver/internal/pkg/model/goserver/v1"
	"github.com/marmotedu/goserver/pkg/auth"
	"github.com/marmotedu/goserver/pkg/core"
	"github.com/marmotedu/goserver/pkg/errno"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// Create add new user to the storage.
func (u *UserController) Create(c *gin.Context) {
	log.L(c).Info("user create function called.")

	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if err := r.Validate(); err != nil {
		core.WriteResponse(c, errno.ErrValidation, nil)

		return
	}

	var err error

	// Encrypt the user password.
	r.Password, err = auth.Encrypt(r.Password)
	if err != nil {
		core.WriteResponse(c, errno.ErrEncrypt, nil)

		return
	}

	// Insert the user to the storage.
	if err := u.srv.Users().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, r)
}
