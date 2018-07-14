package main

import (
	"strings"
)

type Filego struct {
	Content string
	Path    string
}

func CreateFiles() {
	createfile(g_basedir+"/00app/00gae/app.yaml", appyaml)
	createfile(g_basedir+"/00app/00gae/run-dev-server.sh", runDevServersh)
	createfile(g_basedir+"/00app/00gae/run-dev-server.cmd", runDevServersh)

	var files []Filego
	for _, fs := range [][]Filego{
		get00app00Gaefiles(),
		get00app01Muxfiles(),
		get00app02Handlersfiles(),
		get00app03Datamodelfiles(),
		get00app04Viewmodelfiles(),
		get00Commonfiles(),
		get00Typesfiles(),
	} {
		files = append(files, fs...)
	}

	for _, f := range files {
		content := strings.Replace(f.Content, "$PROJECTNAME", g_project, -1)
		content = strings.Replace(content, "$PKGPATH", g_pkg, -1)
		createfile(g_basedir+f.Path, content)
	}

}
