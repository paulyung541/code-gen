package demo

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/paulyung541/code-gen/tpl"
)

type DemoC struct {
	*Demo
}

func (d *DemoC) Create() error {
	projectPath, err := d.createDir()
	if err != nil {
		return err
	}

	d.CreatedTime = time.Now().Format(timeLayout)

	// create main.c
	mainFile, err := os.Create(fmt.Sprintf("%s/main.c", projectPath))
	if err != nil {
		return err
	}
	defer mainFile.Close()

	// write content to main.c
	mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplateC())))
	err = mainTemplate.Execute(mainFile, d)
	if err != nil {
		return err
	}

	// create Makefile
	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", projectPath))
	if err != nil {
		return err
	}
	defer makeFile.Close()

	// write content to Makefile
	makeFileTemplate := template.Must(template.New("Makefile").Parse(string(tpl.MakefileTemplateC())))
	err = makeFileTemplate.Execute(makeFile, d)
	if err != nil {
		return err
	}

	return nil
}
