package main

import (
	"strings"
)

var (
	arabDict = map[string]int{
		"I":1,
		"IV":4,
		"V":5,
		"IX":9,
		"X":10,
		"XL":40,
		"L":50,
		"XC":90,
		"C":100,
		"CD":400,
		"D":500,
		"CM":900,
		"M":1000,
	}
	keys = []string{"M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"}
)


/*

write some code that converts an Arabic numeral to a Roman numeral. For example, 
14 should be converted to XIV. (Wikipedia has a good description of the Roman numeral system. 
For the purposes of this exercise, we're only concerned with the standard numerals - XIIII is invalid.) 

I placed before V or X indicates one less, so four is IV (one less than five) 
and nine is IX (one less than ten)
X placed before L or C indicates ten less, so forty is XL (ten less than fifty) 
and ninety is XC (ten less than a hundred)
C placed before D or M indicates a hundred less, so four hundred is CD 
(a hundred less than five hundred) and nine hundred is CM (a hundred less than a thousand)[

*/


func pickLargestRoman(roman string) (int) {

	var (
		romVal string
	)

	for _, rom := range keys{
		if strings.HasPrefix(roman, rom){
			romVal = rom
			mod := strings.TrimPrefix(roman, romVal)
			return arabDict[romVal] + pickLargestRoman(mod)
		}
	}

	return 0

}