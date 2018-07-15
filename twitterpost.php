<?php

$IPADDR = exec("last -ain 1|sed 's/.*in[^0-9]*\([0-9.]*\)[^0-9]*.*/\\1/'|head -n 1");
$COUNTRY = GeoIP($IPADDR);
$UPTIME = getUptime();

// Load settings
require "settings.php";

// OAuthライブラリの読み込み
require "twitteroauth/autoload.php";
use Abraham\TwitterOAuth\TwitterOAuth;

if (isset($argv[1])) {
  if (trim($argv[1]) == "-r") {
    $status = $rootMessage;
  }
} else {
  $status = $loginMessage." ".getLoginBonus();
}

tweet($status);
logOutput($status);

function logOutput($stat) {
  echo date(DATE_RFC2822)."  ".$stat;
}

function tweet($stat) {
  global $consumerKey, $consumerSecret, $accessToken, $accessTokenSecret;

  //接続
  $connection = new TwitterOAuth($consumerKey, $consumerSecret, $accessToken, $accessTokenSecret);

  //ツイート
  $res = $connection->post("statuses/update", array("status" => $stat ));

  //レスポンス確認
  //var_dump($res);
}

function GeoIP($ipaddr) {
  //Get contry code from IP address
  if(strpos($ipaddr, ':') === false) {
    $contry_code = exec("geoiplookup $ipaddr|cut -b 24-25");
  } else {
    $contry_code = exec("geoiplookup6 $ipaddr|cut -b 27-28");
  }

  //read csv database
  $csv = new SplFileObject(dirname(__FILE__)."/iso3166-1.csv");
  $csv->setFlags(SplFileObject::READ_CSV);
  foreach ($csv as $line) {
    //remove last emply line
    if(!is_null($line[0])) {
      if($line[4] == $contry_code) {
        return $line[0];
      }
      $records[] = $line;
    }
  }
}

function getUptime() {
  return exec ("uptime |cut -f 4-7 -d ' '");
}

function getLoginBonus() {
  global $uptimeMessage, $loginBonusMessage;

  $bonus = array();
  $bonusfile = new SplFileObject(dirname(__FILE__)."/loginbonus");
  $bonusfile->setFlags(SplFileObject::SKIP_EMPTY | SplFileObject::DROP_NEW_LINE);

  foreach($bonusfile as $line) {
    if($line === false) continue;
    array_push( $bonus, $loginBonusMessage.$line );
  }
  array_push( $bonus, $uptimeMessage );

  shuffle( $bonus );
  return $bonus[0];
}
?>
