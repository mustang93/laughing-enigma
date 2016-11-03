package main

import (
	"net"
	"log"
	"sync"
	"time"
	"strings"
	"fmt"
	"math/rand"
)

func main() {

	text := []byte(strings.Repeat("Hey", 10000))

	var waitGroup sync.WaitGroup
	for j:=0; j<10000; j++{
		for i:=0; i<200; i++ {
			waitGroup.Add(1)
			go func() {
				defer waitGroup.Done()

				udpAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("10.5.0.206:%d", rand.Int()%(1<<16)))
				if err != nil {
					log.Fatalln(err)
				}

				udpConn, err := net.DialUDP("udp4", nil, udpAddr)
				if err != nil {
					log.Fatalln(err)
				}
				defer udpConn.Close()
				udpConn.Write(text)
			}()
		}
		time.Sleep(50*time.Millisecond)
	}

	waitGroup.Wait()
}
