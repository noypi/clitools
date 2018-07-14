package main

func get00app00Gaefiles() []Filego {
	return []Filego{
		{maingo, "/00app/00gae/main.go"},
		{templatesHomehtml, "/00app/00gae/view/templates/home.tmpl.html"},
	}
}

const appyaml = `runtime: go
api_version: go1
instance_class: F2
automatic_scaling:
  min_idle_instances: 0
  max_idle_instances: 1
  max_concurrent_requests: 80
  min_pending_latency: 15s
  max_pending_latency: 15s

handlers:
- url: /.*
  script: _go_app

- url: /admin
  script: _go_app
  login: admin

`

const maingo = `package $PROJECTNAME

import (
	"net/http"

	"github.com/noypi/webutil"

	gaeh "github.com/noypi/gae/handlers"
	"$PKGPATH/00app/01mux"
	"$PKGPATH/00app/03datamodel/store"
)

var Σh = webutil.HttpSequence

func init() {

	mux := $PROJECTNAME.NewMux()

	mux.Delims("{{", "}}")

	mux.Use(webutil.EnsureHTTPS,
		gaeh.GetEssentialHandlers(),
		//mux.Sstore.AddSessionHandler("$PROJECTNAME"),
		gaeh.AddServices(gaeh.ServicesOpts{
			Namespace: "",
			AppID:     "$PROJECTNAME",
			//JWTPath:      "./<some file>.json",
			UseDatastore: true,
			UseUrlFetch:  true,
		}),
		datamodel.AddDataModelStore,
		webutil.UseRenderer("./view/templates/*.tmpl.html"),
	)

	mux.Static("/pub/", "./view/pub")

	mux.HandleHome()

	http.Handle("/", mux)
}

`

const templatesHomehtml = `<!DOCTYPE html>
<html>
{{with .oHomePageData}}

<head>
<title>Home</title>
<style>

body {
	font-size:14pt;
}
</style>
</head>

<body>

<h1>It works!</h1>
<br/>
<br/>To update this Home page, modify  -> "myproject/00app/00gae/view/templates/home.tmpl.html"
<br/>
<br/>Dynamic Message: {{ .WelcomeMessage }}

</body>
{{end}}
</html>
`

const runDevServersh = `dev_appserver.py --support_datastore_emulator=True . `
