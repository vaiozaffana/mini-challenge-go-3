package main

import (
	"fmt"
	"os"
	"strings"
)

type Person struct {
	ID                       int
	Nama                     string
	Alamat                   string
	Pekerjaan                string
	AlasanMemilihKelasGolang string
}

func getPeopleData() []Person {
	return []Person{
		{ID: 1, Nama: "Fitri", Alamat: "Jl. Lorem", Pekerjaan: "Backend", AlasanMemilihKelasGolang: "Alasan Fitri"},
		{ID: 2, Nama: "Budi", Alamat: "Jl. Ipsum", Pekerjaan: "Frontend", AlasanMemilihKelasGolang: "Alasan Budi"},
		{ID: 3, Nama: "Fahreza", Alamat: "Bondowoso", Pekerjaan: "PHP Developer", AlasanMemilihKelasGolang: "Alasan Fahreza"},
		{ID: 4, Nama: "Alif", Alamat: "Jakarta Timur", Pekerjaan: "Mahasiswa", AlasanMemilihKelasGolang: "Alasan Alif"},
		{ID: 5, Nama: "Tarmidzi", Alamat: "Jakarta Pusat", Pekerjaan: "Backend", AlasanMemilihKelasGolang: "Alasan Tarmidzi"},
		{ID: 6, Nama: "Samantha", Alamat: "Bandung", Pekerjaan: "Mahasiswi", AlasanMemilihKelasGolang: "Alasan Samantha"},
		{ID: 7, Nama: "Emma", Alamat: "Jakarta", Pekerjaan: "Data Analyst", AlasanMemilihKelasGolang: "Alasan Emma"},
		{ID: 8, Nama: "Bashori", Alamat: "Tangerang", Pekerjaan: "Fullstack", AlasanMemilihKelasGolang: "Alasan Bashori"},
		{ID: 9, Nama: "Zulfikar", Alamat: "Tangerang Selatan", Pekerjaan: "Backend", AlasanMemilihKelasGolang: "Alasan Zulfikar"},
		{ID: 10, Nama: "Syahra", Alamat: "Blitar", Pekerjaan: "Frontend", AlasanMemilihKelasGolang: "Alasan Syahra"},
		{ID: 11, Nama: "Amri", Alamat: "Solo", Pekerjaan: "Fullstack", AlasanMemilihKelasGolang: "Alasan Amri"},
		{ID: 13, Nama: "Yonkie", Alamat: "Tangerang", Pekerjaan: "Assistant Staff", AlasanMemilihKelasGolang: "Alasan Yonkie"},
		{ID: 14, Nama: "Desy", Alamat: "Jakarta", Pekerjaan: "Frontend", AlasanMemilihKelasGolang: "Alasan Desy"},
		{ID: 15, Nama: "Assyifa", Alamat: "Depok", Pekerjaan: "Mahasiswa", AlasanMemilihKelasGolang: "Alasan Assyifa"},
		{ID: 16, Nama: "Mega", Alamat: "Karawang", Pekerjaan: "Admin Staff", AlasanMemilihKelasGolang: "Alasan Mega"},
		{ID: 17, Nama: "Vaio", Alamat: "Sidoarjo", Pekerjaan: "Siswa", AlasanMemilihKelasGolang: "Alasan Vaio"},
	}
}

func displayPerson(people []Person, query string) {
	found := false
	for _, person := range people {
		if strings.EqualFold(fmt.Sprint(person.ID), query) || strings.EqualFold(person.Nama, query) {
			fmt.Printf("ID      : %d\n", person.ID)
			fmt.Printf("Nama    : %s\n", person.Nama)
			fmt.Printf("Alamat  : %s\n", person.Alamat)
			fmt.Printf("Pekerjaan: %s\n", person.Pekerjaan)
			fmt.Printf("Alasan  : %s\n", person.AlasanMemilihKelasGolang)
			fmt.Println("--------------------------")
			found = true
		}
	}
	if !found {
		fmt.Println("Data dengan nama/ID tersebut tidak tersedia")
	}
}

func main() {
	people := getPeopleData()

	if len(os.Args) < 2 {
		fmt.Println("Tolong masukkan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go Fitri' atau 'go run main.go 2'")
		return
	}

	query := os.Args[1]
	displayPerson(people, query)
}
