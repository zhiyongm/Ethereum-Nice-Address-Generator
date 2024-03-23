# 以太坊靓号地址生成器
简介：以太坊靓号地址生成器，在本地生成你喜爱的以太坊规范的靓号地址！

## 特点：
* 多协程运行，充分利用你的CPU资源。
* 无联网代码，安全可靠，代码只有约40行。
* 生成的地址符合以太坊规范，可以直接导入钱包使用。
* 生成的地址可以自定义前缀和后缀，方便找到你喜爱的地址。

## 使用方法：
- 克隆本项目到本地。
- 安装go环境。
- 进入项目目录。
- 修改main.go文件中的prefix和suffix变量，设置你想要的前缀和后缀。
- 修改main.go文件中的goroutineNum变量，设置你想要的协程数量，一般为cpu物理核心数。
- 运行go run main.go，等待屏幕显示公钥和私钥。
- 可直接导入你的钱包使用。

# Ethereum Nice Address Generator
Synopsis: Locally generate your favorite Ethernet-specified pretty address!

## Features:
* Runs on multiple concurrent threads to fully utilize your CPU resources.
* No networking code, safe and reliable, only about 40 lines of code.
* The generated address conforms to the Ethernet specification and can be directly imported into your wallet for use.
* The generated addresses can be customized with prefix and suffix, making it easy to find your favorite addresses.

## Usage:
- Clone this project locally.
- Install the go environment.
- Go to the project directory.
- Modify the prefix and suffix variables in the main.go file to set the prefix and suffix you want.
- Modify the goroutineNum variable in the main.go file to set the number of concurrent programs you want, usually the number of physical cores of the cpu.
- Run go run main.go and wait for the screen to display the public and private keys.
- You can import them directly into your wallet and use them.

