set -e -x

# Install go toolchain
tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

# GO_BINARY="/usr/bin/go"
# GOROOT= /usr/local/go
# GOPATH= $HOME/go
# PATH= $GOPATH/bin:$GOROOT/bin:$PATH