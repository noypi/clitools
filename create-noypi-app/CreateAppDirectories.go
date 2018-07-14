package main

import (
	"os"
)

type Folder struct {
	Name    string
	SubDirs []Folder
}

func CreateAppDirectories() {
	folders := []Folder{
		{"00app",
			[]Folder{
				{"00gae",
					[]Folder{
						{"view",
							[]Folder{
								{"pub", nil},
								{"raw-js", nil},
								{"templates", nil},
							},
						},
					},
				},
				{"01mux", nil},
				{"02handlers", nil},
				{"03datamodel",
					[]Folder{
						{"internal",
							[]Folder{
								{"contexti", nil},
								{"dbi", nil},
								{"sessioni", nil},
							},
						},
						{"store", nil},
					},
				},
				{"04viewmodel", nil},
			},
		},
		{"00common", nil},
		{"00types", nil},
		{"50logic", nil},
	}

	createDir(g_basedir, folders)

}

func createDir(basedir string, folders []Folder) {
	for _, f := range folders {
		installFolder(basedir + "/" + f.Name)
		if 0 < len(f.SubDirs) {
			createDir(basedir+"/"+f.Name, f.SubDirs)
		}
	}
}

func installFolder(dirpath string) {
	V("installing path=", dirpath)
	os.MkdirAll(dirpath, os.ModePerm)

}
