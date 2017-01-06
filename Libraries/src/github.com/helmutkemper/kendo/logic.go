//Comentários Revisado em 20170105
package kendo

// The logical operation to use when the 'filter.filters' option is set.
//
// The supported values are:
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.logic
/*
    Example - set the filter logic
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
     ],
     filter: {
      // leave data items which are "Food" or "Tea"
      logic: "or",
      filters: [
       { field: "category", operator: "eq", value: "Food" },
       { field: "name", operator: "eq", value: "Tea" }
      ]
     }
    });
    dataSource.fetch(function(){
     var view = dataSource.view();
     console.log(view.length); // displays "2"
     console.log(view[0].name); // displays "Tea"
     console.log(view[1].name); // displays "Ham"
    });
    </script>
*/
type LogicEnum int

const(
  FILTER_LOGIC_NULL LogicEnum = iota
  FILTER_LOGIC_AND
  FILTER_LOGIC_OR
)

var LogicEnums = [...]string{
  "",
  "and",
  "or",
}

func (el LogicEnum) String() string {
  return LogicEnums[ el ]
}
