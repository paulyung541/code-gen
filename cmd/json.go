package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/paulyung541/code-gen/cmd/json"
	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "a conversion tool for json to golang struct",
	Long: `eg:
now your directory have a.json and b.json two files
run follow command
	$ code-gen json
then all json file in this directory will be converted to a.go b.go
also you can appoint a file you want to be converted use -f
and golang struct to json is similar to previous`,

	Run: func(cmd *cobra.Command, args []string) {
		if !checkFileName() {
			er(errors.New("file name must different"))
			return
		}

		cFiles, err := getConversionFiles()
		if err != nil {
			er(err)
		}

		cFiles.Convert()

		if len(cFiles.JSONSrcFiles) > 0 {
			fmt.Printf("input  %v files", cFiles.JSONSrcFiles)
			fmt.Printf("output %v files", cFiles.GoSuccessFiles)
		}

		if len(cFiles.GoSrcFiles) > 0 {
			fmt.Printf("input  %v files", cFiles.GoSrcFiles)
			fmt.Printf("output %v files", cFiles.JSONSuccessFiles)
		}
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	jsonCmd.Flags().StringP("file", "f", "", "file path")
}

// if have same name file, this cmd will get a fail
func checkFileName() bool {
	files := getAllFile()
	if files == nil || len(files) <= 0 {
		return false
	}

	files = getPrefix(files)

	set := map[string]struct{}{}
	for _, f := range files {
		set[f] = struct{}{}
	}

	return len(set) == len(files)
}

func getAllFile() []string {
	wd, err := os.Getwd()
	if err != nil {
		er(err)
	}

	files, err := ioutil.ReadDir(wd)
	if err != nil {
		er(err)
	}

	var names []string
	for _, file := range files {
		names = append(names, file.Name())
	}

	return names
}

func getPrefix(files []string) []string {
	var names []string
	for _, file := range files {
		arg := strings.Split(file, ".")
		names = append(names, arg[0])
	}
	return names
}

func getConversionFiles() (*json.ConversionFiles, error) {
	var jsonFiles []string
	var goFiles []string

	files := getAllFile()
	if files == nil || len(files) <= 0 {
		return nil, errors.New("no files")
	}

	for _, file := range files {
		if strings.Contains(file, "json") {
			jsonFiles = append(jsonFiles, file)
		}
		if strings.Contains(file, "go") {
			goFiles = append(goFiles, file)
		}
	}

	if len(jsonFiles) <= 0 && len(goFiles) <= 0 {
		return nil, errors.New("no json or go files")
	}

	return &json.ConversionFiles{
		JSONSrcFiles: jsonFiles,
		GoSrcFiles:   goFiles,
	}, nil
}
