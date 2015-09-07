# Rest Package

This package contains useful tools to build a JSON-based REST API.

## Response Formats

### ListResp

This response format is for returning generic lists, without pagination.

```
{
	"meta": {
		"count": 100
	},
	"results": [
		{
			id: 1,
			foo: "bar"
		}
	]
}
```

### PagedListResp

This response format is for returning generic lists with pagination.

```
{
	"meta": {
		"count": 20,
		"perpage": 20,
		"pages": 5,
		"results": 100,
		"page": 1
	},
	"results": [
		{
			id: 1,
			foo: "bar"
		},
		{
			id: 2,
			foo: "baz"
		},
		...
	]
}
```

## JSON Formatting

Many times you will need to decode an input JSON data structure. A method is provided to do so, taking in the http.Request and destination interface{} (struct).

```
rest.ParseJSON(r, User)
```

You will also need to output JSON for the request, which can be done with:

```
rest.WriteJSON(w, 200, data)
```