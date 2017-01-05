package kendo


type ComplexJavaScriptType struct {
  AsString      string
  AsFunction    string
  AsJSon        interface{}
}

func ( el ComplexJavaScriptType ) String() string {
  if el.AsString != "" {
    return `"` + el.AsString +  `"`
  } else if el.AsFunction != "" {
    return el.AsFunction
  }

  return ``
  //return `JSON.parse( "` + json.Marshal( el.AsJSon ) + `" )`
}