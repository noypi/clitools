package main

func get00Commonfiles() []Filego {
	return []Filego{
		{commonLoggergo, "/00common/logger.go"},
		{commonFakestorego, "/00common/fakestore.go"},
		{commonFakeloggo, "/00common/fakelog.go"},
		{commonFakeemailergo, "/00common/fakemailer.go"},
	}

}

const commonLoggergo = `package common

import (
	"log"

	"github.com/noypi/logfn"
)

type _Logger struct {
	Info, Dbg, Warn, Err logfn.LogFunc
}

func NewLogger() *_Logger {
	o := new(_Logger)
	o.Info, o.Warn, o.Dbg, o.Err = log.Printf, log.Printf, log.Printf, log.Printf
	return o
}

type Logger interface {
	INFO(fmt string, i ...interface{})
	WARN(fmt string, i ...interface{})
	DBG(fmt string, i ...interface{})
	ERR(fmt string, i ...interface{})
}

func (l _Logger) INFO(fmt string, i ...interface{}) { l.Info(fmt, i...) }
func (l _Logger) WARN(fmt string, i ...interface{}) { l.Warn(fmt, i...) }
func (l _Logger) DBG(fmt string, i ...interface{})  { l.Dbg(fmt, i...) }
func (l _Logger) ERR(fmt string, i ...interface{})  { l.Err(fmt, i...) }

func (l _Logger) StackTrace(n int) string {
	return logfn.StackTrace(n)
}

`

const commonFakestorego = `package common

import (
	"fmt"
)

var g_FakeStore = &fakeStore{m: map[string][]byte{}}

func FakeStore() *fakeStore {
	return g_FakeStore
}

type fakeStore struct {
	m map[string][]byte
}

func (l *fakeStore) Delete(k string) (err error) {
	if nil == l.m {
		l.m = map[string][]byte{}
	}

	delete(l.m, k)
	return
}
func (l *fakeStore) Put(k string, bb []byte) (err error) {
	if nil == l.m {
		l.m = map[string][]byte{}
	}
	l.m[k] = bb
	return
}

func (l *fakeStore) Get(k string) (bb []byte, err error) {
	if nil == l.m {
		l.m = map[string][]byte{}
	}

	bb, has := l.m[k]
	if !has {
		err = fmt.Errorf("FakeStore: not found")
	}
	return
}

`

const commonFakeloggo = `package common

import (
	"fmt"
	"io"
	"strings"
)

type FakeLog struct {
	W         io.Writer
	Separator string
}

func (l *FakeLog) Log(f string, s ...interface{}) {
	if 0 == len(l.Separator) {
		l.Separator = "<br/>"
	}
	if strings.HasPrefix(strings.ToLower(f), "<html>") {
		l.W.Write([]byte(fmt.Sprintf(f, s...)))
	} else {
		l.W.Write([]byte(l.Separator + fmt.Sprintf(f, s...)))
	}

}

`

const commonFakeemailergo = `package common

import (
	"io"
)

type EmailerFake struct {
	W io.Writer
}

func (l EmailerFake) Email(subject, body string) {
	l.W.Write([]byte("<br/>" + subject))
	l.W.Write([]byte("<br/>" + body))
}

func (l EmailerFake) EmailTo(ss []string, subject, body string, bHtml bool) {
	l.W.Write([]byte("<br/>" + subject))
	l.W.Write([]byte("<br/>" + body))
}

`
