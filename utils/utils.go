// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

// ******************************************************************************
// Matronator (c) 2024.                                                         *
// ******************************************************************************

package utils

func IsNum(s string) bool {
	dotFound := false

	for _, v := range s {
		if v == '.' {
			if dotFound {
				return false
			}

			dotFound = true
		} else if v < '0' || v > '9' {
			return false
		}
	}

	return true
}

func Zip(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))

	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}

	return r
}
