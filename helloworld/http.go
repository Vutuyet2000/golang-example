package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

c := make(chan string)

func main() {
	resp, err := http.Get("http://google.com")

	fmt.Println(resp)
	fmt.Println(err)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	io.Copy(os.Stdout, resp.Body)

	fmt.Println(bs)
}
