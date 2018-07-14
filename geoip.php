<?php
  //Get contry code from IP address
  if(strpos($argv[1], ':') === false) {
    $contry_code = exec("geoiplookup $argv[1]|cut -b 24-25");
  } else {
    $contry_code = exec("geoiplookup6 $argv[1]|cut -b 27-28");
  }

  //read csv database
  $csv = new SplFileObject(dirname(__FILE__)."/iso3166-1.csv");
  $csv->setFlags(SplFileObject::READ_CSV);
  foreach ($csv as $line) {
    //remove last emply line
    if(!is_null($line[0])) {
      if($line[4] == $contry_code) {
        echo $line[0];
      }
      $records[] = $line;
    }
  }
?>
