package utils

var VALIDARGS = []string{"11", "22", "33"}

func IsValidArg(arg string) bool {
	for _, va := range VALIDARGS {
		if arg == va {
			return true
		}
	}
	return false
}

func GetExpectedDestination(rkey string) string {
	if rkey == "11" {
		return "CONSUMER#1"
	} else if rkey == "22" {
		return "CONSUMER#2"
	} else {
		return "CONSUMER#3"
	}
}
