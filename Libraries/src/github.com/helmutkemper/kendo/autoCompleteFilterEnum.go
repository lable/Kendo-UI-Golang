package kendo

// http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete#configuration-filter
//
// Type: String
// Default: startswith
//
// The filtering method used to determine the suggestions for the current value. The default filter is "startswith" -
// all data items which begin with the current widget value are displayed in the suggestion popup. The supported 'filter' values are 'startswith', 'endswith' and 'contains'.
/*
    Example - set the filter
    <input id="autocomplete" />
    <script>
    $("#autocomplete").kendoAutoComplete({
      filter: "contains"
    });
    </script>
*/
type AutoCompleteFilterEnum int

const(
  AUTOCOMPLETE_FILTER_NIL AutoCompleteFilterEnum   = iota
  AUTOCOMPLETE_FILTER_STARTSWITH
  AUTOCOMPLETE_FILTER_ENDSWITH
  AUTOCOMPLETE_FILTER_CONTAINS
)

var AutoCompleteFilterEnums  = [...]string{
  "",
  "startswith",
  "endswith",
  "contains",
}

func (el AutoCompleteFilterEnum ) String() string {
  return AutoCompleteFilterEnums[ el ]
}
