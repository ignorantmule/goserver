// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package v1

import (
	"context"
	"regexp"
	"sync"

	"github.com/marmotedu/goserver/internal/goserver/store"
	"github.com/marmotedu/goserver/internal/pkg/log"
	v1 "github.com/marmotedu/goserver/internal/pkg/model/goserver/v1"
	"github.com/marmotedu/goserver/pkg/auth"
	"github.com/marmotedu/goserver/pkg/errno"
	metav1 "github.com/marmotedu/goserver/pkg/meta/v1"
)

// UserSrv defines functions used to handle user request.
type UserSrv interface {
	Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error
	Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error)
}

type userService struct {
	store store.Factory
}

var _ UserSrv = (*userService)(nil)

func newUsers(srv *service) *userService {
	return &userService{store: srv.store}
}

func (u *userService) Create(ctx context.Context, user *v1.User, opts metav1.CreateOptions) error {
	if err := u.store.Users().Create(ctx, user, opts); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}

		return err
	}

	return nil
}

func (u *userService) Delete(ctx context.Context, username string, opts metav1.DeleteOptions) error {
	if err := u.store.Users().Delete(ctx, username, opts); err != nil {
		return err
	}

	return nil
}

func (u *userService) Get(ctx context.Context, username string, opts metav1.GetOptions) (*v1.User, error) {
	user, err := u.store.Users().Get(ctx, username, opts)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) Update(ctx context.Context, user *v1.User, opts metav1.UpdateOptions) error {
	if err := u.store.Users().Update(ctx, user, opts); err != nil {
		return err
	}

	return nil
}

// List returns user list in the storage. This function has a good performance.
func (u *userService) List(ctx context.Context, opts metav1.ListOptions) (*v1.UserList, error) {
	users, err := u.store.Users().List(ctx, opts)
	if err != nil {
		log.L(ctx).Errorf("list users from storage failed: %s", err.Error())

		return nil, err
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	var m sync.Map

	// Improve query efficiency in parallel
	for _, user := range users.Items {
		wg.Add(1)

		go func(user *v1.User) {
			defer wg.Done()

			shadowedPass, err := auth.Shadow(user.Password)
			if err != nil {
				errChan <- err

				return
			}
			user.Password = shadowedPass

			m.Store(user.ID, user)
		}(user)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	infos := make([]*v1.User, 0, len(users.Items))
	for _, user := range users.Items {
		info, _ := m.Load(user.ID)
		infos = append(infos, info.(*v1.User))
	}

	log.L(ctx).Debugf("get %d users from backend storage.", len(infos))

	return &v1.UserList{ListMeta: users.ListMeta, Items: infos}, nil
}
