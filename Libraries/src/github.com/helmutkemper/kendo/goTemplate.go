package kendo

import (
  "html/template"
  "io/ioutil"
  "bytes"
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
      "string": func( el interface{} ) string {
        switch el.( type ) {
        case AggregateLine: return el.( AggregateLine ).String()
        case ComplexJavaScriptType: return el.( ComplexJavaScriptType ).String()
        case OperatorEnum: return el.( OperatorEnum ).String()
        case FilterLine: return el.( FilterLine ).String()
        case Schema: return el.( Schema ).String()
        case GroupLine: return el.( GroupLine ).String()
        case SortLine: return el.( SortLine ).String()
        case EffectEnum: return el.( EffectEnum ).String()
        case Open: return el.( Open ).String()
        case Close: return el.( Close ).String()
        case Transport: return el.( Transport ).String()

        }

        return " - error: falta criar o tipo no template - ";
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
