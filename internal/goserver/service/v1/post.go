// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"
	"regexp"

	"github.com/marmotedu/goserver/internal/goserver/store"
	v1 "github.com/marmotedu/goserver/internal/pkg/model/goserver/v1"
	"github.com/marmotedu/goserver/pkg/errno"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// PostSrv defines functions used to handle post request.
type PostSrv interface {
	Create(ctx context.Context, post *v1.Post, opts metav1.CreateOptions) error
	Update(ctx context.Context, post *v1.Post, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, postID string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, postIDs []string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, postID string, opts metav1.GetOptions) (*v1.Post, error)
	List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PostList, error)
}

type postService struct {
	store store.Factory
}

var _ PostSrv = (*postService)(nil)

func newPosts(srv *service) *postService {
	return &postService{store: srv.store}
}

func (p *postService) Create(ctx context.Context, post *v1.Post, opts metav1.CreateOptions) error {
	if err := p.store.Posts().Create(ctx, post, opts); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'title'", err.Error()); match {
			return errno.ErrPostAlreadyExist
		}

		return err
	}

	return nil
}

func (p *postService) Update(ctx context.Context, post *v1.Post, opts metav1.UpdateOptions) error {
	// Save changed fields.
	if err := p.store.Posts().Update(ctx, post, opts); err != nil {
		return errno.ErrDatabase
	}

	return nil
}

func (p *postService) Delete(ctx context.Context, username, postID string, opts metav1.DeleteOptions) error {
	if err := p.store.Posts().Delete(ctx, username, postID, opts); err != nil {
		return err
	}

	return nil
}

func (p *postService) DeleteCollection(
	ctx context.Context,
	username string,
	postIDs []string,
	opts metav1.DeleteOptions,
) error {
	if err := p.store.Posts().DeleteCollection(ctx, username, postIDs, opts); err != nil {
		return errno.ErrDatabase
	}

	return nil
}

func (p *postService) Get(
	ctx context.Context,
	username, postID string,
	opts metav1.GetOptions,
) (*v1.Post, error) {
	post, err := p.store.Posts().Get(ctx, username, postID, opts)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postService) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PostList, error) {
	posts, err := p.store.Posts().List(ctx, username, opts)
	if err != nil {
		return nil, errno.ErrDatabase
	}

	return posts, nil
}
