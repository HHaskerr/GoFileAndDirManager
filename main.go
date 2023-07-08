package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
    "os"
    "path/filepath"
)

func create_dir(dirName string) {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Uyarı: dizin zaten var")
		} else {
			fmt.Println("dizin oluşturma hatası:", err)
		}
		return
	}

	fmt.Println("dizin oluşturuldu.")
}
func create_txt(fileName string) {
	_, err := os.Stat(fileName + ".txt")
	if err == nil {
		fmt.Println("Uyarı: Dosya zaten var")
		return
	}

	file, err := os.Create(fileName + ".txt")
	if err != nil {
		fmt.Printf("Dosya oluşturma hatası: %v", err)
		return
	}
	defer file.Close()

	fmt.Println("Dosya oluşturuldu.")
}
func read_txt(fileName string) {
	file, err := os.Open(fileName + ".txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Dosya bulunamadı")
		} else {
			fmt.Printf("Dosya açma hatası: %v", err)
		}
		return
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Printf("Dosya okuma hatası: %v", err)
		return
	}
}
func write_txt(fileName string) {
	_, err := os.Stat(fileName+".txt")
	if os.IsNotExist(err) {
		fmt.Println("Dosya bulunamadı")
		return
	}

	file, err := os.OpenFile(fileName+".txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Dosya açma hatası:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Lütfen bir metin girin (QUIT yazarak çıkış yapabilirsiniz): ")
		scanner.Scan()
		text := scanner.Text()

		if text == "QUIT" {
			break
		}

		_, err := file.WriteString(text + "\n")
		if err != nil {
			fmt.Println("Yazma hatası:", err)
			return
		}
	}

	fmt.Println("Metin başarıyla dosyaya eklendi.")
}
func remove_dir(dirName string){
    _, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		fmt.Println("Hata: Dizin bulunamadı")
		return
	}

	err = os.RemoveAll(dirName)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Dizin başarıyla silindi.")
}
func remove_txt(fileName string) {
	_, err := os.Stat(fileName+".txt")
	if os.IsNotExist(err) {
		fmt.Println("Hata: Dosya bulunamadı")
		return
	}

	err = os.Remove(fileName+".txt")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}
	fmt.Println("Dosya başarıyla silindi.")
}
func indir(dirName string) {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		fmt.Println("Hata: Dizin bulunamadı")
		return
	}
	err = os.Chdir(dirName)
	if err != nil {
		fmt.Println("Hata: Dizin değiştirilemedi")
		return
	}
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Hata: Şu anki dizin alınamadı")
		return
	}

	fmt.Println("Dizin başarıyla değiştirildi. Şu anki dizin:", currentDir)
}
func outdir() {
	err := os.Chdir("..")
	if err != nil {
		fmt.Println("Dizin değiştirme hatası:", err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Şu anki dizini alınamadı:", err)
		return
	}

	if currentDir == filepath.Dir(currentDir) {
		fmt.Println("En üst dizindeyiz")
	} else {
		fmt.Println("Çıkılan dizin:", currentDir)
	}
}
func list(){
    cmd := exec.Command("ls")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Komut çalıştırma hatası:", err)
		return
	}

	fmt.Println("dizin içeriği:")
	fmt.Println(string(output))
}

    
    func main() {
fmt.Printf("Programa Hoşgeldin\nDosya Komutları:\n  ctxt(create txt):yeni bir txt dosyası oluşturmanıza olanak sağlar\n  rtxt(read txt):txt dosyanızın içeriğinizi görmenize olanak sağlar\n  wtxt(write txt): txt dosyanıza yazı yazmanıza olanak sağlar\n  rmtxt(remove txt):txt dosyanızı silmenize olanak sağlar\nDizin Komutları:\n  cdir(create dir):yeni bir dizin oluşturmanıza olanak sağlar\n  indir(in dir):içinde bulunduğunuz dizinin içindeki dizinlerin içine girmenize olanak sağlar\n  outdir(outdir):bulunduğunuz dizinden bir geri gelmenize olanak sağlar\n  rmdir(remove dir):dizininizi silmenize olanak sağlar\n")
    for {
        hatirlatma:="Dosya Komutları:\n  ctxt(create txt)\n  rtxt(read txt)\n  wtxt(write txt)\n  rmtxt(remove txt)\nDizin Komutları:\n  cdir(create dir)\n  indir(in dir)\n  outdir(outdir)\n  rmdir(remove dir)\nbulunduğunuz dizinin içeriğini listelemek için list yazın\nProgramı kapatmak İçin exit yazın\n"
		fmt.Print("Komut girin(komut hatırlatması için remember yazın)\n>>>")
		var secim string
		fmt.Scanln(&secim)
		switch secim {
        case "list":
            list()
        case "remember":
            fmt.Printf("%s",hatirlatma)
		case "ctxt":
			fmt.Print("Oluşturulacak dosyanın adı: ")
			var fileName string
			fmt.Scanln(&fileName)
			create_txt(fileName)
			fileName = ""
		case "rtxt":
			fmt.Print("Okunacak dosyanın adı: ")
			var fileName string
			fmt.Scanln(&fileName)
			read_txt(fileName)
            fileName = ""
		case "wtxt":
			fmt.Print("Değiştirmek istediğiniz dosyanın adı: ")
			var fileName string
			fmt.Scanln(&fileName)
			write_txt(fileName)
			fileName = ""
		case "rmtxt":
			fmt.Print("Silinecek dosyanın adı: ")
			var fileName string
			fmt.Scanln(&fileName)
			remove_txt(fileName)
			fileName = ""
		case "cdir":
			fmt.Print("Oluşturulacak dizinin adı: ")
			var dirName string
			fmt.Scanln(&dirName)
			create_dir(dirName)
			dirName = ""
		case "indir":
			fmt.Print("Hangi dizinin içine girmek istiyorsunuz: ")
			var dirName string
			fmt.Scanln(&dirName)
			indir(dirName)
			dirName = ""
		case "outdir":
			outdir()
		case "rmdir":
			fmt.Print("Hangi dizini silmek istiyorsunuz: ")
			var dirName string
			fmt.Scanln(&dirName)
			remove_dir(dirName)
			dirName=""
        case "exit":
			fmt.Println("Program kapatılıyor...")
			os.Exit(0)
		default:
			fmt.Println("Hatalı bir giriş yaptınız. Lütfen tekrar deneyin.")
		}
	}
}
