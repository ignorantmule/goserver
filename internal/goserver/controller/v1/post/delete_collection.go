// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/goserver/internal/goserver/constant"
	"github.com/marmotedu/goserver/internal/pkg/log"
	"github.com/marmotedu/goserver/pkg/core"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// DeleteCollection batch delete posts by multiple post ids.
func (p *PostController) DeleteCollection(c *gin.Context) {
	log.L(c).Info("batch delete post function called.")

	postIDs := c.QueryArray("postID")

	if err := p.srv.Posts().DeleteCollection(c, c.GetString(constant.XUsernameKey), postIDs, metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
