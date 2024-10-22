package main

import (
	"context"
	"testing"
)

func TestSender(t *testing.T) {
	client := NewClient(&TcpSender{})
	for i := 0; i <= 200; i++ {
		client.Send(context.Background(), "kus")
	}
}
