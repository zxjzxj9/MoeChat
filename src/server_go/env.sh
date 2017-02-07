#set up the golang build environment
export GOPATH=`pwd`
rm -rf src
go get github.com/mattn/go-sqlite3
