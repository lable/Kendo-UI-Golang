package kendo

import "encoding/json"

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

  js, _ := json.Marshal( el.AsJSon )
  return string( js )
}