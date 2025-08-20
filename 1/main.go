package main

func main() {
	a := "Gopher"

	b := []byte(a)
	b[0] = 'T'
	println(string(b)) // 输出: "Topher"
}
