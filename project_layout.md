# [project-layout](https://github.com/golang-standards/project-layout)

## Organize

### [workspace](https://talks.golang.org/2014/organizeio.slide#9)

```
$GOPATH/
    src/
        github.com/user/repo/
            mypkg/
                mysrc1.go
                mysrc2.go
            cmd/mycmd/
                main.go
    bin/
        mycmd
    pkg/
```
### [Directories You Shouldn't Have](https://github.com/golang-standards/project-layout#src)

- 아래 설명에 따르면 프로젝트는 결국 $workspace/src/ 하위에 존재하게 된다
> The `$GOPATH` environment variable points to your (current) workspace (by default it points to `$HOME/go` on non-windows systems). 
> This workspace includes the top level `/pkg`, `/bin` and `/src` directories. 
> **Your actual project ends up being a sub-directory under `/src`**, 
> so if you have the `/src` directory in your project the project path will look like this: /some/path/to/workspace/src/your_project/src/your_code.go

```
$GOPATH(=workspace)
    ├── bin
    │
    ├── src
    │    ├── $PROJECT1_NAME
    │    │          ├── cmd
    │    │          │   ├── $APP1_NAME
    │    │          │   │           └── main.go
    │    │          │   └── $APP2_NAME
    │    │          │               └── main.go
    │    │          ├── api       
    │    │
    │    ├── $PROJECT2_NAME
```

### Go Module

- Go module 생성
```
/mnt/c/Users/daybreak/GolandProjects/src/learngo$ go mod init github.com/aimpugn/learngo
go: creating new go.mod: module github.com/aimpugn/learngo
```

- go.mod 파일 생성
```
module github.com/aimpugn/learngo

go 1.14
```

- go test로 의존성 확인
```
/mnt/c/Users/daybreak/GolandProjects/src/learngo$ go test mod_test/hello2.go
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
?       command-line-arguments  [no test files]

```

- go.mod 파일 내용 변경 확인
```
module github.com/aimpugn/learngo

go 1.14

require rsc.io/quote v1.5.2 // indirect

```

## 기타 링크
- https://stackoverflow.com/a/46663495/8562273
- https://stackoverflow.com/a/49732445/8562273
- https://stackoverflow.com/a/46650781/8562273
- [$GOPATH/src is required for imports to work](https://stackoverflow.com/a/46650781/8562273)