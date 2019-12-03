package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type JobDefinition struct {
	URL     string
	Product string
	Version string
}

type ProductConfig struct {
	Name     string
	Pipeline string
}

type LinkConfig struct {
	URL   string `toml:"url"`
	Title string `toml:"title"`
}

type Config struct {
	UseCache    bool
	CithURL     string          `toml:"cith_url"`
	Products    []ProductConfig `toml:"products"`
	Links       []LinkConfig    `toml:"links"`
	KickoffJobs []JobDefinition `toml:"kickoff_jobs"`
	OrderedJobs []JobDefinition `toml:"ordered_jobs"`
}

func (c *Config) SetUseCache(val bool) {
	c.UseCache = false
}

func GetConfig() Config {

	tomlData, err := ioutil.ReadFile("conf/config.toml")

	if err != nil {
		panic(err)
	}

	var conf Config

	if _, err := toml.Decode(string(tomlData), &conf); err != nil {
		fmt.Println(err)
	}

	conf.UseCache = true

	return conf
}
