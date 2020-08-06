## Example Apache Thrift 
* Write server, client use language C++

## Install boost
* Step1:
```
sudo apt-get install build-essential curl file git
```
* Step 2:
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```
* Step 3:
```
test -d ~/.linuxbrew && eval $(~/.linuxbrew/bin/brew shellenv)
```
* Step 4:
```
test -d /home/linuxbrew/.linuxbrew && eval $(/home/linuxbrew/.linuxbrew/bin/brew shellenv)
```
* Step 5:
```
test -r ~/.bash_profile && echo eval" ($(brew --prefix)/bin/brew shellenv)" >>~/.bash_profile
```
* Step 6:
```
echo "eval $($(brew --prefix)/bin/brew shellenv)" >>~/.profile
```
* Step 7:
```
brew install boost
```

## Build:
* Run file ./thriftgen.sh
* Run server: /server/LoggerServer.cpp
* Run client: /server/LoggerClient.cpp

## Build use cmake:
* mkdir build
* cd build
* make ..
* make

## Other:
Used languages: C++, Golang, Python 3