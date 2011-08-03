
package configureit

import (
	"testing"
	"os"
)

func makeSimpleConfig() *Config {
	testConfig := New()
	testConfig.Add("key_a", NewStringOption("default 1"))
	testConfig.Add("key_b", NewIntOption(2))
	testConfig.Add("user_test", NewUserOption(""))
	testConfig.Add("user test 2", NewUserOption(""))
	return testConfig
}

func TestConfig(t *testing.T) {
	testConfig := makeSimpleConfig()

	tv := testConfig.Get("key_a")
	if nil == tv {
		t.Errorf("Couldn't find key_a in testConfig")
	} else {
		if !tv.IsDefault() {
			t.Errorf("key_a reported non-default without changes")
		}
		sopt, ok := tv.(*StringOption)
		if !ok {
			t.Errorf("Failed return assertion for key_a back to StringOption")
		}
		if sopt.Value != "default 1" {
			t.Errorf("key_a Value doesn't match initial configured value.")
		}
	}

	tv = testConfig.Get("key_b")
	if nil == tv {
		t.Errorf("Couldn't find key_b in testConfig")
	} else {
		if !tv.IsDefault() {
			t.Errorf("key_b reported non-default without changes")
		}
		iopt, ok := tv.(*IntOption)
		if !ok {
			t.Errorf("Failed return assertion for key_b back to IntOption")
		}
		if iopt.Value != 2 {
			t.Errorf("key_b Value doesn't match initial configured value.")
		}
	}

	tv = testConfig.Get("user_test")
	if nil == tv {
		t.Errorf("Couldn't find user_test in testConfig")
	} else {
		if !tv.IsDefault() {
			t.Errorf("user_test reported non-default without changes")
		}
		uopt, ok := tv.(*UserOption)
		if !ok {
			t.Errorf("Failed return assertion for user_test back to UserOption")
		}
		_, err := uopt.User()
		if err != EmptyUserSet {
			t.Errorf("user_test didn't claim it set empty.")
		}
	}

	tv = testConfig.Get("user test 2")
	if nil == tv {
		t.Errorf("Couldn't find \"user test 2\" in testConfig")
	} else {
		if !tv.IsDefault() {
			t.Errorf("user test 2 reported non-default without changes")
		}
		uopt, ok := tv.(*UserOption)
		if !ok {
			t.Errorf("Failed return assertion for user test 2 back to UserOption")
		}
		_, err := uopt.User()
		if err != EmptyUserSet {
			t.Errorf("user test 2 didn't claim it set empty.")
		}
	}

	tv = testConfig.Get("key_c")
	if nil != tv {
		t.Errorf("Found non-existant key_c in testConfig")
	}
}

func TestFileRead(t *testing.T) {
	testConfig := makeSimpleConfig()
	fh, err := os.Open("sample.conf")
	if err != nil {
		t.Fatalf("Failed to open sample.conf: %s", err)
	}
	err = testConfig.Read(fh, 1)
	if err != nil {
		t.Fatalf("Got error reading config: %s", err)
	}
	fh.Close()

	tv := testConfig.Get("key_a")
	if nil == tv {
		t.Errorf("Couldn't find key_a in testConfig")
	} else {
		if tv.IsDefault() {
			t.Errorf("key_a reported default despite config file")
		}
		sopt, ok := tv.(*StringOption)
		if !ok {
			t.Errorf("Failed return assertion for key_a back to StringOption")
		}
		if sopt.Value != "Alternate Value" {
			t.Errorf("key_a Value doesn't match expected value.")
		}
	}

	tv = testConfig.Get("key_b")
	if nil == tv {
		t.Errorf("Couldn't find key_b in testConfig")
	} else {
		if tv.IsDefault() {
			t.Errorf("key_b reported default despite config file")
		}
		iopt, ok := tv.(*IntOption)
		if !ok {
			t.Errorf("Failed return assertion for key_b back to IntOption")
		}
		if iopt.Value != 27 {
			t.Errorf("key_b Value doesn't match expected value.")
		}
	}

	tv = testConfig.Get("user_test")
	if nil == tv {
		t.Errorf("Couldn't find user_test in testConfig")
	} else {
		if tv.IsDefault() {
			t.Errorf("user_test reported default despite changes")
		}
		uopt, ok := tv.(*UserOption)
		if !ok {
			t.Errorf("Failed return assertion for user_test back to UserOption")
		}
		uinfo, err := uopt.User()
		if err != nil {
			t.Errorf("Error whilst looking up UID: %s", err)
		}
		if uinfo.Uid != 0 {
			t.Errorf("user_test Value doesn't match expected value.")
		}
	}

	tv = testConfig.Get("user test 2")
	if nil == tv {
		t.Errorf("Couldn't find \"user test 2\" in testConfig")
	} else {
		if tv.IsDefault() {
			t.Errorf("user test 2 reported default despite changes")
		}
		uopt, ok := tv.(*UserOption)
		if !ok {
			t.Errorf("Failed return assertion for user test 2 back to UserOption")
		}
		uinfo, err := uopt.User()
		if err != nil {
			t.Errorf("Error whilst looking up UID: %s", err)
		}
		if uinfo.Uid != 1 {
			t.Errorf("user test 2 Value doesn't match expected value.")
		}
		if uinfo.Username != "daemon" {
			t.Errorf("user test 2 name lookup didn't match expected value (ignore if not debian/ubuntu/linux?).")
		}
	}
}
	
