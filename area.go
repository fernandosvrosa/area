package area

import "math"

// Pi é uma proporção numérica defina pela relação entre
// o perimetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ é reponsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é reponsavel por calcular a area de uma retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é visivel!
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
