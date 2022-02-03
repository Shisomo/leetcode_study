package main

func main() {
	data := "oring"
	reversePrefix(data, ' ')
}

func reversePrefix(word string, ch byte) string {
	ex := []byte(word)
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			for j := 0; j <= i/2; j++ {
				ex[j], ex[i-j] = ex[i-j], ex[j]
			}
			break
		}
	}
	return string(ex)
}
