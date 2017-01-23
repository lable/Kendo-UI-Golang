package kendo
















// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close
//
// The animation played when the suggestion popup is closed.
//
//
//
/*
    Example - configure the close animation
    <input id="autocomplete" />
    <script>
    $("#autocomplete").kendoAutoComplete({
      animation: {
       close: {
         effects: "zoom:out",
         duration: 300
       }
      }
    });
    </script>
*/
//















// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close.duration
//
// The duration of the close animation in milliseconds.
//
//
//

//















// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.open.duration
//
// The duration of the open animation in milliseconds.
//
//
//

//















// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-virtual.itemHeight
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













