
package configureit

import (
	"testing"
	"os"
)

func makeSimpleConfig() *Config {
	testConfig := New()
	testConfig.Add(NewStringOption("key_a", "default 1"))
	testConfig.Add(NewIntOption("key_b", 2))

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

}
	
