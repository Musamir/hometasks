package main 
import "fmt"
func main() {
      var n,hour,minute,secons int 
	  fmt.Scan(&n)
	  hour=n/3600;
	  minute=(n-hour*3600)/60
	  secons=n-hour*3600-minute*60
      fmt.Printf(" %d   hours passed :: %d minutes passed :: %d secons passed",hour,minute,secons)   
    



}