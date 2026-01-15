package bob
import "strings"
// Hey should have a comment documenting it.

func Hey(remark string) string {
    remark = strings.TrimSpace(remark)
	main := ""

    status := false
    for _, v := range remark {
        if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z'{
            status = true
        }
    }
	switch{
        case remark == "":
        	main = "Fine. Be that way!"
        case (status == true && remark == strings.ToUpper(remark) ) && remark[len(remark) - 1] == '?':
        	main = "Calm down, I know what I'm doing!"
        case remark[len(remark) - 1] == '?':
        	main = "Sure."
        case status == true && remark == strings.ToUpper(remark):
        	main = "Whoa, chill out!"
        default:
        	main = "Whatever."
    }

	return main
}
