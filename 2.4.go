package main 
import "fmt"
func main() {
      var x int 
      fmt.Scan(&x)
      x=x/1000
      if x == 1  {
      fmt.Printf(" %d  kilometer",x)
      }  else  {
            fmt.Printf(" %d kilometers",x)   
      }



}