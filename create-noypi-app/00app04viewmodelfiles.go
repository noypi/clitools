package main

func get00app04Viewmodelfiles() []Filego {
	return []Filego{
		{viewmodelHomepagego, "/00app/04viewmodel/homepage.go"},
		{viewmodelConfiggo, "/00app/04viewmodel/config.go"},
	}
}

const viewmodelHomepagego = `package viewmodel

type HomePageData struct {
	TPLConfig string ` + "`" + `webtpl:"name=home.tmpl.html"` + "`" + `
	Config    Config
	
	WelcomeMessage string
	
}

`

const viewmodelConfiggo = `package viewmodel

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	// :TODO: 
	// add necessary configurations here
}

func LoadConfig(path string) (o *Config, err error) {
	bb, err := ioutil.ReadFile(path)
	if nil != err {
		return
	}
	o = new(Config)
	err = json.Unmarshal(bb, o)
	return
}

func (cfg Config) ID() string {
	return "/$PROJECTNAME/$config"
}

`
