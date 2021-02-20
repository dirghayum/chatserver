# Programming Project – 1
### Dirghayu Mainali (L20445249)

This is the project to create a chat server using the Go RPC Package which implements the purpose of the remote procedure call, which is to run a program that calls a function on another host (server) to run the code for the function and return the results of the function execution to the original caller host which uses the results to complete a program.  

## REQUIREMENTS
- Go application
- Any Unix or Windows system

## PROJECT FOLDER STRUCTURE
- programmingProject1
```
 |-server
 |   |-server.go
 |-client
 |   |-client.go
 |-client_test.go
```

## HOW TO RUN THE PROGRAM
 
 ## Building and running the application
 -----------------------------------------------------------------------------------------------------------
 #### Start server
 ```
 1. cd into the server directory
    ProgrammingProject1$ cd server
 2. buld server.go (this will create an executable file)
    ProgrammingProject1>server$ go build server.go
 3. run the server
    ProgrammingProject1>server$ ./server
```
 
#### Run the test 
```
 4. go back to main project folder and run the test to make sure server is workimg properly and cliennt can communicate with it
    ProgrammingProject1$ go test client_test.go
if everything is working properly, you will get a OK message printed on screen
```
 
#### Start client
```
 5. cd into the client directory
    ProgrammingProject1$ cd client
 6. buld client.go (this will create an executible file)
    ProgrammingProject1>client$ go build client.go
 7. run the client
    ProgrammingProject1>client$ ./client
 ```

## TESTING FOR CONCURRENCY

 - run the client executable in multiple command prompt window with different username
    ProgrammingProject1>client$ ./client
 - It will prompt for the username. Enter a unique username every time. Username should be one word. 
   - Example:
    ➜  client$ ./client
    Welcome .. Please enter your name and press <enter> to join
    Your Username: Dirghayu
 - Try to send messages from different chat window at once (may be use someone’s help if needed)


This way the program allows you to run multiple clients and send the corresponding message at the same time. Also, the private message can be send to the user with the key word pm <clientname> <message>. We can see in the following screenshot.
 
This application has been tested on all major OS, windows, mac and linux.
