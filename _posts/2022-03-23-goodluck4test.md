---
layout: post
title: "goodluck4test.sh"
subtitle: 'Two shell scripts for simplification'
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
deb=xxx-l4e-agent_${date}

xz=user@127.0.0.${1}
scp ${deb}_arm64.deb ~/dpkgima.sh ${xz}:~/
ssh ${xz}
```  

```shell
#!/bin/bash
if [ $# -eq 1 ]
then
	sudo dpkg -i "$*"
	sudo systemctl daemon-reload
	sudo systemctl restart xxx-agent.service
	sudo systemctl restart media-agent.service
	systemctl status media-agent.service
else
	echo "Usage:$ sudo /bin/bash dpkgima xxx-l4e-agent_0.2.80_arm64.deb"
fi
```