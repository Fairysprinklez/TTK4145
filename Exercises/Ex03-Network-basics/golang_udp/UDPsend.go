//imported from https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/

package main
 
import (
    "fmt"
    "net"
	"time"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    ServerAddr,err := net.ResolveUDPAddr("udp","129.241.187.43:20024")
    CheckError(err)
 
    LocalAddr, err := net.ResolveUDPAddr("udp", "129.241.187.255:20124")
    CheckError(err)
 
    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)
 
    defer Conn.Close()

    for {
        msg := "Hello from Arbeidsplass 24"
        
        buf := []byte(msg)
        _,err := Conn.Write(buf)
        if err != nil {
            fmt.Println(msg, err)
        }
        time.Sleep(time.Second * 1)
    }
}
