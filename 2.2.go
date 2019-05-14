package main 
import "fmt"
func main() {
      var x int 
      fmt.Scan(&x)
      x=x/100
      if x == 1  {
      fmt.Printf(" %d  quintal",x)
      }  else  {
            fmt.Printf(" %d quintals",x)   
      }



}