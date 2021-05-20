ffmpeg
========

ffmpeg -i '$RTMP_URL' -an -vcodec libvpx -cpu-used 5 -deadline 1 -g 10 -error-resilient 1 -auto-alt-ref 1 -f rtp rtp://127.0.0.1:5104 -vn -c:a libopus -f rtp rtp:/127.0.0.1:5106


ffmpeg -r 30 -f lavfi -i testsrc -vf scale=640:480 -framerate 30 -vcodec libvpx -pix_fmt yuv420p -f rtp rtp://127.0.0.1:5104 -c:a libopus -f rtp rtp:/127.0.0.1:5106
