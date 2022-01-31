// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package store

import (
	"context"

	v1 "github.com/marmotedu/goserver/internal/pkg/model/goserver/v1"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// PostStore defines the post storage interface.
type PostStore interface {
	Create(ctx context.Context, post *v1.Post, opts metav1.CreateOptions) error
	Update(ctx context.Context, post *v1.Post, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, postID string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, postIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, postID string, opts metav1.GetOptions) (*v1.Post, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PostList, error)
}
