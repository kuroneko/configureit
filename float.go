// float.go
//
// Float Type

package configureit

import (
	"strconv"
	"strings"
)

type FloatOption struct {
	defaultvalue float64
	isset        bool
	Value        float64
}

func NewFloatOption(defaultValue float64) ConfigNode {
	opt := new(FloatOption)

	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *FloatOption) String() string {
	return strconv.FormatFloat(opt.Value, 'g', -1, 64)
}

func (opt *FloatOption) Parse(newValue string) error {
	nativenv, err := strconv.ParseFloat(strings.TrimSpace(newValue), 64)
	if err != nil {
		return err
	}
	opt.Value = nativenv
	opt.isset = true

	return nil
}

func (opt *FloatOption) IsDefault() bool {
	return !opt.isset
}

func (opt *FloatOption) Reset() {
	opt.Value = opt.defaultvalue
	opt.isset = false
}
