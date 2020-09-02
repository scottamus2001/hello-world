package locale

const (
	// ENGLISH language code
	ENGLISH = "eng"
	// ITALIAN language code
	ITALIAN = "ita"
	// SPANISH language code
	SPANISH = "esp"
)

// Hello prints
func Hello(language string) string {
	switch language {
	case ENGLISH:
		return "hello"
	case ITALIAN:
		return "ciao"
	case SPANISH:
		return "hola"
	}

	return ""
}
