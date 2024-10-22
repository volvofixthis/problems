package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Sender interface {
	Send(s string)
}

type TcpSender struct {
}

func (t *TcpSender) Send(s string) {
	fmt.Println(s)
}

func NewClient(sender Sender) *Client {
	return &Client{
		Sender: sender,
		c:      0,
		m:      sync.Mutex{},
	}
}

type Client struct {
	Sender
	c int
	m sync.Mutex
}

func (c *Client) Send(ctx context.Context, s string) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	cancel()
	cancel()
	cancel()
	c.m.Lock()
	sample := c.c == 100
	c.m.Unlock()
	if sample {
		c.m.Lock()
		c.c = 0
		c.m.Unlock()
		c.Sender.Send(s)
	} else {
		c.m.Lock()
		c.c++
		c.m.Unlock()
	}
}
