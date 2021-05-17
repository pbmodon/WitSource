package main

import "net"

type client struct {
	conn     net.Addr
	pers     string
	room     *room
	commands chan<- command
}
