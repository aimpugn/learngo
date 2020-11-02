package main

/*
	https://stackoverflow.com/a/56015340
	- Install WSL Ubuntu (go1.12.4 linux/amd64), JetBrains GoLand 2019.1 x64
    - Configure the GOPATH in Ubuntu:
		export GOPATH=/mnt/d/dev/golang_ws
	- Configure the Project Goland in the folder: "D:\dev\golang_ws"
	- Build and Test the Golang application in WSL Ubuntu, using Ubuntu Bash.
	or
	-  IntelliJ IDEA uses cmd.exe in the terminal view by default. To replace it with the Ubuntu bash, open up the IntelliJ IDEA settings menu located under “File”  > Settings > Tools > Terminal >
	   Start directory : D:/dev/golang_ws
	   Shell Path : C:\Windows\System32\bash.exe
	   Tab Name: Local
*/
/*
	https://johngrib.github.io/wiki/golang-mod/
	$GOPATH/src 외부의 디렉토리들, 즉 learngo/cmd, learngo/pkg, learngo/internal 등에 접근하려면 go mod 명령어로 모듈 경로를 지정해야 한다
	Go mod provides access to operations on modules.

	Note that support for modules is built into all the go commands,
	not just 'go mod'. For example, day-to-day adding, removing, upgrading,
	and downgrading of dependencies should be done using 'go get'.
	See 'go help modules' for an overview of module functionality.

*/

func main() {

}
