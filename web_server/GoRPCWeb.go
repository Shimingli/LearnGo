package main

import (
	"log"
	"fmt"
	"os"
	"net/rpc"
	"strconv"
)

type ArgsTwo struct {
	A, B int
}

type QuotientTwo struct {
	Quo, Rem int
}

func main() {
	fmt.Println("长度是多少"+strconv.Itoa( len(os.Args) ))
	fmt.Println("os*****************",os.Args,"**********************")
	if len(os.Args) != 2 {
		fmt.Println("老子要退出了哦 傻逼 一号start--------》》》", os.Args[0], "《《《---------------server  end")
		os.Exit(1)
	}

	serverAddress := os.Args[1]
    fmt.Println("severAddress==",serverAddress)
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := ArgsTwo{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot QuotientTwo
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}