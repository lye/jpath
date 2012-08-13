# jpath

Have to deal with JSON APIs that occassionally return strange, inconsistently-structured data? Tired of duplicating your model validation logic with the json decoder? Want to throw away your type-assertion verbosity and just assume zero-values when your assumptions are wrong?

Then jpath is for you!

Feed a `JPath` object your JSON and it will wrap the underlying value -- you can then navigate through the data at whim! If you make incorrect assumptions about the underlying structure, no worries! You'll just get zero values back out. Your data model is already responsible for validating the data anyway, *right*?

## Quick Example

So you're ingesting some crazy API that arbitrarily gives you back different data types in the same object fields --

```json
{ "status" : "success"
, "data" : 
	{ "id" : 4
	, "message" : "woot"
	}
}
```

```json
{ "status" : "error"
, "data" : "invalid authentication"
}
```

You're not really sure what the value of `data` is going to be, so you're stuck with an unfortunate few number of choices:

* Parse it into an `interface{}` and do manual type-checking.
* Create a named [Unmarshaler](http://golang.org/pkg/encoding/json/#Unmarshaler) type that decodes the data properly.

Unfortunately, the first requires manual type assertions, which are quite verbose. The second requires additional verbosity as well. Why not take a page from JavaScript's book and just look at the data (throwing safety to the wind; you have a data validation layer somewhere else, right?):

```go
var jp JPath

if er := jp.ParseString(jsonString) ; er != nil {
	return er
}

status := jp.Field("status").String()
data := jp.Field("data")

if status != "success" {
	return fmt.Errorf("API Error: %s", data.String())
}

id := data.Field("id").Int()
message := data.Field("message").String()
```

You'll note that there is zero type-checking. What happens if your assumptions about the structure of the underlying data is wrong (or is changed out from under you)? No biggie, you just get zero values out.

## Further Reading

[Check the docs](http://go.pkgdoc.org/github.com/lye/jpath)
