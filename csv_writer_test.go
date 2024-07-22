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

func TestKeyWithCapitalizedWords(t *testing.T) {
	b := &bytes.Buffer{}
	wr := NewCSVWriter(b)
	wr.HeaderStyle = InfinitusNotationStyle
	responses := []map[string]interface{}{
		{
			"BCDTest":  1,
			"abcBCD":   "foo",
			"ABC Test": "FOO",
		},
		{
			"BCDTest":  2,
			"abcBCD":   "bar",
			"ABC Test": "BAR",
		},
	}
	csvContent, err := JSON2CSV(responses) // csvContent seems to be complete!
	if err != nil {
		t.Fatal(err)
	}
	wr.WriteCSV(csvContent)
	wr.Flush()

	got := b.String()
	want := `ABC Test,BCD Test,Abc BCD
FOO,1,foo
BAR,2,bar
`

	if got != want {
		t.Errorf("Expected %v, but %v", want, got)
	}
}

func TestKeysWithReadableNotationStyle(t *testing.T) {
	got, err := writeCSVHelper(ReadableNotationStyle)
	if err != nil {
		t.Fatal(err)
	}
	want := `Test,This Is A Test: With A Dot,This Is A Test: With A Dot Two,This Is A Test: With Another: Dot,This Is A Test: With Array #1,This Is A Test: With Array #2,This Is A Test: With Array Object #1: Key1,This Is A Test: With Array Object #1: Key2,This Is A Test: With Array Object #2: Key1,This Is A Test: With Array Object #2: Key2
1,foo,foo2,foo3,foo4,foo5,foo6,foo7,foo8,foo9
2,bar,bar2,bar3,bar4,bar5,bar6,bar7,bar8,bar9
`

	if got != want {
		t.Errorf("Expected %v, but %v", want, got)
	}
}

func TestKeysWithInfinitusNotationStyle(t *testing.T) {
	got, err := writeCSVHelper(InfinitusNotationStyle)
	if err != nil {
		t.Fatal(err)
	}
	want := `Test,With A Dot,With A Dot Two,With Another: Dot,With Array #1,With Array #2,With Array Object #1: Key1,With Array Object #1: Key2,With Array Object #2: Key1,With Array Object #2: Key2
1,foo,foo2,foo3,foo4,foo5,foo6,foo7,foo8,foo9
2,bar,bar2,bar3,bar4,bar5,bar6,bar7,bar8,bar9
`

	if got != want {
		t.Errorf("Expected %v, but %v", want, got)
	}
}

func TestFieldsWithNilValues(t *testing.T) {
	b := &bytes.Buffer{}
	wr := NewCSVWriter(b)
	responses := []map[string]interface{}{
		{
			"A": 1,
			"B": "foo",
			"C": nil,
		},
	}
	csvContent, err := JSON2CSV(responses)
	if err != nil {
		t.Fatal(err)
	}
	wr.WriteCSV(csvContent)
	wr.Flush()

	got := b.String()
	want := `/A,/B,/C
1,foo,null
`
	if got != want {
		t.Errorf("Expected %v, but %v", want, got)
	}
}

func writeCSVHelper(style KeyStyle) (string, error) {
	responses := getTestResponse()
	csvContent, err := JSON2CSV(responses) // csvContent seems to be complete!
	if err != nil {
		return "", err
	}
	b := &bytes.Buffer{}
	wr := NewCSVWriter(b)
	wr.HeaderStyle = style
	wr.WriteCSV(csvContent)
	wr.Flush()

	return b.String(), nil
}

func getTestResponse() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"Test": 1,
			"thisIsATest": map[string]interface{}{
				"withADot":    "foo",
				"withADotTwo": "foo2",
				"withAnother": map[string]interface{}{
					"dot": "foo3",
				},
				"withArray": []interface{}{"foo4", "foo5"},
				"withArrayObject": []interface{}{
					map[string]interface{}{
						"key1": "foo6",
						"key2": "foo7",
					},
					map[string]interface{}{
						"key1": "foo8",
						"key2": "foo9",
					},
				},
			},
		},
		{
			"Test": 2,
			"thisIsATest": map[string]interface{}{
				"withADot":    "bar",
				"withADotTwo": "bar2",
				"withAnother": map[string]interface{}{
					"dot": "bar3",
				},
				"withArray": []interface{}{"bar4", "bar5"},
				"withArrayObject": []interface{}{
					map[string]interface{}{
						"key1": "bar6",
						"key2": "bar7",
					},
					map[string]interface{}{
						"key1": "bar8",
						"key2": "bar9",
					},
				},
			},
		},
	}
}
