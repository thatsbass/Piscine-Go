package piscine

func PrintNbrBase(nb int, base string) {
	R := []rune(base)
	pM := false
	for i := range R {
		if string(R[i]) == "-" || string(R[i]) == "+" {
			pM = true
			break
		} else {
			pM = false
		}
	}
	if pM {
		PrintStr("NV")
	} else if nb == -9223372036854775808 {
		PrintStr("-9223372036854775808")
	} else if len(base) >= 2 {
		r := []rune(base)
		result := ""
		nb2 := nb
		single := true
		for _, motif := range r {
			if IsSingle(string(motif), base) {
				single = true
			} else {
				single = false
				break
			}
		}

		if single {
			if nb2 < 0 {
				nb2 = -nb2
			}
			for !(nb2 == 0) {
				b := nb2 % len(base)
				result = string(r[b]) + result
				a := nb2 / len(base)
				nb2 = a
			}
			if nb < 0 {
				PrintStr("-")
			}
			PrintStr(result)
		} else {
			PrintStr("NV")
		}
	} else {
		PrintStr("NV")
	}
}

func IsSingle(motif, s string) bool {
	cpt := 0
	r := []rune(s)
	for i := range r {
		if motif == string(r[i]) {
			cpt++
		}
		if cpt <= 1 && i == len(r)-1 {
			return true
		}
	}
	return false
}
