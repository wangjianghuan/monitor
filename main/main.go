package main

import (
	"fmt"
	"time"
	"os/exec"
	"io/ioutil"
	"flag"
	"strings"
)

var kafkaIp,kafkaPort,interval string
const key string = "0"

func init() {
	flag.StringVar(&kafkaIp,"kafkaIp","10.2.41.228","Kafka Ip")
	flag.StringVar(&kafkaPort,"port","9092","Listen on port")
	flag.StringVar(&interval,"interval", "300s","The interval time.(unit: seconds)")
}

func main() {

	flag.Parse()
	println(kafkaIp)
	println(kafkaPort)
	println(interval)

	dur, _ := time.ParseDuration(interval)

	for range time.Tick(dur) {
		linuxFlag := ExecCommandWithLinux("echo -e '\n' | telnet 10.2.41.228 9092 2>/dev/null | grep Connected | wc -l")
		linuxFlag = strings.Replace(linuxFlag, "\n", "", -1)
		fmt.Println("Execute finished:" + linuxFlag)
		if strings.ToLower(linuxFlag) == strings.ToLower(key) {
			Sender()
			break;
		} else {
			fmt.Println("kafka is running...")
		}
		//windowsData := ExecCommandWithWindow(port)
		//fmt.Println(windowsData)
		//if strings.Contains(windowsData, "1801") && strings.Contains(windowsData, "LISTENING") {
		//	// nothing to do
		//	fmt.Println("success")
		//} else {
		//	Sender();
		//	break;
		//}
	}
	fmt.Println("Found error,exit!")
}

func ExecCommandWithLinux(strCommand string)(string){
	cmd := exec.Command("/bin/bash", "-c", strCommand)
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil{
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}
	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return string(out_bytes)
}

func ExecCommandWithWindow(strCommand string)(string){
	cmd := exec.Command("cmd", "/C","netstat -an | findStr 1801")
	//cmd := exec.Command("ping", "127.0.0.1")
	stdout, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil{
		fmt.Println("Execute failed when Start:" + err.Error())
		return ""
	}
	out_bytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return ""
	}
	return string(out_bytes)
}


func Sender() {
	serverHost := "smtp.exmail.qq.com"
	serverPort := 25
	fromEmail := "wangjh@emrubik.com"
	fromPasswd := "password"

	myToers := "jiangfanwang@126.com" // 逗号隔开
	myCCers := ""                     //"readchy@163.com"

	subject := "现网KAKFA停止"
	body := `kafka异常停止了<br>
            <h3>OMG</h3>
             快去排查问题<br>`
	// 结构体赋值
	myEmail := &EmailParam{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPasswd,
		Toers:      myToers,
		CCers:      myCCers,
	}

	InitEmail(myEmail)
	fmt.Println("InitEmail done");
	SendEmail(subject, body)
	fmt.Println("SendEmail done");
}