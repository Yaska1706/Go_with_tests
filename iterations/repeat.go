package iterations

func Repeat(char string, count int) string {

	var repeated string

	for i := 0; i < count; i++ {
		repeated = repeated + char
	}
	return repeated
}
