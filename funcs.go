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

func GetLocalIp() string {
	conn,err := net.Dial("udp","8.8.8.8:53")
	if err != nil {
		log.Printf("get local addr err:%v",err)
		return ""
	}
	localIp := strings.Split(conn.LocalAddr().String(),":")[0]
	return localIp
}

func GetHostName()  string	{
	name,_:= os.Hostname()
	return name
}
