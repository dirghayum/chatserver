// Dirghayu Mainali (L20445249)

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc" 
	"os"
	"strconv"
	"strings"
	"time"
)

type Nothing bool

type Message struct {
	User   string
	Target string
	Msg    string
}

type ChatClient struct {
	Username string
	Address  string
	Client   *rpc.Client
}

var (
	DEFAULT_PORT = 8888
	DEFAULT_HOST = "localhost"
)

// connects to server and returns ChatClient type
func (c *ChatClient) getClientConnection() *rpc.Client {
	var err error

	if c.Client == nil {
		c.Client, err = rpc.DialHTTP("tcp", c.Address)
		if err != nil {
			log.Panicf("Error establishing connection with host: %q", err)
		}
	}

	return c.Client
}

// sends join request to server
func (c *ChatClient) Join() {
	var reply string
	c.Client = c.getClientConnection()
	err := c.Client.Call("ChatServer.Join", c.Username, &reply)
	if err != nil {
		log.Printf("Error registering user: %q", err)
	} else {
		log.Printf("Reply: %s", reply)
	}
}

// checks if there are any message for current client by calling CheckMessages method on RPC server
func (c *ChatClient) CheckMessages() {
	var reply []string
	c.Client = c.getClientConnection()

	for {
		err := c.Client.Call("ChatServer.CheckMessages", c.Username, &reply)
		//if any error detected, halt
		if err != nil {
			log.Fatalln("Chat has been shutdown. Goodbye.")
		}
		//if no error is detecred, loop thru all messages if any and print it on screen
		for i := range reply {
			log.Println(reply[i])
		}
		//wait for one second
		time.Sleep(time.Second)
	}
}

// Sends private message to target client

func (c *ChatClient) PM(params []string) {
	var reply Nothing
	c.Client = c.getClientConnection()

	/*
	if message begins with the word "pm", split the messages with the white space.
	the first word param[0] is pm
	the second word param[1] is the recipient
	the third word param[2] and all remaining characters are the intended messages
	*/
	if len(params) > 2 {
		msg := strings.Join(params[2:], " ")
		message := Message{
			User:   c.Username,
			Target: params[1],
			Msg:    msg,
		}

		err := c.Client.Call("ChatServer.PM", message, &reply)
		if err != nil {
			log.Printf("Error telling users something: %q", err)
		}
	} else {
		log.Println("Usage of tell: tell <user> <msg>")
	}
}

// calls Broadcast method on RPC server to send out message to all connected clients
func (c *ChatClient) Broadcast(params string) {
	var reply Nothing
	c.Client = c.getClientConnection()

	message := Message{
		User:   c.Username,
		Target: "",
		Msg:    params,
	}

	err := c.Client.Call("ChatServer.Broadcast", message, &reply)
	if err != nil {
		log.Printf("Error saying something: %q", err)
	}
}

// Logout logs out the current user and shuts down the client
func (c *ChatClient) Logout() {
	var reply Nothing
	c.Client = c.getClientConnection()

	err := c.Client.Call("ChatServer.Logout", c.Username, &reply)
	if err != nil {
		log.Printf("Error logging out: %q", err)
	}
}

// populates ChatClient type and returns its instance
func createClient() *ChatClient {
	var c = &ChatClient{}
	fmt.Println("Welcome .. Please enter your name and press <enter> to join")
	fmt.Print("Your Username : ")
	reader := bufio.NewReader(os.Stdin)
	var name, err = reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")
	if err != nil {
		log.Printf("Error: %q\n", err)
	}

	c.Username = name
	c.Address = net.JoinHostPort(DEFAULT_HOST, strconv.Itoa(DEFAULT_PORT))
	return c
}

// starts chat with connected group
func mainLoop(c *ChatClient) {
	c.Join()
	for {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error: %q\n", err)
		}

		line = strings.TrimSpace(line)
		params := strings.Fields(line)
		//if the message starts with the tet "pm" then it is a private message
		if strings.HasPrefix(line, "pm") {
			c.PM(params)
		}else if strings.HasPrefix(line ,"logout"){
			c.Logout()
			break
		} else {
			//if the message doesnt begin with "pm" or "logout" then it is  a  broadcast message
			c.Broadcast(line)
		}
	}
}

// boots the client
func main() {
	client := createClient()
	go client.CheckMessages()

	mainLoop(client)
}
