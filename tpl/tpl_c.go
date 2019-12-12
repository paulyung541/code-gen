package tpl

func MainTemplateC() []byte {
	return []byte(`/*
 * This is a demo name {{ .Name }} created by {{ .CreatorName }}
 * {{ .CreatedTime }}
 * default lib: /lib /usr/lib /usr/local/lib
*/
#include<stdio.h>
// #include<ctype.h>
// #include<assert.h>
// #include<errno.h>
// #include<float.h>
// #include<limits.h>
// #include<locale.h>
// #include<math.h>
// #include<setjmp.h>
// #include<signal.h>
// #include<stdarg.h>
// #include<stddef.h>
// #include<stdlib.h>
// #include<string.h>
// #include<string.h>
// you can choice you want

int main(int argc, char *argv[]) {
	printf("this is {{ .Name }}\n");
	// your code here
}

`)
}

func MakefileTemplateC() []byte {
	return []byte(`# This is a demo name {{ .Name }} created by {{ .CreatorName }}
# {{ .CreatedTime }}

CC           ?= gcc
PROJECT_NAME ?= {{ .Name }}

all: build run

.PHONY: run
run:
	@echo ">> run binaries"
	@./$(PROJECT_NAME)

.PHONY: build
build:
	@echo ">> building binaries"
	@$(CC) main.c -o $(PROJECT_NAME)

.PHONY: asm
asm:
	@$(CC) -S -fverbose-asm main.c

.PHONY: clean
clean:
	rm -rf $(PROJECT_NAME) *.o *.i

`)
}