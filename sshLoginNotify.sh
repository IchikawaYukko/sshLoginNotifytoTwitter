#!/bin/bash
IPADDR=`last -ain 1|head -n 1|sed -e 's/\s\+/\t/g'|cut -f 10`
echo 百合子さんが $IPADDR\(`php /home/yuriko/sshLoginNotify/geoip.php $IPADDR`\) から ConoHa にsshログインしました♪ サーバのUptime:`uptime |cut -f 4-7 -d ' '`\(自動投稿\)
