package main
import "fmt"
func main(){
    var a,win int
    var b,c string
    fmt.Scan(&a)
    for i:= 0; i < a; i++ {
        fmt.Scan(&b,&c)
        if (b=="C"&&c=="P")||(b=="G"&&c=="C")||(b=="P"&&c=="G") {
            win = win + 1
        }
    }
    fmt.Println(win)
}
