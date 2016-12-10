package advent_of_code

func isABBA(sequence string) bool {
	return sequence[0] == sequence[3] && sequence[1] == sequence[2] && sequence[0] != sequence[1]
}

func containsABBA(sequence string) bool {
	bound := len(sequence)

	for i := 1; i < bound-2; i++ {
		if(sequence[i+2] == ']') {
			return false
		}

		if(isABBA(sequence[i-1:i+3])) {
			return true
		}
	}
	return false
}

func supportTLS(hypernetSplit []string, supernetSplit []string) bool {
	for _, hypernetSequence := range hypernetSplit {
		if(containsABBA(hypernetSequence)) {
			return false
		}
	}

	for _, supernetSequence := range supernetSplit {
		if(containsABBA(supernetSequence)) {
			return true
		}
	}
	return false
}


func isABA(sequence string) bool {
	return sequence[0] == sequence[2] && sequence[0] != sequence[1]
}

func getABAs(split [] string) *List {
	list := NewList()
	for _, sequence := range split {
		bound := len(sequence)

		for i := 1; i < bound - 1; i++ {
			if(sequence[i+1] == ']') {
				break
			}
			slice := sequence[i-1:i+2]
			if(isABA(slice)) {
				list.Add(slice)
			}
		}
	}

	return list
}

func areCorresponding(aba string, bab string) bool {
	return aba[0] == bab[1] && aba[1] == bab[0]
}

func supportSSL(hypernetSplit []string, supernetSplit []string) bool {
	ABAs := getABAs(supernetSplit)
	BABs := getABAs(hypernetSplit)
	lenABAs := ABAs.Size()
	lenBABs := BABs.Size()

	for i := 0; i < lenABAs; i++ {
		for j := 0; j < lenBABs; j++ {
			aba, _ := ABAs.Get(i)
			bab, _ := BABs.Get(j)

			if(areCorresponding(aba.(string), bab.(string))) {
				return true
			}
		}
	}
	return false
}


func supportsTLSSL(hypernetSplit []string, supernetSplit []string, tlsQ bool, sslQ bool) bool {
	tls := false
	ssl := false

	if(tlsQ) {
		tls = supportTLS(hypernetSplit, supernetSplit)
	} else if(sslQ) {
		ssl = supportSSL(hypernetSplit, supernetSplit)
	}

	if(sslQ && tlsQ) {
		return ssl && tls
	} else if(sslQ) {
		return ssl
	} else if(tlsQ) {
		return tls
	}

	return false
}

func NumSupportingIps(inputString string, TLS bool, SSL bool) int {
	lines := GetLines(inputString)

	count := 0
	for _, ip := range lines {
		hypernetSplit := RegSplit(ip, "[a-z]+\\[")
		supernetSplit := RegSplit(ip, "\\[[a-z]+\\]")

		if(supportsTLSSL(hypernetSplit, supernetSplit, TLS, SSL)) {
			count++
		}
	}
	return count
}

