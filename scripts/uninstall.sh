#!/bin/sh

# uninstall Kodi Live
read -p "Please enter root password: " -s rootPassword
echo "Start..."
echo "$rootPassword" | sudo -S systemctl stop kodilive &&
echo "$rootPassword" | sudo -S systemctl disable kodilive &&
echo "$rootPassword" | sudo -S systemctl daemon-reload &&
echo "$rootPassword" | sudo -S rm /usr/bin/kodilive &&
echo "$rootPassword" | sudo -S rm /usr/lib/systemd/system/kodilive.service &&
echo "$rootPassword" | sudo -S rm -rf /etc/kodilive &&
echo "uninstall complete"