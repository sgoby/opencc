package opencc

import (
	"os"
	"bufio"
	"strings"
)

const (
	TYPE_OCD   = "ocd"
	TYPE_GROUP = "group"
)

var dataDir string

type FileOCD string

type Dict struct {
	Type   string  `json:"type"`
	File   FileOCD `json:"file"`
	Dicts  []*Dict  `json:"dicts"`
	CfgMap map[string][]string
}

type Segmentation struct {
	Type string `json:"type"`
	Dict *Dict   `json:"dict"`
}

type ConversionChain struct {
	Dict *Dict `json:"dict"`
}

type Config struct {
	Name            string            `json:"name"`
	Segmentation    Segmentation      `json:"segmentation"`
	ConversionChain []*ConversionChain `json:"conversion_chain"`
}

func (c *Config)convertText(text string) (string, error) {
	var err error
	for _,c := range c.ConversionChain{
		text,err = c.convertText(text)
		if err != nil{
			return text,err
		}
	}
	return text,nil
}
func (c *ConversionChain)convertText(text string) (string, error) {
	var err error
	text,err = c.Dict.convertTextWithMap(text)
	if err != nil{
		return text,err
	}
	return text,nil
}

//
func (d *Dict) getCfgMap() (map[string][]string, error) {
	if d.CfgMap != nil {
		return d.CfgMap, nil
	}
	return d.File.readFile()
}

//
func (fo *FileOCD) readFile() (map[string][]string, error) {
	fileName := dataDir + "/dictionary/" + string(*fo)
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	cfgMap := make(map[string][]string)
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Fields(line)
		if len(fields) > 1 {
			cfgMap[fields[0]] = fields[1:]
		}
	}
	return cfgMap, nil
}

//
func (d *Dict) convertTextWithMap(text string) (string, error) {
	var err error
	if d.CfgMap == nil {
		d.CfgMap,err = d.File.readFile()
		if err != nil{
			return text,err
		}
	}
	//
	for key,val := range d.CfgMap{
		if len(val) > 0 {
			text = strings.Replace(text, key, val[0], -1)
		}
	}
	//
	for _,cd := range d.Dicts{
		text,err = cd.convertTextWithMap(text)
		if err != nil{
			return text,err
		}
	}
	return text,nil
}
