package conf

import (
	"log"
	"testing"
)

func TestParseDotENV(t *testing.T) {
	c := make(map[string]string)
	err := Load(func(e map[string]string, cover bool) {
		log.Println(e)
		for k, v := range e {
			c[k] = v
		}
	}, false, "./testdata/b.env")
	if err != nil {
		t.Error(err)
	}
	if c["k3"] != "value1" {
		t.Errorf("k3: %s", c["k3"])
	}
	if c["K4"] != "100" {
		t.Errorf("K4: %s", c["K4"])
	}
	if c["k5"] != "200" {
		t.Errorf("k5: %s", c["k5"])
	}
	if c["k6"] != "123:" {
		t.Errorf("k6: %s", c["k6"])
	}
	if c["k0"] != "" {
		t.Errorf("k0: %s", c["k0"])
	}
	if c["k7"] != "" {
		t.Errorf("k7: %s", c["k7"])
	}
}

func TestParseJson(t *testing.T) {
	c := make(map[string]string)
	err := Load(func(e map[string]string, cover bool) {
		log.Println(e)
		for k, v := range e {
			c[k] = v
		}
	}, false, "./testdata/a.json")
	if err != nil {
		t.Error(err)
	}
	if c["k1"] != "value1" {
		t.Errorf("k1: %s", c["k1"])
	}
	if c["k2"] != "100" {
		t.Errorf("k2: %s", c["k2"])
	}
	if c["k3"] != "" {
		t.Errorf("k3: %s", c["k3"])
	}
	if c["k4"] != "[1,2,3]" {
		t.Errorf("k4: %s", c["k4"])
	}
	if c["k5"] != "{\"a\":\"b\"}" {
		t.Errorf("k5: %s", c["k5"])
	}
}

func TestParseYML(t *testing.T) {
	c := make(map[string]string)
	err := Load(func(e map[string]string, cover bool) {
		log.Println(e)
		for k, v := range e {
			c[k] = v
		}
	}, false, "./testdata/c.yaml")
	if err != nil {
		t.Error(err)
	}
	if c["k1"] != "v1" {
		t.Errorf("k1: %s", c["k1"])
	}
	if c["k2"] != "k3:v3" {
		t.Errorf("k2: %s", c["k2"])
	}
}
