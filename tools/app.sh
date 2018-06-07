#!/bin/sh

selfdir=$(dirname $0)
export LD_LIBRARY_PATH=$selfidr/usr/lib:$LD_LIBRARY_PATH
exec $selfdir/usr/bin/qofia-ffi "$@"

