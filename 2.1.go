package main 
import "fmt"
func main() {
      var x int 
      fmt.Scan(&x)
      x=x/100
      if x == 1  {
      fmt.Printf(" %d  meter",x)
      }  else  {
            fmt.Printf(" %d meters",x)   
      }



}