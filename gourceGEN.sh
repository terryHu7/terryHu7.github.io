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
     | ffmpeg -y -r 60 -f image2pipe -vcodec ppm -i - -b 2048000 hutygit.mp4
