package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
//
// The aggregates which are calculated when the data source populates with data.
//
// The supported aggregates are:
//
//
// *  "average"
// *  "count"
// *  "max"
// *  "min"
// *  "sum"
//
//
//    The data source calculates aggregates client-side unless the 'serverAggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates option is set to 'true'.
//
/*
  Example - specify aggregates
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Jane Doe", age: 30 },
        { name: "John Doe", age: 33 }
      ],
      aggregate: [
        { field: "age", aggregate: "sum" },
        { field: "age", aggregate: "min" },
        { field: "age", aggregate: "max" }
      ]
    });
    dataSource.fetch(function(){
      var results = dataSource.aggregates().age;
      console.log(results.sum, results.min, results.max); // displays "63 30 33"
    });
    </script>
*/
//
type Datasource struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate
  //
  // Type: Array
  //
  // The aggregates which are calculated when the data source populates with data.
  //
  // The supported aggregates are:
  //
  //
  // *  "average"
  // *  "count"
  // *  "max"
  // *  "min"
  // *  "sum"
  //
  //
  //    The data source calculates aggregates client-side unless the 'serverAggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates option is set to 'true'.
  //
  /*
    Example - specify aggregates
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        aggregate: [
          { field: "age", aggregate: "sum" },
          { field: "age", aggregate: "min" },
          { field: "age", aggregate: "max" }
        ]
      });
      dataSource.fetch(function(){
        var results = dataSource.aggregates().age;
        console.log(results.sum, results.min, results.max); // displays "63 30 33"
      });
      </script>
  */
  //
  Aggregate    []AggregateLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-autoSync
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true' the data source would automatically save any changed data items by calling the 'sync' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-sync  method. By default, changes are not automatically saved.
  //
  //
  //
  /*
    Example - enable auto sync
      <script>
      var dataSource = new kendo.data.DataSource({
        autoSync: true,
        transport: {
          read:  {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          update: {
            url: "http://demos.telerik.com/kendo-ui/service/products/update",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        },
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var product = dataSource.at(0);
        product.set("UnitPrice", 20); // auto-syncs and makes request to http://demos.telerik.com/kendo-ui/service/products/update
      });
      </script>
  */
  //
  AutoSync    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-batch
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will batch CRUD operation requests. For example, updating two data items would cause one HTTP request instead of two. By default, the data source
  // makes a HTTP request for every CRUD operation.
  //
  //    The changed data items are sent as 'models' by default. This can be changed via the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap option.
  //
  /*
    Example - enable the batch mode
      <script>
      var dataSource = new kendo.data.DataSource({
        batch: true,
        transport: {
          read:  {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          update: {
            url: "http://demos.telerik.com/kendo-ui/service/products/update",
            dataType: "jsonp" //"jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        },
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var product = dataSource.at(0);
        product.set("UnitPrice", 20);
        var anotherProduct = dataSource.at(1);
        anotherProduct.set("UnitPrice", 20);
        dataSource.sync(); // causes only one request to "http://demos.telerik.com/kendo-ui/service/products/update"
      });
      </script>
  */
  //
  Batch    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-page
  //
  // Type: Number
  //
  // The page of data which the data source will return when the 'view' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-view  method is invoked or request from the remote service.
  //
  //    The data source will page the data items client-side unless the 'serverPaging' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging option is set to 'true'.
  //
  /*
    Example - set the current page
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages" },
          { name: "Coffee", category: "Beverages" },
          { name: "Ham", category: "Food" }
        ],
        // set the second page as the current page
        page: 2,
        pageSize: 2
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        console.log(view.length); // displays "1"
        console.log(view[0].name); // displays "Ham"
      });
      </script>
  */
  //
  Page    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-pageSize
  //
  // Type: Number
  //
  // The number of data items per page. The property has no default value. That is why to use paging, make sure some 'pageSize' value is set.
  //
  //    The data source will page the data items client-side unless the 'serverPaging' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging option is set to 'true'.
  //
  /*
    Example - set the page size
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages" },
          { name: "Coffee", category: "Beverages" },
          { name: "Ham", category: "Food" }
        ],
        page: 1,
        // a page of data contains two data items
        pageSize: 2
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        console.log(view.length); // displays "2"
        console.log(view[0].name); // displays "Tea"
        console.log(view[1].name); // displays "Coffee"
      });
      </script>
  */
  //
  PageSize    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema
  //
  // Type: Object
  //
  // The configuration used to parse the remote service response.
  //
  //
  //
  /*
    Example - specify the schema of the remote service
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: { q: "html5" } // search for tweets that contain "html5"
          }
        },
        schema: {
          data: function(response) {
            return response.results; // twitter's response is { "results": [ /@ results @/ ] }
          }
        }
      });
      dataSource.fetch(function(){
        var data = this.data();
        console.log(data.length);
      });
      </script>
  */
  //
  Schema    Schema

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the aggregate calculation to the remote service. By default, the data source calculates aggregates client-side.
  //
  //    Configure 'schema.aggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.aggregates if you set 'serverAggregates' to 'true'.
  //
  /*
    Example - enable server aggregates
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverAggregates: true,
        aggregate: [
          { field: "age", aggregate: "sum" }
        ],
        schema: {
          aggregates: "aggregates" // aggregate results are returned in the "aggregates" field of the response
        }
      });
      </script>
  */
  //
  ServerAggregates    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the filtering implementation to the remote service. By default, the data source performs filtering client-side.
  //
  // By default, the 'filter' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the filter '{ logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] }' is sent as:
  //
  //
  // 'filter[logic]: and'
  //
  // 'filter[filters][0][field]: name'
  //
  // 'filter[filters][0][operator]: startswith'
  //
  // 'filter[filters][0][value]: Jane'
  //
  //
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the filter option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  //
  //
  //
  /*
    Example - enable server filtering
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverFiltering: true,
        filter: { logic: "and", filters: [ { field: "name", operator: "startswith", value: "Jane" } ] }
      });
      </script>
  */
  //
  ServerFiltering    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the grouping implementation to the remote service. By default, the data source performs grouping client-side.
  //
  // By default, the 'group' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the group '{ field: "category", dir: "desc" }' is sent as:
  //
  //
  // 'group[0][field]: category'
  //
  // 'group[0][dir]: desc'
  //
  //
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the group option in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  //
  //
  //
  /*
    Example - enable server grouping
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverGrouping: true,
        group: { field: "category", dir: "desc" }
      });
      </script>
  */
  //
  ServerGrouping    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the data item paging implementation to the remote service. By default, the data source performs paging client-side.
  //
  //    Configure 'schema.total' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.total if you set 'serverPaging' to 'true'. In addition, 'pageSize' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-pageSize should be set no matter if paging is performed client-side or server-side.
  //
  /*
    Example - enable server paging
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverPaging: true,
        schema: {
          total: "total" // total is returned in the "total" field of the response
        }
      });
      </script>
  */
  //
  ServerPaging    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting
  //
  // Type: Boolean
  // Defalt: false
  //
  // If set to 'true', the data source will leave the data item sorting implementation to the remote service. By default, the data source performs sorting client-side.
  //
  // By default, the 'sort' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort  is sent to the server following jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  // For example, the sort '{ field: "age", dir: "desc" }' is sent as:
  //
  //
  // 'sort[0][field]: age'
  // 'sort[0][dir]: desc'
  //
  //
  // Use the 'parameterMap' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap  option to send the paging options in a different format.
  //
  // For more information and tips about client and server data operations, refer to the introductory article on the DataSource http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/framework/datasource/overview#mixed-data-operations-mode .
  //
  //
  //
  /*
    Example - enable server sorting
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          /@ transport configuration @/
        },
        serverSorting: true,
        sort: { field: "age", dir: "desc" }
      });
      </script>
  */
  //
  ServerSorting    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport
  //
  // Type: Object
  //
  // The configuration used to load and save the data items. A data source is remote or local based on the way of it retrieves data items.
  //
  // Remote data sources load and save data items from and to a remote end-point (also known as remote service or server). The 'transport' option describes the remote service configuration - URL, HTTP verb, HTTP headers, and others. The 'transport' option can also be used to implement custom data loading and saving.
  //
  // Local data sources are bound to a JavaScript array via the 'data' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-data  option.
  //
  //
  //
  /*
    Example - specify the remote service configuration
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        console.log(products[0].ProductName); // displays "Chai"
      });
      </script>
  */
  //
  Transport    Transport

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-type
  //
  // Type: String
  //
  // If set, the data source will use a predefined 'transport' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport  and/or 'schema' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema .
  //
  // The supported values are:
  //
  //
  //
  // "odata" which supports the OData http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://www.odata.org  v.2 protocol
  //
  // "odata-v4" which partially supports http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttps://github.com/telerik/ui-for-aspnet-mvc-examples/tree/master/grid/odata-v4-web-api-binding
  // odata version 4
  // *  "signalr"
  //
  //
  //
  //
  /*
    Example - enable OData support
      <script>
      var dataSource= new kendo.data.DataSource({
        type: "odata",
        transport: {
          read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
        },
        pageSize: 20,
        serverPaging: true
      });
      dataSource.fetch(function() {
        console.log(dataSource.view().length); // displays "20"
      });
      </script>
  */
  //
  Type    string

  Template    *Template
}

func ( el Datasource ) getTemplate () string {
  return `{{if .Aggregate}}aggregate: [{{range $v := .Aggregate}}{{string $v}},{{end}}],{{end}}
{{if .AutoSync}}batch: true,{{end}}
{{if .Batch}}batch: true,{{end}}
{{if ne (string .Page) "null"}}page: {{string .Page}},{{end}}{{if ne (string .PageSize) "null"}}pageSize: {{string .PageSize}},{{end}}{{if .ServerAggregates}}batch: true,{{end}}
{{if .ServerFiltering}}batch: true,{{end}}
{{if .ServerGrouping}}batch: true,{{end}}
{{if .ServerPaging}}batch: true,{{end}}
{{if .ServerSorting}}batch: true,{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
//
// The name of the aggregate function.
//
// The supported aggregates are:
//
//
// *  "average"
// *  "count"
// *  "max"
// *  "min"
// *  "sum"
//
//
//
//
/*
  Example - specify aggregate function
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Jane Doe", age: 30 },
        { name: "John Doe", age: 33 }
      ],
      aggregate: [
        { field: "age", aggregate: "sum" }
      ]
    });
    dataSource.fetch(function(){
      var results = dataSource.aggregates().age;
      console.log(results.sum); // displays "63"
    });
    </script>
*/
//
type Aggregate struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.aggregate
  //
  // Type: String
  //
  // The name of the aggregate function.
  //
  // The supported aggregates are:
  //
  //
  // *  "average"
  // *  "count"
  // *  "max"
  // *  "min"
  // *  "sum"
  //
  //
  //
  //
  /*
    Example - specify aggregate function
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        aggregate: [
          { field: "age", aggregate: "sum" }
        ]
      });
      dataSource.fetch(function(){
        var results = dataSource.aggregates().age;
        console.log(results.sum); // displays "63"
      });
      </script>
  */
  //
  Aggregate    Aggregate

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-aggregate.field
  //
  // Type: String
  //
  // The data item field which will be used to calculate the aggregates.
  //
  //
  //
  /*
    Example - specify aggregate field
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Jane Doe", age: 30 },
          { name: "John Doe", age: 33 }
        ],
        aggregate: [
          { field: "age", aggregate: "sum" }
        ]
      });
      dataSource.fetch(function(){
        var results = dataSource.aggregates().age;
        console.log(results.sum); // displays "63"
      });
      </script>
  */
  //
  Field    string

  Template    *Template
}

func ( el Aggregate ) getTemplate () string {
  return `{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field
//
// The data item field to which the filter operator is applied.
//
//
//
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
type Filter struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.field
  //
  // Type: String
  //
  // The data item field to which the filter operator is applied.
  //
  //
  //
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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.filters
  //
  // Type: Array
  //
  // The nested filter expressions. Supports the same options as 'filter' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter . Filters can be nested indefinitely.
  //
  //
  //
  /*
    Example - nested filters
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
  Filters    Filters

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-filter.logic
  //
  // Type: String
  //
  // The logical operation to use when the 'filter.filters' option is set.
  //
  // The supported values are:
  //
  //
  // *  "and"
  // *  "or"
  //
  //
  //
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
  Value    Value

  Template    *Template
}

func ( el Filter ) getTemplate () string {
  return `{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}{{if ne (string .Filters) "null"}}filters: {{string .Filters}},{{end}}{{if ne (string .Logic) "null"}}logic: {{string .Logic}},{{end}}{{if ne (string .Value) "null"}}value: {{string .Value}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
//
// The aggregates which are calculated during grouping.
//
// The supported aggregates are:
//
//
// *  "average"
// *  "count"
// *  "max"
// *  "min"
// *  "sum"
//
//
//
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
type Group struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates
  //
  // Type: Array
  //
  // The aggregates which are calculated during grouping.
  //
  // The supported aggregates are:
  //
  //
  // *  "average"
  // *  "count"
  // *  "max"
  // *  "min"
  // *  "sum"
  //
  //
  //
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
  Aggregates    []AggregatesLine

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.dir
  //
  // Type: String
  // Defalt: asc
  //
  // The sort order of the group.
  //
  // The supported values are:
  //
  //
  // "asc" (ascending order)
  //
  // "desc" (descending order)
  //
  //
  // The default sort order is ascending.
  //
  //
  //
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
  Dir    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.field
  //
  // Type: String
  //
  // The data item field to group by.
  //
  //
  //
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

func ( el Group ) getTemplate () string {
  return `{{if .Aggregates}}aggregates: [{{range $v := .Aggregates}}{{string $v}},{{end}}],{{end}}
{{if ne (string .Dir) "null"}}dir: {{string .Dir}},{{end}}{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.aggregate
//
// The name of the aggregate function. Specifies the aggregate function.
//
// The supported aggregates are:
//
//
// *  "average"
// *  "count"
// *  "max"
// *  "min"
// *  "sum"
//
//
//
//
/*
  Example - specify an aggregate function
    <script>
    var dataSource = new kendo.data.DataSource({
      data: [
        { name: "Tea", category: "Beverages", price: 1 },
        { name: "Coffee", category: "Beverages", price: 2 },
        { name: "Ham", category: "Food", price: 3 }
      ],
      group: {
        field: "category",
        aggregates: [
          // calculate max price
          { field: "price", aggregate: "max" }
        ]
      }
    });
    dataSource.fetch(function(){
      var view = dataSource.view();
      var beverages = view[0];
      console.log(beverages.aggregates.price.max); // displays "2"
    });
    </script>
*/
//
type Aggregates struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.aggregate
  //
  // Type: String
  //
  // The name of the aggregate function. Specifies the aggregate function.
  //
  // The supported aggregates are:
  //
  //
  // *  "average"
  // *  "count"
  // *  "max"
  // *  "min"
  // *  "sum"
  //
  //
  //
  //
  /*
    Example - specify an aggregate function
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages", price: 1 },
          { name: "Coffee", category: "Beverages", price: 2 },
          { name: "Ham", category: "Food", price: 3 }
        ],
        group: {
          field: "category",
          aggregates: [
            // calculate max price
            { field: "price", aggregate: "max" }
          ]
        }
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        var beverages = view[0];
        console.log(beverages.aggregates.price.max); // displays "2"
      });
      </script>
  */
  //
  Aggregate    Aggregate

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-group.aggregates.field
  //
  // Type: String
  //
  // The data item field which will be used to calculate the aggregates.
  //
  //
  //
  /*
    Example - specify an aggregate field
      <script>
      var dataSource = new kendo.data.DataSource({
        data: [
          { name: "Tea", category: "Beverages", price: 1 },
          { name: "Coffee", category: "Beverages", price: 2 },
          { name: "Ham", category: "Food", price: 3 }
        ],
        group: {
          field: "category",
          aggregates: [
            // calculate max price
            { field: "price", aggregate: "max" }
          ]
        }
      });
      dataSource.fetch(function(){
        var view = dataSource.view();
        var beverages = view[0];
        console.log(beverages.aggregates.price.max); // displays "2"
      });
      </script>
  */
  //
  Field    string

  Template    *Template
}

func ( el Aggregates ) getTemplate () string {
  return `{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.aggregates
//
// The field from the response which contains the aggregate results. Can be set to a function which is called to return the aggregate results from the response.
//
//    The 'aggregates' option is used only when the 'serverAggregates' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverAggregates option is set to 'true'.
//
/*
  Example - set aggregates as a string
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        /@ transport configuration @/
      }
      serverAggregates: true,
      schema: {
        aggregates: "aggregates" // aggregate results are returned in the "aggregates" field of the response
      }
    });
    </script>


    Example - set aggregates as a function
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        /@ transport configuration @/
      }
      serverAggregates: true,
      schema: {
        aggregates: function(response) {
          return response.aggregates;
        }
      }
    });
    </script>
*/
//
type Schema struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.model
  //
  // Type: Object
  //
  // The data item (model) configuration.
  //
  // If set to an object, the 'Model.define' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/model#model.define  method will be used to initialize the data source model.
  //
  // If set to an existing 'kendo.data.Model' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource/kendo-ui/api/javascript/data/model  instance, the data source will use that instance and will not initialize a new one.
  //
  //
  //
  /*
    Example - set the model as a JavaScript object
      <script>
      var dataSource = new kendo.data.DataSource({
        schema: {
          model: {
            id: "ProductID",
            fields: {
              ProductID: {
                //this field will not be editable (default value is true)
                editable: false,
                // a defaultValue will not be assigned (default value is false)
                nullable: true
              },
              ProductName: {
                //set validation rules
                validation: { required: true }
              },
              UnitPrice: {
                //data type of the field {Number|String|Boolean|Date} default is String
                type: "number",
                // used when new model is created
                defaultValue: 42,
                validation: { required: true, min: 1 }
              }
            }
          }
        }
      });
      </script>


      Example - set the model as an existing <code>kendo.data.Model</code> instance
      <script>
      var Product = kendo.model.define({
        id: "ProductID",
        fields: {
          ProductID: {
            //this field will not be editable (default value is true)
            editable: false,
            // a defaultValue will not be assigned (default value is false)
            nullable: true
          },
          ProductName: {
            //set validation rules
            validation: { required: true }
          },
          UnitPrice: {
            //data type of the field {Number|String|Boolean|Date} default is String
            type: "number",
            // used when new model is created
            defaultValue: 42,
            validation: { required: true, min: 1 }
          }
        }
      });
      var dataSource = new kendo.data.DataSource({
        schema: {
          model: Product
        }
      });
      </script>
  */
  //
  Model    Model

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.parse
//
// Type: Function
//
// Executed before the server response is used. Use it to preprocess or parse the server response.
//
//
//
/*
  Parameters
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        read: {
          url: "http://demos.telerik.com/kendo-ui/service/products",
          dataType: "jsonp"
        }
      },
      schema: {
        parse: function(response) {
          var products = [];
          for (var i = 0; i < response.length; i++) {
            var product = {
              id: response[i].ProductID,
              name: response[i].ProductName
            };
            products.push(product);
          }
          return products;
        }
      }
    });
    dataSource.fetch(function(){
      var data = dataSource.data();
      var product = data[0];
      console.log(product.name); // displays "Chai"
    });
    </script>


    Returns


    Example - data projection
*/
//


<h5>response 'Object |' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object 'Array' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
</h5>

<p>The initially parsed server response that may need additional modifications.</p>


//


<p>'Array' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array  The data items from the response.</p>


//
Parse    ComplexJavaScriptType

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema.type
//
// Type: String
// Defalt: json
//
// The type of the response.
//
// The supported values are:
//
//
// *  "xml"
// *  "json"
//
//
// By default, the schema interprets the server response as JSON.
//
//
//
/*
  Example - use XML data
    <script>
    var dataSource = new kendo.data.DataSource({
      data: '<books><book id="1"><title>Secrets of the JavaScript Ninja</title></book></books>',
      schema: {
        // specify the the schema is XML
        type: "xml",
        // the XML element which represents a single data record
        data: "/books/book",
        // define the model - the object which will represent a single data record
        model: {
          // configure the fields of the object
          fields: {
            // the "title" field is mapped to the text of the "title" XML element
            title: "title/text()",
            // the "id" field is mapped to the "id" attribute of the "book" XML element
            id: "@cover"
          }
        }
      }
    });
    dataSource.fetch(function() {
      var books = dataSource.data();
      console.log(books[0].title); // displays "Secrets of the JavaScript Ninja"
    });
    </script>
*/
//
Type    string

Template    *Template
}

func ( el Schema ) getTemplate () string {
  return `{{if ne (string .Model) "null"}}model: {{string .Model}},{{end}}{{if ne (string .Parse) "null"}}parse: {{string .Parse}},{{end}}{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.dir
//
// The sort order (direction).
//
// The supported values are:
//
//
//
// "asc" (ascending order)
//
// "desc" (descending order)
//
//
//
//
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
//
type Sort struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.dir
  //
  // Type: String
  //
  // The sort order (direction).
  //
  // The supported values are:
  //
  //
  //
  // "asc" (ascending order)
  //
  // "desc" (descending order)
  //
  //
  //
  //
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
  //
  Dir    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-sort.field
  //
  // Type: String
  //
  // The field by which the data items are sorted.
  //
  //
  //
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
  //
  //
  //
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

  Template    *Template
}

func ( el Sort ) getTemplate () string {
  return `{{if ne (string .Dir) "null"}}dir: {{string .Dir}},{{end}}{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}{{if ne (string .Compare) "null"}}compare: {{string .Compare}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create
//
// The configuration used when the data source saves newly created data items. Those are items added to the data source via the 'add' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-add  or 'insert' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#methods-insert  methods.
//
//    The data source uses 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax/ to make a HTTP request to the remote service. The value configured via 'transport.create' is passed to 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jquery.ajax/#jQuery-ajax-settings. This means that you can set all options supported by 'jQuery.ajax' via 'transport.create' except the 'success' and 'error' callback functions which are used by the transport.
//
/*
  Example - set the create remote service
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/create
        create: {
          url: "http://demos.telerik.com/kendo-ui/service/products/create",
          dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
        },
        parameterMap: function(data, type) {
          if (type == "create") {
            // send the created data items as the "models" service parameter encoded in JSON
            return { models: kendo.stringify(data.models) };
          }
        }
      },
      batch: true,
      schema: {
        model: { id: "ProductID" }
      }
    });
    // create a new data item
    dataSource.add( { ProductName: "New Product" });
    // save the created data item
    dataSource.sync(); // server response is [{"ProductID":78,"ProductName":"New Product","UnitPrice":0,"UnitsInStock":0,"Discontinued":false}]
    </script>


    Example - set create as a function
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        read: function(options) {
          /@ implementation omitted for brevity @/
        },
        create: function(options) {
          // make JSONP request to http://demos.telerik.com/kendo-ui/service/products/create
          $.ajax({
            url: "http://demos.telerik.com/kendo-ui/service/products/create",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            // send the created data items as the "models" service parameter encoded in JSON
            data: {
              models: kendo.stringify(options.data.models)
            },
            success: function(result) {
              // notify the data source that the request succeeded
              options.success(result);
            },
            error: function(result) {
              // notify the data source that the request failed
              options.error(result);
            }
          });
        }
      },
      batch: true,
      schema: {
        model: { id: "ProductID" }
      }
    });
    dataSource.add( { ProductName: "New Product" });
    dataSource.sync();
    </script>
*/
//
type Transport struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap
  //
  // Type: Function
  //
  // The function which converts the request parameters to a format suitable for the remote service. By default, the data source sends the parameters using jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
  //
  //    The 'parameterMap' method is often used to encode the parameters in JSON format.
  The 'parameterMap' function will not be called when using custom functions for the read, update, create, and destroy operations.
//
/*
  Parameters
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        read: {
          url: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders?$format=json",
          dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          jsonp: "$callback",
          cache: true
        },
        parameterMap: function(data, type) {
          if (type == "read") {
            // send take as "$top" and skip as "$skip"
            return {
              $top: data.take,
              $skip: data.skip
            }
          }
        }
      },
      schema: {
        data: "d"
      },
      pageSize: 20,
      serverPaging: true // enable serverPaging so take and skip are sent as request parameters
    });
    dataSource.fetch(function() {
      console.log(dataSource.view().length); // displays "20"
    });
    </script>


    Returns
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        create: {
          url: "http://demos.telerik.com/kendo-ui/service/products/create",
          dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
        },
        parameterMap: function(data, type) {
          return kendo.stringify(data);
        }
      },
      batch: true,
      schema: {
        model: { id: "ProductID" }
      }
    });
    dataSource.add( { ProductName: "New Product" });
    dataSource.sync();
    </script>


    Example - convert data source request parameters


    Example - send request parameters as JSON
*/
//


<h5>data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
</h5>

<p>The parameters which will be sent to the remote service. The value specified in the <code>data</code> field of the transport settings (create, read, update or destroy) is included as well. If 'Array' #batch-boolean-default"><code>batch</code></a> is set to <code>false</code>, the fields of the changed data items are also included.</p>

<h5>data.aggregate <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
</h5>

<p>The current aggregate configuration as set via the 'Array' #configuration-aggregate"><code>aggregate</code></a> option. Available if the <a href="#configuration-serverAggregates"><code>serverAggregates</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>data.group <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
</h5>

<p>The current grouping configuration as set via the 'Object' #configuration-group"><code>group</code></a> option. Available if the <a href="#configuration-serverGrouping"><code>serverGrouping</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>data.filter <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
</h5>

<p>The current filter configuration as set via the 'Array' #configuration-filter"><code>filter</code></a> option. Available if the <a href="#configuration-serverFiltering"><code>serverFiltering</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>data.models <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
</h5>

<p>All changed data items. Available if there are any data item changes and the 'Number' #configuration-batch"><code>batch</code></a> option is set to <code>true</code>.</p>

<h5>data.page <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number
</h5>

<p>The current page. Available if the 'Number' #configuration-serverPaging"><code>serverPaging</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>data.pageSize <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number
</h5>

<p>The current page size as set via the 'Number' #configuration-pageSize"><code>pageSize</code></a> option. Available if the <a href="#configuration-serverPaging"><code>serverPaging</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>data.skip <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number
</h5>

<p>The number of data items to skip. Available if the 'Array' #configuration-serverPaging"><code>serverPaging</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>data.sort <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array
</h5>

<p>The current sort configuration as set via the 'Number' #configuration-sort"><code>sort</code></a> option. Available if the <a href="#configuration-serverSorting"><code>serverSorting</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>data.take <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number
</h5>

<p>The number of data items to return (the same as <code>data.pageSize</code>). Available if the 'String' #configuration-serverPaging"><code>serverPaging</code></a> option is set to <code>true</code> and the data source makes a <code>"read"</code> request.</p>

<h5>type <a href="https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String
</h5>

<p>The type of the request which the data source makes.</p>

<p>The supported values are:</p>

<ul>
<li><code>"create"</code></li>
<li><code>"read"</code></li>
<li><code>"update"</code></li>
<li><code>"destroy"</code></li>
</ul>


//


<p>'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object  the request parameters converted to a format required by the remote service.</p>


//
ParameterMap    ComplexJavaScriptType

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.push
//
// Type: Function
//
// The function invoked during transport initialization which sets up push notifications. The data source will call this function only once and provide callbacks which will handle push notifications (data pushed from the server).
//
//
//
/*
  Parameters
    <script src="http://ajax.aspnetcdn.com/ajax/signalr/jquery.signalr-1.1.3.min.js"></script>
    <script>
    var hubUrl = "http://demos.telerik.com/kendo-ui/service/signalr/hubs";
    var connection = $.hubConnection(hubUrl, { useDefaultPath: false});
    var hub = connection.createHubProxy("productHub");
    var hubStart = connection.start({ jsonp: true });
    var dataSource = new kendo.data.DataSource({
    transport: {
      push: function(callbacks) {
        hub.on("create", function(result) {
          console.log("push create");
          callbacks.pushCreate(result);
        });
        hub.on("update", function(result) {
          console.log("push update");
          callbacks.pushUpdate(result);
        });
        hub.on("destroy", function(result) {
          console.log("push destroy");
          callbacks.pushDestroy(result);
        });
      }
    },
    schema: {
      model: {
        id: "ID",
        fields: {
          "ID": { editable: false },
          "CreatedAt": { type: "date" },
          "UnitPrice": { type: "number" }
        }
      }
    }
    });
    </script>


    Example
*/
//


<h5>callbacks 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
</h5>

<p>An object containing callbacks for notifying the data source of push notifications.</p>

<h5>callbacks.pushCreate 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
</h5>

<p>Function that should be invoked to notify the data source about newly created data items that are pushed from the server. Accepts a single argument - the object pushed from the server which should follow the <code>schema.data</code> configuration.</p>

<h5>callbacks.pushDestroy 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
</h5>

<p>Function that should be invoked to notify the data source about destroyed data items that are pushed from the server. Accepts a single argument - the object pushed from the server
which should follow the <code>schema.data</code> configuration.</p>

<h5>callbacks.pushUpdate 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
</h5>

<p>Function that should be invoked to notify the data source about updated data items that are pushed from the server. Accepts a single argument - the object pushed from the server
which should follow the <code>schema.data</code> configuration.</p>


//
Push    ComplexJavaScriptType

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr
//
// Type: Object
//
// The configuration used when 'type' http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-type  is set to "signalr". Configures the SignalR settings - hub, connection promise, server, and client hub methods.
//
// Live demo available at demos.telerik.com/kendo-ui http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://demos.telerik.com/kendo-ui/grid/signalr .
//
// It is recommended to get familiar with the SignalR JavaScript API http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://www.asp.net/signalr/overview/guide-to-the-api/hubs-api-guide-javascript-client .
//
//
//
/*
  Example
    <script src="http://ajax.aspnetcdn.com/ajax/signalr/jquery.signalr-1.1.3.min.js"></script>
    <script>
        var hubUrl = "http://demos.telerik.com/kendo-ui/service/signalr/hubs";
        var connection = $.hubConnection(hubUrl, { useDefaultPath: false});
        var hub = connection.createHubProxy("productHub");
        var hubStart = connection.start({ jsonp: true });

        var dataSource = new kendo.data.DataSource({
            type: "signalr",
            schema: {
                model: {
                    id: "ID",
                    fields: {
                        "ID": { editable: false, nullable: true },
                        "CreatedAt": { type: "date" },
                        "UnitPrice": { type: "number" }
                    }
                }
            },
            transport: {
                signalr: {
                    promise: hubStart,
                    hub: hub,
                    server: {
                        read: "read",
                        update: "update",
                        destroy: "destroy",
                        create: "create"
                    },
                    client: {
                        read: "read",
                        update: "update",
                        destroy: "destroy",
                        create: "create"
                    }
                }
            }
        });
    </script>
*/
//
Signalr    Signalr

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.submit
//
// Type: Function
//
// A function that will handle create, update and delete operations in a single batch.
//
// Typically, you have the 'transport.read' and 'transport.submit' operations defined together. The 'transport.create', 'transport.update' and 'transport.delete' operations will not be executed in this case.
//
//    This function will only be invoked when the DataSource is in batch mode.
//
/*
  Parameters
*/
//


<h5>e.data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
</h5>

<p>An object containing the created (<code>e.data.created</code>), updated (<code>e.data.updated</code>), and destroyed (<code>e.data.destroyed</code>) items.</p>

<h5>e.success 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
</h5>

<p>A callback that should be called for each operation with two parameters - items and operation. See example below.</p>

<h5>e.fail 'Function' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Function
</h5>

<p>A callback that should be called in case of failure of any of the operations.</p>

<h5>Example - declare transport submit function</h5>


//
Submit    ComplexJavaScriptType

Template    *Template
}

func ( el Transport ) getTemplate () string {
  return `{{if ne (string .ParameterMap) "null"}}parameterMap: {{string .ParameterMap}},{{end}}{{if ne (string .Push) "null"}}push: {{string .Push}},{{end}}{{if ne (string .Submit) "null"}}submit: {{string .Submit}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further info.
//
//
//
/*
  Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        create: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Create struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further info.
  //
  //
  //
  /*
    Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - set a content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.type
  //
  // Type: String
  // Defalt: GET
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to '"jsonp"'. JSONP always uses GET requests.
  //
  /*
    Example - set the HTTP verb of the request
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  Template    *Template
}

func ( el Create ) getTemplate () string {
  return `{{if .Cache}}batch: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
//
//
//
/*
  Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        destroy: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Destroy struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the jQuery.ajax http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - set the content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.type
  //
  // Type: String
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to '"jsonp"'. JSONP always uses GET requests.
  //
  /*
    Example
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  Template    *Template
}

func ( el Destroy ) getTemplate () string {
  return `{{if .Cache}}batch: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting cache to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
//
//
//
/*
  Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        read: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Read struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting cache to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - set content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.type
  //
  // Type: String
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to "jsonp". JSONP always uses GET requests.
  //
  /*
    Example - set the HTTP verb of the request
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  Template    *Template
}

func ( el Read ) getTemplate () string {
  return `{{if .Cache}}batch: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client
//
// Specifies the client-side CRUD methods of the SignalR hub.
//
//
//

//
type Signalr struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client
  //
  // Type: Object
  //
  // Specifies the client-side CRUD methods of the SignalR hub.
  //
  //
  //

  //
  Client    Client

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.hub
  //
  // Type: Object
  //
  // The SignalR hub object returned by the 'createHubProxy' method. The 'hub' option is mandatory.
  //
  //
  //

  //
  Hub    Hub

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.promise
  //
  // Type: Object
  //
  // The promise returned by the 'start' method of the SignalR connection. The 'promise' option is mandatory.
  //
  //
  //

  //
  Promise    Promise

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server
  //
  // Type: Object
  //
  // Specifies the server-side CRUD methods of the SignalR hub.
  //
  //
  //

  //
  Server    Server

  Template    *Template
}

func ( el Signalr ) getTemplate () string {
  return `{{if ne (string .Hub) "null"}}hub: {{string .Hub}},{{end}}{{if ne (string .Promise) "null"}}promise: {{string .Promise}},{{end}}`
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.create
//
// Specifies the name of the client-side method of the SignalR hub responsible for creating data items.
//
//
//

//
type Client struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.create
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for creating data items.
  //
  //
  //

  //
  Create    Create

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.destroy
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for destroying data items.
  //
  //
  //

  //
  Destroy    Destroy

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.read
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for reading data items.
  //
  //
  //

  //
  Read    Read

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.client.update
  //
  // Type: String
  //
  // Specifies the name of the client-side method of the SignalR hub responsible for updating data items.
  //
  //
  //

  //
  Update    Update

  Template    *Template
}

func ( el Client ) getTemplate () string {
  return ``
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.create
//
// Specifies the name of the server-side method of the SignalR hub responsible for creating data items.
//
//
//

//
type Server struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.create
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for creating data items.
  //
  //
  //

  //
  Create    Create

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.destroy
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for destroying data items.
  //
  //
  //

  //
  Destroy    Destroy

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.read
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for reading data items.
  //
  //
  //

  //
  Read    Read

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.signalr.server.update
  //
  // Type: String
  //
  // Specifies the name of the server-side method of the SignalR hub responsible for updating data items.
  //
  //
  //

  //
  Update    Update

  Template    *Template
}

func ( el Server ) getTemplate () string {
  return ``
}














// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.cache
//
// If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
//
// Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
//
//
//
/*
  Example - enable request caching
    <script>
    var dataSource = new kendo.data.DataSource({
      transport: {
        update: {
          /@ omitted for brevity @/
          cache: true
        }
      }
    });
    </script>
*/
//
type Update struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - enable request caching
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            cache: true
          }
        }
      });
      </script>
  */
  //
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. Default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON.
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/%60jQuery.ajax%60  documentation for further information.
  //
  //
  //
  /*
    Example - set content type
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            contentType: "application/json"
          }
        }
      });
      </script>
  */
  //
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
    Example - set the data type to JSON
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  //
  DataType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.type
  //
  // Type: String
  //
  // The type of request to make ("POST", "GET", "PUT" or "DELETE"). The default request is "GET".
  //
  //    The 'type' option is ignored if 'dataType' is set to '"jsonp"'. JSONP always uses GET requests.
  //
  /*
    Example - set the HTTP verb of the request
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            type: "POST"
          }
        }
      });
      </script>
  */
  //
  Type    string

  Template    *Template
}

func ( el Update ) getTemplate () string {
  return `{{if .Cache}}batch: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}`
}














