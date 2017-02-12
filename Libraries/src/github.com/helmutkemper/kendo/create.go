package kendo

import "bytes"

type Create struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further info.
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
  Cache    bool

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.contentType
  //
  // Type: String
  //
  // The content-type HTTP header sent to the server. The default is "application/x-www-form-urlencoded". Use "application/json" if the content is JSON. Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
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
  ContentType    string

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.data
  //
  // Type: Object
  //
  // Additional parameters that are sent to the remote service. The parameter names must not match reserved words, which are used by the Kendo UI DataSource for sorting http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverSorting , filtering http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverFiltering , paging http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverPaging , and grouping http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-serverGrouping .
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
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
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.dataType
  //
  // Type: String
  //
  // The type of result expected from the server. Commonly used values are "json" and "jsonp".
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
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
  DataType    TypeDataJSonEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.type
  //
  // Type: String
  // Default: GET
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
  Method    MethodEnum

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.create.url
  //
  // Type: String
  //
  // The URL to which the request is sent.
  //
  // If set to function, the data source will invoke it and use the result as the URL.
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

  GoTemplate    *GoTemplate
}

func ( el Create ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "''"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "''"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Method) "''"}}type: {{string .Method}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
}

func ( el Create ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Create ) String() string {
  out := el.Buffer()
  return out.String()
}
