package main

func get00app02Handlersfiles() []Filego {
	return []Filego{
		{homego, "/00app/02handlers/home.go"},
		{dmsgo, "/00app/02handlers/dms.go"},
	}
}

const homego = `package handlers

import (
	"net/http"

	"$PKGPATH/00app/04viewmodel"
	"github.com/noypi/router"
	"github.com/noypi/webutil"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	c := router.ContextW(w)
	var data viewmodel.HomePageData
	dms := getDMS(c)
	cfg, _ := dms.Db.GetConfig()
	data.Config = cfg
	
	data.WelcomeMessage = "To change this message, modify -> \"$PROJECTNAME/00app/02handlers/home.go\""

	renderer := webutil.GetRenderer(c)
	renderer.Render(200, data)
}

`

const dmsgo = `package handlers

import (
	"$PKGPATH/00app/03datamodel/store"
)

var getDMS = datamodel.GetStore
`
