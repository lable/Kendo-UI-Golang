package kendo

// The type of the response.
// The supported values are: "jsonp" OR "json"
type DataJSonOrJSonpEnum int

const(
  DATA_TYPE_JSON_OR_JSONP_SET_JSONP DataJSonOrJSonpEnum = iota
  DATA_TYPE_JSON_OR_JSONP_SET_JSON
)

var DataJSonOrJSonpEnums = [...]string{
  "jsonp",
  "json",
}

func (el DataJSonOrJSonpEnum) String() string {
  return DataJSonOrJSonpEnums[ el ]
}
