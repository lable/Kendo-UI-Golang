package kendo

import (
  "html/template"
  "io/ioutil"
  "bytes"
  "fmt"
  "reflect"
)

type GoTemplate struct{
  Name      string
  template *template.Template
  isInit    bool
}

func ( GoTemplateAStt *GoTemplate ) Init (){
  GoTemplateAStt.isInit = true

  GoTemplateAStt.template = template.New( GoTemplateAStt.Name ).
    Funcs(
    template.FuncMap{
      "htmlSafe": func(html string) template.HTML {
        return template.HTML(html)
      },
      "openScript": func() template.HTML {
        return template.HTML("<script>")
      },
      "closeScript": func() template.HTML {
        return template.HTML("</script>")
      },
      "nl": func() template.HTML {
        return template.HTML("\n")
      },
      "string": func( el interface{} ) template.HTML {
        switch el.( type ) {
        case string: return template.HTML( `'`+el.( string )+`'` )
        case AggregateLine: return template.HTML( el.( AggregateLine ).String() )
        case ComplexJavaScriptType: return template.HTML( el.( ComplexJavaScriptType ).String() )
        case OperatorEnum: return template.HTML( `'`+el.( OperatorEnum ).String()+`'` )
        case FilterLine: return template.HTML( el.( FilterLine ).String() )
        case Schema: return template.HTML( el.( Schema ).String() )
        case GroupLine: return template.HTML( el.( GroupLine ).String() )
        case SortLine: return template.HTML( el.( SortLine ).String() )
        case EffectEnum: return template.HTML( `'`+el.( EffectEnum ).String()+`'` )
        case Open: return template.HTML( el.( Open ).String() )
        case Close: return template.HTML( el.( Close ).String() )
        case Transport: return template.HTML( el.( Transport ).String() )
        case AggregateEnum: return template.HTML( `'`+el.( AggregateEnum ).String()+`'` )
        case FiltersLine: return template.HTML( el.( FiltersLine ).String() )
        case LogicEnum: return template.HTML( `'`+el.( LogicEnum ).String()+`'` )
        case MethodEnum: return template.HTML( `'`+el.( MethodEnum ).String()+`'` )
        case TypeDataJSonEnum: return template.HTML( `'`+el.( TypeDataJSonEnum ).String()+`'` )
        case Animation: return template.HTML( el.( Animation ).String() )
        case DataSource: return template.HTML( el.( DataSource ).String() )
        case AutoCompleteFilterEnum: return template.HTML( `'`+el.( AutoCompleteFilterEnum ).String()+`'` )
        case DirEnum: return template.HTML( `'`+el.( DirEnum ).String()+`'` )
        case MapValueToEnum: return template.HTML( `'`+el.( MapValueToEnum ).String()+`'` )
        case Client: return template.HTML( el.( Client ).String() )
        case Server: return template.HTML( el.( Server ).String() )
        case Create: return template.HTML( el.( Create ).String() )
        case Destroy: return template.HTML( el.( Destroy ).String() )
        case Read: return template.HTML( el.( Read ).String() )
        case Signalr: return template.HTML( el.( Signalr ).String() )
        case Update: return template.HTML( el.( Update ).String() )
        case FieldLine: return template.HTML( el.( FieldLine ).String() )

        case []FilterLine:
          var buffer bytes.Buffer
          for _, v := range el.( []FilterLine ){
            buffer.WriteString( v.String() )
          }
          return template.HTML( buffer.String() )

        case []FiltersLine:
          var buffer bytes.Buffer
          for _, v := range el.( []FiltersLine ){
            buffer.WriteString( v.String() )
          }
          return template.HTML( buffer.String() )

        case []SortLine:
          var buffer bytes.Buffer
          for _, v := range el.( []SortLine ){
            buffer.WriteString( v.String() )
          }
          return template.HTML( buffer.String() )

        case []GroupLine:
          var buffer bytes.Buffer
          for _, v := range el.( []GroupLine ){
            buffer.WriteString( v.String() )
          }
          return template.HTML( buffer.String() )

        case []string:
          var buffer bytes.Buffer
          for _, v := range el.( []string ){
            buffer.WriteString( `'` + v + `',` )
          }
          return template.HTML( buffer.String() )
        }

        fmt.Printf( "type conversion error: %v\n", reflect.TypeOf( el ) )
        return template.HTML( "" );
      },
    },
  )
}

func ( GoTemplateAStt *GoTemplate ) ParserString ( content string ){
  var err error

  if GoTemplateAStt.isInit == false {
    GoTemplateAStt.Init()
  }

  GoTemplateAStt.template, err = GoTemplateAStt.template.Parse( content )

  if err != nil {
    panic( err )
  }
}


func ( GoTemplateAStt *GoTemplate ) ParserFile ( file string ) {
  content, err := ioutil.ReadFile( file )

  if err != nil {
    panic( err )
  }

  GoTemplateAStt.ParserString( string( content ) )
}

func ( GoTemplateAStt *GoTemplate ) ParserFiles ( file ...string ) {
  var err error

  if GoTemplateAStt.isInit == false {
    GoTemplateAStt.Init()
  }

  GoTemplateAStt.template, err = GoTemplateAStt.template.ParseFiles( file... )

  if err != nil {
    panic( err )
  }
}

func ( GoTemplateAStt *GoTemplate ) Execute( data interface{} ) {
  if GoTemplateAStt.isInit == false {
    GoTemplateAStt.Init()
  }

  /*err := GoTemplateAStt.template.Execute( GoTemplateAStt.Out, data )
  if err != nil {
    panic( err )
  }*/
}

func ( GoTemplateAStt *GoTemplate ) ExecuteTemplate ( out *bytes.Buffer, name string, data interface{} ) {
  if GoTemplateAStt.isInit == false {
    GoTemplateAStt.Init()
  }

  err := GoTemplateAStt.template.ExecuteTemplate( out, name, data )
  if err != nil {
    panic( err )
  }
}
