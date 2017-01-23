package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.type
//
// Type: String
// Default: GET
//
// The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
//
//    The 'type' option is ignored if 'dataType' is set to '"jsonp"'. JSONP always uses GET requests.
//
/*
    Example - set the HTTP verb of the request
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        create: {
          /@ omitted for brevity @/
          type: "POST"
        }
      }
    });
    </script>
*/
type MethodEnum int

const(
  METHOD_NIL MethodEnum   = iota
  METHOD_DELETE
  METHOD_PUT
  METHOD_GET
  METHOD_POST
)

var MethodEnums  = [...]string{
  "",
  "delete",
  "put",
  "get",
  "post",
}

func (el MethodEnum ) String() string {
  return MethodEnums[ el ]
}
