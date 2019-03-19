package space

//Planet provide type planet
type Planet string

//Age Function calculate second in universe
func Age(s float64, p Planet) float64 {
	sp := s / 31557600
	switch p {
	case "Earth":

	case "Mercury":
		sp = sp / 0.2408467

	case "Venus":
		sp = sp / 0.61519726

	case "Mars":
		sp = sp / 1.8808158

	case "Jupiter":
		sp = sp / 11.862615

	case "Saturn":
		sp = sp / 29.447498

	case "Uranus":
		sp = sp / 84.016846

	case "Neptune":
		sp = sp / 164.79132
	}
	return sp
}
