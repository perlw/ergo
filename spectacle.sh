#!/bin/sh
APP=ergo
APPBASE=$HOME/services/ergo

go get -u golang.org/x/vgo &> goget.log
$GOPATH/bin/vgo build -o bin/$APP &> build.log

pkill $APP
rm -rf $HOME/services/$APP
rm -rf $APPBASE
mkdir -p $APPBASE

cp bin/$APP $APPBASE
cp -r static $APPBASE/
mkdir $APPBASE/data
cat schema.sql | sqlite3 $APPBASE/data/data.db

cd $APPBASE
nohup ./ergo &> ergo.log &
