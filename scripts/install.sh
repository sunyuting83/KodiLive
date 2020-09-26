#!/bin/sh

# install Kodi Live
basepath=$(cd `dirname $0`; pwd)
read -p "Please enter root password: " -s rootPassword
echo "Start..."
static_dir="/etc/kodilive"
if [ ! -d $start_dir ]; then
  echo "$rootPassword" | sudo -S mkdir -p -m 755 $start_dir
fi

echo "$rootPassword" | sudo -S install -Dm755 $basepath/kodilive /usr/bin/kodilive &&
echo "$rootPassword" | sudo -S install -Dm644 $basepath/kodilive.service /usr/lib/systemd/system/kodilive.service &&
echo "$rootPassword" | sudo -S install -Dm777 $basepath/kodilive.sh /etc/kodilive/kodilive.sh &&
echo "$rootPassword" | sudo -S systemctl enable kodilive &&
echo "$rootPassword" | sudo -S systemctl daemon-reload &&
echo "$rootPassword" | sudo -S systemctl start kodilive &&
echo "install complete"