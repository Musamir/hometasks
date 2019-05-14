package main 
import "fmt"
func main() {
      var k,b,l,s int 
      var n string
      fmt.Scan(&k)
	  s=k%7
	  b=s
	  l=s+1
	  if(l == 7) {l=0}
	  if(s==0){ n="Sunday" }
      if(s==1){ n="Monday " }
      if(s==2){ n="Tuesday " }
      if(s==3){ n="Wednesday " }
      if(s==4){ n="Thursday"}
      if(s==5){ n="Friday " }
      if(s==6){ n="Saturday" }
	  fmt.Println(b)
	  fmt.Println(l)
	  fmt.Println(" if January 1 is Monday then   ")
	  fmt.Println(" day   ",k, " is ",n)

	  


   



}