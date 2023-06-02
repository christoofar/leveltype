package config

import (
	"os"

	"github.com/kirsle/configdir"
	"github.com/leveltype/src/problemwords"
	"gopkg.in/yaml.v3"
)

var (
	dirroot    = ""
	configpath = ""
)

type Config struct {
	VocabularyLevel int
}

func (e *Config) SaveConfiguration() {
	dirroot = configdir.LocalConfig("leveltype")
	configpath = dirroot + "/config.yaml"

	_, err := os.Stat(dirroot)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dirroot, 0755); err != nil {
			panic(err)
		}
	}

	d, err := yaml.Marshal(e)
	if err != nil {
		println(err)
		panic("Could not marshal LevelType's config settings into a YAML block.")
	}

	// set defaults
	if e.VocabularyLevel == 0 {
		e.VocabularyLevel = 20
		d, _ = yaml.Marshal(e)
	}

	err = os.WriteFile(configpath, d, 0755)
	if err != nil {
		println("Could not save LevelType configuration: ")
		panic(err)
	}
}

func (e *Config) ReadConfiguration() {
	dirroot = configdir.LocalConfig("leveltype")
	configpath = dirroot + "/config.yaml"

	_, err := os.Stat(dirroot)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dirroot, 0755); err != nil {
			panic(err)
		}
	}

	d, err := os.ReadFile(configpath)
	if err != nil { // There is no config file, so make one.
		e.SaveConfiguration()
		return
	}

	yaml.Unmarshal(d, e)

	problemwords.ListSize = e.VocabularyLevel
}

func SaveVocabularyLevel(level int) {

	config := Config{}
	d, err := os.ReadFile(configpath)

	if err != nil {
		yaml.Unmarshal(d, config)
	} else {
		// There was no config file read, so provide a default
		config.VocabularyLevel = 20
	}

	d, err = yaml.Marshal(config)

	config.VocabularyLevel = level
	if config.VocabularyLevel == 0 {
		config.VocabularyLevel = 20
	}
	if err != nil {
		err = os.WriteFile(configpath, d, 0755)
		if err != nil {
			panic(err)
		}
	}
}
