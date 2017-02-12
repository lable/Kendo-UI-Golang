package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/data/datasource#configuration-transport.parameterMap
//
// The function which converts the request parameters to a format suitable for the remote service. By default, the data source sends the parameters using jQuery's conventions http://docs.telerik.com/kendo-ui/api/javascript/data/datasourcehttp://api.jquery.com/jQuery.param/ .
//
//    The 'parameterMap' method is often used to encode the parameters in JSON format.
//    The 'parameterMap' function will not be called when using custom functions for the read, update, create, and destroy operations.
//
/*
    Example - convert data source request parameters
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


    Example - send request parameters as JSON
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
*/
//
// Parameters:
// data 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object
// The parameters which will be sent to the remote service. The value specified in the data field of the transport settings (create, read, update or destroy) is included as well. If 'Array' #batch-boolean-default">batch is set to false, the fields of the changed data items are also included.
// data.aggregate aggregate option. Available if the serverAggregates option is set to true and the data source makes a "read" request.
// data.group group option. Available if the serverGrouping option is set to true and the data source makes a "read" request.
// data.filter filter option. Available if the serverFiltering option is set to true and the data source makes a "read" request.
// data.models batch option is set to true.
// data.page serverPaging option is set to true and the data source makes a "read" request.
// data.pageSize pageSize option. Available if the serverPaging option is set to true and the data source makes a "read" request.
// data.skip serverPaging option is set to true and the data source makes a "read" request.
// data.sort sort option. Available if the serverSorting option is set to true and the data source makes a "read" request.
// data.take serverPaging option is set to true and the data source makes a "read" request.
// type
//
// Return:
// 'Object' https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object  the request parameters converted to a format required by the remote service.
//
type ParameterMapEnum int

const(
  PARAMETERMAP_NIL ParameterMapEnum   = iota
  PARAMETERMAP_CREATE
  PARAMETERMAP_READ
  PARAMETERMAP_UPDATE
  PARAMETERMAP_DESTROY
)

var ParameterMapEnums  = [...]string{
  "",
  "create",
  "read",
  "update",
  "destroy",
}

func (el ParameterMapEnum ) String() string {
  return ParameterMapEnums[ el ]
}
