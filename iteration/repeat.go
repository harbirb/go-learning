package iteration

func Repeat(character string, times int) string {
	// declaring string variable first, not using shortcut :=
	var repeated string
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
