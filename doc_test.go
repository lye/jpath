package jpath

import "testing"

func TestExample1(t *testing.T) {
    var jsonBytes = []byte(`[1, 2, "3", 4, "05", "NaN"]`)
	var jp JPath
	var resultList []int

	if er := jp.ParseBytes(jsonBytes) ; er != nil {
		t.Fatal(er)
	}

	for i := 0; i < jp.Length() ; i += 1 {
		resultList = append(resultList, jp.Index(i).Int())
	}

	// ^ Example, v test

	t.Logf("I: %#v\n", jp.I)

	if len(resultList) != 6 {
		t.Fatalf("resultList does not have 6 elements?\n%#v", resultList)
	}

	for i := 1; i <= 5; i += 1 {
		if resultList[i - 1] != i {
			t.Errorf("resultList[%d] != %d", i, i)
		}
	}

	if resultList[5] != 0 {
		t.Errorf("resultList[5] != 0, got %d", resultList[5])
	}
}

