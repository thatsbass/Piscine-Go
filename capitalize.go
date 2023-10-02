package piscine

func Capitalize(s string) string {
	r := []rune(s)
	if IsAlpha2(string(r[0])) {
		if IsLower(string(r[0])) {
			a := ToUpper(string(r[0]))
			R := []rune(a)
			r[0] = R[0]
		}
	}
	a := 0
	for a != len(r)-1 {
		if !IsAlpha2(string(r[a])) {
			b := ToUpper(string(r[a+1]))
			R := []rune(b)
			r[a+1] = R[0]
		}
		if a != 0 && IsUpper(string(r[a])) && IsAlpha(string(r[a-1])) {
			d := ToLower(string(r[a]))
			R := []rune(d)
			r[a] = R[0]
		}
		if IsUpper(string(r[a])) {
			c := ToLower(string(r[a+1]))
			R := []rune(c)
			r[a+1] = R[0]
		}
		a++
		if !IsAlpha(string(r[a-1])) && a == len(r)-1 {
			c := ToUpper(string(r[a]))
			R := []rune(c)
			r[a] = R[0]
		} else if IsAlpha(string(r[a-1])) && a == len(r)-1 {
			c := ToLower(string(r[a]))
			R := []rune(c)
			r[a] = R[0]
		}
	}
	return string(r)
}

func IsAlpha2(s string) bool {
	cpt := len(s)
	r := []rune(s)
	for i := range r {
		if r[i] >= 65 && r[i] <= 90 || r[i] >= 97 && r[i] <= 122 {
			cpt--
		}
		if cpt == 0 {
			return true
		}
	}
	return false
}
