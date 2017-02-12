package kendo

import "bytes"

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort
//
// The sort order which will be applied over the data items. By default the data items are not sorted.
//
//   The data source sorts the data items client-side unless the serverSorting option is set to true.
//
/*
    EXAMPLE - SORT THE DATA ITEMS
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
*/
//
type SortLine struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.dir
  //
  // Type: String
  //
  // The sort order (direction).
  //
  // The supported values are:
  //
  // "asc" (ascending order)
  // "desc" (descending order)
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
  Dir    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.field
  //
  // Type: String
  //
  // The field by which the data items are sorted.
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
  //
  Field    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.compare
  //
  // Type: Function
  //
  // Function which can be used for custom comparing of the DataSource items.
  /*
      Example - use a custom compare function to compare items in the DataSource
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
  //
  Compare    ComplexJavaScriptType

  GoTemplate    *GoTemplate
}

func ( el SortLine ) getTemplate () string {
  return `{{if ne (string .Dir) "null"}}dir: {{string .Dir}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
{{if ne (string .Compare) "null"}}compare: {{string .Compare}},{{end}}
`
}

func ( el SortLine ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el SortLine ) String() string {
  out := el.Buffer()
  return out.String()
}
