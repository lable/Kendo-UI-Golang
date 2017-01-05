package kendo

// The sort order which will be applied over the data items. By default the data items are not sorted.
//
// The data source sorts the data items client-side unless the serverSorting ( http://docs.telerik.com#configuration-serverSorting ) option is set to 'true'.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort
/*
    Example - sort the data items
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Jane Doe", age: 30 },
      { name: "John Doe", age: 33 }
     ],
     sort: { field: "age", dir: "desc" }
    });
    dataSource.fetch(function(){
     var data = dataSource.view();
     console.log(data[0].age); // displays "33"
    });
    </script>


    Example - sort the data items by multiple fields
    <script>
    var dataSource = new kendo.data.DataSource({
     data: [
      { name: "Tea", category: "Beverages" },
      { name: "Coffee", category: "Beverages" },
      { name: "Ham", category: "Food" }
     ],
     sort: [
      // sort by "category" in descending order and then by "name" in ascending order
      { field: "category", dir: "desc" },
      { field: "name", dir: "asc" }
     ]
    });
    dataSource.fetch(function(){
     var data = dataSource.view();
     console.log(data[1].name); // displays "Coffee"
    });
    </script>
*/
type Sort struct {
  // The sort order (direction).
  //
  // The supported values are:
  //
  // *  "asc" (ascending order)
  // *  "desc" (descending order)
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.dir
  /*
      Example - specify the sort order (direction)
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe", age: 30 },
        { name: "John Doe", age: 33 }
       ],
       // order by "age" in descending order
       sort: { field: "age", dir: "desc" }
      });
      dataSource.fetch(function(){
       var data = dataSource.view();
       console.log(data[0].age); // displays "33"
      });
      </script>
  */
  Dir             DirEnum

  // The field by which the data items are sorted.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.field
  /*
      Example - specify the sort field
      <script>
      var dataSource = new kendo.data.DataSource({
       data: [
        { name: "Jane Doe", age: 30 },
        { name: "John Doe", age: 33 }
       ],
       // order by "age" in descending order
       sort: { field: "age", dir: "desc" }
      });
      dataSource.fetch(function(){
       var data = dataSource.view();
       console.log(data[0].age); // displays "33"
      });
      </script>
  */
  Field           string

  // Function which can be used for custom comparing of the DataSource items.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.compare
  /*
      EXAMPLE - USE A CUSTOM COMPARE FUNCTION TO COMPARE ITEMS IN THE DATASOURCE

      EditPreviewOpen In Dojo
      <script>
        var numbers = {
          "one"  : 1,
          "two"  : 2,
          "three": 3
        };

        var dataSource = new kendo.data.DataSource({
          data: [
            { id: 1, item: "two" },
            { id: 2, item: "one" },
            { id: 3, item: "three" }
          ],
          sort: { field: "item", dir: "asc", compare: function(a, b) {
            return numbers[a.item] - numbers[b.item];
          }
                }
        });

        $("#grid").kendoGrid({
          dataSource: dataSource,
          sortable: true,
          columns: [{
            field: "item",
            sortable: {
              compare: function(a, b) {
                return numbers[a.item] - numbers[b.item];
              }
            }
          }]
        });
      </script>
  */
  Compare         string
}
