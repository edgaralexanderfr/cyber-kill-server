package util

import (
	"testing"
)

func TestSplice(t *testing.T) {
	type StructTest struct {
		Number int16
		String string
	}

	result1 := Splice([]int{9, 8, 7, 6, 5}, 2)
	result2 := Splice([]float32{1.44, 9.99, 3.14}, 0)
	result3 := Splice([]string{"apples", "oranges", "banana", "mango"}, 4)
	result4 := Splice([]byte{0, 255}, -1)
	result5 := Splice([]StructTest{{1, "Text 1"}, {2, "Text 2"}}, 1)

	t.Logf(`Validating Splice([9 8 7 6 5], 2):`)

	if result1 == nil {
		t.Fatal("Splice([9 8 7 6 5], 2) = nil, want [9 8 6 5]")
	}

	if !(len(result1) == 4 &&
		result1[0] == 9 &&
		result1[1] == 8 &&
		result1[2] == 6 &&
		result1[3] == 5) {
		t.Fatal("Splice([9 8 7 6 5], 2) =", result1, ", want [9 8 6 5]")
	}

	t.Logf(`Validating Splice([1.44 9.99 3.14], 0):`)

	if result2 == nil {
		t.Fatal("Splice([1.44 9.99 3.14], 0) = nil, want [9.99 3.14]")
	}

	if !(len(result2) == 2 &&
		result2[0] == 9.99 &&
		result2[1] == 3.14) {
		t.Fatal("Splice([1.44 9.99 3.14], 0) =", result2, ", want [9.99 3.14]")
	}

	t.Logf(`Validating Splice([apples oranges banana mango], 4):`)

	if result3 == nil {
		t.Fatal("Splice([apples oranges banana mango], 4) = nil, want [apples oranges banana mango]")
	}

	if !(len(result3) == 4 &&
		result3[0] == "apples" &&
		result3[1] == "oranges" &&
		result3[2] == "banana" &&
		result3[3] == "mango") {
		t.Fatal("Splice([apples oranges banana mango], 4) =", result3, ", want [apples oranges banana mango]")
	}

	t.Logf(`Validating Splice([0 255], -1):`)

	if result4 == nil {
		t.Fatal("Splice([0 255], -1) = nil, want [0 255]")
	}

	if !(len(result4) == 2 &&
		result4[0] == 0 &&
		result4[1] == 255) {
		t.Fatal("Splice([0 255], -1) =", result4, ", want [0 255]")
	}

	t.Logf(`Validating Splice([{1 Text 1} {2 Text 2}], 1):`)

	if result5 == nil {
		t.Fatal("Splice([{1 Text 1} {2 Text 2}], 1) = nil, want [{1 Text 1}]")
	}

	if !(len(result5) == 1 &&
		result5[0].Number == 1 &&
		result5[0].String == "Text 1") {
		t.Fatal("Splice([{1 Text 1} {2 Text 2}], 1) =", result5, ", want [{1 Text 1}]")
	}
}
