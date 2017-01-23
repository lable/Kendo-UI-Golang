package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.dataType
//
// Type: String
//
// The type of result expected from the server. Commonly used values are "json" and "jsonp".
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
/*
    Example - set the data type to JSON
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        update: {
          /@ omitted for brevity @/
          dataType: "json"
        }
      }
    });
    </script>
*/
type TypeDataJSonEnum int

const(
  TYPE_DATA_NIL TypeDataJSonEnum   = iota
  TYPE_DATA_JSONP
  TYPE_DATA_JSON
)

var TypeDataJSonEnums  = [...]string{
  "",
  "xml",
  "json",
}

func (el TypeDataJSonEnum ) String() string {
  return TypeDataJSonEnums[ el ]
}
