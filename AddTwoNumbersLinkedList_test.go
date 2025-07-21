package LeetCodeSolutions

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddTwoNumbers_TwoEqualDigitsNumbers(t *testing.T) {
	addNumbers(t, "234", "564")
}

func TestAddTwoNumbers_TwoZeros(t *testing.T) {
	addNumbers(t, "0", "0")
}

func TestAddTwoNumbers_TwoUnEqualDigitsNumbers_1(t *testing.T) {
	addNumbers(t, "9999999", "9999")
}

func TestTwoUnEqualDigitsNumbers_2(t *testing.T) {
	addNumbers(t, "9999", "9999999")
}

func addNumbers(t *testing.T, num1 string, num2 string) {
	ln1 := getNumListNode(t, num1)
	ln2 := getNumListNode(t, num2)
	ln3 := addTwoNumbers(ln1, ln2)
	if ln3 == nil {
		//fmt.Println(err)
		t.Fail()
	} else {
		n1, _ := strconv.Atoi(num1)
		n2, _ := strconv.Atoi(num2)
		n3, _ := strconv.Atoi(NumericString(ln3))
		if n1+n2 != n3 {
			t.Fail()
		}
	}
	printNumbers(ln1, ln2, ln3)
}

func getNumListNode(t *testing.T, num string) *ListNode {
	ln1 := NumberStringToListNode(num)
	if ln1 == nil {
		//fmt.Println(err)
		t.Fail()
	}
	return ln1
}

func printNumbers(ln1 *ListNode, ln2 *ListNode, ln3 *ListNode) {
	fmt.Printf("Ln1 = %+v\n", NumericString(ln1))
	fmt.Printf("Ln2 = %+v\n", NumericString(ln2))
	fmt.Printf("Ln1 + Ln2 = %+v\n", NumericString(ln3))
}
