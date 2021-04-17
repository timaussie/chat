package main

import (
	"net"
)

type Client struct {
	socket             net.Conn
	register_channel   chan<- *Client
	unregister_channel chan<- *Client
	username           string
}

func NewClient(socket net.Conn, reg_channel, unreg_channel chan<- *Client) *Client {
	return &Client{
		socket:             socket,
		register_channel:   reg_channel,
		unregister_channel: unreg_channel,
	}
}
