package main
import "fmt"
var nilai float64
func main(){
	fmt.Print("Masukkan Sebuah Bilangan : ")
	fmt.Scan(&nilai)
	decimalNumber := nilai
	fmt.Println("Nilai Bilangan Integer : ", nilai)
	fmt.Printf("Nilai Bilangan Float : %f\n", decimalNumber)
	fmt.Printf("Nilai Bilangan Float Khusus : %.3f\n", decimalNumber)

}