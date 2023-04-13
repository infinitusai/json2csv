package json2csv

import (
	"bytes"
	"testing"
)

func TestKeyWithTrailingSpace(t *testing.T) {
	b := &bytes.Buffer{}
	wr := NewCSVWriter(b)
	responses := []map[string]interface{}{
		{
			" A":  1,
			"B ":  "foo",
			"C  ": "FOO",
		},
		{
			" A":  2,
			"B ":  "bar",
			"C  ": "BAR",
		},
	}
	csvContent, err := JSON2CSV(responses) // csvContent seems to be complete!
	if err != nil {
		t.Fatal(err)
	}
	wr.WriteCSV(csvContent)
	wr.Flush()

	got := b.String()
	want := `/ A,/B ,/C  
1,foo,FOO
2,bar,BAR
`

	if got != want {
		t.Errorf("Expected %v, but %v", want, got)
	}
}

func TestKeysWithReadableNotationStyle(t *testing.T) {
	b := &bytes.Buffer{}
	wr := NewCSVWriter(b)
	responses := []map[string]interface{}{
		{
			"Test":  1,
			"thisIsATest":  map[string]interface{}{
				"withADot": "foo",
				"withADotTwo": "foo2",
				"withAnother": map[string]interface{}{
					"dot": "foo3",
				},
				"withArray": []interface{}{"foo4", "foo5"},
			},
		},
		{
			"Test":  2,
			"thisIsATest":  map[string]interface{}{
				"withADot": "bar",
				"withADotTwo": "bar2",
				"withAnother": map[string]interface{}{
					"dot": "bar3",
				},
				"withArray": []interface{}{"bar4", "bar5"},
			},
		},
	}
	wr.HeaderStyle = ReadableNotationStyle
	csvContent, err := JSON2CSV(responses) // csvContent seems to be complete!
	if err != nil {
		t.Fatal(err)
	}
	wr.WriteCSV(csvContent)
	wr.Flush()

	got := b.String()
	want := `Test,This Is A Test: With A Dot,This Is A Test: With A Dot Two,This Is A Test: With Another: Dot,This Is A Test: With Array #1,This Is A Test: With Array #2
1,foo,foo2,foo3,foo4,foo5
2,bar,bar2,bar3,bar4,bar5
`

	if got != want {
		t.Errorf("Expected %v, but %v", want, got)
	}
}