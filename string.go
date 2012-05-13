// string.go
//
// String Type

package configureit

type StringOption struct {
	defaultvalue string
	isset        bool
	Value        string
}

func NewStringOption(defaultValue string) ConfigNode {
	opt := new(StringOption)

	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *StringOption) String() string {
	return opt.Value
}

func (opt *StringOption) Parse(newValue string) error {
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
