package parsinglogfiles
import "regexp"

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
 	re := regexp.MustCompile(`<[~*=-]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*password.*"`)
	count := 0
    for _, v := range lines {
      if re.MatchString(v){
          count++ 
      }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`\sUser\s+(\S+)`)
    var result []string
    for _, v := range lines {
        find := re.FindStringSubmatch(v)
        if len(find) > 1 {
            v = "[USR] " + find[1] + " " + v
        }
        result = append(result, v)
    }
    return result
}
