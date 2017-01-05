package kendo

import (
  "html/template"
  "io/ioutil"
  "bytes"
)

type Template struct{
  Name      string
  template *template.Template
  isInit    bool
}

func ( TemplateAStt *Template ) Init (){
  TemplateAStt.isInit = true

  TemplateAStt.template = template.New( TemplateAStt.Name ).
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
        case AggregateList: return el.( AggregateList ).String()
        }

        return "error";
      },
    },
  )
}

func ( TemplateAStt *Template ) ParserString ( content string ){
  var err error

  if TemplateAStt.isInit == false {
    TemplateAStt.Init()
  }

  TemplateAStt.template, err = TemplateAStt.template.Parse( content )

  if err != nil {
    panic( err )
  }
}


func ( TemplateAStt *Template ) ParserFile ( file string ) {
  content, err := ioutil.ReadFile( file )

  if err != nil {
    panic( err )
  }

  TemplateAStt.ParserString( string( content ) )
}

func ( TemplateAStt *Template ) ParserFiles ( file ...string ) {
  var err error

  if TemplateAStt.isInit == false {
    TemplateAStt.Init()
  }

  TemplateAStt.template, err = TemplateAStt.template.ParseFiles( file... )

  if err != nil {
    panic( err )
  }
}

func ( TemplateAStt *Template ) Execute( data interface{} ) {
  if TemplateAStt.isInit == false {
    TemplateAStt.Init()
  }

  /*err := TemplateAStt.template.Execute( TemplateAStt.Out, data )
  if err != nil {
    panic( err )
  }*/
}

func ( TemplateAStt *Template ) ExecuteTemplate ( out *bytes.Buffer, name string, data interface{} ) {
  if TemplateAStt.isInit == false {
    TemplateAStt.Init()
  }

  err := TemplateAStt.template.ExecuteTemplate( out, name, data )
  if err != nil {
    panic( err )
  }
}
