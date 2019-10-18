package demo

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/paulyung541/code-gen/tpl"
)

const timeLayout = "2006-01-02 15:04:05"

type Demo struct {
	AbsolutePath string
	Name         string
	CreatorName  string
	CreatedTime  string
}

func (d *Demo) Create() error {
	projectPath := d.AbsolutePath + "/" + d.Name
	d.CreatedTime = time.Now().Format(timeLayout)

	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(projectPath, 0754); err != nil {
			return err
		}
	}

	// create main.go
	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", projectPath))
	if err != nil {
		return err
	}
	defer mainFile.Close()

	// write content to main.go
	mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
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
	makeFileTemplate := template.Must(template.New("Makefile").Parse(string(tpl.MakefileTemplate())))
	err = makeFileTemplate.Execute(makeFile, d)
	if err != nil {
		return err
	}

	return nil
}
