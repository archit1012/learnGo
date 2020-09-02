package main


import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)



func main(){
	var s string
	fmt.Scanln(&s)	//not input just take string from terminal
	fmt.Println(s)


	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter test :")
	str, _ := reader.ReadString('\n') //single quote byte
	fmt.Println(str)

	fmt.Print("Enter number :")
	str, _ = reader.ReadString('\n') //single quote byte
	// f ,err := strconv.ParseFlow(str,64) //this is an error and cannot parse even if int is entered as it adds space
	f , err := strconv.ParseFloat(strings.TrimSpace(str),64)
	
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("value of a number:",f)
	}
}