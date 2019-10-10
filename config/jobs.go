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

type Config struct {
	CithURL     string          `toml:"cith_url"`
	Products    []Product       `toml:"products"`
	KickoffJobs []JobDefinition `toml:"kickoff_jobs"`
	OrderedJobs []JobDefinition `toml:"ordered_jobs"`
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

	return conf
}
