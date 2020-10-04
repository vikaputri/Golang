package main
import "fmt"
func main(){
	var pesan1 string = "halo"
	var pesan2 = `Nama Saya "Dewi".
	Salam Kenal
	Mari belajar bersama-sama tentang "Golang".`

	fmt.Println(pesan1+pesan2)
}