#!/bin/bash
echo "start..."
export PATH=/sbin:/bin:/usr/bin:/usr/local/bin
nohup kodilive -p 5670 > /dev/null 2>&1 &
echo "The Kodi Live is running..."