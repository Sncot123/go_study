package main
import(
	"fmt"
)
//1、编写一个回文检测函数，并为其编写单元测试和基准测试，根据测试的结果逐步对其进行优化。
//（回文：一个字符串正序和逆序一样，如“Madam,I’mAdam”、“油灯少灯油”等。）


func main(){
	s :=[]string{"qe","a","b","c","d","e","f","g"}
	for i := range s{
		fmt.Println(i)
	}
}