package kendo

// The type of the response.
// The supported values are: "xml" OR "json"
type DataXmlOrJSonEnum int

const(
  DATA_TYPE_XML_OR_JSON_SET_XML DataXmlOrJSonEnum = iota
  DATA_TYPE_XML_OR_JSON_SET_JSON
)

var DataXmlOrJSonEnums = [...]string{
  "xml",
  "json",
}

func (el DataXmlOrJSonEnum) String() string {
  return DataXmlOrJSonEnums[ el ]
}