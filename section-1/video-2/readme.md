# Packt Publishing - Hands on Microservices with Go
# Section 1 - Video 2 - Installation of required Software

## - Installing Go

1. Download lastest version of Go from from https://golang.org/dl/
2. Extract it to /usr/local :

```
sudo tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz

sudo tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz

```

3. Add the path to ~/.bashrc

```
export PATH=$PATH:/usr/local/go/bin
```

4. Create Go projects base directory :

```
cd ~
mkdir go
```

5. Add $GOPATH to ~/.bashrc

```
export GOPATH=/home/emiliano/go
export PATH=$PATH:$GOPATH/bin
```

6. Test 

```
echo $PATH
echo $GOPATH
go --version
```

### Learn More

[GO Installation Guide](https://golang.org/doc/install)

## - Installing DEP

```
go get -u github.com/golang/dep/cmd/dep
```

Test:

```
dep version
```

### Learn More

[Dep repository](https://github.com/golang/dep)
[Consice guide to Dep](https://gist.github.com/subfuzion/12342599e26f5094e4e2d08e9d4ad50d)

## - Installing Git

```
sudo apt-get update
sudo apt-get install git-core
git --version
```

### Learn More
[Installing Git on Ubuntu 16.04 LTS](https://www.liquidweb.com/kb/install-git-ubuntu-16-04-lts/)

## - Download Packt Repo

```
cd /home/emiliano/go/src/github.com
mkdir PacktPublishing
cd PacktPublishing
git clone git@github.com:PacktPublishing/Hands-on-Microservices-with-Go.git
```

## - Installing Docker

1. Remove old versions

```
sudo apt-get remove docker docker-engine docker.io docker-ce
```

2. Set up the repository

```
sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
```

3. Install Docker CE

```
sudo apt-get update
sudo apt-get install docker-ce
```

4. Test

```
sudo docker --version
```

### Learn More
[Get Docker CE for Ubuntu](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

## - Installing VS Studio Code and Go Plugin
1. Download Debian Package from: https://code.visualstudio.com/download

2. Install :

```
cd ~/Downloads
sudo dpkg -i code_1.23.1-1525968403_amd64.deb 
sudo apt-get install -f
```

3. Add to toolbar

Open Visual Studio in Ubuntu Finder.
Add to toolbar. (Click on Icon, choose "Lock to Launcher")

4. Add go plugin.

Open Visual Studio Code.
Launch VS Code Quick Open: Ctrl-P
Copy this in the prompt and hit enter: 
```
ext install ms-vscode.Go
```

### Learn More
[Download Visual Studio Code](https://code.visualstudio.com/download)
[Go programming in VS Code](https://code.visualstudio.com/docs/languages/go)

## - Installing Postman

1. Download Postman from: https://www.getpostman.com/apps
2. Decompress
3. Move Postman dir to Soft Dir in home
4. Launch
5. Add to launcher

[Download Postman](https://www.getpostman.com/apps)


## - Installing JMeter

To use JMeter you must have Java 8 or 9 in your System.

### Get Java 8 or 9
If you already have Jav 8 or 9, this is not necessary.

1. Download sdkman:

```
curl -s "https://get.sdkman.io" | bash
source "$HOME/.sdkman/bin/sdkman-init.sh"
```

2. List Java Versions:

```
sdk list java
```

3. Install Jave 9:

```
sdk install 9.0.4-openjdk
```

4. Test

```
 java --version
```

### Install JMeter

1. Download JMeter from https://jmeter.apache.org/download_jmeter.cgi
2. Extract
3. Move to Soft folder in home.
4. Launch:

```
cd ~/Soft/apache-jmeter-4.0/bin
./jmeter
```

5. Lock to launcher
6. Restart computer.

### Learn More

[SDKMAN](https://sdkman.io/)
[Download Jmeter](https://jmeter.apache.org/download_jmeter.cgi)
