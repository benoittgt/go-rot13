package rot13

import "testing"

func TestRot13(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a", expected: "n"},
		{input: "b", expected: "o"},
		{input: "o", expected: "b"},
		{input: "A", expected: "N"},
		{input: "mot", expected: "zbg"},
		{input: "deux mots", expected: "qrhk zbgf"},
	}

	for _, testCase := range tests {
		actual := Rot13(testCase.input)
		expected := testCase.expected
		if actual != expected {
			t.Errorf("%s != %s", actual, expected)
		}
	}
}

func BenchmarkRot13Sentence(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Rot13("Test avec une longue phrase !")
	}
}
func BenchmarkRot13Complicated(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Rot13("456 -ERrt ertTTTTe 5-R  ! 3R")
	}
}
func BenchmarkRot13Simple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Rot13("hop")
	}
}
