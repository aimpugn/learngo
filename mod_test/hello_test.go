package hello

import "testing"

/*
go tool: no such tool "compile" 에러 발생
> https://stackoverflow.com/a/51182684

Window10의 WSL이어서, GOROOT와 GOTOOLDIR를 WSL의 Go에 맞춰 수정
$ export GOROOT=/usr/lib/go-1.13/
$ export GOTOOLDIR=/usr/lib/go-1.13/pkg/tool/linux_amd64/

PASS
ok      _/mnt/c/Users/daybreak/GolandProjects/learngo/mod_test  0.004s  // 전체 패키지 테스트 요약

*/
func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
