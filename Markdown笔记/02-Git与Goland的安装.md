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
   #因为我们国内墙的原因,建议使用 ssl通信方式进行代码的拉取
   $git clone git@github.com:zhouzhijie3607154/ReviewGoLang.git
# 查看我们本地当前的所有分支,如果远程分支有自己之前创建的分支,建议还是用之前的分支名
   $git branch -a
   # 切换到自己的本地分支进行开发
   $git checkout -b yugu
   #开发了一段时间后,我需要提交新写的代码了,则先提交本地代码，然后拉取远程main分支的最新代码
   $git commit -m "提交说明"
   $git pull --rebase origin main
   #将main分支最新的代码合并到我们的本地分支yugu中,有冲突的话需要手动排除一下
   $git merge main
   #提交我们的本地分支yugu到远程仓库中
   $git push origin yugu
   #去远程仓库再发送一个分支的合并请求给你的leader,剩下的交给他就行了
   
   #假设你是leader,则需要在远程仓库中对代码进行审核,检查分支提交上来的代码是否符合编码规范,有无冲突,有些没写注释的地方让他先给你补个注释或者文档,再确认合并提交的分支.
   
   ```
   
   

#### Git的常用命令

```
$git clone 克隆存储库到新目录 
$git init  创建一个空的Git存储库或重新初始化一个现有的存储库
$git add 将工作区的文件内容添加到索引中
$git mv 移动或重命名文件、目录或符号链接
$git restore 恢复工作树文件
$git rm 从工作树和索引中删除文件
$git diff 显示提交、提交和工作树之间的变化
$git log 显示提交日志
$git show 显示各种类型的对象
$git status 显示工作树状态
$git branch 列出、创建或删除分支
$git checkout 切换到新分支
$git commit 记录对存储库的更改
$git merge 将两个或多个开发历史连接在一起
$git rebase 在另一个基本提示上重新应用提交
$git reset 将当前HEAD复位到指定状态
$git switch 交换分支
$git tag 创建、列出、删除或验证用GPG签名的标记对象
$git fetch 从另一个存储库下载对象和引用
$git pull 从另一个存储库或本地分支中提取并与之集成
$git push 更新远程引用和相关对象
```

##### 选择  merge 还是 rebase ?

merge 是 合并,融合的意思，rebase 是 复位基底。在git中使用 merge 会将两个不同的分支内容合并为一个新的节点，之前的提交分开显示，也就是说git的历史工作提交树会产生分叉，在代码review时，需要进行三方节点的对比。而rebase进行合并则不会生成新节点，而是使工作树呈线性的方式笔直的前进，这样可以得到更好的提交树。

在多人协作开发，提交代码遇到版本冲突时，使用merge操作只要解决冲突内容，add 修改内容，commit就可以了。而使用rebase在遇到冲突时会中断合并操作，同时会提示去解决冲突。解决冲突后，add 修改 后需要执行git rebase –continue 继续合并操作，或者 执行 git rebase –skip 忽视已解决的这个冲突。

我个人更喜欢使用

