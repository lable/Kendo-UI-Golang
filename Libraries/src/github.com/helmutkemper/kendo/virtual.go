package kendo

type Virtual struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual.itemHeight
  //
  // Type: Number
  // Default: null
  //
  // Specifies the height of the virtual item. All items in the virtualized list must have the same height.
  // If the developer does not specify one, the framework will automatically set 'itemHeight' based on the current theme and font size.
  //
  //
  //
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
  //
  ItemHeight    int64

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual.mapValueTo
  //
  // Type: String
  // Default: index
  //
  // The changes introduced with the Kendo UI R3 2016 release enable you to determine if the 'valueMapper' must resolve a value to an 'index' or a value to a 'dataItem'. This is configured through the 'mapValueTo' option that accepts two possible values - "index" or "dataItem". By default, the 'mapValueTo' is set to "index", which does not affect the current behavior of the virtualization process.
  //
  // For more information, refer to the article on virtualization http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete/kendo-ui/controls/editors/combobox/virtualization#value-mapping .
  //
  //
  //

  //
  MapValueTo    string

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual.valueMapper
  //
  // Type: Function
  // Default: null
  //
  //
  //
  //    Important
  //    As of the Kendo UI R3 2016 release, the implementation of the 'valueMapper' function is not necessary.
  //
  /*
      Example - AutoComplete with a virtualized list
      // the example is relevant to Kendo UI versions prior to 2016.3.914
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
                                  //Example:
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


      Example - add a data item to the data source
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        dataSource: [
          { name: "Apples" },
          { name: "Oranges" }
        ],
        dataTextField: "name"
      });
      var autocomplete = $("#autocomplete").data("kendoAutoComplete");
      autocomplete.dataSource.read();
      autocomplete.dataSource.add({ name: "Appricot" });
      autocomplete.search("A");
      </script>


      Example - get options of the widget
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete();

      var autocomplete = $("#autocomplete").data("kendoAutoComplete");

      var element = autocomplete.element;

      var options = autocomplete.options;
      </script>


      Example - get list element
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete();

      var autocomplete = $("#autocomplete").data("kendoAutoComplete");

      var list = autocomplete.list;
      </script>


      Example - get ul element
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete();

      var autocomplete = $("#autocomplete").data("kendoAutoComplete");

      var ul = autocomplete.ul;
      </script>
  */
  //
  ValueMapper    ComplexJavaScriptType

  Template    *Template
}

func ( el Virtual ) getTemplate () string {
  return `{{if .ItemHeight }}itemHeight: {{.ItemHeight}},{{end}}
{{if ne (string .MapValueTo) "null"}}mapValueTo: {{string .MapValueTo}},{{end}}
{{if ne (string .ValueMapper) "null"}}valueMapper: {{string .ValueMapper}},{{end}}
`
}

