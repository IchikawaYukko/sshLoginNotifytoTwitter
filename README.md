# Overview
This command `login-tweet` will tweet notify message when user logged into server through SSH.

Tested under CentOS 7 and go1.14.2 linux/amd64

This command is depends on `geoiplookup` command. So you need install it first.

Package `GeoIP` is required for CentOS/RHEL. Package `geoip-bin` is required for Debian/Ubuntu.

# Go Version
## Install
1. Clone this repository and build. Or download compiled release.
1. Get Twitter App key and fill that into go/twitter-token.sh
1. Put your own notify message into go/settings.json
1. Add below script to your .bash_profile (not .bashrc)

```shell:.bash_profile
source ~/path/to/go/twitter-token.sh
~/path/to/go/login-tweet
```

or (with logging)

```shell:.bash_profile
source ~/path/to/go/twitter-token.sh
~/path/to/go/login-tweet -v >> ~/tweet.log
```

or on root's .bash_profile

```shell:.bash_profile
source ~/path/to/go/twitter-token.sh
~/path/to/go/login-tweet -r
```

# PHP Version
*PHP version is NOT maintained anymore (tested under PHP 7.0 (rh-php70))

## Install
1. Clone this repository.
2. Clone [twitteroauth](https://github.com/abraham/twitteroauth.git) to sshLoginNotifytoTwitter.
3. Get Twitter App key and fill that into php/settings.php
4. Put your own notify message into php/settings.php
5. Add below script to /etc/ssh/sshrc

```shell:sshrc
source /opt/rh/rh-php70/enable
php /path/to/sshLoginNotify/twitterpost.php >> /path/to/sshLoginNotify/tweet.log
```

6. If you want notify when some one get root privlege add below to /root/.bashrc

```shell:.bashrc
source /opt/rh/rh-php70/enable
php /path/to/sshLoginNotify/twitterpost.php -r >> /path/to/sshLoginNotify/tweet.log
```
