//Dirghayu Mainali(L20445249)

package main

//this file test the client functionality by creating two person and sending message to each other
//if no error is detected in doing so, then we ca say the app is working properly
//the server should be running before executing this test as this is only for the client
import (
	"log"
	"net/rpc"
	"testing"
)

//
func TestJoin(t *testing.T) {
	reply := ""
	name := "Dirghayu"
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8888")
	if err != nil {
		t.Fatalf("error registering user")
	}
	err = conn.Call("ChatServer.Join", name, &reply)
	if err != nil {
		t.Fatalf("error registering user")
	}
}

//declare a message struct
type Message struct {
	User   string
	Target string
	Msg    string
}
type Nothing bool
func TestPM(t *testing.T) {
	var reply string
	nameA := "PersonX"
	nameB := "PersonY"
	msg := "i am awesome"
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8888")
	if err != nil {
		t.Fatalf("error registering user")
	}

	err = conn.Call("ChatServer.Join", nameA, &reply)
	if err != nil {
		t.Fatalf("error registering user %s",err)
	}
	err = conn.Call("ChatServer.Join", nameB, &reply)
	if err != nil {
		t.Fatalf("error registering user")
	}

	message := Message{
		User:   nameA,
		Target: nameB,
		Msg:    msg,
	}

	err1 := conn.Call("ChatServer.PM", message, &reply)
	if err1 != nil {
		log.Printf("Error telling users something: %q", err)
	}
}

