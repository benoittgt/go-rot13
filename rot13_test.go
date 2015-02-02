package rot13

import "testing"

func TestRot13(t *testing.T) {
  tests := []struct {
    input string
    expected string
  } {
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
