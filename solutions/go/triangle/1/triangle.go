package triangle

type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind= "NaT"// not a triangle
	Equ Kind= "Equ"// equilateral
	Iso Kind= "Iso"// isosceles
	Sca Kind= "Sca"// scalene
)

func KindFromSides(a, b, c float64) Kind {
    var k Kind
    if (a <= 0 || b <= 0 || c <= 0 ) || (a + b < c || a + c < b || b + c < a){
        k = NaT
    	return k
    }
    
	switch {
        case a == b && b == c:
        	k = Equ
   
        case a == b || a == c || b == c:
        	k = Iso

        default :
        	k = Sca
    }
	
	return k
}
