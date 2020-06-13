# Overview [English](README.md)
コマンド `login-tweet` はサーバに SSH ログインしたときに、ツイートして通知します。

CentOS 7 と go1.14.2 linux/amd64 でテスト済み。

このコマンドは `geoiplookup` コマンドに依存しています。まず、これをインストールしましょう。

CentOS/RHEL には `GeoIP` パッケージが、Debian/Ubuntu には `geoip-bin` パッケージが必要です。

# Go 版
## インストール
1. このリポジトリをクローンして、ビルド。または release からビルド済みをダウンロード
1. Twitter API キーを取得して go/twitter-token.sh に指定する
1. 好きな通知メッセージを go/settings.json に指定する
1. 以下のスクリプトを自分の .bash_profile (.bashrc ではない) に追加する (/path/to/ はクローンしたディレクトリに置き換えること)

```shell:.bash_profile
source ~/path/to/go/twitter-token.sh
~/path/to/go/login-tweet
```

または (ログつき)

```shell:.bash_profile
source ~/path/to/go/twitter-token.sh
~/path/to/go/login-tweet -v >> ~/tweet.log
```

または root の .bash_profile の場合

```shell:.bash_profile
source ~/path/to/go/twitter-token.sh
~/path/to/go/login-tweet -r
```

# PHP 版
※PHP 版は、もうメンテナンスされていません。(PHP 7.0 (rh-php70) でテスト済み)

## インストール
1. このリポジトリをクローンする
2. [twitteroauth](https://github.com/abraham/twitteroauth.git) を1でクローンした中にクローン
3. Twitter API キーを取得して php/settings.php に指定する
1. 好きな通知メッセージを php/settings.php に指定する
5. 以下のスクリプトを /etc/ssh/sshrc に追加する

```shell:sshrc
source /opt/rh/rh-php70/enable
php /path/to/sshLoginNotify/twitterpost.php >> /path/to/sshLoginNotify/tweet.log
```

6. root に権限昇格したことも通知したければ、以下を /root/.bashrc に追加

```shell:.bashrc
source /opt/rh/rh-php70/enable
php /path/to/sshLoginNotify/twitterpost.php -r >> /path/to/sshLoginNotify/tweet.log
```
