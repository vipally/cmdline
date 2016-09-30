// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally

package cmdline

func isSpace(c byte) bool {
	return (c == ' ' || c == '\t')
}

//SplitLine splits a command-line text separated with any ' ' or '\t'
func SplitLine(s string) []string {
	n := len(s) / 2
	len_sep := 1
	start := 0
	a := make([]string, n)
	na := 0
	inString := 0
	escape := 0
	for i := 0; i+len_sep <= len(s) && na+1 < n; i++ {
		if s[i] == '\'' || s[i] == '"' { //" xxxx yyyy " case, do not include \"
			inString++
			escape = 1
		} else {
			if !isSpace(s[i]) {
				escape = 0
			}
		}
		if inString%2 == 0 && isSpace(s[i]) {
			if start == i { //escape continuous space
				start += len_sep
			} else {
				a[na] = s[start+escape : i-escape]
				na++
				start = i + len_sep
				i += len_sep - 1
			}
		}
	}
	if start < len(s) {
		a[na] = s[start+escape : len(s)-escape]
	} else {
		na--
	}

	return a[0 : na+1]
}
