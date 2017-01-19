package kendo

// Enables the virtualization feature of the widget. The configuration can be set on an object, which contains two properties - 'itemHeight' and 'valueMapper'.
//
// For detailed information, refer to the article on virtualization ( http://docs.telerik.com/kendo-ui/controls/editors/combobox/virtualization ).
//
// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual
/*
    Example - AutoComplete with a virtualized list
    <input id="orders" style="width: 400px" />
    <script>
      $(document).ready(function() {
        $("#orders").kendoAutoComplete({
          template: "#= OrderID # | For: #= ShipName #, #= ShipCountry #",
          dataTextField: "ShipName",
          virtual: true,
          height: 520,
          dataSource: {
            type: "odata",
            transport: {
              read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
            },
            schema: {
              model: {
                fields: {
                  OrderID: { type: "number" },
                  Freight: { type: "number" },
                  ShipName: { type: "string" },
                  OrderDate: { type: "date" },
                  ShipCity: { type: "string" }
                }
              }
            },
            pageSize: 80,
            serverPaging: true,
            serverFiltering: true
          }
        });
      });
    </script>
*/
type Virtual struct {
  // default: null
  // Specifies the height of the virtual item. All items in the virtualized list *must* have the same height. If the developer does not specify one, the framework will automatically set 'itemHeight' based on the current theme and font size.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual.itemHeight
  /*
      Example - AutoComplete with a virtualized list
      <input id="orders" style="width: 400px" />
      <script>
        $(document).ready(function() {
          $("#orders").kendoAutoComplete({
            template: "#= OrderID # | For: #= ShipName #, #= ShipCountry #",
            dataTextField: "ShipName",
            virtual: {
              itemHeight: 26
            },
            height: 520,
            dataSource: {
              type: "odata",
              transport: {
                read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
              },
              schema: {
                model: {
                  fields: {
                    OrderID: { type: "number" },
                    Freight: { type: "number" },
                    ShipName: { type: "string" },
                    OrderDate: { type: "date" },
                    ShipCity: { type: "string" }
                  }
                }
              },
              pageSize: 80,
              serverPaging: true,
              serverFiltering: true
            }
          });
        });
      </script>
  */
  ItemHeight  int64

  // default: "index"
  // The changes introduced with the Kendo UI R3 2016 release enable you to determine if the 'valueMapper' must resolve a <em>value to an 'index'</em> or a <em>value to a 'dataItem'</em>. This is configured through the 'mapValueTo' option that accepts two possible values - '"index"' or '"dataItem"'. By default, the 'mapValueTo' is set to '"index"', which does not affect the current behavior of the virtualization process.
  //
  // For more information, refer to the article on virtualization ( http://docs.telerik.com/kendo-ui/controls/editors/combobox/virtualization#value-mapping ).
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual.mapValueTo
  /*
      Example - AutoComplete with a virtualized list
      <input id="orders" style="width: 400px" />
      <script>
        $(document).ready(function() {
          $("#orders").kendoAutoComplete({
            template: '<span class="order-id">#= OrderID #</span> #= ShipName #, #= ShipCountry #',
            dataTextField: "ShipName",
            virtual: {
              itemHeight: 26,
              valueMapper: function(options) {
                $.ajax({
                  url: "http://demos.telerik.com/kendo-ui/service/Orders/ValueMapper",
                  type: "GET",
                  dataType: "jsonp",
                  data: convertValues(options.value),
                  success: function (data) {
                    //the **data** is either index or array of indices.
                    //<b>Example</b>:
                    // "Ernst Handel" -> 10 (index in the Orders collection)
                    // ["Ernst Handel", "Que Delícia"] -> [10, 14] (indices in the Orders collection)

                    options.success(data);
                  }
                })
              }
            },
            height: 520,
            dataSource: {
              type: "odata",
              transport: {
                read: "http://demos.telerik.com/kendo-ui/service/Northwind.svc/Orders"
              },
              pageSize: 80,
              serverPaging: true,
              serverFiltering: true
            }
          });
        });

        function convertValues(value) {
          var data = {};

          value = $.isArray(value) ? value : [value];

          for (var idx = 0; idx < value.length; idx++) {
            data["values[" + idx + "]"] = value[idx];
          }

          return data;
        }
      </script>
  */
  MapValueTo  string

  // default: null
  // The 'valueMapper' function is *mandatory* for the functionality of the virtualized widget. The widget calls the 'valueMapper' function when the widget receives a value, that is not fetched from the remote server yet. The widget will pass the selected value(s) in the 'valueMapper' function. In turn, the valueMapper implementation should return the *respective data item(s) index/indices*.
  //
  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual.valueMapper
  //ValueMapper  Function
}