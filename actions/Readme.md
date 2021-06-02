
Testing Videos
================
https://test-videos.co.uk/bigbuckbunny/webm-vp9

https://commons.wikimedia.org/wiki/File:Big_Buck_Bunny_4K.webm


Stream
=========

netstat -nau | awk -F'[[:space:]]+|:' 'NR>2 && $5>=6000 && $5<=7000'


ffmpeg -protocol_whitelist file,udp,rtp -i subscribe.sdp -c:v libx264 -preset veryfast -b:v 3000k -maxrate 3000k -bufsize 6000k -pix_fmt yuv420p -g 50 -c:a aac -b:a 160k -ac 2 -ar 44100 -f flv rtmp://bom01.contribute.live-video.net/app/live_666332364_5791UvimKkDZW8edq8DAi4011wc4cR

ffmpeg -protocol_whitelist file,rtp,udp,https,tls,tcp -i subscribe.sdp -c:v libx264 -pix_fmt yuv420p -c:a aac -ar 16k -ac 1 -preset ultrafast -tune zerolatency -f flv rtmp://bom01.contribute.live-video.net/app/live_666332364_5791UvimKkDZW8edq8DAi4011wc4cR

