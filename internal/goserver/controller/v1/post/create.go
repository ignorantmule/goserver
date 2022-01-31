// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/internal/pkg/constant"
	"github.com/marmotedu/goserver/internal/pkg/log"
	v1 "github.com/marmotedu/goserver/internal/pkg/model/goserver/v1"
	"github.com/marmotedu/goserver/pkg/core"
	"github.com/marmotedu/goserver/pkg/errno"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// Create add new post to the storage.
func (p *PostController) Create(c *gin.Context) {
	log.L(c).Info("post create function called.")

	var r v1.Post

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	// Add username
	r.Username = c.GetString(constant.XUsernameKey)

	if err := r.Validate(); err != nil {
		core.WriteResponse(c, errno.ErrValidation, nil)

		return
	}

	// Insert the post to the storage.
	if err := p.srv.Posts().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, map[string]string{"postID": r.PostID})
}
