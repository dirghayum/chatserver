// Dirghayu Mainali-(L20445249)

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Nothing bool

type Message struct {
	User   string
	Target string
	Msg    string
}

type ChatServer struct {
	port         string
	messageQueue map[string][]string
	users        []string
	shutdown     chan bool
}

/*
Join is responsible for performing two jobs
1) register users
2) send message to all existing users that new users has joined server
 */
func (c *ChatServer) Join(username string, reply *string) error {
	*reply = "Welcome to Dirghayu Chatroom, " + username + "\ntype pm <username of another user> <msg> to send private message"

	c.messageQueue[username] = nil
	c.users = append(c.users, username)
	for k, _ := range c.messageQueue {
		c.messageQueue[k] = append(c.messageQueue[k], username+" has joined the group.")
	}
	log.Printf("%s has joined the chat.\n", username)

	return nil
}

// this method checks for new message is pool
func (c *ChatServer) CheckMessages(username string, reply *[]string) error {
	*reply = c.messageQueue[username]
	c.messageQueue[username] = nil
	return nil
}

// PM sends private message to target client defined by type "msg"
func (c *ChatServer) PM(msg Message, reply *Nothing) error {

	if queue, ok := c.messageQueue[msg.Target]; ok {
		m := msg.User + " says : " + msg.Msg
		c.messageQueue[msg.Target] = append(queue, m)
	} else {
		m := msg.Target + " does not exist"
		c.messageQueue[msg.User] = append(queue, m)
	}

	*reply = false

	return nil
}

// Broadcast sends out message to all joined clients
func (c *ChatServer) Broadcast(msg Message, reply *Nothing) error {
	for k, v := range c.messageQueue {
		m := msg.User + " says " + msg.Msg
		c.messageQueue[k] = append(v, m)
	}

	*reply = true

	return nil
}

//remove the user from the messagequeue
func (c *ChatServer) Logout(username string, reply *Nothing) error {

	delete(c.messageQueue, username)

	for i := range c.users {
		if c.users[i] == username {
			c.users = append(c.users[:i], c.users[i+1:]...)
		}
	}

	for k, v := range c.messageQueue {
		c.messageQueue[k] = append(v, username+" has logged out.")
	}

	fmt.Println("User " + username + " has logged out.")

	*reply = false

	return nil
}

// Shutsdown server and disconnects all clients
func (elt *ChatServer) Shutdown(nothing Nothing, reply *Nothing) error {

	log.Println("Server shutdown..")
	*reply = false
	elt.shutdown <- true

	return nil
}

// runs the server
func RunServer(cs *ChatServer) {
	rpc.Register(cs)
	rpc.HandleHTTP()

	log.Printf("Listening on port 8888")

	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Panicf("Can't bind port to listen. %q", err)
	}

	go http.Serve(l, nil)
}

// main function
func main() {
	cs := new(ChatServer)
	cs.messageQueue = make(map[string][]string)
	cs.shutdown = make(chan bool, 1)

	RunServer(cs)

	<-cs.shutdown
}
