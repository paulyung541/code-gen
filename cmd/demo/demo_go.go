package demo

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/paulyung541/code-gen/tpl"
)

type DemoGo struct {
	PkgName           string
	DisableVSCODEFile bool
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

	return d.createBaseGOMOD(projectPath, d.PkgName)
}

// create go.mod file
func (d *DemoGo) createBaseGOMOD(projectPath, name string) error {
	modFile, err := os.Create(projectPath + "/go.mod")
	if err != nil {
		return err
	}
	defer modFile.Close()

	if _, err := modFile.WriteString(fmt.Sprintf("module %s", name)); err != nil {
		return err
	}

	if !d.DisableVSCODEFile {
		return d.createVSCODEFile(projectPath)
	}

	return nil
}

// create .vscode dir and launch.json
func (d *DemoGo) createVSCODEFile(projectPath string) error {
	vscodeDir := projectPath + "/.vscode"
	if err := os.Mkdir(vscodeDir, 0754); err != nil {
		return err
	}

	vscodeFile, err := os.Create(vscodeDir + "/launch.json")
	if err != nil {
		return err
	}
	defer vscodeFile.Close()

	if _, err := vscodeFile.WriteString(`{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go",
            "env": {},
            "args": []
        }
    ]
}`); err != nil {
		return err
	}

	return nil
}
