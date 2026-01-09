package strand

func ToRNA(dna string) string {
    RNA := ""
	for _, v := range dna{
        switch {
            case v == 'G':
            	RNA += "C"
            case v == 'C':
            	RNA += "G"
            case v == 'T':
            	RNA += "A"
            case v == 'A':
            	RNA += "U"
        }
    }
    	return RNA
}
