package main
import(
	"strings"
    "fmt"
)
func isPalindrome(s string) bool {
    var temp []byte
    s = strings.ToLower(s)
    fmt.Println(s)
    for i:=0;i<len(s);i++{
        if 'a'> s[i] || s[i] > 'z'{
            continue
        }
        temp = append(temp,s[i])
    }
    fmt.Println(temp)
    m,n := 0,len(temp)-1
    for  m<=n {
        if temp[m]!=temp[n]{
            return false
        }
        m++
        n--
    }
    return true
}
func main() {
    res := isPalindrome("A man, a plan, a canal: Panama")
    fmt.Println(res)
}