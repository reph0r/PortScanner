package scanner

import (
	"fmt"
	"net"
	"time"

)

func Connect(ip string,port int)(string,int,error){
	t := fmt.Sprintf("%v:%v",ip,port)
	//fmt.Println(t)
	conn,err := net.DialTimeout("tcp",t,2*time.Second)
	defer func ()  {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return ip,port,err
}
