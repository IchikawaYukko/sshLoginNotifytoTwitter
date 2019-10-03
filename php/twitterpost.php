<?php

$IPADDR = exec("last -ain 1|sed 's/.*in[^0-9]*\([0-9.]*\)[^0-9]*.*/\\1/'|head -n 1");
$COUNTRY = GeoIP($IPADDR);
$UPTIME = getUptime();

// Load settings
require dirname(__FILE__)."/"."settings.php";

// OAuthライブラリの読み込み
require dirname(__FILE__)."/"."twitteroauth/autoload.php";
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

function isIPv6(string $ipaddr) {
  if(strpos($ipaddr, ':')) {
    return true;
  } else {
    return false;
  }
}

function GeoIP($ipaddr) {
  //Get contry code from IP address
  if(!isIPv6($ipaddr)) {
    $contry_code = exec("geoiplookup $ipaddr|cut -b 24-25");  // TODO use composer geoip/geoip instead exec();
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
  $str   = file_get_contents('/proc/uptime');
  $num   = floatval($str);
  $secs  = fmod($num, 60); $num = (int)($num / 60);
  $mins  = $num % 60;      $num = (int)($num / 60);
  $hours = $num % 24;      $num = (int)($num / 24);
  $days  = $num;

  return "$days days";
}

function getLoginBonus() {
  global $uptimeMessage, $loginBonusMessage;

  $bonus = [];
  $bonusfile = new SplFileObject(dirname(__FILE__)."/loginbonus");
  $bonusfile->setFlags(SplFileObject::SKIP_EMPTY | SplFileObject::DROP_NEW_LINE);

  foreach($bonusfile as $line) {
    if($line === false) continue;
    $bonus[] = $loginBonusMessage.$line;
  }
  $bonus[] = $uptimeMessage;

  shuffle( $bonus );
  return $bonus[0];
}
