package usrgen

import (
	"errors"

	"github.com/gosimple/slug"
)

const (
	errMsg = "usrgen: Exceeded length of first name"

	// We use 1 for the first letter and not 0 because we're slicing the
	// first name string. Slicing takes up to but not including the number passed.
	firstLetter = 1
)

type ug struct {
	first, last, userName string

	// n represents the number of letters taken
	// from the first name to produce the user name
	n int
}

// New creates an instance of usrgen with the first and
// last name and the language for character substitution.
func New(firstName, lastName, lang string) *ug { //nolint: revive
	slug.CustomSub = map[string]string{
		" ": "", // spaces are originally replaced by "-" in slug
		"_": "", // remove underscores from names if any
	}

	return &ug{
		first: slug.MakeLang(firstName, lang),
		last:  slug.MakeLang(lastName, lang),
		n:     firstLetter,
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
		return "", errors.New(errMsg)
	}
	defer func() { u.n++ }()

	u.userName = u.first[:u.n] + u.last

	return u.userName, nil
}

// Stringer method returns the latest generated user name.
func (u *ug) String() string {
	return u.userName
}
