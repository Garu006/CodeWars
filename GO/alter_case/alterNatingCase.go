package kata

func ToAlternatingCase(str string) string {
	var result string
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(char - ('a' - 'A'))
		} else if char >= 'A' && char <= 'Z' {
			result += string(char + ('a' - 'A'))
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	toAlternatingCase("hello world")
	toAlternatingCase("HELLO WORLD")
	toAlternatingCase("hello WORLD")
	toAlternatingCase("HeLLo WoRLD")
	toAlternatingCase("12345")
	toAlternatingCase("1a2b3c4d5e")
	toAlternatingCase("String.prototype.toAlternatingCase")
}