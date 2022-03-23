---
layout: post
title: "goodluck4test.sh"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Git
---

```shell
#!/bin/bash
date=`date +%m%d%H%M`
version=${date}
make deb ARCH=arm64 VERSION=${version}
cd ./bin_arm64
deb=eswin-l4e-agent_${date}

xz=user@10.12.224.${1}
scp ${deb}_arm64.deb ~/dpkgima.sh ${xz}:~/
ssh ${xz}
```