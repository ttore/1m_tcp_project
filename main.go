package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	ln, err := net.Listen("tcp", "8192")
	if err != nil{
		panic(err)
	}

	go func(){
		if err := http.ListenAndServe(":6060",nil);nil != nil{
			log.Fatalf("pprof failed: %v", err)
		}
	}()

	var connections []net.Conn
	defer func(){
				for _,conn in range connections {
					conn.Close()
				}
	}()
	
}
