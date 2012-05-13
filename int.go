// int.go
//
// Integer Type

package configureit

import (
	"fmt"
	"strconv"
	"strings"
)

type IntOption struct {
	defaultvalue int
	isset        bool
	Value        int
}

func NewIntOption(defaultValue int) ConfigNode {
	opt := new(IntOption)

	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *IntOption) String() string {
	return fmt.Sprintf("%d", opt.Value)
}

func (opt *IntOption) Parse(newValue string) error {
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
