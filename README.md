# OpenCC 中文简繁体转换

基于OpenCC中文简繁体转换的golang开发包

#### 转换类型

* `s2t` Simplified Chinese to Traditional Chinese 简体到繁體
* `t2s` Traditional Chinese to Simplified Chinese 繁體到简体
* `s2tw` Simplified Chinese to Traditional Chinese (Taiwan Standard) 简体到臺灣正體
* `tw2s` Traditional Chinese (Taiwan Standard) to Simplified Chinese 臺灣正體到简体
* `s2hk` Simplified Chinese to Traditional Chinese (Hong Kong Standard) 简体到香港繁體（香港小學學習字詞表標準）
* `hk2s` Traditional Chinese (Hong Kong Standard) to Simplified Chinese 香港繁體（香港小學學習字詞表標準）到简体
* `s2twp` Simplified Chinese to Traditional Chinese (Taiwan Standard) with Taiwanese idiom 简体到繁體（臺灣正體標準）並轉換爲臺灣常用詞彙
* `tw2sp` Traditional Chinese (Taiwan Standard) to Simplified Chinese with Mainland Chinese idiom 繁體（臺灣正體標準）到简体並轉換爲中國大陸常用詞彙
* `t2tw` Traditional Chinese (OpenCC Standard) to Taiwan Standard 繁體（OpenCC 標準）到臺灣正體
* `t2hk` Traditional Chinese (OpenCC Standard) to Hong Kong Standard 繁體（OpenCC 標準）到香港繁體（香港小學學習字詞表標準）

### 使用

```go
import "github.com/sgoby/opencc"

...

//简体到繁體（臺灣正體標準）
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
// 输出:
// 杜拜（阿拉伯語：دبي，英語：Dubai），是阿拉伯聯合大公國人口最多的城市，位於波斯灣東南海岸，杜拜也是組成阿聯酋七個酋長國之一——杜拜酋長國的首都。
```