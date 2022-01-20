package models

import (
	"testing"
)

func TestNull(t *testing.T) {
	tasks := List()
	if tasks == nil {
		t.Log("it is null")
		return
	}
	t.Log("it is not null")
}

func f() (arr []string) {
	return
}

func TestNull1(t *testing.T) {
	null := f()
	if null==nil {
		t.Fatal("it is null")
	}
}
