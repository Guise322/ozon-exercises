package infra

const (
	yandexStr = "yandex.ru"
	yaStr     = "ya.ru"
)

func YandexCheck(email string) bool {
	emailRunes := []rune(email)
	emailLen := len(emailRunes)
	if isYa(email, emailRunes, emailLen) {
		return true
	}
	if isYandex(email, emailRunes, emailLen) {
		return true
	}
	return false
}

func isYandex(email string, emailRunes []rune, emailLen int) bool {
	yandexRunes := []rune(yandexStr)
	yandexLen := len(yandexRunes)
	runesToComp := emailRunes[emailLen-yandexLen : emailLen-1]
	if emailLen > yandexLen {
		for i, r := range runesToComp {
			if r != yandexRunes[i] {
				return false
			}
		}
		return true
	}
	return false
}

func isYa(email string, emailRunes []rune, emailLen int) bool {
	yaRunes := []rune(yaStr)
	yaLen := len(yaRunes)
	runesToComp := emailRunes[emailLen-yaLen : emailLen-1]
	if emailLen > yaLen {
		for i, r := range runesToComp {
			if r != yaRunes[i] {
				return false
			}
		}
		return true
	}
	return false
}
