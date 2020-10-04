package main
import "fmt"
func main(){
	var n,m int
	fmt.Print("Masukkan Bilangan 1 : ")
	fmt.Scan(&n)
	fmt.Print("Masukkan Bilangan 2 : ")
	fmt.Scan(&m)
	for i:=0;i<n;i++{
		for j:=i;j<m;j++{
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
