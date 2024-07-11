package usrgen

import (
	"fmt"

	"github.com/gosimple/slug"
)

const errMsg = "User Generator error: Exceeded length of first name."

type ug struct {
	first, last string
	n           int
}

// New creates an instance of usrgen with the first and
// last name and the language for character substitution.
func New(firstName, lastName, lang string) *ug {
	slug.CustomSub = map[string]string{
		" ": "", // spaces are originally replaced by "-" in slug
		"_": "", // remove underscores from names if any
	}

	return &ug{
		first: slug.MakeLang(firstName, lang),
		last:  slug.MakeLang(lastName, lang),
		n:     1,
	}
}

// Generate will return a username and an error.
// For each time the function is called, it will return a new
// username with an additional letter from the first name.
//
// An error will be returned if the function is called more
// times than the length of the first name.
func (u *ug) Generate() (string, error) {
	if u.n > len(u.first) {
		return "", fmt.Errorf(errMsg)
	}
	defer u.increment()

	return u.first[:u.n] + u.last, nil
}

func (u *ug) increment() {
	u.n++
}
