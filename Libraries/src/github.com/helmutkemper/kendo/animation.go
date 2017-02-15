package kendo

import (
  "bytes"
)

type Animation struct{

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.close
  //
  // Type: Object
  //
  // The animation played when the suggestion popup is closed.
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
  Close    Close

  // http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-animation.open
  //
  // Type: Object
  //
  // The animation played when the suggestion popup is opened.
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
  Open    Open

  GoTemplate    *GoTemplate
}

func ( el Animation ) getTemplate () string {
  return `{{if ne (string .Open) ""}}open: { {{string .Open}} },{{end}}{{if ne (string .Close) ""}}close: { {{string .Close}} }{{end}}`
}

func ( el Animation ) Buffer() bytes.Buffer {
  var buffer bytes.Buffer

  if el.GoTemplate == nil {
    buffer.WriteString( "" )
    return buffer
  }

  el.GoTemplate.ParserString( el.getTemplate() )
  el.GoTemplate.ExecuteTemplate( &buffer, "", el )

  return buffer
}

func ( el Animation ) String() string {
  out := el.Buffer()
  return out.String()
}