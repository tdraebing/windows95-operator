#!/bin/bash

mkdir -p ~/.vnc
echo $VNC_PWD | opt/TurboVNC/bin/vncpasswd -f > ~/.vnc/passwd
chmod 600 ~/.vnc/passwd

export DISPLAY=:1
opt/TurboVNC/bin/vncserver $DISPLAY -geometry 1280x800 &
sleep 15
windows95 &
sleep 5
xdotool key F11

/var/novnc/utils/launch.sh --vnc localhost:5901