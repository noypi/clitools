package main

func get00Typesfiles() []Filego {
	return []Filego{
		{typesLoggergo, "/00types/logger.go"},
		{typesStorego, "/00types/store.go"},
	}

}

const typesLoggergo = `package types

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

const typesStorego = `package types

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io/ioutil"

	"github.com/noypi/kv"
	"github.com/pkg/errors"
)

type Store interface {
	Put(k string, bb []byte) (err error)
	Get(k string) (bb []byte, err error)
	Delete(k string) error
}

type StoreKv struct {
	Store kv.KVStore
}

type StoreUser struct {
	Store
}

type HasID interface {
	ID() string
}

type _gob struct {
	V interface{}
}

func init() {
	gob.Register(_gob{})
}

func SerializeRaw(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	o := &_gob{v}
	err := gob.NewEncoder(buf).Encode(o)

	var bufgz bytes.Buffer
	w := gzip.NewWriter(&bufgz)
	if _, err = w.Write(buf.Bytes()); nil != err {
		return nil, err
	}
	w.Flush()
	w.Close()

	return bufgz.Bytes(), err
}

func DeserializeRaw(bb []byte) (v interface{}, err error) {
	bufgz := bytes.NewBuffer(bb)
	r, err := gzip.NewReader(bufgz)
	if nil != err {
		err = errors.WithMessage(err, "gz reader")
		return nil, err
	}
	defer r.Close()

	bb, err = ioutil.ReadAll(r)
	if nil != err {
		err = errors.WithMessage(err, "io readall")
		return
	}

	o := new(_gob)
	buf := bytes.NewBuffer(bb)
	if err := gob.NewDecoder(buf).Decode(o); nil != err {
		err = errors.WithMessage(err, "decoder")
		return nil, err
	}

	return o.V, nil
}

func (o *StoreUser) Put(v HasID) (err error) {
	bb, err := SerializeRaw(v)
	if nil != err {
		err = errors.WithStack(err)
		return
	}

	return o.Store.Put(v.ID(), bb)
}

func (o *StoreUser) PutID(id string, v interface{}) (err error) {
	bb, err := SerializeRaw(v)
	if nil != err {
		err = errors.WithStack(err)
		return
	}

	return o.Store.Put(id, bb)
}

func (store *StoreKv) Put(k string, bb []byte) (err error) {
	w, err := store.Store.Writer()
	if nil != err {
		return
	}
	defer w.Close()
	batch := w.NewBatch()
	defer batch.Close()
	batch.Set([]byte(k), bb)

	return w.ExecuteBatch(batch)
}

func (store *StoreKv) Get(k string) (bb []byte, err error) {
	r, err := store.Store.Reader()
	if nil != err {
		return
	}
	defer r.Close()

	return r.Get([]byte(k))
}

func (store *StoreKv) Delete(k string) error {
	w, err := store.Store.Writer()
	if nil != err {
		return err
	}
	defer w.Close()
	batch := w.NewBatch()
	defer batch.Close()

	batch.Delete([]byte(k))
	return w.ExecuteBatch(batch)

}

`
