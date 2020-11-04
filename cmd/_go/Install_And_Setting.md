# INSTALLATION
## UBUNTU
### apt(Advanced Package Tool) 업데이트
```
$ sudo apt-get update  
$ sudo apt-get -y upgrade  
```

### Go 버전 확인
```
$ sudo apt-cache search golang-1.1
```

### Go 1.13 > Go 1.14
```
$ sudo apt install golang-1.14-go

Reading package lists... Done
Building dependency tree
Reading state information... Done
The following additional packages will be installed:
  golang-1.14-src
Suggested packages:
  bzr | brz mercurial subversion
The following NEW packages will be installed:
  golang-1.14-go golang-1.14-src
0 upgraded, 2 newly installed, 0 to remove and 3 not upgraded.
Need to get 63.2 MB of archives.
After this operation, 327 MB of additional disk space will be used.
Do you want to continue? [Y/n] Y
Get:1 http://archive.ubuntu.com/ubuntu focal-updates/main amd64 golang-1.14-src amd64 1.14.3-2ubuntu2~20.04.1 [13.3 MB]
Get:2 http://archive.ubuntu.com/ubuntu focal-updates/main amd64 golang-1.14-go amd64 1.14.3-2ubuntu2~20.04.1 [49.9 MB]
Fetched 63.2 MB in 4min 47s (220 kB/s)
debconf: unable to initialize frontend: Dialog
debconf: (Dialog frontend requires a screen at least 13 lines tall and 31 columns wide.)
debconf: falling back to frontend: Readline
Selecting previously unselected package golang-1.14-src.
(Reading database ... 68432 files and directories currently installed.)
Preparing to unpack .../golang-1.14-src_1.14.3-2ubuntu2~20.04.1_amd64.deb ...
Unpacking golang-1.14-src (1.14.3-2ubuntu2~20.04.1) ...
Setting up golang-1.14-src (1.14.3-2ubuntu2~20.04.1) ...
Setting up golang-1.14-go (1.14.3-2ubuntu2~20.04.1) ...
```

### Go 1.14 적용
#### GOROOT 수정
```
$ export GOROOT=/usr/lib/go-1.14/
```
#### PATH 수정
```
$ export PATH=$GOROOT/bin:$PATH
```
#### 버전 확인
```
$ go version
go version go1.14.3 linux/amd64
```

# Go
## [GOPATH](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/01.2.html)
### `$GOPATH`
- Go는 머신 내의 **모든** `go` 코드를 포함하는 `$GOPATH`라는 디렉토리에서 코드를 관리
- 머신 내에 `go`가 설치된 경로를 나타내는 `$GOROOT`라는 환경 변수와는 다르다는 것에 주의!
- 언어를 사용하기 전에 `$GOPATH`를 정의해야 한다. *nix(unix, linux 등) 시스템들에서는 `.profile`이라 불리는 파일이 있고, 해당 파일에 추가해야 한다
- `$GOPATH`의 배경 컨셉은 모호함 없이 언제든지 어떤 go 코드에 연결할 수 있다는 새로운 개념
