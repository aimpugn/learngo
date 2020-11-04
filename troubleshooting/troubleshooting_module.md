# Module
## `$GOPATH/go.mod exists but should not`
### 문제
- 프로젝트 루트(`/mnt/c/Users/daybreak/GolandProjects/learngo`)에서 `go test` 수행 시 발생
- `$GOPATH`에 `go.mod`가 존재하지만, 그러지 않아야 한다.
- 그렇다면 `$GOPATH`바로 하위가 아닌 다른 경로에 존재해야 한다? 
### 원인
- [kubernetes](https://github.com/aimpugn/kubernetes) 참고하면
```
You have a working Go environment.

mkdir -p $GOPATH/src/k8s.io
cd $GOPATH/src/k8s.io
git clone https://github.com/kubernetes/kubernetes
cd kubernetes
make
```
- `$GOPATH/src/k8s.io` 디렉토리를 먼저 생성하고, 해당 디렉토리에 kubernetes 코드를 클론한다
### 해결