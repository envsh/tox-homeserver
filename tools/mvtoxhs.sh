#!/bin/sh

# move to/from remote toxhs process run state
# usage: cd /path/to/tox-homeserver/ && ./mvtoxhs.sh <to|from> <passwd>

function move_from()
{
    true;
}

function move_to()
{
    true;
}

passwd=$2

set -x
# TODO 改用rsync应该更好些
PTARG="-P 222"
if [ x"$1" == x"from" ]; then
    sshpass -p "$passwd" scp $PTARG root@10.0.0.7:/mnt/sda5/vmos/x64root/home/android/toxhs/toxhs.sqlite* .
    sshpass -p "$passwd" scp $PTARG root@10.0.0.7:/mnt/sda5/vmos/x64root/home/android/toxhs/toxhs.tsbin .
elif [ x"$1" == x"to" ];then
    sshpass -p "$passwd" scp $PTARG bin/toxhs root@10.0.0.7:/mnt/sda5/vmos/x64root/home/android/toxhs/
    sshpass -p "$passwd" scp $PTARG toxhs.tsbin root@10.0.0.7:/mnt/sda5/vmos/x64root/home/android/toxhs/
    sshpass -p "$passwd" scp $PTARG toxhs.sqlite* root@10.0.0.7:/mnt/sda5/vmos/x64root/home/android/toxhs/
    sshpass -p "$passwd" scp $PTARG toxhsfiles/* root@10.0.0.7:/mnt/sda5/vmos/x64root/home/android/toxhs/toxhsfiles/
else
    echo "not support command $1"
fi
echo "Done move $1"


