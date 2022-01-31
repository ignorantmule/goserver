// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package post

import (
	srvv1 "github.com/marmotedu/goserver/internal/goserver/service/v1"
	"github.com/marmotedu/goserver/internal/goserver/store"
)

// PostController create a post handler used to handle request for post resource.
type PostController struct {
	srv srvv1.Service
}

// NewPostController creates a post handler.
func NewPostController(store store.Factory) *PostController {
	return &PostController{
		srv: srvv1.NewService(store),
	}
}
