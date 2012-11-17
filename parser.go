package main

func parser(s string) ([]string, string) {
	const STATE_NULL int = 0
	const STATE_WORD int = 1
	const STATE_STR int = 2

	state := STATE_NULL
	var bf [100]string
	var bfi int
	slen := len(s)
	sii := 0
	si := 0
	for si = 0; si < slen; si = si + 1 {
		ch := string(s[si])
		if ch != " " && ch != "\"" && ch != "\n" && state == STATE_NULL {
			sii = si
			state = STATE_WORD
		}
		if (ch == " " || ch == "\"" || ch == "\n") && state == STATE_WORD {
			bf[bfi] = s[sii:si]
			bfi = bfi + 1
			state = STATE_NULL
		}
		if ch == "\"" && state == STATE_NULL {
			sii = si + 1
			state = STATE_STR
		} else if ch == "\"" && state == STATE_STR {
			bf[bfi] = s[sii:si]
			bfi = bfi + 1
			state = STATE_NULL
		}
	}
	if state != STATE_NULL {
		bf[bfi] = s[sii:si]
		bfi = bfi + 1
		state = STATE_NULL
	}
	str := bf[0:bfi]
	return str, s
}
