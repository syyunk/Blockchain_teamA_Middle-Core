package main
import(
	"fmt"
	"bytes"
	"crypto/sha256"
)
func main(){
	s := "Hello, world!"
	h1 := sha256.Sum256([]byte(s))
	fmt.Printf("%x\n",h1)
}