#!/usr/bin/env bash

WORK_DIR=`dirname $0`
if [ "$WORK_DIR" = "." ];then
    WORK_DIR=$PWD
fi

PNAME=hackathon
usage() {
cat <<EOF
Usage:
  start with daemon:
    \$ $0
    \$ $0 daemon
  start without daemon:
    \$ PROGRAM_NO_DAEMON=true $0
    \$ $0 nodaemon
EOF
}

prepare() {
    export LD_LIBRARY_PATH=$WORK_DIR/lib:$LD_LIBRARY_PATH
}

run_as_no_daemon() {
    cd $WORK_DIR
    ./bin/${PNAME}
}

run_as_daemon() {
    [ -d $WORK_DIR/log ] || mkdir $WORK_DIR/log
    cd $WORK_DIR
    nohup ./bin/${PNAME} >> $WORK_DIR/logs/stdout.log 2>>$WORK_DIR/logs/stderr.log &
}

prepare

case $1 in
    "nodaemon")
        run_as_no_daemon
    ;;
    "daemon")
        run_as_daemon
    ;;
    "")
        if [ -z "$PROGRAM_NO_DAEMON" ];then
            run_as_daemon
        else
            run_as_no_daemon
        fi
    ;;
    *)
        usage
    ;;
esac

