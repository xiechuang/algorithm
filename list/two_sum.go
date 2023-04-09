package list

func TwoSum(l1 *IntNode, l2 *IntNode) *IntNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	len1, len2 := getIntListLen(l1), getIntListLen(l2)
	if len1 > len2 {

	} else {

	}
}

func getIntListLen(l *IntNode) int {
	var length int
	for l != nil {
		length += 1
		l = l.Next
	}
	return length
}
