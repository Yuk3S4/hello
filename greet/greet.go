package greet

// Variables a nivel de paquetes deben ser declaradas con var
var emoji = "🙋‍♂️"

// identificadores con letra mayúscula se exportan y con letra minúscula no se exportan

// Retorna saludo en ingles
func English() string {
	return "Hi " + emoji
}

// Retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
