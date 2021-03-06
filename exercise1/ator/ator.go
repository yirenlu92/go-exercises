package main

var (
	romanNumDict = map[int]string{
		1:"I",
		4:"IV",
		5:"V",
		9:"IX",
		10:"X",
		40:"XL",
		50:"L",
		90:"XC",
		100:"C",
		400:"CD",
		500:"D",
		900:"CM",
		1000:"M",
	}
	keys = []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
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


func pickLargest(arabic int) (string) {

	var (
		romVal int
		emptyString = ""
	)

	for _ , arab := range keys {
		if arab <= arabic {
			romVal = arab
			mod := arabic - romVal
			return romanNumDict[romVal] + pickLargest(mod)
		}
	}

	return emptyString

}