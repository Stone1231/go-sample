package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//https://www.jianshu.com/p/ae72ad58ecb6

func notifyAll() {
	c := make(chan os.Signal)
	//監聽所有信號
	signal.Notify(c)
	//阻塞直到有信號傳入
	fmt.Println("啟動")
	s := <-c
	fmt.Println("退出信號", s)
}

func notify() {
	//合建chan
	c := make(chan os.Signal)
	//監聽指定信號 ctrl+c kill
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	//阻塞直到有信號傳入
	fmt.Println("啟動")
	//阻塞直至有信號傳入
	s := <-c
	fmt.Println("退出信號", s)
}

// 優雅退出GO守護進程
func exitFunc() {
	c := make(chan os.Signal)
	//監聽指定信號 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				fmt.Println("開始退出...")
				fmt.Println("執行清理...")
				fmt.Println("結束退出...")
				os.Exit(0)
			case syscall.SIGUSR1:
				fmt.Println("usr1", s)
			case syscall.SIGUSR2:
				fmt.Println("usr2", s)
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("進程啟動...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}
}

// 優雅退出GO守護進程
func exitFunc2() {
	go func() {
		fmt.Println("進程啟動...")
		sum := 0
		for {
			sum++
			fmt.Println("sum:", sum)
			time.Sleep(time.Second)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM)
	<-stopChan
	fmt.Println("main: shutting down ...")
	// 處理剩下工作與相關清理
	os.Exit(0)
}

func main() {
	// notifyAll()
	// notify()
	// exitFunc()
	exitFunc2()
}
