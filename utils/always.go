package utils

func AlwaysCommand(always string, level int) string {
	if always == "1" && level == 2 {
		return " --always"
	}
	return ""
}
