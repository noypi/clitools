package main

func get00app01Muxfiles() []Filego {
	return []Filego{
		{muxgo, "/00app/01mux/mux.go"},
		{handleHomego, "/00app/01mux/handleHome.go"},
	}
}

const muxgo = `package $PROJECTNAME

import (
	"github.com/noypi/router"
	"github.com/noypi/webutil"
)

type Mux struct {
	*router.EngineStd
	Sstore *webutil.SessionStore
}

func NewMux() *Mux {
	router.DisableDebugging()

	o := &Mux{
		EngineStd: router.New(),
	}

	//var keys [][]byte
	//if bb, err := ioutil.ReadFile("./keys.bin"); nil == err {
	//	keys, _ = webutil.UnmarshalKeys(bb)
	//}
	//o.Sstore = webutil.NewCookieSession(nil, keys...)
	//o.Sstore.Path("/")

	return o
}

`

const handleHomego = `package $PROJECTNAME

import (
	. "$PKGPATH/00app/02handlers"
)

func (mux *Mux) HandleHome() {
	mux.GET("/", HandleHome)
}
`
