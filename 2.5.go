package main 
import "fmt"
func main() {
	  var x int 
	  x=234 
      fmt.Scan(&x)
      x=x/7
      if x == 1  {
      fmt.Printf(" %d  week",x)
      }  else  {
            fmt.Printf(" %d weeks",x)   
      }



}