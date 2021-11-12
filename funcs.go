package common_tools

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func GetNowTimestr() string {
	return time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")
}

func GetLocalIp() (string,error) {
	conn,err := net.Dial("udp","8.8.8.8:53")
	if err != nil {
		log.Printf("get local addr err:%v",err)
		return "",err
	}
	localIp := strings.Split(conn.LocalAddr().String(),":")[0]
	return localIp,nil
}

func GetHostName()  string	{
	name,_:= os.Hostname()
	return name + "v2.0.1"
}
