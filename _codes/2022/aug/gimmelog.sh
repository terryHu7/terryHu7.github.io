#!/bin/bash

cd ~/Desktop

time=`date +"%H-%M-%S"`
mkdir ${time}

cd /var/log/ems/
emslog=$(ls -lt|grep ems-|head -n 1|awk '{print $9}')
cp ./${emslog} ~/Desktop/${time}

cd ../mediaAgent/
restart=$(cat magent.log |grep -E -n  "Media Agent start... "|tail -n 1|awk '{print $1}'|cut -d ':' -f 1)
tail -n +${restart} magent.log >> ~/Desktop/${time}/magent.log

echo "~/Desktop/${time}"