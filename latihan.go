package main

import "fmt"

func uang(masukan int) {
    var x, y, z, x1, y1, z1 int
    x = 10000
    y = 2000
    z = 1000

    for masukan > 0 {
        if masukan >= x {
            x1 = masukan / x // Hitung jumlah lembar 10000
            masukan %= x     // Sisa uang setelah diambil pecahan 10000
        }
        if masukan >= y {
            y1 = masukan / y // Hitung jumlah lembar 2000
            masukan %= y     // Sisa uang setelah diambil pecahan 2000
        }
        if masukan >= z {
            z1 = masukan / z // Hitung jumlah lembar 1000
            masukan %= z     // Sisa uang setelah diambil pecahan 1000
        }
    }

    fmt.Println(x1, "Lembar 10000")
    fmt.Println(y1, "Lembar 2000")
    fmt.Println(z1, "Lembar 1000")
}

func main() {
    var masukan1 int
    fmt.Print("Masukkan jumlah uang: ")
    fmt.Scan(&masukan1)
    uang(masukan1)
}