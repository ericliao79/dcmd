package dcmd

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	if _, e := IsEmpty("./"); e != nil { //try a unit test on function
		t.Error("IsEmpty error.")
	} else {
		t.Log("IsEmpty pass.")
	}
}

func TestSetConfig(t *testing.T) {
	MyAppConfig.StorePath = "./"

	path := "./"
	var con *[]Container
	c := []Container{
		{
			Name:    "Eric",
			Service: []string{"workspace", "php-fpm", "mysql", "nginx", "redis"},
		},
		{
			Name:    "Mai",
			Service: []string{"workspace", "php-fpm", "mysql", "nginx", "redis"},
		},
		{
			Name:    "Test",
			Service: []string{"workspace"},
		},
	}

	con = &c

	if _, err := SetConfig(path, con); err != nil {
		t.Error("SetConfig error.")
	} else {
		t.Log("SetConfig pass.")
	}
}

func TestGetComposePath(t *testing.T) {
	cp, err := GetComposePath()
	if err != nil {
		t.Error("GetComposePath error.")
	}
	if cp.PATH != "./" {
		t.Error(cp.PATH)
	}
}

func TestLoadComposes(t *testing.T) {
	// TODO:: LoadComposes test
}

func TestLoadComposeYaml(t *testing.T) {
	// TODO:: LoadComposeYaml test
}
