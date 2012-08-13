package jpath

import "testing"

func TestBigIDs(t *testing.T) {
	var jp JPath
	jsonBlob := `{
		"id": 12738165059,
	    "id_str": "12738165059",
		"retweet_count": 0,
		"geo": null
	}`

	if er := jp.ParseString(jsonBlob) ; er != nil {
		t.Fatal(er)
	}

	t.Logf("I: %#v\n", jp.I)

	val := jp.Field("id_str").Uint64()
	if val != 12738165059 {
		t.Errorf("Expected %d, got %d", uint64(12738165059), val)
	}
}

func TestObjectFields(t *testing.T) {
	var jp JPath
	jsonBlob := `{
		"numField": 42,
		"strField": "hello",
		"objField": {
			"one": 1,
			"two": 2
		}
	}`

	if er := jp.ParseString(jsonBlob) ; er != nil {
		t.Fatal(er)
	}

	t.Logf("I: %#v\n", jp.I)

	fields := jp.Fields()

	if len(fields) != 3 {
		t.Fatalf("Expected 3 fields, got %d", len(fields))
	}

	expectedFields := []string{"numField", "strField", "objField"}

	for _, expected := range expectedFields {
		found := false

		for _, actual := range fields {
			if expected == actual {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Field list did not contain %s", expected)
		}
	}
}
