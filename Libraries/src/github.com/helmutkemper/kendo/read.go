package kendo

import "bytes"

type Read struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.cache
  //
  // Type: Boolean
  //
  // If set to 'false', the request result will not be cached by the browser. Setting 'cache' to 'false' will only work correctly with HEAD and GET requests. It works by appending "_={timestamp}" to the GET parameters. By default, "jsonp" requests are not cached.
  //
  // Refer to the 'jQuery.ajax' http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.ajax  documentation for further information.
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
  Data    ComplexJavaScriptType

  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.update.dataType
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
          update: {
            /@ omitted for brevity @/
            dataType: "json"
          }
        }
      });
      </script>
  */
  DataType    TypeDataJSonEnum

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
  Method    MethodEnum

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

  GoTemplate    *GoTemplate
}

func ( el Read ) getTemplate () string {
  return `{{if .Cache}}cache: true,{{end}}
{{if ne (string .ContentType) "''"}}contentType: {{string .ContentType}},{{end}}
{{if ne (string .Data) "null"}}data: {{string .Data}},{{end}}
{{if ne (string .DataType) "''"}}dataType: {{string .DataType}},{{end}}
{{if ne (string .Method) "''"}}type: {{string .Method}},{{end}}
{{if ne (string .Url) "null"}}url: {{string .Url}},{{end}}
`
}

func ( el Read ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "null" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Read ) String() string {
  out := el.Buffer()
  return out.String()
}
