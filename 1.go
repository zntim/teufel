package main
 imprort(
	"fmt"
	"os"
	"strings"
	)
	func main(){
		who :="World!"
		if len (os.Args) > 1 {
		who = strings.Join(os.Args[1:],"")
		}
		fmt.println("Hello",who)
		}