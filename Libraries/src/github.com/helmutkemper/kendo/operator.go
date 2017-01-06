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
  FILTER_OPERATOR_NULL  OperatorEnum = iota
  FILTER_OPERATOR_EQ
  FILTER_OPERATOR_NEQ
  FILTER_OPERATOR_ISNULL
  FILTER_OPERATOR_ISNOTNULL
  FILTER_OPERATOR_LT
  FILTER_OPERATOR_LTE
  FILTER_OPERATOR_GT
  FILTER_OPERATOR_GTE
  FILTER_OPERATOR_STARTSWITH
  FILTER_OPERATOR_ENDSWITH
  FILTER_OPERATOR_CONTAINS
  FILTER_OPERATOR_DOESNOTCONTAIN
  FILTER_OPERATOR_ISEMPTY
  FILTER_OPERATOR_ISNOTEMPTY
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
