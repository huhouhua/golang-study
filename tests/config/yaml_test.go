package config

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

const path = "../../assets/config.yaml"

func TestRead(t *testing.T) {
	yamlfile, err := ioutil.ReadFile(path)
	assert.NoError(t, err)
	var cg *config
	err = yaml.Unmarshal(yamlfile, &cg)
	assert.NoError(t, err)
	t.Logf("config.app: %#v\\n", cg.App)
	t.Logf("config.log: %#v\\n", cg.Log)
}

func TestWrite(t *testing.T) {
	yamlfile, err := ioutil.ReadFile(path)
	assert.NoError(t, err)
	var cg *config
	err = yaml.Unmarshal(yamlfile, &cg)
	assert.NoError(t, err)
	cg.App.UserName = "李四"
	m, err := yaml.Marshal(cg)
	assert.NoError(t, err)

	file, err := os.OpenFile(path, os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	assert.NoError(t, err)
	file.Write(m)
	file.Close()
}

type config struct {
	App *app `yaml:"app"`
	Log *log `yaml:"log"`
}

type app struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

type log struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}
