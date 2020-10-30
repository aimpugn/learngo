/*
첫 줄은 반드시 package 라인
main은 go 프로그램이 실행되는 특별한 패키지
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
모든 Go 프로그램은 main 함수에서 실행된다
*/
func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		/*
			WriteHeader가 호출된 적 없으며, 데이터 출력 전에 Write가 WriteHeader(http.StatusOK) 호출
			Header에 Content-Type 줄이 없으면, Write는 func DetectConentType로 전달되는 쓰여진 데이터의 초기 512 bytes의 결과에 Content-Type 집합(set)을 추가
			추가적으로, 쓰여진 데이터가 몇 KB 미만이고 Flush가 호출되지 않았다면, Content-Length 헤더가 자동으로 추가된다
			HTTP 프로토콜 버전과 클라이언트에 따라서, Write 또는 WriteHeader를 호출하는 것은 향후 Request.Body에 대한 읽기를 막을 수 있다.

			# HTTP/1.x 요청
			handlers는 응답을 쓰기 전에 필요한 요청 바디 데이터를 읽어야 한다.
			한번 헤더가 flush되고 나면(due to either an explicit Flusher.Flush call or writing enough data to trigger a flush), 요청 바디는 사용할 수 없을 수 있다.
			# HTTP/2 요청
			Go HTTP 서버는 handler가 응답을 쓰는 동시에 요청 바디를 읽을 수 있게 한다
			하지만 그러한 동작을 모든 HTTP/2 클라이언트 지원하지 않을 수 있다

			가용성을 극대화하기 위해서 가능하면 Handlers는 데이터를 쓰기 전에 읽어야 한다
		*/

		/*
			GO는 try/catch/finally를 지원하지 않는 대신, 그냥 error가 함수의 반환 값이 된다
			Go에서 error는 예상치 못한 일이 발생 했을 때 함수가 반환할 수 있는 값
			error는 Go의 내장(build-in) 타입이며, 초기값은 nil이다

			https://stackoverflow.com/questions/43976140/check-errors-when-calling-http-responsewriter-write
		*/
		rw = LogWriter{rw}
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})

	http.ListenAndServe(":8080", nil)
}

/*
자동으로 로그 남기기 위해 http.ResponseWriter 한번 감싼다(wrapping)
*/
type LogWriter struct {
	http.ResponseWriter
}

func (w LogWriter) Write(p []byte) (n int, err error) {
	n, err = w.ResponseWriter.Write(p)
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
	return
}
