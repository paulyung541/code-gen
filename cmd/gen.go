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
	Short: "code generator",
	Long: `生成demo项目模板，使用方法如下

example:
$code-gen gen -d -n project 会在当前目录生成名称为 project 的目录；目录结构如下

./project
	main.go
	Makefile`,
	Run: func(cmd *cobra.Command, args []string) {
		// 暂时这样写，以后再扩展其它模板
		isGenDemo, _ := cmd.Flags().GetBool("gen-demo")
		if !isGenDemo {
			er(errors.New("you should add '-d' flag to create demo"))
		}

		projectName, err := cmd.Flags().GetString("name")
		if err != nil {
			er(err)
		}

		pkgName, err := cmd.Flags().GetString("pkg-name")
		if err != nil {
			er(err)
		}

		userName := os.Getenv("USER")
		if pkgName == "" {
			pkgName = fmt.Sprintf("demo.com/%s/%s", userName, projectName)
		}

		wd, err := os.Getwd()
		if err != nil {
			er(err)
		}

		codeType, err := cmd.Flags().GetString("demo-type")
		if err != nil {
			er(err)
		}
		switch codeType {
		case "golang":
			demoPro := &demo.DemoGo{
				Demo: &demo.Demo{
					AbsolutePath: wd,
					Name:         projectName,
					CreatorName:  userName,
				},
				PkgName: pkgName,
			}
			if err = demoPro.Create(); err != nil {
				er(err)
			}
		case "clang":
			demoPro := &demo.DemoC{
				Demo: &demo.Demo{
					AbsolutePath: wd,
					Name:         projectName,
					CreatorName:  userName,
				},
			}
			if err = demoPro.Create(); err != nil {
				er(err)
			}
		default:
			fmt.Println("--demo-type value should be one of [golang, clang]")
			return
		}

		fmt.Printf("your [%s] code demo project <%s> is ready!!!\n", codeType, projectName)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// 是否是生成 demo 代码
	genCmd.Flags().BoolP("gen-demo", "d", false, "create demo project")
	// demo project name
	genCmd.Flags().StringP("name", "n", "demo", "demo project name")
	// pkg name
	genCmd.Flags().StringP("pkg-name", "", "", "demo pkg name")

	genCmd.Flags().String("demo-type", "golang", "the code type of demo use")
}
