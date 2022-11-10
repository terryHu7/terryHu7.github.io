---
layout: post
title: "Top Trace"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Bash
  - Gstreamer
---

```bash
#!/bin/bash
  
file=$1
splitsize=$2
date=`date +%m-%d-%H:%M:%S`

gst-launch-1.0 filesrc location=${file} ! qtdemux ! h264parse ! splitmuxsink location="split%02d.mp4" max-size-bytes=${splitsize} muxer="qtmux faststart=true" &

echo "pid: $! "

top -c -d 1 -n 20 -b -p $! > ./toptrace_${date}.txt
```
