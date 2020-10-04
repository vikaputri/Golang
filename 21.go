package main
import "fmt"
func main(){
	for i:=1;i<10;i++{
		fmt.Print("Masukkan Bilangan : ")
		fmt.Scan(&i)
		if i<=10{
			if i%2==0{
				fmt.Print(i, " adalah Bilangan Genap \n")
			}else{
				fmt.Print(i, " adalah Bilangan Ganjil \n")
			}
		}else{
			fmt.Print("Pertanyaan selesai, karena angka sudah lebih dari 10.Terima Kasih \n")
		}
	}
}