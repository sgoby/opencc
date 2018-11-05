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
	Dicts  []*Dict `json:"dicts"`
	CfgMap map[string][]string
	maxLen int //最大长度
	minLen int //最小长度
}

type Segmentation struct {
	Type string `json:"type"`
	Dict *Dict  `json:"dict"`
}

type ConversionChain struct {
	Dict *Dict `json:"dict"`
}

type Config struct {
	Name            string             `json:"name"`
	Segmentation    Segmentation       `json:"segmentation"`
	ConversionChain []*ConversionChain `json:"conversion_chain"`
}

func (c *Config) init() error {
	var err error
	for _, cv := range c.ConversionChain {
		err = cv.init()
		if err != nil {
			return err
		}
	}
	return nil
}
func (cv *ConversionChain) init() error {
	return cv.Dict.init()
}

//
func (d *Dict) init() (err error) {
	if len(d.File) > 0 {
		d.CfgMap, d.maxLen, d.minLen, err = d.File.readFile()
		//fmt.Println("File = ", string(d.File), d.maxLen, d.minLen)
		if err != nil {
			return err
		}
	}
	//
	if d.Dicts != nil && len(d.Dicts) > 0 {
		for _, childDict := range d.Dicts {
			err = childDict.init()
			if err != nil {
				return
			}
		}
	}
	return nil
}

//
func (fo *FileOCD) readFile() (map[string][]string, int, int, error) {
	fileName := dataDir + "/dictionary/" + string(*fo)
	f, err := os.Open(fileName)
	if err != nil {
		return nil, 0, 0, err
	}
	cfgMap := make(map[string][]string)
	buf := bufio.NewReader(f)
	//
	max := 0;
	min := 0;
	//
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			return cfgMap, max, min, nil
		}
		fields := strings.Fields(line)
		if len(fields) > 1 {
			if len([]rune(fields[0])) > max {
				max = len([]rune(fields[0]))
			}
			if min <= 0 || len([]rune(fields[0])) < min {
				min = len([]rune(fields[0]))
			}
			cfgMap[fields[0]] = fields[1:]
		}
	}
	return cfgMap, max, min, nil
}

//=============================================================
func (c *Config) convertText(text string) (string, error) {
	var err error
	for _, cv := range c.ConversionChain {
		text, err = cv.convertText(text)
		if err != nil {
			return text, err
		}
	}
	return text, nil
}
func (c *ConversionChain) convertText(text string) (string, error) {
	var err error
	text, err = c.Dict.convertTextWithMap(text)
	if err != nil {
		return text, err
	}
	return text, nil
}

//
func (d *Dict) convertTextWithMap(text string) (string, error) {
	var err error
	newText := text
	runes := []rune(text)
	//
	if d.CfgMap != nil {
		if len(runes) < d.minLen {
			return text, nil
		}
		//
		maxL := d.maxLen;
		if maxL > len(runes) {
			maxL = len(runes)
		}
		//
		for i := maxL; i >= d.minLen; i-- {
			for j := 0; j <= len(runes)-i; j++ {
				if i == 0 || j+i > len(runes) {
					continue
				}
				old := string(runes[j:j+i]);
				if newStr, ok := d.CfgMap[old]; ok {
					newText = strings.Replace(newText, old, newStr[0], 1)
					j = j + i - 1;
				}
			}
		}
	}
	//
	for _, cd := range d.Dicts {
		newText, err = cd.convertTextWithMap(newText)
		if err != nil {
			return text, err
		}

	}
	return newText, nil
}
