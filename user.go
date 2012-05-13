// user.go
//
// User Type

package configureit

import (
	"errors"
	"os/user"
	"strconv"
	"strings"
)

// User options represent user specifications in a config file.
//
// They can be either a literal username, or a number.
type UserOption struct {
	defaultvalue string
	isset        bool
	// literal value
	Value string
}

var EmptyUserSet = errors.New("No value set")

func NewUserOption(defaultValue string) ConfigNode {
	opt := new(UserOption)

	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *UserOption) String() string {
	return opt.Value
}

func (opt *UserOption) Parse(newValue string) error {
	nvs := strings.TrimSpace(newValue)
	if nvs != "" {
		// validate the input.
		_, err := strconv.Atoi(nvs)
		if err != nil {
			switch err.(type) {
			case *strconv.NumError:
				// not a number.  do a lookup.
				_, err = user.Lookup(nvs)
				if err != nil {
					return err
				}
			default:
				return err
			}
		}
	}
	opt.Value = newValue
	opt.isset = true

	return nil
}

func (opt *UserOption) IsDefault() bool {
	return !opt.isset
}

func (opt *UserOption) Reset() {
	opt.Value = opt.defaultvalue
	opt.isset = false
}

func (opt *UserOption) User() (userinfo *user.User, err error) {
	nvs := strings.TrimSpace(opt.Value)
	if nvs == "" {
		// special case: empty string is the current euid.
		return nil, EmptyUserSet
	}
	// attempt to map this as a number first, in case a numeric UID 
	// was provided.
	_, err = strconv.Atoi(nvs)
	if err != nil {
		switch err.(type) {
		case *strconv.NumError:
			// not a number.  do a user table lookup.
			userinfo, err = user.Lookup(nvs)
			if err != nil {
				return nil, err
			}
			return userinfo, nil
		default:
			return nil, err
		}
	}
	userinfo, err = user.LookupId(nvs)
	return userinfo, err
}
