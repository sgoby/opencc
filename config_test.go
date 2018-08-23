package opencc

import (
	"testing"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"path/filepath"
	"log"
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
	nText,err := cc.ConvertText(`迪拜（阿拉伯语：دبي，英语：Dubai），是阿拉伯联合酋长国人口最多的城市，位于波斯湾东南海岸，迪拜也是组成阿联酋七个酋长国之一——迪拜酋长国的首都。`)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(nText)

}
func Test_openccFile(t *testing.T){
	localdir := filepath.Dir(os.Args[0])

	inFile,err := os.Open(localdir +"/online.txt")
	if err != nil{
		fmt.Println("in:",err)
		return
	}
	outFile,err := os.OpenFile(localdir +"/online_2.txt",os.O_CREATE|os.O_APPEND,0644)
	if err != nil{
		fmt.Println("out:",err)
		return
	}
	cc,err := NewOpenCC("s2t")
	if err != nil{
		fmt.Println("cc:",err)
		return
	}
	log.Println("start...")
	err = cc.ConvertFile(inFile,outFile)
	if err != nil{
		fmt.Println("ccf",err)
		return
	}
	log.Println("end...")
}