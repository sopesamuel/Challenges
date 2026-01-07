package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    result := float64(0.0)
	switch {
        case planet == "Mercury":
        	result = float64(seconds) / (0.2408467 * float64(31557600))
        case planet == "Venus":
        	result = float64(seconds) / (0.61519726 * float64(31557600))
        case planet == "Earth":
        	result = float64(seconds) / (1.0 * float64(31557600))
        case planet == "Mars":
        	result = float64(seconds) / (1.8808158 * float64(31557600))
        case planet == "Jupiter":
        	result = float64(seconds) / (11.862615 * float64(31557600))
        case planet == "Saturn":
        	result = float64(seconds) / (29.447498 * float64(31557600))
        case planet == "Uranus":
       		result = float64(seconds) / (84.016846 * float64(31557600))
        case planet == "Neptune":
        	result = float64(seconds) / (164.79132 * float64(31557600))
        default:
        	result = -1.00
    }
    return result
}

