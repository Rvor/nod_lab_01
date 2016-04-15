package repos_test

import (
	"nhaoday.com/repos"
	"testing"
)

func TestPostList(t *testing.T) {
	l, e := repos.PostList()
	if e != nil {
		panic(e)
	}
	if l == nil {
		t.Error("expected NOT nil")
	}
}