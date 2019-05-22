#!/bin/bash
cd /home/pi
rm allout.log
/home/pi/app_arm > allout.log 2>&1 &
xset s noblank
xset s off
xset -dpms
#unclutter -idle 0.5 -root &
sed -i 's/"exited_cleanly":false/"exited_cleanly":true/' /home/pi/.config/chromium/Default/Preferences
sed -i 's/"exit_type":"Crashed"/"exit_type":"Normal"/' /home/pi/.config/chromium/Default/Preferences
pressEnter()
{
sleep 10
xdotool key Return
}
pressEnter &
chromium-browser --start-fullscreen --noerrdialogs --disable-infobars http://localhost:58000
