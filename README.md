这只是一个玩具，并且大部分代码并不是原创的。
参见[https://www.infoq.com/articles/build-a-container-golang/](https://www.infoq.com/articles/build-a-container-golang/)

一、使用（我是用的是ubuntu 20.04）
1. 安装go
sudo apt install golang-go
2. 安装debootstrap
sudo apt install debootstrap
3. git clone
git clone https://github.com/xvlincaigou/mytoydocker.git
4. 拉取
go run main.go pull <distribution>
5. 进入某个distribution的终端
go run main.go activate <distribution>
6. 只执行一个命令
go run main.go run <distribution> <command> <params>