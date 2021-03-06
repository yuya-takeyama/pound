package main

import (
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/lestrrat/go-tcptest"
)

func TestRun(t *testing.T) {
	fin := make(chan bool)

	run := func(port int) {
		server := Server{Host: "localhost", Port: port}
		go func() {
			err := server.Run()
			if err != nil {
				t.Error("Failed to run the server.")
			}
		}()
		<-fin
		server.Shutdown()
	}

	tcptestinfo, err := tcptest.Start(run, 10*time.Second)
	if err != nil {
		t.Error("An error has occured in the tcputil.")
	}

	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(tcptestinfo.Port()))
	if err != nil {
		t.Errorf("Failed to connect the server localhost:%d", tcptestinfo.Port())
	}

	conn.Close()

	fin <- true
	tcptestinfo.Wait()
}
