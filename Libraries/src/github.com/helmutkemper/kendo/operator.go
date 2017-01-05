//Comentários Revisado em 20170105
package kendo

// The filter operator (comparison).
//
// The supported operators are:
//
// '"isnotempty"'
//
// The last five are supported only for string fields.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.operator
/*
    Example - set the filter operator
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe" },
      { name: "John Doe" }
     ],
     filter: { field: "name", operator: "startswith", value: "Jane" }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "1"
     console.log(view[0].name); // displays "Jane Doe"
    });
    </script>
*/
type OperatorEnum int

const(
  FILTER_LOGIC_EQ  OperatorEnum = iota
  FILTER_LOGIC_NEQ
  FILTER_LOGIC_ISNULL
  FILTER_LOGIC_ISNOTNULL
  FILTER_LOGIC_LT
  FILTER_LOGIC_LTE
  FILTER_LOGIC_GT
  FILTER_LOGIC_GTE
  FILTER_LOGIC_STARTSWITH
  FILTER_LOGIC_ENDSWITH
  FILTER_LOGIC_CONTAINS
  FILTER_LOGIC_DOESNOTCONTAIN
  FILTER_LOGIC_ISEMPTY
  FILTER_LOGIC_ISNOTEMPTY
)

var OperatorEnums = [...]string{
  "eq",
  "neq",
  "isnull",
  "isnotnull",
  "lt",
  "lte",
  "gt",
  "gte",
  "startswith",
  "endswith",
  "contains",
  "doesnotcontain",
  "isempty",
  "isnotempty",
}

func (el OperatorEnum) String() string {
  return OperatorEnums[ el ]
}
