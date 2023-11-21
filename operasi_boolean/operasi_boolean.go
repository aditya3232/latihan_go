package main

import "fmt"

func main() {
	/*
		- operasi boolean itu untuk membandingkan boolean
		- hasilnya pun juga boolean
		- &&
		- ||
		- !
	*/

	var (
		rataRataNilaiAkhirAdit = 80
		absensiAdit            = 90

		// penilaian kelulusan
		lulusNilaiAkhir bool = rataRataNilaiAkhirAdit >= 80
		lulusAbsensi    bool = absensiAdit >= 80

		penentualKelulusan bool = lulusNilaiAkhir && lulusAbsensi // syarat lulus adalah sesuai variabel penilaian kelulusan
	)

	fmt.Println(penentualKelulusan)

}
