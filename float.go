// float.go
//
// Float Type

package configureit

import (
	"strings"
	"strconv"
	"os"
)

type FloatOption struct {	
	defaultvalue	float64
	isset		bool
	Value		float64
}

func NewFloatOption(defaultValue float64) ConfigNode {
	opt := new(FloatOption)

	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *FloatOption) String() string {
	return strconv.Ftoa64(opt.Value, 'g', -1)
}

func (opt *FloatOption) Parse(newValue string) os.Error {
	nativenv, err := strconv.Atof64(strings.TrimSpace(newValue))
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
