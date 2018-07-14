<?php

// Load settings
require "setting.php";

// OAuthライブラリの読み込み
require "twitteroauth/autoload.php";
use Abraham\TwitterOAuth\TwitterOAuth;

//接続
$connection = new TwitterOAuth($consumerKey, $consumerSecret, $accessToken, $accessTokenSecret);

$status = trim(fgets(STDIN));

//ツイート
$res = $connection->post("statuses/update", array("status" => $status ));

echo $status;
//レスポンス確認
//var_dump($res);

?>
