package main

import "fmt"

func main() {

	// 만들 것
	// 월요일 기준으로 resource 생성 금요일 기준으로 resource 삭제 혹은
	// 출근시 resource 생성, 퇴근시 resource 삭제가 필요할 것으로 보인다.
	// 일단은 멀티쓰레드는 필요 없고...
	// dynamodb? test 용도로 ?

	// 1. aws 계정에 연동한다.
	// 1.1 config.yml 파일 read
	// 1.2 aws 접속
	// 2. resource 생성
	// 3. resource 파괴... -> 계속 떠 있어야 하나 ?.. 혹은 배치 실행?
	// 배치 실행으로(옵션) 을 가지고.... 생성 및 파괴를 하면 되겠따.

	fmt.Println("hello world")
}
