package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	client := &http.Client{
		Timeout: time.Nanosecond * 200,
	}

	_, err := client.Get("https://google.com")
	// solution 1 using os.IsTimeout -> ใช้ได้
	isTimeout := false
	if os.IsTimeout(err) {
		isTimeout = true
	}

	// solution 2 using net.Error -> ใช้ได้
	isTimeOut2 := false
	if e, ok := err.(net.Error); ok && e.Timeout() {
		isTimeOut2 = true
	}

	// solution 3 using StatuCode -> ใช้ไม่ได้ เพราะ เวลา err timeout มันไม่มี response มันเข้า error
	// isTimeOut3 := false
	// fmt.Printf("%d", res.StatusCode)
	// if res.StatusCode == 408 {
	// 	fmt.Println(isTimeOut3)
	// }

	fmt.Println(isTimeout)
	fmt.Println(isTimeOut2)
	// fmt.Println(isTimeOut3)
}
