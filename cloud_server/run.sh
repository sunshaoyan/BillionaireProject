#!/bin/sh


WORK_DIR=`dirname $0`
if [ "$WORK_DIR" = "." ];then
    WORK_DIR=$PWD
fi

PNAME=hackathon

usage(){

cat <<EOF
Usage:
  start:
    \$ $0 start
  restart:
    \$ $0 restart
  stop:
    \$ $0 stop
EOF
}

start(){
[ -d $WORK_DIR/logs ] || mkdir $WORK_DIR/logs
    cd $WORK_DIR
    make
    echo starting
    nohup ./bin/${PNAME}  >> $WORK_DIR/logs/stdout.log 2>&1 &

    ps aux | grep $PNAME  | grep -v grep |awk '{print $2}'
    echo start ok
}

stop(){
    echo stopping
    aid=`ps aux | grep $PNAME  | grep -v grep |awk '{print $2}'`
    if [[ $aid ]]; then
        echo $aid
        kill -9 $aid
    fi
    echo stopped
}

restart(){
    stop
    sleep  3
    start
}

case $1 in
    "start")
        start
    ;;
    "restart")
        restart
    ;;
    "stop")
        stop
    ;;
    *)
        usage
    ;;
esac

