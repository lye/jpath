/*
Package jpath provides an easy way to navigate arbitrary JSON blobs.

Often, you'll encounter JSON blobs with a varying or undefined format that don't 
readily conform to a statically-typed struct. In these cases, you often end up
parsing JSON objects into a map[string]interface{} (or worse, interface{}), which
necessitates a significant amount of manual type-checking. jpath allows you to
defer error checking to make code more succient:

Presume you're dealing with a list of numbers, encoded heterogeneously as JSON 
doubles and strings (you know you've seen unfortunates APIs like this). Your code
may look like the following:

    var jsonBytes = []byte(`[1, 2, "3", 4, "05", "NaN"]`)
	var crazyList []interface{}
	var resultList []int

	if er = json.Unmarshal(jsonBytes, &crazyList) ; er != nil {
		return er
	}

	for _, crazyVal := range(crazyList) {
		var intVal int = 0

		if floatVal, ok := crazyVal.(float64) ; ok {
			intVal = int(floatVal)

		} else if strVal, ok := crazyVal.(string) ; ok {
			if intVal, er = strconv.ParseInt(strVal, 10, 64) ; er != nil {
				intVal = 0
			}
		}

		resultList = append(resultList, intVal)
	}

Navigating a tree of JSON objects is even worse. Many times, as above, you're not
overly concerned with invalid values coming from someone else's API. If they throw
something weird at you, all you want is a zero-value which will fail a sanity check
somewhere else in the code. Enter jpath:

    var jsonBytes = []byte(`[1, 2, "3", 4, "05", "NaN"]`)
	var jp JPath
	var resultList []int

	if er := jp.ParseBytes(jsonBytes) ; er != nil {
		return er
	}

	for i := 0; i < jp.Length() ; i += 1 {
		resultList = append(resultList, jp.Index(i).Int())
	}

If you're not passed a list, nothing blows up -- Length will be 0 and you'll be
square. If something isn't an integer, it'll attempt to coerce it to one, or return
a zero-value. jpath allows you to completely defer all error checking to a later
validation phase, which can be decoupled from the parsing/decoding layer.

All jpath operations are performed by-value. Anything that could return an object
(e.g., Index and Field) return a new JPath object which can be further inspected 
without modifying the state of the original object.
*/
package jpath
