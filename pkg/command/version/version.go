package version

import (
	"crypto/md5"
	"fmt"
	"io"

	"gopkg.in/src-d/go-git.v4"
)

type Version interface {
	ID() int
	Hash() string
}

var (
	repository *git.Repository
	err        error
)

func init() {
	repository, err = git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
}

func ID() int {
	return 0
}

func Hash() string {
	ref, _ := repository.Head()
	h := md5.New()
	io.WriteString(h, ref.Name().String())

	return fmt.Sprintf("%x", h.Sum(nil))[:8]
}
