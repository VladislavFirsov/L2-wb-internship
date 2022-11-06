package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func NewTelnet(host, port string, timeout time.Duration) *Telnet {
	return &Telnet{
		host:    host,
		port:    port,
		timeout: timeout * time.Second,
	}
}

type Telnet struct {
	host    string
	port    string
	timeout time.Duration
	conn    net.Conn
}

func (t *Telnet) dial() error {
	c, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", t.host, t.port), t.timeout)
	if err != nil {
		return err
	}
	t.conn = c
	return nil
}

func (t *Telnet) write() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		_, err = t.conn.Write([]byte(msg))
		if err != nil {
			return err
		}
	}
}

func (t *Telnet) getTimeout(s string) error {
	if strings.Contains(s, "timeout") {
		re := regexp.MustCompile("[0-9]+")
		tt := re.FindAllString(s, -1)[0]
		num, err := strconv.Atoi(tt)
		if err != nil {
			return err
		}
		t.timeout = time.Duration(num)
	} else {
		t.timeout = 10
	}
	return nil
}

func (t *Telnet) getSocket(host, port string) {
	t.port = port
	t.host = host
}
