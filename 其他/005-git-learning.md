# Git 学习笔记

<!-- TOC -->

- [基本操作命令](#基本操作命令)
    - [仓库初始化](#仓库初始化)
    - [获取远程仓库](#获取远程仓库)
    - [查看状态](#查看状态)
    - [查看变更前后的差异](#查看变更前后的差异)
    - [提交文件](#提交文件)
    - [文件重命名或移动](#文件重命名或移动)
    - [删除文件](#删除文件)
    - [查看提交记录](#查看提交记录)
    - [查看分支](#查看分支)
    - [新建分支](#新建分支)
    - [分支重命名](#分支重命名)
    - [删除分支](#删除分支)
    - [本地分支推送至远程仓库](#本地分支推送至远程仓库)
    - [获取远程分支](#获取远程分支)
    - [分支合并](#分支合并)
    - [冲突解决](#冲突解决)
    - [撤销修改](#撤销修改)
    - [版本回退](#版本回退)
    - [查看远程库信息](#查看远程库信息)
    - [暂存变更信息](#暂存变更信息)
    - [忽略文件和目录](#忽略文件和目录)
- [举个栗子](#举个栗子)
- [参考文献](#参考文献)

<!-- /TOC -->

关于`git`的介绍实在太多，Google 或 百度一下 就有，不再介绍。  

个人接触 `git` 已经差不多三年，第一次使用`git`作为版本管理工具时，网上查查资料也能使用，但习惯了`svn`的客户端，感觉学习`git`的曲线很陡，很多命令总是很容易搞混或忘记，后来换工作又切回`svn`了。直到最近加入新的team，完全使用`git`，经过不断请教和学习，该把成果整理归纳一下，以备今后查询。  

----

## 基本操作命令
### 仓库初始化
个人认为实际用到的不多，一般都是现有的代码仓库，直接克隆下来，里面就有`.git` 目录
```
git init # 生成 .git 目录
git remote add origin git@github.com:my-name/proj.git # 本地和远程仓库关联
```

### 获取远程仓库  
````
git clone git@github.com:my-name/proj.git
````

### 查看状态
提交之前，先看一下本次 修改、删除、新增 了哪些文件
````
git status
````

### 查看变更前后的差异
````
git diff [/path/to/file]  # 可指定文件，展示工作区和最新提交之间的差异，commit之前建议使用
git diff --stat  # 展示工作区和最新提交之间的差异，只有文件名和行数
git diff HEAD  # 比较本次提交和上次提交之间的差异，push之前建议使用
git diff 2265d4a86 4b94416057 # 比较两个历史版本之间的差异
````

### 提交文件  
提交仓库的变更，新增和修改文件命令相同
````
git add /path/to/file1 [/path/to/file2...]  # 提交到本地暂存区，可同时提交多个文件
git commit -m '提交文件'  # 保存暂存区的变更到当前分支
git push origin branch-name  # 把变更文件推到当前分支对应的远程仓库
````  

提交所有文件，可以使用  
````
git add -A
# or
git add .
git commit -m '提交所有文件'
````

或简写为
````
git commit -am '简写:提交所有文件'
````  

如提交之后，想修改最近一次提交信息，可使用
````
git commit --amend
````

### 文件重命名或移动
````
git mv /path/to/old-file /path/to/new-file
git commit -m '文件重命名'
git push origin branch-name
````

### 删除文件  
````
git rm /path/to/file
git commit -m '删除文件'
git push origin branch-name
````

### 查看提交记录  
log 会已提交时间由近及远展示（不能看已删除的 commit log）  
````
git log # 展示提交的哈希值、合并信息、作者、提交时间和备注信息
git log /path/to/file # 展示指定文件/目录提交的哈希值、合并信息、作者、提交时间和备注信息
git log --pretty=short # 展示提交的哈希值、合并信息、作者和备注信息的第一行
git log --pretty=online # 一行展示提交的哈希值、备注信息
git log --graph # 简单图形展示每次提交的分支及其分化衍合情况
git log -p [/path/to/file] # -p 会展示前后差异
git log --stat # 展示每次提交新增或删除的文件行数
git log --grep=test # 展示包含test关键字的提交
git log --author=myname # 展示作者是myname的提交
````  

那么，有时想查看包括已删除在内的所有commit log，则使用
````
git reflog
````

### 查看分支 
````
git branch # 查看本地分支
git branch -r # 查看远程分支
git branch -a # 查看所有分支
````

### 新建分支
````
git branch new-branch-1 # 新建分支
git checkout new-branch-1 # 切换新分支
git checkout master # 切换回到master
````  
可简写为
````
git checkout -b new-branch-1
````  

使用以下命令即可把本地分支推送到远程
````
git push origin new-branch-1
````

### 分支重命名  
````
git branch -m old-branch-name new-branch-name
git branch -M old-branch-name new-branch-name # 如已存在则强行覆盖
````
推送到远程，
````
git push origin new-branch-name
````
此时，远程同时存在两个分支 `old-branch-name` 和 `new-branch-name`

### 删除分支 
````
git checkout master # 不能删除当前所在分支
git branch -d old-branch-name # 删除本地分支，如有未合并的代码，将会删除失败
git branch -D old-branch-name # 强行删除，即使有未合并的代码
git push origin --delete old-branch-name # 删除远程分支
git push origin :old-branch-name # 把空推送到远程，也相当于删除远程分支
````

### 本地分支推送至远程仓库  
在推送之前，首先要确保本地已经跟远程仓库关联 `仓库初始化`，再推送
````
git push origin new-branch-1
````

### 获取远程分支  
建议本地和远程分支的名称保持一致
````
git checkout -b new-branch-2 origin/new-branch-2
````  

将远程分支的最新代码同步到本地
````
git pull origin new-branch-2 # 会自动merge  
# 等同于如下命令，注意git pull 和 git fetch 的区别
git fetch origin
git merge origin/new-branch-2
````

### 分支合并  
````
git checkout master
git merge --no-ff new-branch-1 # --no-ff 会把本次合并记录到历史记录中
git branch --merged # 查看已合并的分支
git branch --no-merged # 查看未合并的分支
````

### 冲突解决  
当发生冲突时，打开冲突文件编辑，再参考 `提交文件` 提交即可

### 撤销修改  
在提交到暂存区之前（即`未 git add`），需撤销某个文件的修改，则使用
````
git checkout -- filename
````

如已经提交到暂存区，但 `未git commit`，则使用
````
git reset HEAD filename
git checkout -- filename
````

如 `已 git commit`，则使用
````
git revert HEAD
````

还有一种情况，就是 `恢复已删除的文件`，则使用  
````
rm filename # 发现误删了
git checkout -- filename # 即可恢复
````

### 版本回退 
`--hard`表示彻底回退  
````
git reset --hard HEAD~1 # 本地回退到上一个历史版本
git reset --hard 3b40bcb448c5b3 # 本地回到指定版本，可以是当前版本之前或之后的某一个版本
git reset --hard origin/master # 本地版本回到跟远程一样
git push -f origin your-branch # 回滚后的代码同步到远程
````

### 查看远程库信息
````
git remote -v
git remote show origin # 查看某个远程仓库的详细信息
````

### 暂存变更信息 
开发过程中，难免遇到紧急插入的需求，不得不停下手头去处理问题，这时候就需要保留现场（包括
暂存区和工作区的内容），可使用如下命令
````
git stash # 对当前的暂存区和工作区状态进行保存
git stash list # 查看保存的现场
git stash pop # 恢复保存的现场并删除现场快照
````

### 忽略文件和目录 
一些文件不想提交，可配置为忽略，在根目录下的文件`.gitignore` 加入要忽略的文件和目录即可。

--- 

## 举个栗子  
在实际的使用过程中，很多team 一般使用 Github、Gitlab、Bitbucket等进行管理。下面以 `Github`为例。  

现在`Github` 上已有项目 `test_proj`，地址为 `git@github.com:my-org/test_proj.git`  

> <small>项目名称都是虚构</small>  

- 首先，在 `Github`上`fork`这个项目，完成之后你的开发分支就变成
`git@github.com:your-name/test_proj.git`。  

- 将自己的开发分支克隆到本地
```
git clone git@github.com:your-name/test_proj.git # 克隆项目
cd test_proj/
git remote add upstream git@github.com:my-org/test_proj.git # 很关键，设置个人分支的源仓库
```
- 在本地开发，期间的新增、修改、删除源文件以及提交等操作（参考`基本操作命令`即可），都不会对源仓库有影响。

- 开发完成并通过测试之后，代码需合入源仓库，提个 `Pull Request`（简称`PR`），就可以等待项目的owner 进行代码审查和合并了。

- 需要注意的是，实际开发过程中，很可能是多人协作，每个人都`fork`自己的分支，提`PR`也有先后顺序，也许我们提`PR`时，源仓库的代码已经经过了几次迭代，那么发生冲突的可能性就非常大。
为了尽可能避免代码冲突，`设置个人分支的源仓库` 这个步骤就显得非常关键了，我们要在设置好源仓库之后，在开发过程中，经常同步源仓库的代码，使用如下命令即可：
```
 git pull upstream master # 同步master的最新代码到本地
```

目前团队使用的这种模式，这也跟开源项目的模式是一致的，我个人也非常喜欢，省事！
  
之前使用 `Git-flow`模式开发，每次开发都需要从`master`新建分支，`merge` 之后再删除分支，每次需求来了就反复`新建分支，删除分支`，感觉比较繁琐。  

----

学习`git` 确实曲线比`svn`陡，查了不少资料，还专门买了两本书来啃，花了大半个月，累死我了！
不过确实非常有收获，也非常开心，周末没出门，宅家里两天写完这篇博客，算是一个小小的总结吧！^_^

----

## 参考文献
1. [《Pro Git》](https://git-scm.com/book/zh/v1)
1. 《GitHub入门与实践》大塚弘记 (作者)