package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; KSN")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))

	for i := 0; i < len(melding); i++ {
		indeks := sokIAlfabetet(melding[i], alphabet) //fmt alphabet nummer
		if indeks+chiffer >= len(alphabet) {          // if  0 + 4 >=26
			kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]

		} else {
			kryptertMelding[i] = alphabet[indeks+chiffer] // 0 + 4
		}
	}
	return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
			break
		}
	}
	return -1
}
