package iteration

func Repeat(char string, amount int) string {
	var repeated string
	for i := 0; i < amount; i++ {
		repeated += char
	}
	return repeated
}
