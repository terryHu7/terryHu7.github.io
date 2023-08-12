#!/bin/bash
if [ $# -eq 0 ]
then
        message=`date +%Y-%m-%d`
else    
        message="$*"
fi      

echo ${message}

git add -A;
git commit -m "${message}";
git push;
