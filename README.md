# Overview
This script will tweet notify message when user logged into server through SSH.
Tested under CentOS 7 and rh-php70

# Install
1. Clone this repository.
2. Clone [twitteroauth](https://github.com/abraham/twitteroauth.git) to sshLoginNotifytoTwitter.
3. Get Twitter App key and fill that into settings.php
4. Put your own notify message into settings.php
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
