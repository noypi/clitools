package main

func get00app03Datamodelfiles() []Filego {
	return []Filego{
		{datamodelContextiGaego, "/00app/03datamodel/internal/contexti/gae.go"},
		{datamodelContextiStorego, "/00app/03datamodel/internal/contexti/store.go"},

		{datamodelDbiConfiggo, "/00app/03datamodel/internal/dbi/config.go"},
		{datamodelDbiStorego, "/00app/03datamodel/internal/dbi/store.go"},

		{datamodelSessioniStorego, "/00app/03datamodel/internal/sessioni/store.go"},
		{datamodelSessioniUsergo, "/00app/03datamodel/internal/sessioni/user.go"},

		{datamodelStoreAddDataModelStoresgo, "/00app/03datamodel/store/h_AddDataModelStores.go"},
		{datamodelStoreStorego, "/00app/03datamodel/store/store.go"},
	}
}

const datamodelContextiGaego = `package contexti

import (
	"net/http"

	"github.com/noypi/gae"
	logih "github.com/noypi/gae/logi/handler"
	webih "github.com/noypi/gae/webi/handler"
)

func (this *ContextStore) GetUrlFetch() *http.Client {
	return webih.GetUrlFetchClient(this.c)
}

func (this *ContextStore) GetGaeLogger() gae.LogInt {
	return logih.GetGAELogi(this.c)
}

`

const datamodelContextiStorego = `package contexti

import (
	"context"

	"github.com/noypi/logfn"
	"github.com/noypi/router"
	"github.com/noypi/webutil"
)

type ContextStore struct {
	c     context.Context
	store webutil.Store

	ERR, INFO, WARN, DBG logfn.LogFunc
}

func NewContextStore(c context.Context) *ContextStore {
	ERR := router.GetErrLog(c)
	INFO := router.GetInfoLog(c)
	WARN := router.GetWarnLog(c)
	DBG := router.GetDebugLog(c)

	return &ContextStore{
		c:     c,
		store: webutil.ToStore(c),
		ERR:   ERR,
		INFO:  INFO,
		WARN:  WARN,
		DBG:   DBG,
	}
}
`

const datamodelDbiConfiggo = `package dbi

import (
	. "$PKGPATH/00app/04viewmodel"
)

func (db *DBStore) GetConfig() (cfg Config, has bool) {
	err := db.Store.GetObject(cfg.ID(), &cfg)
	has = (err == nil)
	return
}

func (db *DBStore) PutConfig(cfg *Config) {
	db.Store.PutObject(cfg.ID(), cfg)
	return
}

`

const datamodelDbiStorego = `package dbi

import (
	"context"

	"github.com/noypi/gae"
	dbih "github.com/noypi/gae/dbi/handler"
	"$PKGPATH/00app/03datamodel/internal/contexti"
	"$PKGPATH/00app/03datamodel/internal/sessioni"

	"github.com/noypi/logfn"
	"github.com/noypi/router"
)

type DBStore struct {
	c     context.Context
	Store gae.DbExInt
	Ctx   *contexti.ContextStore
	Sess  *sessioni.SessionStore

	ERR, INFO, WARN, DBG logfn.LogFunc
}

func NewDBStore(c context.Context, ctxstore *contexti.ContextStore, sessstore *sessioni.SessionStore) *DBStore {
	ERR := router.GetErrLog(c)
	INFO := router.GetInfoLog(c)
	WARN := router.GetWarnLog(c)
	DBG := router.GetDebugLog(c)

	return &DBStore{
		Ctx:   ctxstore,
		Sess:  sessstore,
		Store: dbih.GetGAEDbi(c),
		c:     c,
		ERR:   ERR,
		INFO:  INFO,
		WARN:  WARN,
		DBG:   DBG,
	}
}

`

const datamodelSessioniStorego = `package sessioni

import (
	"context"

	"github.com/gorilla/sessions"
	"github.com/noypi/logfn"
	"github.com/noypi/router"
	"github.com/noypi/webutil"
)

type SessionStore struct {
	c     context.Context
	wsess *sessions.Session

	ERR, INFO, WARN, DBG logfn.LogFunc
}

func NewSessionStore(c context.Context) *SessionStore {
	ERR := router.GetErrLog(c)
	INFO := router.GetInfoLog(c)
	WARN := router.GetWarnLog(c)
	DBG := router.GetDebugLog(c)

	return &SessionStore{
		c:     c,
		wsess: webutil.GetSession(c),
		ERR:   ERR,
		INFO:  INFO,
		WARN:  WARN,
		DBG:   DBG,
	}
}

func (this *SessionStore) Save() error {
	return this.wsess.Save(router.GetRequest(this.c), router.GetWriter(this.c))
}

`

const datamodelSessioniUsergo = `package sessioni

import (
	"encoding/gob"
)

type _userinfokey int

const (
	UserIDKey _userinfokey = iota
)

func init() {
	gob.Register(UserIDKey)
}

func (this *SessionStore) GetUserID() string {
	s, exists := this.wsess.Values[UserIDKey].(string)
	if !exists {
		this.ERR.PrintStackTrace(10)
		this.ERR("no userID found.")
	}
	return s
}

func (this *SessionStore) PutUserID(userID string) {
	this.wsess.Values[UserIDKey] = userID
}

`

const datamodelStoreAddDataModelStoresgo = `package datamodel

import (
	"net/http"

	"github.com/noypi/router"
)

func AddDataModelStore(w http.ResponseWriter, r *http.Request) {
	c := router.ContextW(w)
	AddStore(c)
}

`

const datamodelStoreStorego = `package datamodel

import (
	"context"

	"$PKGPATH/00app/03datamodel/internal/contexti"
	"$PKGPATH/00app/03datamodel/internal/dbi"
	"$PKGPATH/00app/03datamodel/internal/sessioni"
	"github.com/noypi/logfn"
	"github.com/noypi/router"
	"github.com/noypi/webutil"
)

type Store struct {
	Db   *dbi.DBStore
	Ctx  *contexti.ContextStore
	Sess *sessioni.SessionStore
	c    context.Context

	ERR, INFO, WARN, DBG logfn.LogFunc
}

type _datamodelkey int

const (
	DataModelStore _datamodelkey = iota
)

func newStore(c context.Context) *Store {
	Ctx := contexti.NewContextStore(c)
	Sess := sessioni.NewSessionStore(c)

	return &Store{
		Db:   dbi.NewDBStore(c, Ctx, Sess),
		Ctx:  Ctx,
		Sess: Sess,
		ERR:  router.GetErrLog(c),
		INFO: router.GetInfoLog(c),
		WARN: router.GetWarnLog(c),
		DBG:  router.GetDebugLog(c),
	}

}

func AddStore(ctx context.Context) {
	c := webutil.ToStore(ctx)
	c.Set(DataModelStore, newStore(ctx))
}

func GetStore(ctx context.Context) *Store {
	c := webutil.ToStore(ctx)
	o, exists := c.Get(DataModelStore)
	if exists {
		return o.(*Store)
	}

	return nil
}

`
