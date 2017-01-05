package kendo

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "regexp"
)







func main() {
  res, err := http.Get("http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete")
  if err != nil {
    log.Fatal(err)
  }
  robots, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
    log.Fatal(err)
  }
  var pageHtmlLStr string = string( robots )

  re := regexp.MustCompile( "(?si)<h3 id=\"configuration.*?>.*?(?:<h3)" )
  matchesMainLArr := re.FindAllStringSubmatch( pageHtmlLStr, -1 )
  for _, v := range matchesMainLArr{
    fmt.Println( "1>------------------------------------------------------------------------------------------" )
    fmt.Println( v )
    fmt.Println( "1<------------------------------------------------------------------------------------------" )
  }

  re  = regexp.MustCompile( "(?si)<article>.*?id=[\"'](.*?)[\"'].*?href=[\"'](#.*?)[\"'].*?>(.*?)</a>.*?<p>(.*?)</p>" )
  matchesMainDataLArr := re.FindAllStringSubmatch( pageHtmlLStr, -1 )
  for _, v := range matchesMainDataLArr{
    fmt.Println( "2>------------------------------------------------------------------------------------------" )
    fmt.Println( v )
    fmt.Println( "2<------------------------------------------------------------------------------------------" )
  }

  re  = regexp.MustCompile( "(?si)<article>.*?(<blockquote>.*?</blockquote>).*?(?:<h3)" )
  matchesBlockQuoteDataLArr := re.FindAllStringSubmatch( pageHtmlLStr, -1 )
  for _, v := range matchesBlockQuoteDataLArr{
    fmt.Println( "3>------------------------------------------------------------------------------------------" )
    fmt.Println( v )
    fmt.Println( "3<------------------------------------------------------------------------------------------" )
  }
  classNameCStr := matchesMainDataLArr[0][3]
  fmt.Printf( "class name: %v\n", classNameCStr )

  //todo strip_tags
  classDescriptionLStr := matchesMainDataLArr[0][4];
  fmt.Printf( "class description: %v\n", classDescriptionLStr )

  //todo matchesBlockQuoteDataLArr[ 0 ][ 1 ] pode não existir
  warningLStr := matchesBlockQuoteDataLArr[ 0 ][ 1 ]
  fmt.Printf( "class warning: %v\n", warningLStr )

  for _, matchesMainValueLStr := range matchesMainLArr[0]{

    //preg_match_all ( "%<a.*?href=[\"'](#.*?)[\"']>(.*?)</a>%s", $matchesMainValueLStr, $matchesLinkLArr );
    re  = regexp.MustCompile( "(?si)<a.*?href=[\"'](#.*?)[\"']>(.*?)</a>" )
    matchesLinkLArr := re.FindAllStringSubmatch( matchesMainValueLStr, -1 )
    fmt.Printf( "matchesLinkLArr: %v\n", matchesLinkLArr )

    //preg_match_all ( "%<h3.*?h3>%s", $matchesMainValueLStr, $matchesH3LArr );
    re  = regexp.MustCompile( "(?si)<h3.*?h3>" )
    matchesH3LArr := re.FindAllStringSubmatch( matchesMainValueLStr, -1 )
    fmt.Printf( "matchesH3LArr: %v\n", matchesH3LArr )

    //preg_match_all ( "%<code>(.*?)</code>%s", $matchesH3LArr[ 0 ][ 0 ], $matchesCodeLArr );
    re  = regexp.MustCompile( "(?si)<code>(.*?)</code>" )
    matchesCodeLArr := re.FindAllStringSubmatch( matchesH3LArr[0][0], -1 )
    fmt.Printf( "matchesCodeLArr: %v\n", matchesCodeLArr )

    //preg_match_all ( "%<h4>(.*?)</h4>%s", $matchesMainValueLStr, $matchesH4LArr );
    re  = regexp.MustCompile( "(?si)<h4>(.*?)</h4>" )
    matchesH4LArr := re.FindAllStringSubmatch( matchesMainValueLStr, -1 )
    fmt.Printf( "matchesH4LArr: %v\n", matchesH4LArr )

    //preg_match_all ( "%(<h4>.*?</h4>)*[\r\n]*<pre><code>(.*?)</code></pre>%s", $matchesMainValueLStr, $matchesExampleLArr );
    re  = regexp.MustCompile( "(?si)(<h4>.*?</h4>)*[\r\n]*<pre><code>(.*?)</code></pre>" )
    matchesExampleLArr := re.FindAllStringSubmatch( matchesMainValueLStr, -1 )
    fmt.Printf( "matchesH4LArr: %v\n", matchesExampleLArr )

    //preg_match_all ( "%<blockquote>.*?<p>(.*?)</p>.*?</blockquote>%s", $matchesMainValueLStr, $matchesBlockQuoteLArr );
    re  = regexp.MustCompile( "(?si)<blockquote>.*?<p>(.*?)</p>.*?</blockquote>" )
    matchesBlockQuoteLArr := re.FindAllStringSubmatch( matchesMainValueLStr, -1 )
    fmt.Printf( "matchesBlockQuoteLArr: %v\n", matchesBlockQuoteLArr )

    //preg_match_all ( "%<p>(.*?)</p>%s", $matchesMainValueLStr, $matchesPLArr );
    re  = regexp.MustCompile( "(?si)<p>(.*?)</p>" )
    matchesPLArr := re.FindAllStringSubmatch( matchesMainValueLStr, -1 )
    fmt.Printf( "matchesPLArr: %v\n", matchesPLArr )

    //preg_match_all ( "%<em>.*?default: *(.*?)\)</em>%s", $matchesH3LArr[ 0 ][ 0 ], $matchesDefaultLArr );
    re  = regexp.MustCompile( "(?si)<em>.*?default: *(.*?)[)]</em>" )
    matchesDefaultLArr := re.FindAllStringSubmatch( matchesH3LArr[0][0], -1 )
    fmt.Printf( "matchesDefaultLArr: %v\n", matchesDefaultLArr )

    //preg_match_all ( "%<h4><em>(.*?)</em></h4>[ \r\n]*<p>(.*?)</p>%s", $matchesMainValueLStr, $matchesPH4LArr );
    re  = regexp.MustCompile( "(?si)<h4><em>(.*?)</em></h4>[ \r\n]*<p>(.*?)</p>" )
    matchesPH4LArr := re.FindAllStringSubmatch( matchesMainValueLStr, -1 )
    fmt.Printf( "matchesPH4LArr: %v\n", matchesPH4LArr )
  }
}