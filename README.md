# mytoydocker

这只是一个玩具，并且大部分代码并不是原创的。
参见[https://www.infoq.com/articles/build-a-container-golang/](https://www.infoq.com/articles/build-a-container-golang/)

## 使用方法（我是用的是Ubuntu 20.04）

1. 安装Go语言环境：

   ```shell
   sudo apt install golang-go
   ```

2. 安装debootstrap：

   ```shell
   sudo apt install debootstrap
   ```

3. 克隆代码库：

   ```shell
   git clone https://github.com/xvlincaigou/mytoydocker.git
   ```

4. 拉取发行版（distribution）：

   ```shell
   go run main.go pull <distribution>
   ```

   例如，运行以下命令来拉取 `focal` 版本的发行版：

   ```shell
   go run main.go pull focal
   ```

5. 进入特定发行版的终端：

   ```shell
   go run main.go activate <distribution>
   ```

6. 执行单个命令：

   ```shell
   go run main.go run <distribution> <command> <params>
   ```

**请注意：** 该代码仅用于演示和玩耍，部分功能可能尚不完善。这里面我写的部分可能还比较稚嫩。另外，`pull` 功能可能有限，实际上只是使用 `debootstrap` 直接拉取 Debian/Ubuntu 发行版。我自己进行了一些测试，还是能用的。祝你玩得愉快！