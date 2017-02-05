package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
//
// The aggregates which are calculated during grouping.
//
// The supported aggregates are:
//
// *  "average"
// *  "count"
// *  "max"
// *  "min"
// *  "sum"
//
/*
    Example - set group aggregates
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Tea", category: "Beverages", price: 1 },
        { name: "Coffee", category: "Beverages", price: 2 },
        { name: "Ham", category: "Food", price: 3 },
      ],
      group: {
        field: "category",
        aggregates: [
          { field: "price", aggregate: "max" },
          { field: "price", aggregate: "min" }
        ]
      }
    });
    dataSource.fetch(function(){
      var view = dataSource.view();
      var beverages = view[0];
      console.log(beverages.aggregates.price.max); // displays "2"
      console.log(beverages.aggregates.price.min); // displays "1"
      var food = view[1];
      console.log(food.aggregates.price.max); // displays "3"
      console.log(food.aggregates.price.min); // displays "3"
    });
    </script>
*/
//
type GroupEnum int

const(
  GROUP_NIL GroupEnum   = iota
  GROUP_AVERAGE
  GROUP_COUNT
  GROUP_MAX
  GROUP_MIN
  GROUP_SUM
)

var GroupEnums  = [...]string{
  "",
  "average",
  "count",
  "max",
  "min",
  "sum",
}

func (el GroupEnum ) String() string {
  return GroupEnums[ el ]
}

