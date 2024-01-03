package main

// countBits adalah fungsi yang digunakan untuk menghitung jumlah bit yang bernilai 1 pada suatu bilangan.
// Fungsi ini menerima parameter n bertipe uint32 dan mengembalikan jumlah bit yang bernilai 1 bertipe int32.
func countBits(n uint32) int32 {
	count := int32(0)
	for n > 0 {
		count += int32(n & 1) // Mengkonversi hasil dari n & 1 menjadi int32
		n >>= 1
	}
	return count
}

func main() {
	println(countBits(127))
}
