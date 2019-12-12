package demo

import "os"

type Demo struct {
	AbsolutePath string
	Name         string
	CreatorName  string
	CreatedTime  string
}

const timeLayout = "2006-01-02 15:04:05"

func (d *Demo) createDir() (string, error) {
	projectPath := d.AbsolutePath + "/" + d.Name
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(projectPath, 0754); err != nil {
			return "", err
		}
	}
	return projectPath, nil
}
