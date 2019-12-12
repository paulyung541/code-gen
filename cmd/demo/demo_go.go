package demo

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/paulyung541/code-gen/tpl"
)

type DemoGo struct {
	PkgName string
	*Demo
}

func (d *DemoGo) Create() error {
	projectPath, err := d.createDir()
	if err != nil {
		return err
	}

	d.CreatedTime = time.Now().Format(timeLayout)

	// create main.go
	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", projectPath))
	if err != nil {
		return err
	}
	defer mainFile.Close()

	// write content to main.go
	mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplateGo())))
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
	makeFileTemplate := template.Must(template.New("Makefile").Parse(string(tpl.MakefileTemplateGo())))
	err = makeFileTemplate.Execute(makeFile, d)
	if err != nil {
		return err
	}

	return createBaseGOMOD(projectPath, d.PkgName)
}

// create go.mod file
func createBaseGOMOD(projectPath, name string) error {
	modFile, err := os.Create(projectPath + "/go.mod")
	if err != nil {
		return err
	}
	defer modFile.Close()

	if _, err := modFile.WriteString(fmt.Sprintf("module %s", name)); err != nil {
		return err
	}

	return nil
}
