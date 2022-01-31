// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/internal/pkg/constant"
	"github.com/marmotedu/goserver/internal/pkg/log"
	"github.com/marmotedu/goserver/pkg/core"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// Delete delete a post by the post identifier.
func (p *PostController) Delete(c *gin.Context) {
	log.L(c).Info("delete post function called.")

	if err := p.srv.Posts().Delete(c, c.GetString(constant.XUsernameKey), c.Param("postID"), metav1.DeleteOptions{Unscoped: true}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
