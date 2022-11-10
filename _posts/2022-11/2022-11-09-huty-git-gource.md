---
layout: post
title: "Git Visualization Using Gource"
subtitle: ''
author: "Terry"
header-style: text
tags:
  - Git
---


![](../hutygit.mp4)


```bash 
#!/bin/bash
  
gource \
    -s 0.05 \
    -1280x720 \
    --auto-skip-seconds .1 \
    --multi-sampling \
    --stop-at-end \
    --file-idle-time 0 \
    --output-ppm-stream - \
    --output-framerate 60 \
     | ffmpeg -y -r 60 -f image2pipe -vcodec ppm -i - -b 65536K hutygit.mp4
```


```shell
gource --hide dirnames,filenames --seconds-per-day 0.1 --auto-skip-seconds 1 -1280x720 -o - | ffmpeg -y -r 60 -f image2pipe -vcodec ppm -i - -vcodec libx264 -preset ultrafast -pix_fmt yuv420p -crf 1 -threads 0 -bf 0 gource.mp4
```