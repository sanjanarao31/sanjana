package main
import (
"fmt"
)
func main(){
   //reading an integer
    var x,y int
    fmt.Println("Enter two numbers")
    fmt.Scanf("%d",&x)
    fmt.Scanf("%d",&y)
     if x<=y{
   fmt.Println(x, " is lesser than or equal ",y)
   fmt.Printf("%T\n",x) //Prints datatype of x
 }else{   fmt.Println("  is greater than ",y,x)

}
}
