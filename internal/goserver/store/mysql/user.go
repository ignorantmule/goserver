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

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

// Create creates a new user account.
func (u *users) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	return u.db.Create(&user).Error
}

// Update updates an user account information.
func (u *users) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	return u.db.Save(user).Error
}

// Delete deletes the user by the user identifier.
func (u *users) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	err := u.db.Where("username = ?", username).Delete(&v1.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errno.ErrUserNotFound
	}

	return nil
}

// Get return an user by the user identifier.
func (u *users) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (u *users) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	ret := &v1.UserList{}
	ol := gormutil.Unpointer(opts.Offset, opts.Limit)

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, _ := selector.RequiresExactMatch("username")
	d := u.db.Where("username like ?", "%"+username+"%").
		Offset(ol.Offset).
		Limit(ol.Limit).
		Order("id desc").
		Find(&ret.Items).
		Offset(-1).
		Limit(-1).
		Count(&ret.TotalCount)

	return ret, d.Error
}
