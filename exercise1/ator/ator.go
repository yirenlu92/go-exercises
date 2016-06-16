package main

var (
	romanNumDict = map[int]string{
		1:"I",
		5:"V",
		10:"X",
		50:"L",
		100:"C",
		500:"D",
		1000:"M",
	}
	keys = []int{1,5,10,50,100,500,1000}
)


func pickLargest(arabic int) (string) {

	emptyString := ""

	if arabic == 0 {
		return emptyString
	}

	var (
		romVal int
		k int
	)

	for arabic >= keys[k] {
		romVal = keys[k]
		k = k + 1
	}

	mod := arabic - romVal
	return romanNumDict[romVal] + pickLargest(mod) 

}