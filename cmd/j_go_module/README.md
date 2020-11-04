# Using Go Modules
- [Source of Contents](https://blog.golang.org/using-go-modules)

## 개요
- Go 1.11과 1.12 버전에서 [모듈에 대한 사전(preliminary) 지원](https://golang.org/doc/go1.11#modules) 포함
- [Go의 새로운 의존성 관리 시스템](https://blog.golang.org/versioning-proposal) 
- 모듈은 루트 경로에 위치한 go.mod과 함께 파일 트리에 저장된 [Go 패키지](https://golang.org/ref/spec#Packages) 모음 
- $GOPATH/src 외부에 위치한 경우
    - 현재 또는 어떤 상위 디렉토리가 go.mod 파일을 가지고 있을 때 모듈을 사용할 수 있다
- $GOPATH/src 내부에 위치한 경우
    - 호환성을 위해서, go.mod 파일이 반견된다고 하더라도 go command는 예전 GOPATH 모드로 실행된다. [이 문서](https://golang.org/cmd/go/#hdr-Preliminary_module_support) 참조
- Go 1.13부터 module 모드가 모든 개발의 기본이 됐다

## 새로운 모듈 생성
- $GOPATH/src 외부에 hello.go 생성
```go
package hello

func Hello() string {
    return "Hello, world."
}
```

- hello_test.go 생성. 이때, go.mod 파일이 없으므로 디렉토리는 패키지를 포함하지만, 모듈을 포함하진 않는다. 
```go
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```
- go test 수행하면 아래와 같은 결과 얻는다.
```
PASS
ok      _/mnt/c/Users/daybreak/GolandProjects/learngo/mod_test  0.004s  // 전체 패키지 테스트 요약
```
   - $GOPATH와 다른 모듈 외부에서 작업하고 있으므로, go 커맨드는 현재 디렉토리에 대한 임포트 경로를 모르기 때문에 디렉토리 이름에 기반하여 언더스코어(_) 붙은 가상의 경로를 생성한다
- go mod init 사용하여 현재 디렉토리를 모듈의 루트로 만들면,
```
/mnt/c/Users/daybreak/GolandProjects/learngo/mod_test$ go mod init github.com/mod_test
go: creating new go.mod: module github.com/mod_test
```
- 아래와 같은 go.mod 파일이 생성되며
```
module github.com/mod_test

go 1.13
```

- go test를 다시 실행하면 아래와 같은 결과를 얻을 수 있다
```
/mnt/c/Users/daybreak/GolandProjects/learngo/mod_test$ go test
PASS
ok      github.com/mod_test     0.004s
```

- go.mod 파일은 모듈의 루트에만 나타난다. 
    - 하위 디렉토리에 있는 패키지들은 `모듈 경로` + `하위 디렉토리 경로`로 구성된 임포트 경로를 갖는다
    - 가령 world라는 하위 디렉토리 생성한다면,
        - 모듈 경로 `github.com/mod_test` + 하위 디렉토리 경로 `/world`를 자동으로 `github.com/mod_test` 모듈의 일부로 파악하며
        - 해당 디렉토리에서 go mod init 명령어를 실행할 필요 없다

## 의존성 추가
- hello.go 파일에서 `rsc.io/quote` 임포트하여 Hello 구현에 사용. (IDE에서 빨간 줄이 나와도 그대로 진행한다)
```
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```

- go test 실행하면 아래와 같은 결과가 나오고, IDE에서 빨간 줄도 사라지며 제대로 import가 되는 것을 확인할 수 있다

```
/mnt/c/Users/daybreak/GolandProjects/learngo/mod_test$ go test
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
PASS
ok      github.com/mod_test     0.005s
```
- `go` 커맨드는 `go.mod`에 나열된 특정 의존성 모듈 버전을 사용하여 임포트를 해소(resolve)한다
- `go.mod`에 나열되지 않은 패키지의 `import`를 만나게 되면, `go` 커맨드는 자동으로 해당 패키지를 포함하고 있는 모듈을 탐색하고 `go.mod`에 최신 버전으로 추가한다
```
module github.com/mod_test

go 1.13

require rsc.io/quote v1.5.2
```
- 위에서 `최신`(Latest)이라 함은 다음과 같이 정의된다
    1. [prerelease](https://semver.org/#spec-item-9) 아닌 latest로 태그된 stable 버전
    2. latest로 태그된 prerelease 버전
    3. leates 태그되지 않은 버전
- 다시 go test를 수행해도 `go.mod` 파일이 최신이고 다운로드된 모듈이 `$GOPATH/pkg/mod`에 캐시되기 때문에 위의 의존성 해소 작업을 다시 하지는 않는다
- 단, `go` 커맨드가 새로운 의존성을 빠르고 쉽게 추가할 수 있도록 하지만, 비용이 없는 것은 아니다. 정확성(correctness), 보안성 그리고 적절한 라이센스 등에 있어서 나의 모듈이 말 그대로 새로운 모듈에 `의존`하게 된다
    - [Our Software Dependency Problem.](https://research.swtch.com/deps) by Russ Cox 참조
- 직접적인 의존성을 추가함으로써, 간접적인 의존성도 가져오게 된다. `go list -m all`은 현재 모듈과 모든 의존성을 보여준다
```
/mnt/c/Users/daybreak/GolandProjects/learngo/mod_test$ go list -m all
github.com/mod_test
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```
- `go list` 출력에서, `main module`로도 알려진 현재 모듈은 언제나 첫줄에 위치하고, 그 다음에는 모듈 경로로 정렬된 의존성들이 따른다
- `golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c`에서 `v0.0.0-20170915032832-14c0d48ead0c`은 태그가 지정되지 않은 특정 커밋에 대한 `go` 커맨드의 버전 구문(syntax)인 [pseudo-version](https://golang.org/cmd/go/#hdr-Pseudo_versions) 예제
- `go.mod` 외에, `go` 커맨드는 특정 모듈 버전의 컨텐츠의 예상되는 [cryptographic hashes](https://golang.org/cmd/go/#hdr-Module_downloading_and_verification) 포함하는 `go.sum` 파일을 유지한다
- `go` 커맨드는 `go.sum` 파일을 사용해서 장래에 해당 모듈의 다운로드 시 처음 다운로드 했던 것과 동일한 비트를 가져오도록 하며, 예상치 못하게 프로젝트가 의존한느 모듈이 변하지 않는 것을 보장한다.
- `go.mod`와 `go.sum` 파일은 버전 컨트롤에서 체크되어야 한다

## 의존성 업그레이드

## 새로운 주요 버전에 대한 의존성 추가

## 새로운 주요 버전으로 의존성 업그레이드

## 사용하지 않는 의존성 삭제
