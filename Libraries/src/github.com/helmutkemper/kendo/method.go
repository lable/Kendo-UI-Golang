package kendo

// The type of request to make ("POST", "GET", "PUT" or "DELETE").
// The type option is ignored if dataType is set to "jsonp". JSONP always uses GET requests.
// Refer to the jQuery.ajax documentation for further information.
type MethodEnum int

const(
  METHOD_NULL MethodEnum = iota
  METHOD_POST
  METHOD_GET
  METHOD_PUT
  METHOD_DELETE
)

var MethodEnums = [...]string{
  "",
  "POST",
  "GET",
  "PUT",
  "DELETE",
}

func (el MethodEnum) String() string {
  return MethodEnums[ el ]
}