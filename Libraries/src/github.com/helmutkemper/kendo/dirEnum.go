package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.dir
//
// The sort order of the group.
//
// The supported values are:
//
// "asc" (ascending order)
// "desc" (descending order)
//
// The default sort order is ascending.
/*
    Example - sort the groups in descending order
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Tea", category: "Beverages"},
        { name: "Ham", category: "Food"},
      ],
      // group by "category" in descending order
      group: { field: "category", dir: "desc" }
    });
    dataSource.fetch(function(){
      var view = dataSource.view();
      var food = view[0];
      console.log(food.value); // displays "Food"
      var beverages = view[1];
      console.log(beverages.value); // displays "Beverages"
    });
    </script>
*/
//
type DirEnum int

const(
  DIR_NIL DirEnum   = iota
  DIR_ASC
  DIR_DESC
)

var DirEnums  = [...]string{
  "",
  "asc",
  "desc",
}

func (el DirEnum ) String() string {
  return DirEnums[ el ]
}
