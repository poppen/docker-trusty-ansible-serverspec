package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Version struct {
	OS         string
	AnsibleVer string
}

func main() {
	os_names := []string{"trusty", "xenial"}
	ansible_versions := []string{"1.9", "2.0", "2.1", "2.2"}

	tpl := template.Must(template.ParseFiles("Dockerfile.tpl"))

	var member Version
	for _, os_name := range os_names {
		for _, ansible_version := range ansible_versions {
			dir_name := os_name + "_" + ansible_version
			_, err := os.Stat(dir_name)
			if err != nil {
				err := os.Mkdir(dir_name, 0755)
				if err != nil {
					log.Println(err)
					return
				}
			}

			f, err := os.Create(filepath.Join(dir_name, "Dockerfile"))
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()

			member = Version{os_name, ansible_version}
			err = tpl.Execute(f, member)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
