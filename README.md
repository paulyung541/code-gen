# code-gen
a code generator CLI tool

## how to install
> go get -u github.com/paulyung541/code-gen

## create demo project
### golang
default is golang
```bash
$ code-gen gen -d -n {your_project_name}
```
then the `{your_project_name}` project will created in your current directory like follow
```
./your_project_name
    main.go
    Makefile
```

### clang
use `--demo-type` flag
```bash
$ code-gen gen -d -n ysy --demo-type clang
```

## json to golang struct
let your json file in current directory first, and run this command
```shell
$ code-gen json
```
then the go file will be created

## License
[license](https://github.com/paulyung541/code-gen/blob/master/LICENSE)