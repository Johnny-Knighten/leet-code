package main

func isValid(s string) bool {
	paras := []string{}
	for _, char := range s {
		strChar := string(char)
		if strChar == "{" || strChar == "(" || strChar == "[" {
			paras = append(paras, strChar)
		} else {

			if len(paras) == 0 {
				return false
			}

			if strChar == "}" && paras[len(paras)-1] != "{" {
				return false
			}

			if strChar == "]" && paras[len(paras)-1] != "[" {
				return false
			}

			if strChar == ")" && paras[len(paras)-1] != "(" {
				return false
			}

			paras = paras[:len(paras)-1]
		}
	}

	return len(paras) == 0
}
