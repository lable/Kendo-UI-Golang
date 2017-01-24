package kendo

import "bytes"

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group
//
// The grouping configuration of the data source. If set, the data items will be grouped when the data source is populated. By default, grouping is not applied.
//
//   The data source groups the data items client-side unless the 'serverGrouping' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping option is set to true.
//
/*
    EXAMPLE - SET A GROUP AS AN OBJECT
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Tea", category: "Beverages" },
        { name: "Coffee", category: "Beverages" },
        { name: "Ham", category: "Food" }
      ],
      // group by the "category" field
      group: { field: "category" }
    });
    dataSource.fetch(function(){
      var view = dataSource.view();
      console.log(view.length); // displays "2"
      var beverages = view[0];
      console.log(beverages.value); // displays "Beverages"
      console.log(beverages.items[0].name); // displays "Tea"
      console.log(beverages.items[1].name); // displays "Coffee"
      var food = view[1];
      console.log(food.value); // displays "Food"
      console.log(food.items[0].name); // displays "Ham"
    });
    </script>
*/
type GroupLine struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
  //
  // Type: Array
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
  Aggregates    []AggregateEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.dir
  //
  // Type: String
  // Default: asc
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
  Dir    DirEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.field
  //
  // Type: String
  //
  // The data item field to group by.
  /*
      Example - set the field
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages" },
          { name: "Coffee", category: "Beverages" },
          { name: "Ham", category: "Food" }
        ],
        // group by the "category" field
        group: { field: "category" }
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        var beverages = view[0];
        console.log(beverages.items[0].name); // displays "Tea"
        console.log(beverages.items[1].name); // displays "Coffee"
        var food = view[1];
        console.log(food.items[0].name); // displays "Ham"
      });
      </script>
  */
  //
  Field    string

  Template    *Template
}

func ( el GroupLine ) getTemplate () string {
  return `{{if .Aggregates}}aggregates: [{{range $v := .Aggregates}}{{string $v}},{{end}}],{{end}}
{{if ne (string .Dir) "null"}}dir: {{string .Dir}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
`
}

func ( el GroupLine ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.Template == nil {
    buffer.WriteString( "null" )
    return buffer
  }

  el.Template.ParserString( el.getTemplate() )
  el.Template.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el GroupLine ) String() string {
  out := el.Buffer()
  return out.String()
}
