#### Git是什么?

Git是一个十分优秀的分布式版本控制系统，经常被用于于公司以及个人项目的代码管理。Git具有以下特点：

- 免费开源，使用成本低
- 分布式版本控制，支持本地不联网使用
- 多人协同开发简单，且合并也不吃带宽

#### Git的诞生背景

Git是由Linux的创建者Linus开发并免费开源的，很多人都知道Linus在1991年创建开源的Linux系统，从此，Linux不断发展，到现在已经成为了最大服务器软件了，无数的公司将自己的服务和应用部署到了Linux系统上。

最早Linux社区是采用手工的方式进行代码合并，世界各地的志愿者通过diff将自己的源代码文件发给Linus，再由Linus一个个进行合并。

到了2002年，Linux经过了长达10年的发展，代码库也变得十分之大了，原先手工合并代码的方式效率已经无法满足Linux社区的需求了，很多社区的兄弟也和Linus反应了这个问题，于是Linus选择了一个商业的版本控制系统BitKeeper，BitKeeper的东家BitMover公司出于人道主义精神，授权Linux社区免费使用这个版本控制系统。 

安定团结的大好局面在2005年就被打破了，原因是Linux社区牛人聚集，不免沾染了一些梁山好汉的江湖习气。开发Samba的Andrew试图破解BitKeeper的协议（这么干的其实也不只他一个），被BitMover公司发现了（监控工作做得不错！），于是BitMover公司怒了，要收回Linux社区的免费使用权。

Linus可以向BitMover公司道个歉，保证以后严格管教弟兄们，嗯，这是不可能的。实际情况是这样的：

Linus花了两周时间自己用C写了一个分布式版本控制系统，这就是Git！一个月之内，Linux系统的源码已经由Git管理了！牛是怎么定义的呢？大家可以体会一下。

Git迅速成为最流行的分布式版本控制系统，尤其是2008年，GitHub网站上线了，它为开源项目免费提供Git存储，无数开源项目开始迁移至GitHub，包括jQuery，PHP，Ruby等等。

历史就是这么偶然，如果不是当年BitMover公司威胁Linux社区，可能现在我们就没有免费而超级好用的Git了。

#### Git的安装以及配置

[Git下载地址](https://git-scm.com/download/windows)

1. 双击安装包执行安装引导程序

2. 根据需要对可选功能进行打钩

3. 指定本地的安装路径(后续需要配置环境变量)

4. 配置Path环境变量,添加 : `你的安装路径\cmd`

   - `E:\Utilities\Git\cmd`

5. 配置 Git Global Config  (需要有GitHub账户,没有要先注册一个)

   - ```cmd
     git config --global user.name “你真实的GitHub用户名”
     git config --global user.email “你对应的邮箱地址”
     ```

6. 配置ssl链接秘钥

   1. `cmd执行 :  ssh-keygen -t rsa -C "自己的github邮箱 “`
   2. 执行完成后,在用户主目录下 .ssh文件夹会生成一对rsa秘钥,将rsa.pub公钥全选复制
   3.  Github中进入  settings -->SSH Keys -->Add SSH Key ->new 
   4. 将之前复制的rsa.pub内容进行粘贴并保存,本地就可以使用ssh方式与GitHub进行通信了

#### Git仓库的初始化

一般我们使用Git进行仓库初始化分两种

- 一种是在我们个人在本地创建一个git初始化仓库,向GitHub远程仓库进行推送
- 另一种则是公司或者我们个人之前已经在GitHub或组织内部的仓库中提交了项目代码,我们需要拉下来,进行自己分支上的开发,再合并

先讲第一种怎么操作,首先Git Bash进入当前项目目录

1. ```cmd
   # 初始化本地仓库
   $git init 
   # 添加项目中的所有文件到暂存区
   $git add .
   # 将暂存区的文件进行提交, -m 是提交的注释信息,对本次提交进行说明
   $git commit -m "The First Commit"
   # 由于是第一次提交,无需对远程仓库进行pull拉去再进行分支合并,这里我们新建本地分支yugu
   $git checkout -b yugu
   # 将新建的yugu分支推送到远程代码仓库中去
   $git push origin yugu
   ```

2. 第二种,进入一个目录,打开Git Bash,(我们拉取下来的项目会保存在这个目录下)

   ```cmd
   
   ```

   

​	

