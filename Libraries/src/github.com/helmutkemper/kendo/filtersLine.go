package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter
//
// The filters which are applied over the data items. By default, no filter is applied.
/*
    Example - set the filter field
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
//
type FiltersLine struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field
  //
  // Type: String
  //
  // The data item field to which the filter operator is applied.
  /*
      Example - set the filter field
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
  //
  Field    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.logic
  //
  // Type: String
  //
  // The logical operation to use when the 'filter.filters' option is set.
  //
  // The supported values are:
  //
  // *  "and"
  // *  "or"
  //
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
  //
  Logic    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.operator
  //
  // Type: String
  //
  // The filter operator (comparison).
  //
  // The supported operators are:
  //
  // *  "eq" (equal to)
  // *  "neq" (not equal to)
  // *  "isnull" (is equal to null)
  // *  "isnotnull" (is not equal to null)
  // *  "lt" (less than)
  // *  "lte" (less than or equal to)
  // *  "gt" (greater than)
  // *  "gte" (greater than or equal to)
  // *  "startswith"
  // *  "endswith"
  // *  "contains"
  // *  "doesnotcontain"
  // *  "isempty"
  // *  "isnotempty"
  //
  // The last five are supported only for string fields.
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
  Operator    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.value
  //
  // Type: Object
  //
  // The value to which the 'field' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field  is compared. The value has to be of the same type as the field.
  //
  //    By design, the '"\n"' is removed from the filter before the filtering is performed. That is why an '"\n"' identifier from the filter will not match data items whose corresponding fields contain new lines.
  //
  /*
      Example - specify the filter value
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", birthday: new Date(1983, 1, 1) },
          { name: "John Doe", birthday: new Date(1980, 1, 1)}
        ],
        filter: { field: "birthday", operator: "gt", value: new Date(1980, 1, 1) }
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        console.log(view.length); // displays "1"
        console.log(view[0].name); // displays "Jane Doe"
      });
      </script>
  */
  //
  Value    ComplexJavaScriptType

  Template    *Template
}

func ( el FiltersLine ) getTemplate () string {
  return `{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
{{if ne (string .Filters) "null"}}filters: {{string .Filters}},{{end}}
{{if ne (string .Logic) "null"}}logic: {{string .Logic}},{{end}}
{{if ne (string .Operator) "null"}}operator: {{string .Operator}},{{end}}
{{if ne (string .Value) "null"}}value: {{string .Value}},{{end}}
`
}


