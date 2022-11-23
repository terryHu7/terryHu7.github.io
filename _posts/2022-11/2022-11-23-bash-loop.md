---
layout: post
title: "for var in ``;do <cmd>;done"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Bash
---

```bash
#!/bin/bash 

for i in {1..5};do
if [ $i -eq 1 ];then
        echo 1
else
        echo 2
fi
done

echo -e "\n" 

for file in ./*;do
        echo $file
done

echo -e "\n"

for file in `ls`;do
        echo $file
done
```