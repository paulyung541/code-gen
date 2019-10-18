/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/paulyung541/code-gen/cmd/demo"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `生成demo项目模板，使用方法如下

example:
	$ysydtp gen -d -n project 会在当前目录生成名称为 project 的目录；目录结构如下

	./project
		main.go
		Makefile`,
	Run: func(cmd *cobra.Command, args []string) {

		// 暂时这样写，以后再扩展其它模板
		isGenDemo, _ := cmd.Flags().GetBool("gen-demo")
		if !isGenDemo {
			er(errors.New("you should add '-d' flag to create demo"))
		}

		wd, err := os.Getwd()
		if err != nil {
			er(err)
		}
		fmt.Println("wd: ", wd)

		projectName, err := cmd.Flags().GetString("name")
		if err != nil {
			er(err)
		}

		demoPro := &demo.Demo{
			AbsolutePath: wd,
			Name:         projectName,
			CreatorName:  os.Getenv("USER"),
		}
		if err = demoPro.Create(); err != nil {
			er(err)
		}

		fmt.Printf("your demo project <%s> is ready!!!", projectName)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// 是否是生成 demo 代码
	genCmd.Flags().BoolP("gen-demo", "d", false, "create demo project")
	// 模板文件位置
	genCmd.Flags().StringP("name", "n", "demo", "demo project name")
}
