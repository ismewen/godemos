package main

import (
	"errors"
	"fmt"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return fmt.Sprintf("Op: %s, Path: %s, Err: %s", e.Op, e.Path, e.Err.Error())
}

func test1() error {
	return &PathError{Op: "ls", Path: "/", Err: errors.New("ls not found")}
}

type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

func main() {
	err := test1()
	switch v := err.(type) {
	case nil:
		return
	case *PathError:
		fmt.Printf("Op: %s", v.Op)
	}

	if IsTemporary(err) {
		fmt.Printf("temporary error, try later")
	}
}
