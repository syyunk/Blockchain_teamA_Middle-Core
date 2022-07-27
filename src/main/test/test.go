package main
import(
	"fmt"
)
func main(){
	var a *int // 포인터 변수. a의 역참조값은 int여야 함
	b := 3
	a = &b // 주소를 할당해줘야 함
    
	fmt.Println(a) // 주소
	fmt.Println(*a) // 역참조
}