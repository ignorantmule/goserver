// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"errors"

	"gorm.io/gorm"

	v1 "github.com/marmotedu/goserver/internal/pkg/model/goserver/v1"
	"github.com/marmotedu/goserver/internal/pkg/util/gormutil"
	"github.com/marmotedu/goserver/pkg/errno"
	"github.com/marmotedu/goserver/pkg/fields"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

type posts struct {
	db *gorm.DB
}

func newPosts(ds *datastore) *posts {
	return &posts{ds.db}
}

// Create creates a new post.
func (p *posts) Create(ctx context.Context, post *v1.Post, opts metav1.CreateOptions) error {
	return p.db.Create(&post).Error
}

// Update updates an post information by the post identifier.
func (p *posts) Update(ctx context.Context, post *v1.Post, opts metav1.UpdateOptions) error {
	return p.db.Save(post).Error
}

// Delete deletes the post by the post identifier.
func (p *posts) Delete(ctx context.Context, username, postID string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	err := p.db.Where("username = ? and postID = ?", username, postID).Delete(&v1.Post{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errno.ErrPostNotFound
	}

	return nil
}

// DeleteCollection batch deletes the posts.
func (p *posts) DeleteCollection(
	ctx context.Context,
	username string,
	postIDs []string,
	opts metav1.DeleteOptions,
) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username = ? and postID in (?)", username, postIDs).Delete(&v1.Post{}).Error
}

// Get return an post by the post identifier.
func (p *posts) Get(ctx context.Context, username, postID string, opts metav1.GetOptions) (*v1.Post, error) {
	post := &v1.Post{}
	err := p.db.Where("username = ? and postID = ?", username, postID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPostNotFound
		}

		return nil, err
	}

	return post, nil
}

// List return all posts.
func (p *posts) List(ctx context.Context, username string, opts metav1.ListOptions) (*v1.PostList, error) {
	ret := &v1.PostList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	if username != "" {
		p.db = p.db.Where("username = ?", username)
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	title, _ := selector.RequiresExactMatch("title")

	d := p.db.Where("title like ?", "%"+title+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
