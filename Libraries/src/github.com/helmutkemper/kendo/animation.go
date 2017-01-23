package kendo

type Animation struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close
  //
  // Type: Object
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
  Close    Close

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.open
  //
  // Type: Object
  //
  // The animation played when the suggestion popup is opened.
  //
  //
  //
  /*
      Example - configure the open animation
      <input id="autocomplete" />
      <script>
      $("#autocomplete").kendoAutoComplete({
        animation: {
         open: {
           effects: "zoom:in",
           duration: 300
         }
        }
      });
      </script>
  */
  //
  Open    Open

  Template    *Template
}

func ( el Animation ) getTemplate () string {
  return `{{if ne (string .Close) "null"}}close: {{string .Close}},{{end}}
{{if ne (string .Open) "null"}}open: {{string .Open}},{{end}}
`
}
