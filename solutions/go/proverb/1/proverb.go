
package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
    if len(rhyme) == 0 {
        return []string{}
    }
    
	Proverbs := []string{}
    if len(rhyme) == 1 {
        Proverbs = append(Proverbs, "And all for the want of a " + rhyme[0] + ".")
        return Proverbs
    }
	
	for i := 0 ; i < len(rhyme) - 1; i++ {
            Proverb := "For want of a " + rhyme[i] + " the " + rhyme[i + 1] + " was lost."
        	Proverbs = append(Proverbs, Proverb)
    }
    lastP := "And all for the want of a " + rhyme[0] + "."
    Proverbs = append(Proverbs, lastP) 
    return Proverbs
}
