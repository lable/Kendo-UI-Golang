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
      $configurationReturnsParameters[ $methodsBlocksKey ] = preg_replace( "%<a href=\"(.*?)\" class=\"type-link\"><code>(.*?)</code></a>%si", "'$2' $1 ", $matchesMethodsTypes[ 1 ][ 0 ] );
      
      preg_match_all( "%<h4>Returns</h4>(.*?)(?:<h4>Example|<script)%si", $methodsBlocksLine . "<endLine>", $matchesMethodsTypes );
      $configurationReturnsLink[ $methodsBlocksKey ] = preg_replace( "%<a href=\"(.*?)\" class=\"type-link\"><code>(.*?)</code></a>%si", "'$2' $1 ", $matchesMethodsTypes[ 1 ][ 0 ] );
      
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
              "default" => $configurationDefault[ $keyToGet ][ 1 ],
              "typeLink" => $configurationType[ $keyToGet ][ 0 ],
              "helpLink" => $configurationLink[ $keyToGet ],
              "parameters" => $configurationReturnsParameters[ $keyToGet ],
              "return" => $configurationReturnsLink[ $keyToGet ],
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
              "default" => $configurationDefault[ $keyToGet ][ 1 ],
              "typeLink" => $configurationType[ $keyToGet ][ 0 ],
              "helpLink" => $configurationLink[ $keyToGet ],
              "parameters" => $configurationReturnsParameters[ $keyToGet ],
              "return" => $configurationReturnsLink[ $keyToGet ],
              "description" => $configurationDescription[ $keyToGet ],
              "spotlight" => $configurationSpotlight[ $keyToGet ],
              "example" => $configurationExample[ $keyToGet ],
            );
          }
        }
      }
    }
  
    $output = array();
    foreach( $structList as $structName => $structData ){
      foreach( $structData as $itemName => $itemData ){
        
        if( !isset( $output[ $itemName ] ) ){
          $output[ $itemName ]  = "";
          $output[ $itemName ] .= "// " . $itemData[ "helpLink" ] . "\n// \n";
          //$output[ $itemName ] .= "// Type: " . $itemData[ "type" ][ 0 ];
  
          //if( $itemData[ "default" ] ){
          //  $output[ $itemName ] .= "Defalt: " . $itemData[ "default" ] . "\n  // \n";
          //}
          //else{
          //  $output[ $itemName ] .= "\n// \n";
          //}
  
          $output[ $itemName ] .= str_replace( "  //", "//", $itemData[ "description" ] ) . "\n// \n";
          $output[ $itemName ] .= str_replace( "  //", "//", $itemData[ "spotlight" ] ) . "\n// \n";
          $output[ $itemName ] .= str_replace( "\n  ", "\n", substr( $itemData[ "example" ], 2 ) ) . "\n// \n";
  
          if( $itemData[ "parameters" ] ) {
            $output[ $itemName ] .= $itemData["parameters"] . "\n// \n";
          }
          if( $itemData[ "return" ] ) {
            $output[ $itemName ] .= $itemData["return"] . "\n// \n";
          }
          $output[ $itemName ] .= "type " . ucfirst( $itemName ) . " struct{\n\n";
        }
        
        if( isset( $structList[ $itemName ] ) ){
          if( count( $itemData[ "type" ] ) == 1 ){
              $output[ $itemName ] .= "  // " . $itemData[ "helpLink" ] . "\n  // \n";
              $output[ $itemName ] .= "  // Type: " . $itemData[ "type" ][ 0 ];
              
              if( $itemData[ "default" ] ){
                $output[ $itemName ] .= "Defalt: " . $itemData[ "default" ] . "\n  // \n";
              }
              else{
                $output[ $itemName ] .= "\n  // \n";
              }
              
              $output[ $itemName ] .= $itemData[ "description" ] . "\n  // \n";
              $output[ $itemName ] .= $itemData[ "spotlight" ] . "\n  // \n";
              $output[ $itemName ] .= $itemData[ "example" ] . "\n  // \n";
              
              if( $itemData[ "parameters" ] ) {
                $output[ $itemName ] .= $itemData["parameters"] . "\n  // \n";
              }
              if( $itemData[ "return" ] ) {
                $output[ $itemName ] .= $itemData["return"] . "\n  // \n";
              }
  
            if( $itemData[ "type" ][ 0 ] == "Array" ) {
              $output[$itemName] .= ucfirst($itemName) . "    []" . ucfirst($itemName) . "Line\n";
            }
            else {
              $output[$itemName] .= ucfirst($itemName) . "    " . ucfirst($itemName) . "\n";
            }
          }
        }
        else {
          if( count( $itemData[ "type" ] ) == 1 ){
            
            switch ( $itemData[ "type" ][ 0 ] ){
              case "String":
                $type = "string";
                break;
              
              case "Boolean":
                $type = "bool";
                break;
                
              case "Number":
                $type = "int64";
                break;
                
              case "Object":
              case "Array":
                die( "ainda tem problemas com os tipos" );
                
              default:
                die( "falta criar um tipo de variável" );
            }
            
            $output[ $itemName ] .= "  // " . $itemData[ "helpLink" ] . "\n  // \n";
            $output[ $itemName ] .= "  // Type: " . $type;
    
            if( $itemData[ "default" ] ){
              $output[ $itemName ] .= "Defalt: " . $itemData[ "default" ] . "\n  // \n";
            }
            else{
              $output[ $itemName ] .= "\n  // \n";
            }
    
            $output[ $itemName ] .= $itemData[ "description" ] . "\n  // \n";
            $output[ $itemName ] .= $itemData[ "spotlight" ] . "\n  // \n";
            $output[ $itemName ] .= $itemData[ "example" ] . "\n  // \n";
    
            if( $itemData[ "parameters" ] ) {
              $output[ $itemName ] .= $itemData["parameters"] . "\n  // \n";
            }
            if( $itemData[ "return" ] ) {
              $output[ $itemName ] .= $itemData["return"] . "\n  // \n";
            }
    
            if( $itemData[ "type" ][ 0 ] == "Array" ) {
              $output[$itemName] .= ucfirst($itemName) . "    []" . ucfirst($itemName) . "Line\n";
            }
            else {
              $output[$itemName] .= ucfirst($itemName) . "    " . $type . "\n";
            }
          }
        }
      }
    }
  }
}
