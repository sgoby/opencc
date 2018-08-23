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
	cc,err := NewOpenCC("s2t")
	if err != nil{
		fmt.Println(err)
		return
	}
	nText,err := cc.ConvertText(`台湾（Taiwan），位于中国大陆东南沿海的大陆架上，东临太平洋，东北邻琉球群岛，南界巴士海峡与菲律宾群岛相对，西隔台湾海峡与福建省相望， [1]  总面积约3.6万平方千米，包括台湾岛及兰屿、绿岛、钓鱼岛等21个附属岛屿和澎湖列岛64个岛屿。台湾岛面积35882.6258平方千米，是中国第一大岛， [2]  7成为山地与丘陵，平原主要集中于西部沿海，地形海拔变化大。由于地处热带及亚热带气候之交界，自然景观与生态资源丰富。人口约2350万，逾7成集中于西部5大都会区，其中以首要都市台北为中心的台北都会区最大。`)
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