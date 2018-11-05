package opencc

import (
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"path/filepath"
	"log"
	"time"
)
//
func Test_config(t *testing.T){
	fileName := `s2t.json`
	body, err := ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Println(err)
		return
	}
	var conf *Config
	err = json.Unmarshal(body, &conf)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(conf)

}
//
func Test_opencc(t *testing.T){
	cc,err := NewOpenCC("s2twp")
	if err != nil{
		fmt.Println(err)
		return
	}
	nText,err := cc.ConvertText(`保税工厂声明：本书为无限小说网(txt53.com)以下作品内容之版权与本站无任何关系`)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(nText)

}
func Test_openccFile(t *testing.T){
	localdir := filepath.Dir(os.Args[0])

	inFile,err := os.Open(localdir +"/神剑渡魔.txt")
	if err != nil{
		fmt.Println("in:",err)
		return
	}
	outFile,err := os.OpenFile(localdir +"/神剑渡魔_3.txt",os.O_CREATE|os.O_APPEND,0644)
	if err != nil{
		fmt.Println("out:",err)
		return
	}
	cc,err := NewOpenCC("s2t")
	if err != nil{
		fmt.Println("cc:",err)
		return
	}
	startTime := time.Now()
	log.Println("start...")
	err = cc.ConvertFile(inFile,outFile)
	if err != nil{
		fmt.Println("ccf",err)
		return
	}
	log.Println("end...",time.Now().Unix() - startTime.Unix())
}