
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go;
export PATH=$PATH:$GOPATH/bin;
echo $PATH
echo $GOPATH

go mod init github.com/anubhavitis/Xmeme
go mod vendor
go build 
go run main.go