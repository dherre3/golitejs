package main
//Variable Declarations
var a int 
var (
	ToBe   bool       = false
	c float64 = 123.12312
	d []int
	
)
//Type declarations
type intAlias int
func main() {
  //Printing
  print("Hello World");
  //Appending slice
  d = append(d,3)
  d = append(d,5)
  d = append(d,10)
  print(d[0],d[1],d[2])
}