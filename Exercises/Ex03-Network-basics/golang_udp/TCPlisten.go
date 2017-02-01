//copied from: https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "129.241.187.43:33546")
  for {
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Text to send: ")
	text, _ := 'test av TCP\x00'
    // send to socket
    fmt.Fprintf(conn, text)
    // listen for reply
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message from server: "+message)
  }
}


