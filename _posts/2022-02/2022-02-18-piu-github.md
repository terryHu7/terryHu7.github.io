---
layout: post
title: "A bash script for git push simplification"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Git
  - Bash
---

After getting tired of entering three-way git commands (add, commit and then push), I write a bash script just for the convience, named `piu.sh`. 


```shell
cd /bin/  
sudo vim piu.sh
```



```shell
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
```


```shell
sudo chmod 755 piu.sh
```


Done. Just hit `piu.sh` any time you're in need! 😘:
