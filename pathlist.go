// string.go
//
// String Type

package configureit

import (
	"os"
	"strings"
)

var PathListSeparator = os.PathListSeparator

func init() {
	if PathListSeparator == 0 {
		// Bloody Plan9.
		PathListSeparator = '!'
	}
}

// PathList is a list of paths.  Paths are assuemd to be separated by
// os.PathListSeparator ('!' if undefined.)
//
// White spaces are valid within terms, but leading and trailing whitespace 
// are discarded from the whole input, not from terms!
type PathListOption struct {	
	defaultvalue	[]string
	isset		bool
	Values		[]string
}

func NewPathListOption(defaultValue []string) ConfigNode {
	opt := new(PathListOption)

	opt.defaultvalue = defaultValue
	opt.Reset()

	return opt
}

func (opt *PathListOption) String() string {
	return strings.Join(opt.Values, string(PathListSeparator))
}

func (opt *PathListOption) Parse(newValue string) os.Error {
	newValue = strings.TrimSpace(newValue)

	opt.Values = strings.Split(newValue, string(PathListSeparator))
	opt.isset = true

	return nil
}

func (opt *PathListOption) IsDefault() bool {
	return !opt.isset
}

func (opt *PathListOption) Reset() {
	opt.Values = opt.defaultvalue
	opt.isset = false
}
