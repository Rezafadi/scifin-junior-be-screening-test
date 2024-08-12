package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Map yang menghubungkan setiap huruf dengan representasi Braille-nya.
var brailleMap = map[rune]string{
	'a': "100000", 'b': "101000", 'c': "110000", 'd': "110100", 'e': "100100",
	'f': "111000", 'g': "111100", 'h': "101100", 'i': "011000", 'j': "011100",
	'k': "100010", 'l': "101010", 'm': "110010", 'n': "110110", 'o': "100110",
	'p': "111010", 'q': "111110", 'r': "101110", 's': "011010", 't': "011110",
	'u': "100011", 'v': "101011", 'w': "011101", 'x': "110011", 'y': "110111",
	'z': "100111",
}

// Karakter Braille untuk penanda huruf kapital.
const capitalPrefix = "000001"

// Karakter Braille untuk spasi.
const space = "000000"

func BrailleTranslation(input string) string {
	var result strings.Builder

	for _, char := range input {
		if char == ' ' {
			// Tambahkan Braille untuk spasi.
			result.WriteString(space)
		} else if char >= 'A' && char <= 'Z' {
			// Tambahkan penanda kapital dan Braille untuk huruf kapital.
			result.WriteString(capitalPrefix)
			result.WriteString(brailleMap[char+('a'-'A')]) // Konversi ke huruf kecil
		} else if char >= 'a' && char <= 'z' {
			// Tambahkan Braille untuk huruf kecil.
			result.WriteString(brailleMap[char])
		}
	}

	return result.String()
}

func main() {
	// Membaca input dari pengguna.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan teks: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := BrailleTranslation(input)

	fmt.Println("Output :", result)
}
