#!/bin/sh
go get -u golang.org/x/vgo &> goget.log
$GOPATH/bin/vgo build -o bin/ergo &> build.log

pkill ergo
cp bin/ergo $HOME/services/
cp -r index.html $HOME/services/
mkdir $HOME/services/data
cat schema.sql | sqlite3 $HOME/services/data/data.db
cd $HOME/services
nohup ./ergo &> ergo.log &
