// configureit.go
//
// configureit:  A library for parsing configuration files.
//
// Copyright (C) 2011, Chris Collins <chris.collins@anchor.net.au>

package configureit

import (
	"strings"
	"io"
	"os"
	"bufio"
	"unicode"
	"fmt"
)

// ParseErrors are returned by ConfigNodes when they encounter a
// problem with their input, or by the config reader when it
// has problems.
type ParseError struct {
	LineNumber	int
	InnerError	os.Error
}

var MissingEqualsOperator = os.NewError("No equals (=) sign on non-blank line")

func (err *ParseError) String() string {
	return fmt.Sprintf("%s (at line %d)", err.InnerError, err.LineNumber)
}

func NewParseError(lineNumber int, inner os.Error) os.Error {
	err := new(ParseError)
	
	err.LineNumber = lineNumber
	err.InnerError = inner

	return err
}

// Unknown option errors are thrown when the key name (left-hand side
// of a config item) is unknown.
type UnknownOptionError struct {
	LineNumber	int
	Key		string
}

func (err *UnknownOptionError) String() string {
	return fmt.Sprintf("Unknown Key \"%s\" at line %d", err.Key, err.LineNumber)
}

func NewUnknownOptionError(lineNumber int, key string) os.Error {
	err := new(UnknownOptionError)
	
	err.LineNumber = lineNumber
	err.Key = key

	return err
}

// A configuration is made up of many ConfigNodes.
//
// ConfigNodes are typed, and are handled by their own node
// implementations.
type ConfigNode interface {
	// attribute name.  Must be all lower case.
	Name()		string

	// returns the value formatted as a string.  Must be parsable with
	// Parse() to produce the same value.
	String()	string

	// parses the string and set the value.  Clears default.  
	// Returns errors if the results can't be read.
	Parse(newValue string)	os.Error

	// is the current value the default?
	IsDefault()	bool

	// reset to the default value.
	Reset()
}

// This represents a configuration.
type Config struct {
	structure	[]ConfigNode
}

// Create a new configuration object.
func New() (config *Config) {
	config = new(Config)

	return config
}

// Add the specified ConfigNode to the configuration
func (config *Config) Add(newNode ConfigNode) {
	config.structure = append(config.structure, newNode)
}

// Reset the entire configuration.
func (config *Config) Reset() {
	for _, v := range config.structure {
		v.Reset()
	}
}

// Get the named node
func (config *Config) Get(keyName string) ConfigNode {
	keyName = strings.ToLower(keyName)
	for _, citem := range config.structure {
		if citem.Name() == keyName {
			return citem
		}
	}
	return nil
}

// Save spits out the configuration to the nominated writer.
// if emitDefaults is true, values that are set to the default
// will be omitted, otherwise they will be omitted.
//
// When in doubt, you probably want emitDefaults == false.
func (config *Config) Write(out io.Writer, emitDefaults bool) {
	for _,v := range config.structure {
		if !v.IsDefault() || emitDefaults {
			// non-default value, must write!
			line := fmt.Sprintf("%s=%s\n", v.Name(), v)
			io.WriteString(out, line)
		}
	}
}

// Read the configuration from the specified reader.
//
// Special behaviour to note:
//
//   Lines beginning with '#' or ';' are treated as comments.  They are 
//   not comments anywhere else on the line unless the config node parser
//   handles it itself.
//
//   Whitespace surrounding the name of a configuration key will be ignored.
//
//   Configuration key names will be tested case insensitively.
//
// firstLineNumber specifies the actual first line number in the file (for
// partial file reads, or resume from error)
func (config *Config) Read(in io.Reader, firstLineNumber int) os.Error {
	bufin := bufio.NewReader(in)

	// position the line number before the 'first' line.
	var lineNumber int = (firstLineNumber-1)
	
	for {
		var bline []byte = nil
		var isPrefix bool
		var err os.Error

		// get a whole line of input, and handle buffer exhausation
		// correctly.
		bline, isPrefix, err = bufin.ReadLine()
		if err != nil {
			if err == os.EOF {
				break
			} else {
				return err
			}
		}
		for isPrefix {
			var contline []byte
			
			contline, isPrefix, err = bufin.ReadLine()
			if err != nil {
				return err
			}
			bline = append(bline, contline...)
		}
		// advance the line number
		lineNumber++

		// back convert the bytearray to a native string.
		line := string(bline)
		
		// now, start doing unspreakable things to it! (bwahaha)

		// remove left space
		line = strings.TrimLeftFunc(line, unicode.IsSpace)

		// if empty, skip.
		if line == "" {
			continue
		}

		// if a comment, skip.
		if line[0] == '#' || line[0] == ';' {
			continue
		}

		// since it is neither, look for an equals sign.
		epos := strings.Index(line, "=")
		if epos < 0 {
			// no =.  Throw a parse error.
			return NewParseError(lineNumber, MissingEqualsOperator)
		}

		// take the two slices.
		keyname := line[0:epos]
		rawvalue := line[epos+1:len(line)]

		// clean up the keyname
		keyname = strings.TrimRightFunc(keyname,unicode.IsSpace)
		keyname = strings.ToLower(keyname)

		// find the correct key in the config.
		cnode := config.Get(keyname)
		if nil == cnode {
			return NewUnknownOptionError(lineNumber, keyname)
		} else {
			err := cnode.Parse(rawvalue)
			if (err != nil) {
				return NewParseError(lineNumber, err)
			}
		}
		// and we're done!
	}
	return nil
}