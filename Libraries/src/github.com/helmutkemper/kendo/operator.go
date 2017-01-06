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
  OPERATOR_NULL  OperatorEnum = iota
  OPERATOR_EQ
  OPERATOR_NEQ
  OPERATOR_ISNULL
  OPERATOR_ISNOTNULL
  OPERATOR_LT
  OPERATOR_LTE
  OPERATOR_GT
  OPERATOR_GTE
  OPERATOR_STARTSWITH
  OPERATOR_ENDSWITH
  OPERATOR_CONTAINS
  OPERATOR_DOESNOTCONTAIN
  OPERATOR_ISEMPTY
  OPERATOR_ISNOTEMPTY
)

var OperatorEnums = [...]string{
  "",
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
