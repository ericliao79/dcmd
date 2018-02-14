package dcmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	Name        = "Dcmd"
	Usage       = "fast do docker-compose something"
	Version     = "0.0.1"
	ConfigName  = "/config.json"
	CheckSymbol = "\u2714 "
	CrossSymbol = "\u2716 "
	EditSymbol  = "\u2710 "
	up          = "up"
	down        = "stop"
	Detached    = "-d"
	yamlName    = "/docker-compose.yml"
)

var (
	// StorePath is the default dcmd config
	StorePath = filepath.Join(os.Getenv("HOME"), ".dcmd")
)

//check config is empty.
func IsEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

//Get Docker-composes path
func GetComposePath() Config {
	var c Config
	raw, err := ioutil.ReadFile(StorePath + ConfigName)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(raw, &c)

	return c
}

//set Docker-composes path into Config
func SetConfig(p string, con *[]Container) (int, error) {
	//fmt.Println(con[].Name)
	//return 1, nil

	f, err := os.Create(StorePath + ConfigName)
	defer f.Close()
	if err != nil {
		return 0, err
	}

	config := Config{
		PATH:       p,
		CONTAINERS: *con,
	}

	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return 0, err
	}
	s, err := f.Write(b)
	if err != nil {
		return 0, err
	}

	return s, nil
}

//load all Composes
func LoadComposes(s ...string) map[string]string {
	keys := map[string]string{}
	var path string
	if len(s) > 0 {
		path = s[0]
	} else {
		path = GetComposePath().PATH
	}

	list, error := ioutil.ReadDir(path)
	if error != nil {
		fmt.Println(error.Error())
	}

	for _, f := range list {
		if f.IsDir() {
			keys[f.Name()] = f.Name()
		}
	}

	return keys
}

//up docker containers
func Start(s string) {
	composes := GetComposePath()
	l := composes.GetService(s).Service
	if len(l) > 0 {
		color.Green("Start services from config.")
	}
	args := append([]string{up, Detached}, l...)
	runCmd("docker-compose", composes.PATH+"/"+s, down)
	runCmd("docker-compose", composes.PATH+"/"+s, args...)
}

//stop all containers
func Stop() {
	//docker stop $(docker ps -q); cd -;
	cmd := exec.Command("docker", "ps", "-q")
	output, _ := cmd.Output()
	var temp string
	for _, output := range output {
		out := string(output)
		if out == "\n" {
			cmd := exec.Command("docker", down, temp)
			o, _ := cmd.Output()
			color.Green("%s %s stopping", CheckSymbol, strings.Replace(string(o), "\n", "", -1))
			temp = ""
			continue
		}
		temp += out
	}
}

func runCmd(name string, path string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Dir = path
	cmd.Run()
	return cmd
}

//
func LoadComposeYaml(path string, proj string) map[string]string {
	var y DockerYAML
	services := map[string]string{}

	raw, _ := ioutil.ReadFile(path + "/" + proj + yamlName)
	yaml.Unmarshal(raw, &y)
	for k := range y.Services {
		services[k] = k
	}

	return services
}
