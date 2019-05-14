package main 
import "fmt"
func main() {
      var n int 
      fmt.Scan(&n)
	  fmt.Println("number of tens in it  ",n/10," :: number of units in it  ",n%10," :: sum of its digits ",n/10+n%10," :: product of its digits  ",(n/10)*(n%10))  
    



}