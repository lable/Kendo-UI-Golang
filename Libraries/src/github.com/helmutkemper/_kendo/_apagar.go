package kendo

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
type AggregatesEnum int

const(
  AGGREGATES_NIL AggregatesEnum   = iota
  AGGREGATES_AVERAGE
  AGGREGATES_COUNT
  AGGREGATES_MAX
  AGGREGATES_MIN
  AGGREGATES_SUM
)

var AggregatesEnums  = [...]string{
  "",
  "average",
  "count",
  "max",
  "min",
  "sum",
}

func (el AggregatesEnum ) String() string {
  return AggregatesEnums[ el ]
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
  return `{{if ne (string .Aggregate) "null"}}aggregate: {{string .Aggregate}},{{end}}
{{if ne (string .Field) "null"}}field: {{string .Field}},{{end}}
`
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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.data
  //
  // Type: Object
  //
  // Additional parameters that are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            data: {
              name: "Jane Doe",
              age: 30
            }
          }
        }
      });
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            /@ omitted for brevity @/
            data: function() {
              return {
                name: "Jane Doe",
                age: 30
              }
            }
          }
        }
      });
      </script>
  */
  //
  Data    ComplexJavaScriptType

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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify the URL as a string
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            url: "http://demos.telerik.com/kendo-ui/service/products/create",
            cache: true,
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "create") {
              return { models: kendo.stringify(data.models) }
            }
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


      Example - specify the URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          create: {
            url: function(options) {
              return "http://demos.telerik.com/kendo-ui/service/products/create"
            },
            cache: true,
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "create") {
              return { models: kendo.stringify(data.models) }
            }
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
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Create ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.data
  //
  // Type: Object
  //
  // Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            data: {
              name: "Jane Doe",
              age: 30
            }
          }
        }
      });
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          destroy: {
            /@ omitted for brevity @/
            data: function() {
              return {
                name: "Jane Doe",
                age: 30
              }
            }
          }
        }
      });
      </script>
  */
  //
  Data    ComplexJavaScriptType

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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.destroy.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify the URL as a string
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp"
          },
          destroy: {
            url: "http://demos.telerik.com/kendo-ui/service/products/destroy",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "destroy") {
              return { models: kendo.stringify(data.models) }
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        dataSource.remove(products[0]);
        dataSource.sync();
      });
      </script>


      Example - specify the URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp"
          },
          destroy: {
            url: function (options) {
              return "http://demos.telerik.com/kendo-ui/service/products/destroy",
            },
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          parameterMap: function(data, type) {
            if (type == "destroy") {
              return { models: kendo.stringify(data.models) }
            }
          }
        },
        batch: true,
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var products = dataSource.data();
        dataSource.remove(products[0]);
        dataSource.sync();
      });
      </script>
  */
  //
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Destroy ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.data
  //
  // Type: Object
  //
  // Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: {
              q: "html5" // send "html5" as the "q" parameter
            }
          }
        }
      });
      dataSource.fetch();
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: "http://demos.telerik.com/kendo-ui/service/twitter/search",
            dataType: "jsonp", // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
            data: function() {
              return {
                q: "html5" // send "html5" as the "q" parameter
              };
            }
          }
        }
      });
      dataSource.fetch();
      </script>
  */
  //
  Data    ComplexJavaScriptType

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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.read.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify URL as a string
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
        console.log(dataSource.view().length); // displays "77"
      });
      </script>


      Example - specify URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read: {
            url: function(options) {
              return "http://demos.telerik.com/kendo-ui/service/products",
            }
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        }
      });
      dataSource.fetch(function() {
        console.log(dataSource.view().length); // displays "77"
      });
      </script>
  */
  //
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Read ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
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
  return `{{if ne (string .Create) "null"}}create: {{string .Create}},{{end}}
{{if ne (string .Destroy) "null"}}destroy: {{string .Destroy}},{{end}}
{{if ne (string .Read) "null"}}read: {{string .Read}},{{end}}
{{if ne (string .Update) "null"}}update: {{string .Update}},{{end}}
`
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
  return `{{if ne (string .Create) "null"}}create: {{string .Create}},{{end}}
{{if ne (string .Destroy) "null"}}destroy: {{string .Destroy}},{{end}}
{{if ne (string .Read) "null"}}read: {{string .Read}},{{end}}
{{if ne (string .Update) "null"}}update: {{string .Update}},{{end}}
`
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
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.data
  //
  // Type: Object
  //
  // Additional parameters which are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for
  // sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
  //
  //
  //
  /*
      Example - send additional parameters as an object
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            data: {
              name: "Jane Doe",
              age: 30
            }
          }
        }
      });
      </script>


      Example - send additional parameters by returning them from a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          update: {
            /@ omitted for brevity @/
            data: function() {
              return {
                name: "Jane Doe",
                age: 30
              }
            }
          }
        }
      });
      </script>
  */
  //
  Data    ComplexJavaScriptType

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

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
  //
  //
  //
  /*
      Example - specify URL as a string
      <script>
      var dataSource = new kendo.data.DataSource({
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
        product.set("UnitPrice", 20);
        dataSource.sync();
      });
      </script>


      Example - specify URL as a function
      <script>
      var dataSource = new kendo.data.DataSource({
        transport: {
          read:  {
            url: "http://demos.telerik.com/kendo-ui/service/products",
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          },
          update: {
            url: function(options) {
              return "http://demos.telerik.com/kendo-ui/service/products/update",
            },
            dataType: "jsonp" // "jsonp" is required for cross-domain requests; use "json" for same-domain requests
          }
        },
        schema: {
          model: { id: "ProductID" }
        }
      });
      dataSource.fetch(function() {
        var product = dataSource.at(0);
        product.set("UnitPrice", 20);
        dataSource.sync();
      });
      </script>
  */
  //
  Url    ComplexJavaScriptType

  Template    *Template
}

func ( el Update ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "null"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "null"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Type) "null"}}type: {{string .Type}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
}













