package src

import ("fmt"
"regexp")

// func DetailsRegex(name string, sourceStr string) (regexRes string){
// 	regexs := fmt.Sprintf("sid: (\d+), qcat: ''})'>%s", name)
// 	fmt.Println(regexs)
// 	r1:=regexp.MustCompile(regexs)
// 	b1:=r1.FindString(sourceStr)
// 	fmt.Println(b1)
// 	return b1
// }

func compressStr(str string) string {
    if str == "" {
        return ""
    }
    //匹配一个或多个空白符的正则表达式
    reg := regexp.MustCompile("\\s+")
    return reg.ReplaceAllString(str, "")
}

func main(){
	// url := fmt.Sprintf("sda%v", 64)
	fmt.Println(compressStr("  nihao"))

}