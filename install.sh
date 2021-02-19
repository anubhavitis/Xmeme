set -e -x

# Install go toolchain
sudo wget -c https://golang.org/dl/go1.15.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xvzf go1.15.2.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin