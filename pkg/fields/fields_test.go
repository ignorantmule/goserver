// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package fields

import (
	"testing"
)

func matches(t *testing.T, ls Set, want string) {
	if ls.String() != want {
		t.Errorf("Expected '%s', but got '%s'", want, ls.String())
	}
}

func TestSetString(t *testing.T) {
	matches(t, Set{"x": "y"}, "x=y")
	matches(t, Set{"foo": "bar"}, "foo=bar")
	matches(t, Set{"foo": "bar", "baz": "qup"}, "baz=qup,foo=bar")
}

func TestFieldHas(t *testing.T) {
	fieldHasTests := []struct {
		Ls  Fields
		Key string
		Has bool
	}{
		{Set{"x": "y"}, "x", true},
		{Set{"x": ""}, "x", true},
		{Set{"x": "y"}, "foo", false},
	}
	for _, lh := range fieldHasTests {
		if has := lh.Ls.Has(lh.Key); has != lh.Has {
			t.Errorf("%#v.Has(%#v) => %v, expected %v", lh.Ls, lh.Key, has, lh.Has)
		}
	}
}

func TestFieldGet(t *testing.T) {
	ls := Set{"x": "y"}
	if ls.Get("x") != "y" {
		t.Errorf("Set.Get is broken")
	}
}
