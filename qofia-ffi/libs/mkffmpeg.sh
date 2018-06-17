#!/bin/sh

# videokit?

# this has static lib, seems good
# but it use ndk17, API>=19, mmap64, link failed
# change some and successful build armeabi-v7a static lib
# note: need manually git clone x264 and ffmpeg
# https://github.com/cmeng-git/ffmpeg-android/tree/master/build/ffmpeg/android/armeabi-v7a/lib

# seems this prebuild version fine. try first.
# this is a ffmpeg execute, not library...
# https://github.com/Khang-NT/ffmpeg-binary-android

