// string.go
//
// String Type

package configureit

import (
	"strings"
	"os"
)

type StringOption struct {	
	keyname		string
	defaultvalue	string
	isset		bool
	Value		string
}

func NewStringOption(keyName string, defaultValue string) ConfigNode {
	opt := new(StringOption)

	opt.keyname = strings.ToLower(keyName)
	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *StringOption) Name() string {
	return opt.keyname
}

func (opt *StringOption) String() string {
	return opt.Value
}

func (opt *StringOption) Parse(newValue string) os.Error {
	opt.Value = newValue
	opt.isset = true

	return nil
}

func (opt *StringOption) IsDefault() bool {
	return !opt.isset
}

func (opt *StringOption) Reset() {
	opt.Value = opt.defaultvalue
	opt.isset = false
}
