import subprocess

subprocess.call(
    " cd assets/media/2/hls && ffmpeg -f v4l2 -framerate 25 -video_size 640x480 -i /dev/video0 -profile:v baseline -level 3.0 -pix_fmt yuv420p -s 640x360 -start_number 0 -segment_list_flags live -hls_time 10 -hls_list_size 0 -f hls index.m3u8",
    shell=True,
)
