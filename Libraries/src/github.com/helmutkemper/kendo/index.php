<?php

error_reporting(0);

// This class use dir ""./class/output/" to make files
//$o = new classMaker( "http://docs.telerik.com/kendo-ui/api/javascript/ui/autocomplete" );
$o = new classMaker( "http://docs.telerik.com/kendo-ui/api/javascript/data/datasource" );

class classMaker
{
  private $linkCStr;
  private $toXmlCheckCArr;
  private $classNameCStr;
  public $return;
  
  function classMaker ( $linkAStr )
  {
    $this->toXmlCheckCArr = array ();
    
    // Caminho do arquivo
    $this->linkCStr = $linkAStr;
    
    $pageHtmlLArr = file ( $linkAStr );
    $pageHtmlLStr = implode ( "", $pageHtmlLArr );
    $pageHtmlLStr = preg_replace( "%(^.*?)(<h3 id=\"configuration-.*)%si", "$2", $pageHtmlLStr );
    
    preg_match_all( "%(<h3 id=\"configuration-.*?)(?:<h2 id=\"methods)%si", $pageHtmlLStr, $matchesMethodsBlocks );
    $matchesMethodsBlocks = $matchesMethodsBlocks[ 1 ][ 0 ] . "<h3 id=\"configuration-";
    $matchesMethodsBlocks = preg_split( "%(<h3 id=\"configuration)%si", $matchesMethodsBlocks );
    unset( $matchesMethodsBlocks[0], $matchesMethodsBlocks[count($matchesMethodsBlocks)] );
    
    foreach( $matchesMethodsBlocks as $methodsBlocksKey => $methodsBlocksLine ){
      $configurationName[ $methodsBlocksKey ] = preg_replace( "%^-(.*?)(?:\">.*)%si", "$1", $methodsBlocksLine );
      preg_match_all( "%<a href=\"(.*?)\"%si", $methodsBlocksLine, $matchesMethodsLinks );
      $configurationLink[ $methodsBlocksKey ] = $linkAStr . $matchesMethodsLinks[ 1 ][ 0 ];
      
      preg_match_all( "%^-.*?\">\s*<a href=\".*?\">.*?</a>(.*?)</h3>%si", $methodsBlocksLine, $matchesMethodsTypes );
  
      preg_match_all( "%<em>\s*\(\s*default:\s*\"*(.*?)\"*\s*\)\s*</em>%si", $matchesMethodsTypes[ 1 ][ 0 ], $matchesMethodsDefault );
      $configurationDefault[ $methodsBlocksKey ] = $matchesMethodsDefault[ 1 ][ 0 ];
      
      preg_match_all( "%<a href=\"(.*?)\".*?>\s*<code>(.*?)\s*\|*</code>\s*</a>%si", $matchesMethodsTypes[ 1 ][ 0 ], $matchesMethodsTypes );
      $configurationType[ $methodsBlocksKey ] = array( $matchesMethodsTypes[ 1 ], $matchesMethodsTypes[ 2 ] );
  
      preg_match_all( "%<h4>Parameters</h4>(.*?)(?:<h4>Example|<script|<h4>Returns)%si", $methodsBlocksLine . "<endLine>", $matchesMethodsTypes );
      $configurationParameters[ $methodsBlocksKey ] = trim( strip_tags( preg_replace( "%<a href=\"(.*?)\" class=\"type-link\"><code>(.*?)\s*\|*</code></a>%si", "'$2' $1 ", $matchesMethodsTypes[ 1 ][ 0 ] ) ) );
      if( $configurationParameters[ $methodsBlocksKey ] ) {
        $configurationParameters[$methodsBlocksKey] = preg_split("%[\n\r]+%", $configurationParameters[$methodsBlocksKey]);
        foreach ($configurationParameters[$methodsBlocksKey] as $parameterKey => $parameterValue) {
          $configurationParameters[$methodsBlocksKey][$parameterKey] = "// " . $parameterValue;
        }
        $configurationParameters[$methodsBlocksKey] = implode("\n", $configurationParameters[$methodsBlocksKey]);
      }
      
      preg_match_all( "%<h4>Returns</h4>(.*?)(?:<h4>Example|<script)%si", $methodsBlocksLine . "<endLine>", $matchesMethodsTypes );
      $configurationReturns[ $methodsBlocksKey ] = trim( strip_tags( preg_replace( "%<a href=\"(.*?)\" class=\"type-link\"><code>(.*?)\s*\|*</code></a>%si", "'$2' $1 ", $matchesMethodsTypes[ 1 ][ 0 ] ) ) );
      if( $configurationReturns[ $methodsBlocksKey ] ) {
        $configurationReturns[$methodsBlocksKey] = preg_split("%[\n\r]+%", $configurationReturns[$methodsBlocksKey]);
        foreach ($configurationReturns[$methodsBlocksKey] as $parameterKey => $parameterValue) {
          $configurationReturns[$methodsBlocksKey][$parameterKey] = "// " . $parameterValue;
        }
        $configurationReturns[$methodsBlocksKey] = implode("\n", $configurationReturns[$methodsBlocksKey]);
      }
      
      preg_match_all( "%</h3>(.*?)(?:<blockquote>|<h4>|<endLine>)%si", $methodsBlocksLine . "<endLine>", $matchesMethodsTypes );
      $configurationDescription[ $methodsBlocksKey ] = $matchesMethodsTypes[ 1 ][ 0 ];
      $configurationDescription[ $methodsBlocksKey ] = trim( $configurationDescription[ $methodsBlocksKey ] );
      $configurationDescription[ $methodsBlocksKey ] = preg_replace( "%(<code>)([^\"].*?)(</code>)%si", "'$2'", $configurationDescription[ $methodsBlocksKey ] );
      $configurationDescription[ $methodsBlocksKey ] = preg_replace( "%(<a href=\")(.*?)(\">)(.*?)(</a>)%si", "$4 {$linkAStr}$2 ", $configurationDescription[ $methodsBlocksKey ] );
      $configurationDescription[ $methodsBlocksKey ] = preg_split( "%[\r\n]%si", $configurationDescription[ $methodsBlocksKey ] );
      foreach( $configurationDescription[ $methodsBlocksKey ] as $descriptionKey => $descriptionLine ){
        $configurationDescription[ $methodsBlocksKey ][ $descriptionKey ] = "  // " . trim( strip_tags( str_replace( "<li><code>", " *  ", $configurationDescription[ $methodsBlocksKey ][ $descriptionKey ] ) ) );
      }
      $configurationDescription[ $methodsBlocksKey ] = implode( "\n", $configurationDescription[ $methodsBlocksKey ] );
      $configurationDescription[ $methodsBlocksKey ] = preg_replace( "%// \n// (?:\n)%si", "", $configurationDescription[ $methodsBlocksKey ] );
  
      preg_match_all( "%<blockquote>(.*?)</blockquote>%si", $methodsBlocksLine, $matchesMethodsTypes );
      
      preg_match_all( "%<a href=\"(.*?)\"><code>(.*?)\s*</code></a>\s*option\s+is\s+set\s+to\s*<code>\s*true\s*</code>%si", $matchesMethodsTypes[ 1 ][ 0 ], $matchesEspecialType );
      $configurationSpecialType[ $methodsBlocksKey ] = $matchesEspecialType[ 2 ][ 0 ];
      
      $configurationSpotlight[ $methodsBlocksKey ] = preg_replace( "%<a href=\"(.*?)\"><code>(.*?)</code></a>%si", "'$2' {$linkAStr}$1", $matchesMethodsTypes[ 1 ][ 0 ] );
      $configurationSpotlight[ $methodsBlocksKey ] = "  //    " . trim( strip_tags( preg_replace( "%<code>(.*?)</code>%si", "'$1'", $configurationSpotlight[ $methodsBlocksKey ] ) ) );
  
      preg_match_all( "%<h4>(.*?)</h4>%si", $methodsBlocksLine, $matchesMethodsExamplesTitles );
      preg_match_all( "%<pre><code>(.*?)</code></pre>%si", $methodsBlocksLine, $matchesMethodsExamplesCode );
      foreach( $matchesMethodsExamplesTitles[ 1 ] as $matchesMethodsExamplesTitlesKey => $matchesMethodsExamplesTitlesLine ){
        if( !isset( $configurationExample[ $methodsBlocksKey ] ) ){
          $configurationExample[ $methodsBlocksKey ] = "";
        }
        $configurationExample[ $methodsBlocksKey ] .= $matchesMethodsExamplesTitles[ 1 ][ $matchesMethodsExamplesTitlesKey ] . "\n";
        $configurationExample[ $methodsBlocksKey ] .= $matchesMethodsExamplesCode[ 1 ][ $matchesMethodsExamplesTitlesKey ] . "\n\n";
      }
      $configurationExample[ $methodsBlocksKey ]  = preg_split( "%[\r\n]%", $configurationExample[ $methodsBlocksKey ] );
      foreach( $configurationExample[ $methodsBlocksKey ] as $exampleKey => $exampleLine ){
        $configurationExample[ $methodsBlocksKey ][ $exampleKey ] = "      " . $exampleLine;
      }
      $configurationExample[ $methodsBlocksKey ]  = implode( "\n", $configurationExample[ $methodsBlocksKey ] );
      $configurationExample[ $methodsBlocksKey ]  = str_replace( array( "/*", "*/" ), array( "/@", "@/" ), $configurationExample[ $methodsBlocksKey ] );
      $configurationExample[ $methodsBlocksKey ]  = trim( htmlspecialchars_decode( $configurationExample[ $methodsBlocksKey ] ) );
      if( $configurationExample[ $methodsBlocksKey ] ) {
        $configurationExample[$methodsBlocksKey] = "  /*\n    " . $configurationExample[$methodsBlocksKey] . "\n  */";
      }
    }
    
    $structList = [];
    foreach( $configurationName as $keyToGet => $nameToStruct ){
      $structs = explode( ".", $nameToStruct );
      foreach( $structs as $structsKey => $structsValue ){
        if( $structsKey == 0 ){
          if( !isset( $structList[ basename( $linkAStr ) ][ $structsValue ] ) ) {
            $structList[basename($linkAStr)][$structsValue] = array(
              "type" => $configurationType[ $keyToGet ][ 1 ],
              "default" => $configurationDefault[ $keyToGet ],
              "specialType" => $configurationSpecialType[ $keyToGet ],
              "typeLink" => $configurationType[ $keyToGet ][ 0 ],
              "helpLink" => $configurationLink[ $keyToGet ],
              "parameters" => $configurationParameters[ $keyToGet ],
              "return" => $configurationReturns[ $keyToGet ],
              "description" => $configurationDescription[ $keyToGet ],
              "spotlight" => $configurationSpotlight[ $keyToGet ],
              "example" => $configurationExample[ $keyToGet ],
            );
          }
        }
        else{
          if( !isset( $structList[ $structs[$structsKey - 1 ] ][ $structsValue ] ) ) {
            $structList[ $structs[$structsKey - 1 ] ][ $structsValue ] = array(
              "type" => $configurationType[ $keyToGet ][ 1 ],
              "default" => $configurationDefault[ $keyToGet ],
              "specialType" => $configurationSpecialType[ $keyToGet ],
              "typeLink" => $configurationType[ $keyToGet ][ 0 ],
              "helpLink" => $configurationLink[ $keyToGet ],
              "parameters" => $configurationParameters[ $keyToGet ],
              "return" => $configurationReturns[ $keyToGet ],
              "description" => $configurationDescription[ $keyToGet ],
              "spotlight" => $configurationSpotlight[ $keyToGet ],
              "example" => $configurationExample[ $keyToGet ],
            );
          }
        }
      }
    }
  
    $output = array();
    $outputRef = null;
    foreach( $structList as $structName => $structData ){
      $model = "";
      foreach( $structData as $itemName => $itemData ){
  
        if( !isset( $output[ $structName ] ) ) {
          $output[ $structName ] = "";
          $outputRef = &$output[ $structName ];
          
          $outputRef .= "// " . $itemData[ "helpLink" ] . "\n// \n";
  
          $outputRef .= str_replace( "  //", "//", $itemData[ "description" ] ) . "\n// \n";
          $outputRef .= str_replace( "  //", "//", $itemData[ "spotlight" ] ) . "\n// \n";
          $outputRef .= str_replace( "\n  ", "\n", substr( $itemData[ "example" ], 2 ) ) . "\n// \n";
  
          if( $itemData[ "parameters" ] ) {
            $outputRef .= "// Parameters:\n";
            $outputRef .= $itemData["parameters"] . "\n// \n";
          }
          if( $itemData[ "return" ] ) {
            $outputRef .= "// Return:\n";
            $outputRef .= $itemData["return"] . "\n// \n";
          }
          $outputRef .= "type " . ucfirst( $structName ) . " struct{\n\n";
        }
        
        if( count( $itemData[ "type" ] ) == 1 ){

          if( isset( $structList[ $itemName ] ) ){
            $type = ucfirst( $itemName );
            $nameUcFirst = ucfirst( $itemName );
            if( $itemData[ "type" ][ 0 ] == "Array" ){
              $type = "[]{$type}Line";
  
              if( $itemData[ "specialType" ] ){
                $testName = ucfirst( $itemData[ "specialType" ] );
              }
              else{
                $testName = $nameUcFirst;
              }
  
              $model .= "{{if .{$testName}}}{$itemName}: [{{range \$v := .{$nameUcFirst}}}{{string \$v}},{{end}}],{{end}}\n";
            }
            else{
  
              if( $itemData[ "specialType" ] ){
                $testName = ucfirst( $itemData[ "specialType" ] );
                $model .= "{{if .{$testName}}}{$itemName}: {{string .{$type}}},{{end}}\n";
              }
              else{
                $model .= "{{if ne (string .{$type}) \"null\"}}{$itemName}: {{string .{$type}}},{{end}}\n";
              }
            }
          }
          else{
            $nameUcFirst = ucfirst( $itemName );
            switch ( $itemData[ "type" ][ 0 ] ){
              case "String":
                $type = "string";
                
                if( $itemData[ "specialType" ] ){
                  $testName = ucfirst( $itemData[ "specialType" ] );
                  $model .= "{{if ne (string .{$nameUcFirst}) \"null\" and .{{{$testName}}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
                }
                else{
                  $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
                }
                break;
  
              case "Boolean":
                $type = "bool";
                $model .= "{{if .{$nameUcFirst}}}{$itemName}: true,{{end}}\n";
                break;
  
              case "Number":
                $type = "int64";
  
                if( $itemData[ "specialType" ] ){
                  $testName = ucfirst( $itemData[ "specialType" ] );
                }
                else{
                  $testName = $nameUcFirst;
                }
  
                $model .= "{{if .{$testName} and .{$nameUcFirst}}}{$itemName}: {{.{$nameUcFirst}}},{{end}}\n";
                break;
  
              case "Function":
                $type = "ComplexJavaScriptType";
  
                if( $itemData[ "specialType" ] ){
                  $testName = ucfirst( $itemData[ "specialType" ] );
                  $model .= "{{if ne (string .{$nameUcFirst}) \"null\" and .{$testName}}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
                }
                else{
                  $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
                }
                break;
  
              case "Object":
              case "Array":
                $type = ucfirst( $itemName );
  
                if( $itemData[ "specialType" ] ){
                  $testName = ucfirst( $itemData[ "specialType" ] );
                  $model .= "{{if ne (string .{$nameUcFirst}) \"null\" and .{$testName}}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
                }
                else{
                  $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
                }
                break;
  
              default:
                die( "falta criar um tipo de variável" );
            }
          }
                      
          $outputRef .= "  // " . $itemData[ "helpLink" ] . "\n  // \n";
          $outputRef .= "  // Type: " . $itemData[ "type" ][ 0 ] . "\n";
  
          if( $itemData[ "default" ] ){
            $outputRef .= "  // Defalt: " . $itemData[ "default" ] . "\n  // \n";
          }
          else{
            $outputRef .= "  // \n";
          }
  
          $outputRef .= $itemData[ "description" ] . "\n  // \n";
          $outputRef .= $itemData[ "spotlight" ] . "\n  // \n";
          $outputRef .= $itemData[ "example" ] . "\n  // \n";
  
          if( $itemData[ "parameters" ] ) {
            $outputRef .= $itemData["parameters"] . "\n  // \n";
          }
          if( $itemData[ "return" ] ) {
            $outputRef .= $itemData["return"] . "\n  // \n";
          }

          $outputRef .= "  " . ucfirst($itemName) . "    " . $type . "\n\n";
        }
        else{
          $nameUcFirst = ucfirst( $itemName );
          if( count( $itemData[ "type" ] ) == 2 ){
            
            if( in_array( "Array", $itemData[ "type" ] ) && in_array( "String", $itemData[ "type" ] ) ){
              $type = "ComplexJavaScriptType";
              $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
            }
            else if( in_array( "Function", $itemData[ "type" ] ) && in_array( "String", $itemData[ "type" ] ) ){
              $type = "ComplexJavaScriptType";
              $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
            }
            else if( in_array( "Object", $itemData[ "type" ] ) && in_array( "String", $itemData[ "type" ] ) ){
              $type = "ComplexJavaScriptType";
              $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
            }
            else if( in_array( "Array", $itemData[ "type" ] ) && in_array( "Object", $itemData[ "type" ] ) ){
              $type = ucfirst( $itemName );
  
              if( $itemData[ "specialType" ] ){
                $testName = ucfirst( $itemData[ "specialType" ] );
                $model .= "{{if ne (string .{$nameUcFirst}) \"null\" and .{$testName}}}{{\$length := len .$type}}{{if le \$length 1}}{$itemName}: { {{string .{$nameUcFirst}}} },{{else}}{$itemName}: [ {{string .{$nameUcFirst}}} ],{{end}},{{end}}\n";
              }
              else{
                $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{{\$length := len .$type}}{{if le \$length 1}}{$itemName}: { {{string .{$nameUcFirst}}} },{{else}}{$itemName}: [ {{string .{$nameUcFirst}}} ],{{end}},{{end}}\n";
              }
            }
            else if( in_array( "Object", $itemData[ "type" ] ) && in_array( "Function", $itemData[ "type" ] ) ){
              $type = "ComplexJavaScriptType";
              $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
            }
            else{
              print "";
            }
          }
          else if( count( $itemData[ "type" ] ) == 3 ){
            if( in_array( "Object", $itemData[ "type" ] ) && in_array( "Function", $itemData[ "type" ] ) && in_array( "String", $itemData[ "type" ] ) ){
              $type = "ComplexJavaScriptType";
              $model .= "{{if ne (string .{$nameUcFirst}) \"null\"}}{$itemName}: {{string .{$nameUcFirst}}},{{end}}\n";
            }
          }
          else{
            print "";
          }
  
          $outputRef .= "  // " . $itemData[ "helpLink" ] . "\n  // \n";
          $outputRef .= "  // Type: " . $itemData[ "type" ][ 0 ] . "\n";
  
          if( $itemData[ "default" ] ){
            $outputRef .= "  // Defalt: " . $itemData[ "default" ] . "\n  // \n";
          }
          else{
            $outputRef .= "  // \n";
          }
  
          $outputRef .= $itemData[ "description" ] . "\n  // \n";
          $outputRef .= $itemData[ "spotlight" ] . "\n  // \n";
          $outputRef .= $itemData[ "example" ] . "\n  // \n";
  
          if( $itemData[ "parameters" ] ) {
            $outputRef .= $itemData["parameters"] . "\n  // \n";
          }
          if( $itemData[ "return" ] ) {
            $outputRef .= $itemData["return"] . "\n  // \n";
          }
  
          $outputRef .= "  " . ucfirst($itemName) . "    " . $type . "\n\n";
        }
      }
      $outputRef .= "  Template    *Template\n";
      $outputRef .= "}\n\n";
      $outputRef .= "func ( el " . ucfirst( $structName ) . " ) getTemplate () string {\n  return `{$model}`\n}\n\n";
      $outputRef .= "\n\n\n\n\n\n\n\n\n\n\n\n";
    }
    print implode( "\n", $output );
  }
}
