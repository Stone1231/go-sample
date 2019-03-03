package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		//得到輸入的字串
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" || temp == "stop" {
			break
		}

		callAPI(temp, &c)
	}
	c.Close()
}

func callAPI(str string, c *net.Conn) {
	txt := url.QueryEscape(str)
	encodeurl := "http://localhost:" + API_PORT + "/?txt=" + txt
	urls := []string{encodeurl}

	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch) // start a goroutine
	}
	for range urls {
		(*c).Write([]byte(<-ch))
	}
}

func startTCP() {
	//func Test_tcp(t *testing.T) {

	PORT := ":" + TCP_PORT
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
