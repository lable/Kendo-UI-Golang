package kendo

import "bytes"

// The configuration used to parse the remote service response.
//
// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema
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
       return response.results; // twitter's response is { "results": [ // results ] }
      }
     }
    });
    dataSource.fetch(function(){
     var data = this.data();
     console.log(data.length);
    });
    </script>
*/
type Schema struct{
  // The configuration used to parse the remote service response.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-schema
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
         return response.results; // twitter's response is { "results": [ // results ] }
        }
       }
      });
      dataSource.fetch(function(){
       var data = this.data();
       console.log(data.length);
      });
      </script>
  */
  Data    ComplexJavaScriptType

  Template *Template
}

func ( el Schema ) getTemplate () string {
  return `data: {{string .Data}},`
}

func ( el Schema ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.Template == nil {
    buffer.WriteString( "null" )
    return buffer
  }

  el.Template.ParserString( el.getTemplate() )
  el.Template.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Schema ) String() string {
  out := el.Buffer()
  return out.String()
}