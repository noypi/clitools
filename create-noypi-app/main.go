package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	g_basedir string
	g_project string
	g_pkg     string
	g_verbose bool
)

func init() {
	flag.StringVar(&g_basedir, "basedir", "./", "the base directory")
	flag.StringVar(&g_pkg, "pkg", "github.com/mygit/myproject", "package path")
	pbHelp := flag.Bool("help", false, "display usage")
	flag.BoolVar(&g_verbose, "v", false, "verbose")

	flag.Parse()

	if *pbHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	g_basedir = path.Clean(g_basedir)
	g_project = path.Base(g_pkg)

	V("basedir =", g_basedir)
	V("package =", g_pkg)
	V("project =", g_project)
}

func main() {

	CreateAppDirectories()
	CreateFiles()

}

func createfile(fpath, content string) {
	V("creating file=", fpath)

	err := ioutil.WriteFile(fpath, []byte(content), os.ModePerm)
	if nil != err {
		log.Fatal(err)
	}
}

func Vf(s string, as ...interface{}) {
	if g_verbose {
		log.Println(fmt.Sprintf(s, as...))
	}
}

func V(as ...interface{}) {
	if g_verbose {
		log.Println(as...)
	}
}
