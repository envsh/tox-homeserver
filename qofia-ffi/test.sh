n=0
while true; do
    if [ $n -gt 30 ]; then
        exit
    fi
    n=`expr $n + 1`
    ./qofia-ffi
done
