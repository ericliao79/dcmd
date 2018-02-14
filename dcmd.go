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
	name        = "Dcmd"
	usage       = "fast do docker-compose something"
	version     = "0.0.1"
	configName  = "/config.json"
	checkSymbol = "\u2714 "
	crossSymbol = "\u2716 "
	editSymbol  = "\u2710 "
	up          = "up"
	down        = "stop"
	detached    = "-d"
	yamlName    = "/docker-compose.yml"
)

var (
	MyAppConfig = AppConfig{
		Name:        name,
		Usage:       usage,
		Version:     version,
		ConfigName:  configName,
		CheckSymbol: checkSymbol,
		CrossSymbol: crossSymbol,
		EditSymbol:  editSymbol,
		up:          up,
		down:        down,
		Detached:    detached,
		yamlName:    yamlName,
		StorePath:   filepath.Join(os.Getenv("HOME"), ".dcmd"),
	}
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
func GetComposePath() (Config, error) {
	var c Config
	raw, err := ioutil.ReadFile(MyAppConfig.StorePath + MyAppConfig.ConfigName)
	if err != nil {
		return c, err
	}
	json.Unmarshal(raw, &c)
	return c, nil
}

//set Docker-composes path into Config
func SetConfig(p string, con *[]Container) (int, error) {
	f, err := os.Create(MyAppConfig.StorePath + MyAppConfig.ConfigName)
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
		g, _ := GetComposePath()
		path = g.PATH
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
	composes, _ := GetComposePath()
	l := composes.GetService(s).Service
	if len(l) > 0 {
		color.Green("Start services from config.")
	}
	args := append([]string{up, MyAppConfig.Detached}, l...)
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
			color.Green("%s %s stopping", MyAppConfig.CheckSymbol, strings.Replace(string(o), "\n", "", -1))
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
