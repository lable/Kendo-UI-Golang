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
    
    unset ( $pageHtmlLArr );
    
    // Facilita a exp reg [ POG ]
    $pageHtmlLStr = str_replace ( "<h3", "<h3<h3", $pageHtmlLStr );
    
    //id="configuration... limita a pesquisa apenas para as configurações do sistema
    preg_match_all ( "%<h3 id=\"configuration.*?>.*?(?:<h3)%s", $pageHtmlLStr, $matchesMainLArr );
    preg_match_all ( "%<article>.*?id=[\"'](.*?)[\"'].*?href=[\"'](#.*?)[\"'].*?>(.*?)</a>.*?<p>(.*?)</p>%s", $pageHtmlLStr, $matchesMainDataLArr );
    preg_match_all ( "%<article>.*?(<blockquote>.*?</blockquote>).*?(?:<h3)%s", $pageHtmlLStr, $matchesBlockQuoteDataLArr );
    
    // Nome da classe
    $this->classNameCStr = $matchesMainDataLArr[ 3 ][ 0 ];
    $classNameLArr = explode ( ".", $this->classNameCStr );
    $this->classNameCStr = "";
    foreach ( $classNameLArr as $classNameValueLStr )
    {
      $this->classNameCStr .= ucwords( $classNameValueLStr );
    }
    
    if ( is_null ( $this->classNameCStr ) )
    {
      die ( "falha de conexão a internet" );
    }
    
    // Caso o arquivo exista, abre o arquivo e procura por coisas que não vão ser apagadas.
    $matchesDocumentLArr = array ();
    if ( is_readable( "./class/output/{$this->classNameCStr}.class.php" ) )
    {
      $fileSizeLUInt = filesize ( "./class/output/{$this->classNameCStr}.class.php" );
      if ( $fileSizeLUInt > 0 )
      {
        $resourceLObj = fopen ( "./class/output/{$this->classNameCStr}.class.php", "r" );
        $fileLStr = fread ( $resourceLObj, filesize ( "./class/output/{$this->classNameCStr}.class.php" ) );
        preg_match_all ( "%(//-+ *inicio codigo importante.*?//-+ *fim codigo importante.*?)[\r\n]*%si", $fileLStr, $matchesDocumentLArr );
        fclose ( $resourceLObj );
      }
      unset( $fileSizeLUInt );
    }
    
    // Pega a descrição da classe
    $classDescriptionLStr = $matchesMainDataLArr[ 4 ][ 0 ];
    /*if ( preg_match ( "%<a href=[\"'](.*?)[\"']%", $classDescriptionLStr, $matchesSeeLArr ) )
    {
      $seeLStr = "   * @see http://docs.telerik.com{$matchesSeeLArr[ 1 ]}\r\n";
    }
    else
    {
      $seeLStr = "";
    }*/
    $classDescriptionLStr = strip_tags( $classDescriptionLStr );
    
    // Pega notas de advertência da classe, caso haja.
    $warningLStr = "";
    if ( is_array ( $matchesBlockQuoteDataLArr[ 1 ] ) )
    {
      $warningLStr = trim ( strip_tags( $matchesBlockQuoteDataLArr[ 1 ][ 0 ] ) );
    }
    
    $outputLArr = array ();
    $toXmlLArr  = array ();
    
    foreach ( $matchesMainLArr[ 0 ] as $matchesMainKeyLUInt => $matchesMainValueLStr )
    {
      // Este bloco é baseado no formato padrão de documento da Telerik para o Kendo UI
      preg_match_all ( "%<a.*?href=[\"'](#.*?)[\"']>(.*?)</a>%s", $matchesMainValueLStr, $matchesLinkLArr );
      preg_match_all ( "%<h3.*?h3>%s", $matchesMainValueLStr, $matchesH3LArr );
      preg_match_all ( "%<code>(.*?)</code>%s", $matchesH3LArr[ 0 ][ 0 ], $matchesCodeLArr );
      preg_match_all ( "%<h4>(.*?)</h4>%s", $matchesMainValueLStr, $matchesH4LArr );
      preg_match_all ( "%(<h4>.*?</h4>)*[\r\n]*<pre><code>(.*?)</code></pre>%s", $matchesMainValueLStr, $matchesExampleLArr );
      preg_match_all ( "%<blockquote>.*?<p>(.*?)</p>.*?</blockquote>%s", $matchesMainValueLStr, $matchesBlockQuoteLArr );
      preg_match_all ( "%<p>(.*?)</p>%s", $matchesMainValueLStr, $matchesPLArr );
      preg_match_all ( "%<em>.*?default: *(.*?)\)</em>%s", $matchesH3LArr[ 0 ][ 0 ], $matchesDefaultLArr );
      preg_match_all ( "%<h4><em>(.*?)</em></h4>[ \r\n]*<p>(.*?)</p>%s", $matchesMainValueLStr, $matchesPH4LArr );
      
      // Separa a função dos atributos
      $objectKeyNameLArr = trim ( $matchesLinkLArr[ 2 ][ 0 ] );
      $objectKeyNameLArr = explode ( ".", $objectKeyNameLArr );
      $outputReferenceLPt = &$outputLArr;
      
      $javaScripObjectPathLStr  = "";
      foreach ( $objectKeyNameLArr as $objectKeyNameKeyLUInt => $objectKeyNameValueLStr )
      {
        if ( $objectKeyNameKeyLUInt == 0 )
        {
          // Faz a navegação do array por referencia
          $outputReferenceLPt = &$outputReferenceLPt[ $objectKeyNameValueLStr ];
        }
        
        // Montador do objeto js
        $javaScripObjectPathLStr .= "[ '{$objectKeyNameValueLStr}' ]";
        
        if ( count ( $objectKeyNameLArr ) - 1 != $objectKeyNameKeyLUInt )
        {
          continue;
        }
        
        // Montador do objeto js
        $outputReferenceLPt[ "js_object" ][ $objectKeyNameKeyLUInt ][] = $javaScripObjectPathLStr;
        
        // procura pelo padrão do tipo
        // "%<em>.*?default: *(.*?)\)*</em>%s"
        if ( !is_null ( $matchesDefaultLArr[ 1 ][ 0 ] ) )
        {
          $outputReferenceLPt[ "default" ][ $objectKeyNameKeyLUInt ][] = trim ( $matchesDefaultLArr[ 1 ][ 0 ] );
        }
        else
        {
          $outputReferenceLPt[ "default" ][ $objectKeyNameKeyLUInt ][] = null;
        }
        
        // procura pelo link do manual
        if ( !is_null ( $matchesLinkLArr[ 1 ][ 0 ] ) )
        {
          $matchesLinkLArr[ 1 ][ 0 ] = trim ( $matchesLinkLArr[ 1 ][ 0 ] );
          $seeLStr = "{$this->linkCStr}{$matchesLinkLArr[ 1 ][ 0 ]}";
          $outputReferenceLPt[ "see" ][ $objectKeyNameKeyLUInt ][] = "@see {$seeLStr}";
          $outputReferenceLPt[ "name" ][ $objectKeyNameKeyLUInt ][] = trim ( $matchesLinkLArr[ 2 ][ 0 ] );
          
          $tmpLArr = explode ( ".", trim ( $matchesLinkLArr[ 2 ][ 0 ] ) );
          $tmpRefToFinal  = &$toXmlLArr;
          foreach ( $tmpLArr as $tmpKeyLUInt => $tmpValueLStr )
          {
            if ( $tmpKeyLUInt == count( $tmpLArr ) - 1 )
            {
              $tmpRefToFinal[ $tmpValueLStr ] = null;
            }
            else
            {
              $tmpRefToFinal = &$tmpRefToFinal[ $tmpValueLStr ];
            }
          }
        }
        
        $outputReferenceLPt[ "textToCorrect" ][ $objectKeyNameKeyLUInt ][] = $matchesPH4LArr;
        
        // procura pelo text explicativo
        foreach ( $matchesPLArr[ 1 ] as $matchesPKeyLStr => $matchesPValueLStr )
        {
          $passLBol = true;
          foreach ( $matchesBlockQuoteLArr[ 1 ] as $matchesBlockQuoteValueLStr )
          {
            // compara o texto explicativo pelo parágrafo de advertência
            if ( $matchesPValueLStr == $matchesBlockQuoteValueLStr )
            {
              $passLBol = false;
            }
          }
          
          if ( $passLBol == false )
          {
            continue;
          }
          
          $matchesPValueLStr = str_replace ( "<code>", "<b><u>", $matchesPValueLStr );
          $matchesPValueLStr = str_replace ( "</code>", "</u></b>", $matchesPValueLStr );
          $matchesPValueLStr = str_replace ( "\r\n", "\r\n     * ", $matchesPValueLStr );
          $matchesPValueLStr = str_replace ( "\n", "\n     * ", $matchesPValueLStr );
          $matchesPValueLStr = str_replace ( "\r", "\r     * ", $matchesPValueLStr );
          
          $matchesPLArr[ 1 ][ $matchesPKeyLStr ] = $matchesPValueLStr;
          
          foreach ( $outputReferenceLPt[ "textToCorrect" ][ $objectKeyNameKeyLUInt ][ 0 ][ 1 ] as $textToCorrectKeyLUInt => $ignoredValue )
          {
            if ( $outputReferenceLPt[ "textToCorrect" ][ $objectKeyNameKeyLUInt ][ 0 ][ 2 ][ $textToCorrectKeyLUInt ] == $matchesPValueLStr )
            {
              $matchesPValueLStr = "<b>" . $outputReferenceLPt[ "textToCorrect" ][ $objectKeyNameKeyLUInt ][ 0 ][ 1 ][ $textToCorrectKeyLUInt ] . "</b> - " . $outputReferenceLPt[ "textToCorrect" ][ $objectKeyNameKeyLUInt ][ 0 ][ 2 ][ $textToCorrectKeyLUInt ];
            }
          }
          
          $matchesPValueLStr = preg_replace ( "%[\r\n* ]+%si", " ", $matchesPValueLStr );
          $matchesPValueLStr = preg_replace ( "%(<a href=['\"])(#.*?)(['\"]>)%", "$1{$this->linkCStr}$2$3", $matchesPValueLStr );
          
          $matchesPLArr[ 1 ][ $matchesPKeyLStr ] = $matchesPValueLStr;
        }
        
        $outputReferenceLPt[ "text" ][ $objectKeyNameKeyLUInt ][] = $matchesPLArr[ 1 ];
        
        unset ( $outputReferenceLPt[ "textToCorrect" ] );
        
        // procura pelo texto de advertência
        if ( is_array ( $matchesBlockQuoteLArr[ 1 ] ) )
        {
          foreach ( $matchesBlockQuoteLArr[ 1 ] as $matchesBlockQuoteKeyLStr => $matchesBlockQuoteValueLStr )
          {
            $matchesBlockQuoteValueLStr = str_replace ( "<code>", "<b><u>", $matchesBlockQuoteValueLStr );
            $matchesBlockQuoteValueLStr = str_replace ( "</code>", "</u></b>", $matchesBlockQuoteValueLStr );
            $matchesBlockQuoteValueLStr = str_replace ( "\r\n", "\r\n     * ", $matchesBlockQuoteValueLStr );
            $matchesBlockQuoteValueLStr = str_replace ( "\n", "\n     * ", $matchesBlockQuoteValueLStr );
            $matchesBlockQuoteValueLStr = str_replace ( "\r", "\r     * ", $matchesBlockQuoteValueLStr );
            
            $matchesBlockQuoteLArr[ 1 ][ $matchesBlockQuoteKeyLStr ] = $matchesBlockQuoteValueLStr;
            
            $outputReferenceLPt[ "note" ][ $objectKeyNameKeyLUInt ][] = $matchesBlockQuoteValueLStr;
          }
        }
        
        // procura pelo exemplo de uso
        if ( is_array ( $matchesExampleLArr[ 2 ] ) )
        {
          foreach ( $matchesExampleLArr[ 2 ] as $matchesExampleKeyLStr => $matchesExampleValueLStr )
          {
            //Pega a descrição do exemplo
            $exampleLabelLStr = strip_tags( $matchesExampleLArr[ 1 ][ $matchesExampleKeyLStr ] );
            $exampleLabelLStr = trim ( $exampleLabelLStr );
            
            //Remove quebras de linha no final do script
            $matchesExampleValueLStr = preg_replace ( "%([ \r\n]+)$%", "", $matchesExampleValueLStr );
            
            //Remove comentários da turma do Kendo
            $matchesExampleValueLStr = preg_replace ( "%(&lt;!--.*?--&gt;[\r\n]+)%si", "", $matchesExampleValueLStr );
            $matchesExampleValueLStr = preg_replace ( "%(/\*)(.*?)(\*/)(.*)%i", "//$2$4", $matchesExampleValueLStr );
            
            //Corrige a barra invertida que aparece em alguns exemplos
            $matchesExampleValueLStr = str_replace ( "\\", "\\\\", $matchesExampleValueLStr );
            
            //Não sei qual foi a necessidade de fazer isto com os comentários
            //$matchesExampleValueLStr = str_replace ( "/*", "//", $matchesExampleValueLStr );
            //$matchesExampleValueLStr = str_replace ( "*/", "", $matchesExampleValueLStr );
            
            $matchesExampleToHtmlToolTipLStr = str_replace( "\n", "<br>", $matchesExampleValueLStr );
            $matchesExampleToHtmlToolTipLStr = str_replace( "  ", "&nbsp;", $matchesExampleToHtmlToolTipLStr );
            $matchesExampleToHtmlToolTipLStr = preg_replace ( "%(example)%i", "<b>$1</b>", $matchesExampleToHtmlToolTipLStr );
            
            $outputReferenceLPt[ "exampleToHtmlToolTip" ][ $objectKeyNameKeyLUInt ][] = $exampleLabelLStr . "<br>" . $matchesExampleToHtmlToolTipLStr;
            
            //Devolve o html a sua forma original
            $matchesExampleToClassFileLStr = str_replace ( "&lt;", "<", $matchesExampleValueLStr );
            $matchesExampleToClassFileLStr = str_replace ( "&gt;", ">", $matchesExampleToClassFileLStr );
            
            //Adiciona os espaços para as linhas de comentários do código ficarem alinhadas de forma correta
            $matchesExampleToClassFileLStr = str_replace ( "\r\n", "\r\n        ", $matchesExampleToClassFileLStr );
            $matchesExampleToClassFileLStr = str_replace ( "\r",   "\r        ",   $matchesExampleToClassFileLStr );
            $matchesExampleToClassFileLStr = str_replace ( "\n",   "\n        ",   $matchesExampleToClassFileLStr );
            
            $outputReferenceLPt[ "exampleToClass" ][ $objectKeyNameKeyLUInt ][] = $exampleLabelLStr . "\r\n        " . $matchesExampleToClassFileLStr;
            
            $namePropertyLStr = $outputReferenceLPt[ "name" ][ $objectKeyNameKeyLUInt ][ count( $outputReferenceLPt[ "name" ][ $objectKeyNameKeyLUInt ] ) - 1 ];
            $outputReferenceLPt[ "exampleValue" ][ $objectKeyNameKeyLUInt ][] = $this->compileText( $namePropertyLStr, $matchesExampleToClassFileLStr );
          }
        }
        
        // procura pelo tipo
        // "%<code>(.*?)</code>%s"
        if ( !is_null ( $matchesCodeLArr[ 1 ][ 0 ] ) )
        {
          foreach ( $matchesCodeLArr[ 1 ] as $matchesCodeKeyLArr => $matchesCodeValueLArr )
          {
            $matchesCodeValueLArr = preg_replace ( "%([a-zA-Z_.]+).*%", "$1", trim ( $matchesCodeValueLArr ) );
            $matchesCodeLArr[ 1 ][ $matchesCodeKeyLArr ] = $matchesCodeValueLArr;
            
            //$outputReferenceLPt[ "type" ][ $objectKeyNameKeyLUInt ][] = $matchesCodeValueLArr;
          }
          $outputReferenceLPt[ "type" ][ $objectKeyNameKeyLUInt ][] = $matchesCodeLArr[ 1 ];
        }
      }
    }
    
    //print_r ( $outputReferenceLPt );
    
    // Conta quantos níveis tem o array de dados
    foreach ( $outputLArr as $functionNameLStr => $functionDataLArr )
    {
      $outputLArr[ $functionNameLStr ][ "count" ][ "max" ] = 0;
      if ( is_array ( $functionDataLArr[ "name" ] ) )
      {
        foreach ( $functionDataLArr[ "name" ] as $outputSubElementKeyLArr => $outputSubElementValueLArr )
        {
          foreach ( $outputSubElementValueLArr as $outputTypeKeyLUInt => $outputTypeValueLUInt )
          {
            $countLUInt = count ( explode ( ".", $outputTypeValueLUInt ) );
            if ( $countLUInt >= $outputLArr[ $functionNameLStr ][ "count" ][ "max" ] )
            {
              $outputLArr[ $functionNameLStr ][ "count" ][ "max" ] = $countLUInt;
            }
            $outputLArr[ $functionNameLStr ][ "count" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] = $countLUInt;
          }
        }
      }
    }
    
    // Cria um ordenador pala quantidade dos níveis do array de dados, do maior para o menor
    // pois, isto ajuda para não montar um dado repetido na resposta final
    // exemplo $dado[ "animation" ][ "close" ][ "duration" ] = 300;
    // após o uso, unset( $dado[ "animation" ][ "close" ][ "duration" ] ) deixa como lixo o array
    // $dado[ "animation" ][ "close" ] = null. Esta ordenação ajuda no coletor de lixo e na
    // montagem dos dados
    foreach ( $outputLArr as $functionNameLStr => $functionDataLArr )
    {
      for ( $i = $outputLArr[ $functionNameLStr ][ "count" ][ "max" ]; $i != -1; $i -= 1 )
      {
        foreach ( $functionDataLArr[ "name" ] as $outputSubElementKeyLArr => $outputSubElementValueLArr )
        {
          foreach ( $outputSubElementValueLArr as $outputTypeKeyLUInt => $outputTypeValueLUInt )
          {
            if ( $outputLArr[ $functionNameLStr ][ "count" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] == $i )
            {
              $outputLArr[ $functionNameLStr ][ "order" ][] = array ( $outputSubElementKeyLArr => $outputTypeKeyLUInt );
            }
          }
        }
      }
    }
    
    //print_r ( $outputLArr );
    //print_r ( $toXmlLArr );
    //exit;
    
    $outputToControlPanelHtmlLStr = "";
    $outputToControlPanelJavaScriptLStr = "";
    $outputToToolTipPanelScriptLStr = "";
    $outputToControlToolTipJavaScriptLStr = "";
    print "<pre>";
    $outputGolangLStr = "";
    foreach ( $outputLArr as $functionNameLStr => $functionDataLArr )
    {
      foreach($functionDataLArr["name"] as $nameKey => $nameList ){
        foreach( $nameList as $line => $value ){
          $outputGolangLStr .= "\n";
          if( !is_null( $functionDataLArr["default"][$nameKey][$line] ) ) {
            $outputGolangLStr .= "  // default: " . $functionDataLArr["default"][$nameKey][$line] . "\n";
          }
          $outputGolangLStr .= "  // " . preg_replace(
            array(
              "%(<a.*?href=\")(.*?)(\">)(.*?)(</a>)%si",
              "%(<b><u>)%si",
              "%(</u></b>)%si",
              "%(</strong>)%si",
              "%(<strong>)%si",
            ),
            array(
              "$4 ( http://docs.telerik.com$2 )",
              "'",
              "'",
              "*",
              "*"
            ),
              implode( "\n  //  \n  // ", $functionDataLArr["text"][$nameKey][$line] ) ) . "\n  //\n";
          $outputGolangLStr .= "  // " . str_replace( "@see ", "", $functionDataLArr["see"][$nameKey][$line] ) . "\n";
          if( !is_null( $functionDataLArr["exampleToHtmlToolTip"][$nameKey][$line] ) ) {
            //$outputGolangLStr .= "  /*\n  ```\n  " . str_replace("<br>", "<br>\n  ", $functionDataLArr["exampleToHtmlToolTip"][$nameKey][$line]) . "\n  ```\n  */\n\n";
            $outputGolangLStr .= "  /*\n      " . str_replace("<br>", "\n      ", html_entity_decode($functionDataLArr["exampleToHtmlToolTip"][$nameKey][$line])) . "\n*/\n";
          }
          $outputGolangLStr .= "  " . str_replace( " . ", ".", ucwords( str_replace( ".", " . ", $functionDataLArr["name"][$nameKey][$line] ) ) ) . "  ";
          if( is_array( $functionDataLArr["type"][$nameKey][$line] ) ) {
            $outputGolangLStr .= implode( " - ", $functionDataLArr["type"][$nameKey][$line] ) . "\n";
          } else {
            $outputGolangLStr .= $functionDataLArr["type"][$nameKey][$line] . "\n";
          }
          $outputGolangLStr .= "\n\n\n";
        }
        
        continue;
      }
      
      
      continue;
      if ( !is_array ( $functionDataLArr[ "type" ] ) )
      {
        continue;
      }
      
      foreach ( $functionDataLArr[ "order" ] as $outputSubElementValueLArr )
      {
        foreach ( $outputSubElementValueLArr as $outputTypeKeyLUInt => $outputTypeValueLUInt )
        {
          $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] = array ();
          
          foreach ( $functionDataLArr[ "type" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ] as $typeKeyLStr => $typeValueLStr )
          {
            $nameToControlPanelElementStr  = $functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ];
            $nameToControlPanelElementStr  = str_replace ( ".", " ", $nameToControlPanelElementStr );
            $nameToControlPanelElementStr  = $this->classNameCStr . " " . $nameToControlPanelElementStr;
            $nameToControlPanelElementStr  = ucwords ( $nameToControlPanelElementStr );
            $nameToControlPanelElementStr  = str_replace ( " ", "", $nameToControlPanelElementStr );
            
            $helpToolTip  = "<p>" . implode ( "</p><p>", $functionDataLArr[ "text" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ] ) . "</p>\r\n";
            $helpToolTip .= "<p>Property: {$functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}</p>\r\n";
            $helpToolTip .= "<p>Type: " . implode ( " | ", $functionDataLArr[ "type" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ] ) . "</p>\r\n";
            if ( !is_null ( $functionDataLArr[ "default" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ] ) )
            {
              $helpToolTip .= "<p>Default: {$functionDataLArr[ "default" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}</p>\r\n";
            }
            
            if ( isset ( $functionDataLArr[ "exampleToHtmlToolTip" ][ $outputTypeKeyLUInt ] ) )
            {
              $helpToolTip .= "<p>&nbsp;</p><code>" . implode ( "</code>\r\n<p>&nbsp;</p><code>", $functionDataLArr[ "exampleToHtmlToolTip" ][ $outputTypeKeyLUInt ] ) . "</code>\r\n";
            }
            $helpToolTip = "<div align=\"left\">{$helpToolTip}</div>";
            $helpToolTip = str_replace ( "\"", "\\\"", $helpToolTip );
            $helpToolTip = preg_replace ( "%([\r\n])%", "", $helpToolTip );
            
            $labelToControlPanelElementStr = $functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ];
            
            switch ( $typeValueLStr )
            {
              case "Boolean":
                
                if ( in_array ( "Boolean", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "Boolean";
                
                $outputToControlPanelHtmlLStr .= "<div id=\"id_{$nameToControlPanelElementStr}ToolTipCheckBox__replace__\"><label for=\"name_{$nameToControlPanelElementStr}CheckBox__replace__\">{$labelToControlPanelElementStr}</label>\r\n";
                $outputToControlPanelHtmlLStr .= "<input name=\"name_{$nameToControlPanelElementStr}CheckBox__replace__\" id=\"id_{$nameToControlPanelElementStr}CheckBox__replace__\" type=\"checkbox\"></div>\r\n";
                
                $outputToControlToolTipJavaScriptLStr .= "$(\"#id_{$nameToControlPanelElementStr}ToolTipCheckBox__replace__\").kendoTooltip({\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  filter: \"label\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  content: \"{$helpToolTip}\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  position: \"left\"\r\n";
                $outputToControlToolTipJavaScriptLStr .= "});\r\n";
                
                break;
              
              case "Date":
                
                if ( in_array ( "Date", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "Date";
                
                $outputToControlPanelHtmlLStr .= "<div id=\"id_{$nameToControlPanelElementStr}DatePicker__replace__\"><label for=\"name_{$nameToControlPanelElementStr}DatePicker__replace__\">{$labelToControlPanelElementStr}</label>\r\n";
                $outputToControlPanelHtmlLStr .= "<input name=\"name_{$nameToControlPanelElementStr}DatePicker__replace__\" id=\"id_{$nameToControlPanelElementStr}DatePicker__replace__\" ></div>\r\n";
                
                $outputToControlToolTipJavaScriptLStr .= "$(\"#id_{$nameToControlPanelElementStr}DatePicker__replace__\").kendoTooltip({\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  filter: \"label\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  content: \"{$helpToolTip}\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  position: \"left\"\r\n";
                $outputToControlToolTipJavaScriptLStr .= "});\r\n";
                
                $outputToControlPanelJavaScriptLStr .= "$(\"#id_{$nameToControlPanelElementStr}DatePicker__replace__\").kendoDatePicker();\r\n";
                
                break;
              
              case "Number":
                if ( in_array ( "Number", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "Number";
                
                $outputToControlPanelHtmlLStr .= "<div id=\"id_{$nameToControlPanelElementStr}ToolTipNumberBox__replace__\"><label for=\"name_{$nameToControlPanelElementStr}NumberBox__replace__\">{$labelToControlPanelElementStr}</label>\r\n";
                $outputToControlPanelHtmlLStr .= "<input name=\"name_{$nameToControlPanelElementStr}NumberBox__replace__\" id=\"id_{$nameToControlPanelElementStr}NumberBox__replace__\" type=\"number\" ></div>\r\n";
                
                $outputToControlToolTipJavaScriptLStr .= "$(\"#id_{$nameToControlPanelElementStr}ToolTipNumberBox__replace__\").kendoTooltip({\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  filter: \"label\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  content: \"{$helpToolTip}\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  position: \"left\"\r\n";
                $outputToControlToolTipJavaScriptLStr .= "});\r\n";
                
                $outputToControlPanelJavaScriptLStr .= "$(\"#id_{$nameToControlPanelElementStr}NumberBox__replace__\").kendoNumericTextBox();\r\n";
                
                break;
              
              case "Color":
              case "Element":
              case "jQuery":
              case "Selector":
              case "String":
              case "Function":
              case "Array":
              case "Object":
                if ( in_array ( "String", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "String";
                
                $outputToControlPanelHtmlLStr .= "<div id=\"id_{$nameToControlPanelElementStr}ToolTipTextBox__replace__\"><label for=\"name_{$nameToControlPanelElementStr}TextBox__replace__\">{$labelToControlPanelElementStr}</label>\r\n";
                //$outputToControlPanelHtmlLStr .= "<textarea name=\"name_{$nameToControlPanelElementStr}TextBox__replace__\" id=\"id_{$nameToControlPanelElementStr}TextBox__replace__\">{$functionDataLArr[ "exampleValue" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ][ $typeKeyLStr ][ "data" ]}</textarea></div>\r\n";
                $outputToControlPanelHtmlLStr .= "<textarea name=\"name_{$nameToControlPanelElementStr}TextBox__replace__\" id=\"id_{$nameToControlPanelElementStr}TextBox__replace__\"></textarea></div>\r\n";
                
                $outputToControlPanelJavaScriptLStr .= "var tags={};\r\n";
                $outputToControlPanelJavaScriptLStr .= "CodeMirror.fromTextArea( $(\"#id_{$nameToControlPanelElementStr}TextBox__replace__\")[ 0 ], {\r\n";
                $outputToControlPanelJavaScriptLStr .= "  lineNumbers: true,\r\n";
                $outputToControlPanelJavaScriptLStr .= "  theme: 'default',\r\n";
                $outputToControlPanelJavaScriptLStr .= "  smartIndent: true,\r\n";
                $outputToControlPanelJavaScriptLStr .= "  tabSize: 2,\r\n";
                $outputToControlPanelJavaScriptLStr .= "  indentWithTabs: false,\r\n";
                $outputToControlPanelJavaScriptLStr .= "  electricChars: true,\r\n";
                $outputToControlPanelJavaScriptLStr .= "  viewportMargin: 1,\r\n";
                $outputToControlPanelJavaScriptLStr .= "  mode: 'text/html',\r\n";
                //$outputToControlPanelJavaScriptLStr .= "  extraKeys: {'Ctrl-Space': 'autocomplete'}\r\n";
                $outputToControlPanelJavaScriptLStr .= "  extraKeys: {\r\n";
                $outputToControlPanelJavaScriptLStr .= "    \"'<'\": completeAfter,\r\n";
                $outputToControlPanelJavaScriptLStr .= "    \"'/'\": completeIfAfterLt,\r\n";
                $outputToControlPanelJavaScriptLStr .= "    \"' '\": completeIfInTag,\r\n";
                $outputToControlPanelJavaScriptLStr .= "    \"'='\": completeIfInTag,\r\n";
                $outputToControlPanelJavaScriptLStr .= "    \"Ctrl-Space\": function(cm) {\r\n";
                $outputToControlPanelJavaScriptLStr .= "      CodeMirror.showHint(cm, CodeMirror.hint.xml, {schemaInfo: tags} );\r\n";
                $outputToControlPanelJavaScriptLStr .= "    }\r\n";
                $outputToControlPanelJavaScriptLStr .= "  }\r\n";
                $outputToControlPanelJavaScriptLStr .= "});\r\n";
                
                $outputToControlToolTipJavaScriptLStr .= "$(\"#id_{$nameToControlPanelElementStr}ToolTipTextBox__replace__\").kendoTooltip({\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  filter: \"label\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  content: \"{$helpToolTip}\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  position: \"left\"\r\n";
                $outputToControlToolTipJavaScriptLStr .= "});\r\n";
                break;
              
              case "kendo.data.DataSource":
              case "kendo.data.HierarchicalDataSource":
              case "kendo.data.TreeListDataSource":
              case "kendo.data.SchedulerDataSource":
              case "kendo.data.GanttDataSource":
              case "kendo.data.GanttDependencyDataSource":
              case "kendo.data.PivotDataSource":
                if ( in_array ( "kendo.data.DataSource", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "kendo.data.DataSource";
                
                $outputToControlPanelHtmlLStr .= "<div id=\"id_{$nameToControlPanelElementStr}ToolTipText__replace__\"><label for=\"name_{$nameToControlPanelElementStr}Text__replace__\">{$labelToControlPanelElementStr}</label>\r\n";
                $outputToControlPanelHtmlLStr .= "<input name=\"name_{$nameToControlPanelElementStr}Text__replace__\" id=\"id_{$nameToControlPanelElementStr}Text__replace__\" value=\"http://\" type=\"text\" ></div>\r\n";
                
                $outputToControlToolTipJavaScriptLStr .= "$(\"#id_{$nameToControlPanelElementStr}ToolTipText__replace__\").kendoTooltip({\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  filter: \"label\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  content: \"{$helpToolTip}\",\r\n";
                $outputToControlToolTipJavaScriptLStr .= "  position: \"left\"\r\n";
                $outputToControlToolTipJavaScriptLStr .= "});\r\n";
                
                break;
            }
          }
        }
      }
    }
  
    print $outputGolangLStr . "\n\n\n\n\n";
    print "\n\n\n\n";
    die();
    $this->return = array (
      "toolTipPanel" => $outputToToolTipPanelScriptLStr,
      "toolTipScript" => $outputToControlToolTipJavaScriptLStr,
      "html" => $outputToControlPanelHtmlLStr,
      "script" => $outputToControlPanelJavaScriptLStr
    );
    //print_r ( $outputLArr );
    //return;
    
    // Varre os tipos e monta a estrutura de if else throw baseado na ordem dos elementos calculada acima
    foreach ( $outputLArr as $functionNameLStr => $functionDataLArr )
    {
      if ( !is_array ( $functionDataLArr[ "type" ] ) )
      {
        continue;
      }
      
      //foreach ( $functionDataLArr[ "order" ] as $outputSubElementKeyLArr => $outputSubElementValueLArr )
      foreach ( $functionDataLArr[ "order" ] as $outputSubElementValueLArr )
      {
        $firstPassLBol = true;
        $complementLStr = "";
        
        foreach ( $outputSubElementValueLArr as $outputTypeKeyLUInt => $outputTypeValueLUInt )
        {
          $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] = array ();
          
          foreach ( $functionDataLArr[ "type" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ] as $typeValueLStr )
          {
            switch ( $typeValueLStr )
            {
              default:
                print "$typeValueLStr\r\n";
                die ( $functionNameLStr );
              
              case "Boolean":
                
                if ( in_array ( "Boolean", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "Boolean";
                
                $outputLArr[ $functionNameLStr ][ "function" ] .= "      {$complementLStr}if ( is_bool ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} ) )\r\n" .
                  "      {\r\n" .
                  "        parent::addData ( \"{$functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}\", \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        unset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        parent::garbageCollector ( \$dataAX );\r\n" .
                  "      }\r\n";
                
                $outputReferenceLPt[ "function_tmp_pass" ] = true;
                break;
              
              case "Date":
                
                if ( in_array ( "Date", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "Date";
                
                $outputLArr[ $functionNameLStr ][ "function" ] .= "      {$complementLStr}if ( is_array ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} ) )\r\n" .
                  "      {\r\n" .
                  "        \$dateAsString = \"\";\r\n" .
                  "        if ( isset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Year\" ] ) )\r\n" .
                  "        {\r\n" .
                  "          \$dateAsString .= \"{\$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Year\" ]}, \";\r\n" .
                  "        }\r\n" .
                  "        if ( isset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Month\" ] ) )\r\n" .
                  "        {\r\n" .
                  "          \$dateAsString .= \"{\$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Month\" ]}, \";\r\n" .
                  "        }\r\n" .
                  "        if ( isset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Day\" ] ) )\r\n" .
                  "        {\r\n" .
                  "          \$dateAsString .= \"{\$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Day\" ]}, \";\r\n" .
                  "        }\r\n" .
                  "        if ( isset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Hour\" ] ) )\r\n" .
                  "        {\r\n" .
                  "          \$dateAsString .= \"{\$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Hour\" ]}, \";\r\n" .
                  "        }\r\n" .
                  "        if ( isset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Minute\" ] ) )\r\n" .
                  "        {\r\n" .
                  "          \$dateAsString .= \"{\$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Minute\" ]}, \";\r\n" .
                  "        }\r\n" .
                  "        if ( isset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Second\" ] ) )\r\n" .
                  "        {\r\n" .
                  "          \$dateAsString .= \"{\$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}[ \"Second\" ]}\";\r\n" .
                  "        }\r\n" .
                  "        \$dateAsString = \"new Date ({\$dateAsString});\";\r\n" .
                  "        parent::addData ( \"{$functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}\", \$dateAsString );\r\n" .
                  "        unset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        parent::garbageCollector ( \$dataAX );\r\n" .
                  "      }\r\n";
                
                $outputReferenceLPt[ "function_tmp_pass" ] = true;
                break;
              
              case "Number":
                if ( in_array ( "Number", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "Number";
                
                $outputLArr[ $functionNameLStr ][ "function" ] .= "      {$complementLStr}if ( is_numeric ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} ) )\r\n" .
                  "      {\r\n" .
                  "        parent::addData ( \"{$functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}\", \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        unset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        parent::garbageCollector ( \$dataAX );\r\n" .
                  "      }\r\n";
                
                $outputReferenceLPt[ "function_tmp_pass" ] = true;
                break;
              
              case "Color":
              case "Element":
              case "jQuery":
              case "Selector":
              case "String":
              case "Function":
                if ( in_array ( "String", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "String";
                
                $outputLArr[ $functionNameLStr ][ "function" ] .= "      {$complementLStr}if ( is_string ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} ) )\r\n" .
                  "      {\r\n" .
                  "        parent::addData ( \"{$functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}\", \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        unset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        parent::garbageCollector ( \$dataAX );\r\n" .
                  "      }\r\n";
                
                $outputReferenceLPt[ "function_pass" ] = true;
                break;
              
              case "kendo.data.DataSource":
              case "kendo.data.HierarchicalDataSource":
              case "kendo.data.TreeListDataSource":
              case "kendo.data.SchedulerDataSource":
              case "kendo.data.GanttDataSource":
              case "kendo.data.GanttDependencyDataSource":
              case "kendo.data.PivotDataSource":
              case "Array":
              case "Object":
                if ( in_array ( "kendo.data.DataSource", $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ] ) )
                {
                  break;
                }
                $outputLArr[ $functionNameLStr ][ "it_has_been_done" ][ $outputTypeValueLUInt ][] = "kendo.data.DataSource";
                
                $outputLArr[ $functionNameLStr ][ "function" ] .= "      {$complementLStr}if ( is_array ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} ) )\r\n" .
                  "      {\r\n" .
                  "        parent::addData ( \"{$functionDataLArr[ "name" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]}\", \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        unset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} );\r\n" .
                  "        parent::garbageCollector ( \$dataAX );\r\n" .
                  "      }\r\n";
                
                $outputReferenceLPt[ "function_pass" ] = true;
                break;
            }
            
            if ( $firstPassLBol == true )
            {
              $firstPassLBol = false;
              $complementLStr = "else ";
            }
          }
          
          if ( $firstPassLBol == false )
          {
            $outputLArr[ $functionNameLStr ][ "function" ] .= "      else if ( isset ( \$dataAX{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} ) )\r\n" .
              "      {\r\n" .
              "        throw new Exception ( \"Type error - {$this->classNameCStr}::{$functionNameLStr} ( \\\$data{$functionDataLArr[ "js_object" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ]} type must be '" . implode ( "' | '", $functionDataLArr[ "type" ][ $outputTypeKeyLUInt ][ $outputTypeValueLUInt ] ) . "'\" );\r\n" .
              "      }\r\n" .
              "      \r\n";
          }
        }
      }
    }
    
    //print_r ( $outputLArr );
    //exit;
    
    $toFileLStr = "" .
      "<?php\r\n" .
      "\r\n" .
      "  /**\r\n" .
      "   * " . preg_replace ( "%(\r*\n)%s", "\r\n   * ", $classDescriptionLStr ) . "\r\n" .
      "   * \r\n" .
      "   * This class has been automatically generated by the class \"classMaker.class.php\" written by\r\n" .
      "   * Helmut Kemper <helmut.kemper@gmail.com> based into Kendo UI documents.\r\n" .
      "   * {$this->linkCStr}\r\n" .
      "   * {$warningLStr}\r\n" .
      "   */\r\n" .
      "  class {$this->classNameCStr} extends javascript\r\n" .
      "  {\r\n" .
      "    private \$outputCStr;\r\n" .
      "    private \$idElementCStr;\r\n" .
      "    private \$nameVarCStr;\r\n" .
      "    private \$htmlTagExtraCStr;\r\n" .
      "    \r\n" .
      "    function __construct ( \$idElementAStr, \$nameVarAStr = null, \$htmlTagExtraAStr = null )\r\n" .
      "    {\r\n" .
      "      \$this->idElementCStr = \$idElementAStr;\r\n" .
      "      \$this->nameVarCStr = \$nameVarAStr;\r\n" .
      "      \$this->htmlTagExtraCStr = \$htmlTagExtraAStr;\r\n" .
      "    }\r\n" .
      "    \r\n";
    
    if ( count ( $matchesDocumentLArr[ 1 ] ) )
    {
      foreach ( $matchesDocumentLArr[ 1 ] as $matchesDocumentKeyLUInt => $matchesDocumentValueLStr )
      {
        $toFileLStr .= "" .
          "    {$matchesDocumentValueLStr}\r\n";
        
        if ( ( count ( $matchesDocumentLArr[ 1 ] ) > 1 ) && ( $matchesDocumentKeyLUInt != count ( $matchesDocumentLArr[ 1 ] ) - 1 ) )
        {
          $toFileLStr .= "    \r\n";
        }
      }
    }
    else
    {
      $toFileLStr .= "" .
        "    //---------- inicio codigo importante\r\n" .
        "    //\r\n" .
        "    //Coisas que não serão apagadas podem ser digitadas aqui dentro.\r\n" .
        "    //\r\n" .
        "    \r\n" .
        "    public function toHtmlOutput ()\r\n" .
        "    {\r\n" .
        "      self::\$mainHtmlCodeCStr .= \"<input id=\\\"{\$this->idElementCStr}\\\" {\$this->htmlTagExtraCStr}>\";\r\n" .
        "      return self::\$mainHtmlCodeCStr;\r\n" .
        "    }\r\n" .
        "    \r\n" .
        "    \r\n" .
        "    //---------- fim codigo importante\r\n";
    }
    
    $toFileLStr .= "" .
      "    \r\n" .
      "    public function toJavaScriptOutput ()\r\n" .
      "    {\r\n" .
      "      if ( !is_null ( \$this->nameVarCStr ) )\r\n" .
      "      {\r\n" .
      "        self::\$mainJavaScriptCodeCStr .= \"var {\$this->nameVarCStr} = \";\r\n" .
      "      }\r\n" .
      "      self::\$mainJavaScriptCodeCStr .= \"\$(\\\"#{\$this->idElementCStr}\\\").{$this->classNameCStr}({\";\r\n" .
      "      self::\$mainJavaScriptCodeCStr .= \$this->toObject( \$this->dataCArr );\r\n" .
      "      self::\$mainJavaScriptCodeCStr .= \"});\\r\\n\";\r\n" .
      "      \$this->dataCArr = null;\r\n" .
      "      return self::\$mainJavaScriptCodeCStr;\r\n" .
      "    }\r\n" .
      "    \r\n";
    
    
    foreach ( $outputLArr as $functionNameLStr => $functionDataLArr )
    {
      $toFileLStr .= "    /**\r\n";
      $firstPassLBol = true;
      
      if ( is_array ( $functionDataLArr[ "name" ] ) )
      {
        foreach ( $functionDataLArr[ "name" ] as $outputSubElementKeyLArr => $outputSubElementValueLArr )
        {
          foreach ( $outputSubElementValueLArr as $outputTypeKeyLUInt => $outputTypeValueLUInt )
          {
            if ( isset ( $outputLArr[ $functionNameLStr ][ "text" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) )
            {
              if ( $firstPassLBol == false )
              {
                $toFileLStr .= "     * \r\n";
              }
              $firstPassLBol = false;
              
              foreach ( $outputLArr[ $functionNameLStr ][ "text" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] as $testLStr )
              {
                $toFileLStr .= "     * {$testLStr}\r\n";
              }
            }
            
            if ( isset ( $outputLArr[ $functionNameLStr ][ "note" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) )
            {
              $toFileLStr .= "     * \r\n";
              $toFileLStr .= "     * ( ! ) {$outputLArr[ $functionNameLStr ][ "note" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]}\r\n";
            }
            
            if ( ( isset ( $outputLArr[ $functionNameLStr ][ "name" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) ) && ( isset ( $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) ) && ( isset ( $outputLArr[ $functionNameLStr ][ "default" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) ) )
            {
              if ( !is_array ( $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) )
              {
                $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] = array ( $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] );
              }
              $toFileLStr .= "     * @var \$dataAX{$outputLArr[ $functionNameLStr ][ "js_object" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]}: ( " . implode ( " | ", $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) . " ) ( default: {$outputLArr[ $functionNameLStr ][ "default" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]} )\r\n";
            }
            
            else if ( ( isset ( $outputLArr[ $functionNameLStr ][ "name" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) ) && ( isset ( $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) ) )
            {
              if ( !is_array ( $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) )
              {
                $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] = array ( $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] );
              }
              $toFileLStr .= "     * @var \$dataAX{$outputLArr[ $functionNameLStr ][ "js_object" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]}: ( " . implode ( " | ", $outputLArr[ $functionNameLStr ][ "type" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) . " )\r\n";
            }
            
            else if ( ( isset ( $outputLArr[ $functionNameLStr ][ "name" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) ) && ( isset ( $outputLArr[ $functionNameLStr ][ "default" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) ) )
            {
              $toFileLStr .= "     * @var \$dataAX{$outputLArr[ $functionNameLStr ][ "js_object" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]} ( default: {$outputLArr[ $functionNameLStr ][ "default" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]} )\r\n";
            }
            else
            {
              $toFileLStr .= "     * @var \$dataAX{$outputLArr[ $functionNameLStr ][ "js_object" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]}\r\n";
            }
            
            if ( isset ( $outputLArr[ $functionNameLStr ][ "see" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ] ) )
            {
              $toFileLStr .= "     * {$outputLArr[ $functionNameLStr ][ "see" ][ $outputSubElementKeyLArr ][ $outputTypeKeyLUInt ]}\r\n";
            }
            
            if ( count ( $outputLArr[ $functionNameLStr ][ "exampleToClass" ][ $outputSubElementKeyLArr ] ) )
            {
              $toFileLStr .= "     * \r\n";
              $toFileLStr .= "     * @code\r\n";
              $toFileLStr .= "       " . implode ( "\r\n     * @endcode\r\n     * @code\r\n        ", $outputLArr[ $functionNameLStr ][ "exampleToClass" ][ $outputSubElementKeyLArr ] ) . "\r\n";
              $toFileLStr .= "     * @endcode\r\n";
              //$toFileLStr .= "     *\r\n";
            }
          }
        }
        $toFileLStr .=  "     * @throws string type error\r\n" .
          "     */\r\n" .
          "    public function {$functionNameLStr} ( \$dataAX )\r\n" .
          "    {\r\n" .
          "      if ( !is_string ( \$this->outputCStr ) )\r\n" .
          "      {\r\n" .
          "        \$this->outputCStr = '';\r\n" .
          "      }\r\n" .
          "      \r\n" .
          "      \$dataAX = array ( \"{$functionNameLStr}\" => \$dataAX );\r\n" .
          "      \r\n" .
          $outputLArr[ $functionNameLStr ][ "function" ] .
          "    }\r\n" .
          "    \r\n" .
          "    \r\n";
      }
    }
    
    $toFileLStr .= "  }";
    
    $resourceLObj = fopen ( "./class/output/{$this->classNameCStr}.class.php", "w" );
    fwrite ( $resourceLObj, $toFileLStr );
    fclose ( $resourceLObj );
    
    //$resourceLObj = fopen ( "./class/output/toXml.js", "a" );
    //fwrite ( $resourceLObj, $this->toMakeXmlSubElementOutput ( $toXmlLArr ) );
    //fwrite ( $resourceLObj, $this->toMakeXmlMainElementOutput ( $toXmlLArr ) );
    //fclose ( $resourceLObj );
  }
  
  function toMakeXmlMainElementOutput ( $toXmlAArr )
  {
    $toXmlLStr = "";
    if ( is_array ( $toXmlAArr ) )
    {
      $toXmlLStr .= "  \"ui:{$this->classNameCStr}\": {\r\n    \"children\": [\r\n";
      foreach ( $toXmlAArr as $keyLStr => $ignoredValue )
      {
        $toXmlLStr .= "      \"{$this->classNameCStr}:{$keyLStr}\",\r\n";
      }
      
      $toXmlLStr = substr ( $toXmlLStr, 0, -3 );
      $toXmlLStr .= "\r\n    ]\r\n  },\r\n";
    }
    
    return $toXmlLStr;
  }
  
  function toMakeXmlSubElementOutput ( $toXmlAArr )
  {
    $toXmlLStr = "";
    $toConcatenateLStr = "";
    foreach ( $toXmlAArr as $toXmlKeyLX => $toXmlValueLX )
    {
      
      if ( in_array ( $toXmlKeyLX , $this->toXmlCheckCArr ) )
      {
        if ( !is_null ( $toXmlValueLX ) )
        {
          $toConcatenateLStr .= $this->toMakeXmlSubElementOutput ( $toXmlValueLX );
        }
        continue;
      }
      
      $this->toXmlCheckCArr[] = $toXmlKeyLX;
      
      $toXmlLStr .= "  \"{$this->classNameCStr}:{$toXmlKeyLX}\": {\r\n    \"children\": [";
      
      if ( !is_null ( $toXmlValueLX ) )
      {
        $toXmlLStr .= "\r\n";
        foreach ( $toXmlValueLX as $toXmlValueKeyLX => $ignoredValue )
        {
          $toXmlLStr .= "      \"{$this->classNameCStr}:{$toXmlValueKeyLX}\",\r\n";
        }
        
        $toXmlLStr = substr ( $toXmlLStr, 0, -3 );
        $toXmlLStr .= "\r\n    ]\r\n  },\r\n";
        
        $toConcatenateLStr .= $this->toMakeXmlSubElementOutput ( $toXmlValueLX );
      }
      else
      {
        $toXmlLStr .= "]\r\n  },\r\n";
      }
    }
    return $toXmlLStr . $toConcatenateLStr;
  }
  
  function compileText( $namePropertyAStr, $exampleAStr )
  {
    $namePropertyAStr = explode ( ".", $namePropertyAStr );
    $namePropertyAStr = implode ( ".*?", $namePropertyAStr );
    
    $returnLArr = array ();
    
    preg_match_all ( "%<script.*?>(.*?)</script.*?>%si", $exampleAStr, $matchesExampleLArr );
    
    foreach ( $matchesExampleLArr[ 1 ] as $codeKeyLUInt => $codeValueLStr )
    {
      //boolean
      if ( preg_match( "%$namePropertyAStr\s*:\s*(true|false)[\r\n;,]%si", $codeValueLStr, $matchesLArr ) )
      {
        $returnLArr[] = array (
          "data" => $matchesLArr[ 1 ],
          "source" => null
        );
      }
      
      //number
      else if ( preg_match( "%$namePropertyAStr\s*:\s*(-*[0-9.]+)[\r\n;,]%si", $codeValueLStr, $matchesLArr ) )
      {
        $returnLArr[] = array (
          "data" => $matchesLArr[ 1 ],
          "source" => null
        );
      }
      
      //var
      else if ( preg_match( "%$namePropertyAStr\s*:\s*([a-zA-Z][a-zA-Z0-9]+)[\r\n;,]%si", $codeValueLStr, $matchesLArr ) )
      {
        $variableLStr  = "";
        $variableLStr .= "/*\r\nvar ";
        $variableLStr .= $this->getObject( $matchesLArr[ 1 ], "var", $codeValueLStr );
        $variableLStr .= "\r\n*/\r\n";
        
        $returnLArr[] = array (
          "data" => $this->getObject( $namePropertyAStr . "\s*:", "", $codeValueLStr ),
          "source" => $variableLStr
        );
      }
      
      //kendo
      else if ( preg_match( "%$namePropertyAStr\s*:\s*(kendo[.A-Za-z ]+)%si", $codeValueLStr, $matchesLArr ) )
      {
        $returnLArr[] = array (
          "data" => $matchesLArr[ 1 ] . $this->getObject( $namePropertyAStr, "kendo.", $codeValueLStr ),
          "source" => null
        );
      }
      
      //function
      else if ( preg_match( "%$namePropertyAStr\s*:\s*function%si", $codeValueLStr ) )
      {
        $returnLArr[] = array (
          "data" => $this->codeIndentation ( $this->getObject( $namePropertyAStr . "\s*:", "", $codeValueLStr ) ),
          "source" => null
        );
      }
      
      //string
      else if ( preg_match( "%$namePropertyAStr\s*:\s*[\"']%si", $codeValueLStr ) )
      {
        $codeLStr = preg_replace ( "%:\s%", "", $this->getObject( $namePropertyAStr, "", $codeValueLStr ) );
        $codeLStr = preg_replace ( "%(,)$%", "", $codeLStr );
        $returnLArr[] = array (
          "data" => $codeLStr,
          "source" => null
        );
      }
      
      //object | array
      else if ( preg_match( "%$namePropertyAStr\s*:\s*[\[\{]%si", $codeValueLStr ) )
      {
        $codeLStr = preg_replace ( "%:\s%", "", $this->getObject( $namePropertyAStr, "", $codeValueLStr ) );
        $returnLArr[] = array (
          "data" => $this->codeIndentation ( $codeLStr ),
          "source" => null
        );
      }
      
      
      else
      {
        //print "$codeValueLStr\r\n";
        //print "$namePropertyAStr: ??\r\n";
        //print "\r\n\r\n\r\n\r\n\r\n";
      }
      //print "\r\n";
    }
    
    //print_r ( $returnLArr );
    //print "\r\n";
    return $returnLArr;
  }
  
  function codeIndentation ( $codeAStr )
  {
    preg_match_all ( "%[^ ]([\t ]{2,})[^ ]%s", $codeAStr, $matchLArr );
    
    $countLUInt = 1000;
    $replaceLStr = "";
    if ( is_array ( $matchLArr[ 1 ] ) )
    {
      foreach ( $matchLArr[ 1 ] as $spacesLStr )
      {
        if ( strlen ( $spacesLStr ) < $countLUInt )
        {
          $countLUInt = strlen ( $spacesLStr );
          $replaceLStr = $spacesLStr;
        }
      }
      
      $codeAStr = str_replace( $replaceLStr, "", $codeAStr );
    }
    return $codeAStr;
  }
  
  function getObject( $namePropertyAStr, $typeLStr, $exampleAStr )
  {
    $parenthesisCountLUInt = 0;
    preg_match( "%({$typeLStr}\s*{$namePropertyAStr})%si", $exampleAStr, $matchesLArr );
    $positionLUInt = strpos( $exampleAStr, $matchesLArr[ 1 ] ) + strlen( $matchesLArr[ 1 ] );
    $returnLStr = "";
    $quotesOpenLBol = false;
    
    for( $i = $positionLUInt; $i != strlen( $exampleAStr ); $i += 1 )
    {
      if ( ( $exampleAStr[ $i ] == "'" ) || ( $exampleAStr[ $i ] == '"' ) )
      {
        if ( $quotesOpenLBol == false )
        {
          $parenthesisCountLUInt += 1;
        }
        else
        {
          $parenthesisCountLUInt -= 1;
        }
        
        $quotesOpenLBol = !$quotesOpenLBol;
      }
      else if( ( $exampleAStr[ $i ] == "{" ) || ( $exampleAStr[ $i ] == "[" ) || ( $exampleAStr[ $i ] == "(" ) )
      {
        $parenthesisCountLUInt += 1;
      }
      else if( ( $exampleAStr[ $i ] == "}" ) || ( $exampleAStr[ $i ] == "]" ) || ( $exampleAStr[ $i ] == ")" ) )
      {
        $parenthesisCountLUInt -= 1;
      }
      
      $returnLStr .= $exampleAStr[ $i ];
      
      if ( ( $parenthesisCountLUInt == 0 ) && ( ( $exampleAStr[ $i ] == "\r" ) || ( $exampleAStr[ $i ] == "\n" ) || ( $exampleAStr[ $i ] == "," ) || ( $exampleAStr[ $i ] == ";" ) ) )
      {
        //print "{$returnLStr}\r\n";
        return trim ( $returnLStr );
      }
    }
    
    print "{$namePropertyAStr}\r\n";
    print "{$typeLStr}\r\n";
    print "{$exampleAStr}\r\n";
    die ( ">" . $matchesLArr[ 1 ] . "<" );
  }
}
