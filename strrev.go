package piscine

func StrRev(s string) string {
	// Преобразуем строку в срез рун для обработки Unicode
	runes := []rune(s)
	// Обходим срез рун и меняем местами символы
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Возвращаем строку, сформированную из развернутого среза рун
	return string(runes)
}
