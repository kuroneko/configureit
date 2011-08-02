// int.go
//
// Integer Type

package configureit

import (
	"strings"
	"strconv"
	"os"
	"fmt"
)

type IntOption struct {	
	keyname		string
	defaultvalue	int
	isset		bool
	Value		int
}

func NewIntOption(keyName string, defaultValue int) ConfigNode {
	opt := new(IntOption)

	opt.keyname = strings.ToLower(keyName)
	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *IntOption) Name() string {
	return opt.keyname
}

func (opt *IntOption) String() string {
	return fmt.Sprintf("%d", opt.Value)
}

func (opt *IntOption) Parse(newValue string) os.Error {
	nativenv, err := strconv.Atoi(strings.TrimSpace(newValue))
	if err != nil {
		return err
	}
	opt.Value = nativenv
	opt.isset = true

	return nil
}

func (opt *IntOption) IsDefault() bool {
	return !opt.isset
}

func (opt *IntOption) Reset() {
	opt.Value = opt.defaultvalue
	opt.isset = false
}
