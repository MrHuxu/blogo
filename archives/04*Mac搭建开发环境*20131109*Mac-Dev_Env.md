# Mac下搭建开发环境

### Ruby开发环境

1. 显示Mac隐藏文件的命令:  
```
defaults write com.apple.finder AppleShowAllFiles -bool true```  
隐藏Mac隐藏文件的命令：  
```
defaults write com.apple.finder AppleShowAllFiles -bool false```  
两者都要运行一下命令```KillAll Finder```后才能有效    

2. 卸载MacPorts：  

        sudo port -fp uninstall installed
        sudo rm -rf \
            /opt/local \
            /Applications/DarwinPorts \
            /Applications/MacPorts \
            /Library/LaunchDaemons/org.macports.* \
            /Library/Receipts/DarwinPorts*.pkg \
            /Library/Receipts/MacPorts*.pkg \
            /Library/StartupItems/DarwinPortsStartup \
            /Library/Tcl/darwinports1.0 \
            /Library/Tcl/macports1.0 \
            ~/.macports

3. - 安装homebrew：  

        ruby -e "$(curl -fsSL https://raw.github.com/mxcl/homebrew/go)"

   - 通过homebrew安装软件，比如ctags：  
   
        homebrew search ctag
        homebrew install ctags

4. 然后进入Game/Mac app目录，安装Command Line Tools，OSX-GCC等工具

5. 安装rvm以及Ruby，一下全文抄字ruby-china的Wiki：
RVM 是干什么的这里就不解释了，后面你将会慢慢搞明白。  
```$ curl -L https://get.rvm.io | bash -s stable```  
等待一段时间后就可以成功安装好 RVM。  
然后，载入 RVM 环境（新开 Termal 就不用这么做了，会自动重新载入的）  
```$ source ~/.rvm/scripts/rvm```  
检查一下是否安装正确 
 
		$ rvm -v  
		rvm 1.17.3 (stable) by Wayne E. 	Seguin<wayneeseguin@gmail.com>,  ... 
###### 用 RVM 安装 Ruby 环境
   - 替换 Ruby 下载地址到国内淘宝镜像服务器  
	for Mac  
```$ sed -i .bak 's!ftp.ruby-lang.org/pub/ruby!ruby.taobao.org/mirrors/ruby!' $rvm_path/config/db```  
    for Linux  
```$ sed -i 's!ftp.ruby-lang.org/pub/ruby!ruby.taobao.org/mirrors/ruby!' $rvm_path/config/db```
   - 安装 readline 包  
```$ rvm pkg install readline```  
     安装 Ruby 2.0.0  
```$ rvm install 2.0.0 --with-readline-dir=$rvm_path/usr```  
或者可以安装 1.8.7 版本，也可以是 1.9.3，只要将后面的版本号跟换一下就可以了  
同样继续等待漫长的下载，编译过程，完成以后，Ruby, Ruby Gems 就安装好了。  
   - 设置 Ruby 版本  
RVM 装好以后，需要执行下面的命令将指定版本的 Ruby 设置为系统默认版本  
```$ rvm 2.0.0 --default```  
同样，也可以用其他版本号，前提是你有用 rvm install 安装过那个版本  
```$ gem source -r https://rubygems.org/```  
```$ gem source -a http://ruby.taobao.org```   
   - 安装常用gem  
```$ gem install sinatra slim shotgun heroku heroku-api thin pry yajl-ruby```  
然后测试安装是否正确

### Java开发环境
Intellij IDEA在安装好JDK后还要下载一个蛋疼的java 6 SE运行环境，妈蛋，JDK默认安装路径：  
```Mac系统盘/资源库/java/JavaVirtualMachines/jdk1.7.0_40.jdk```

### 安装fish-shell
1. 安装fish-shell  
```brew install fish```
2. 安装oh-my-fish  
```curl -L https://github.com/bpinto/oh-my-fish/raw/master/tools/install.sh | sh```
3. 安装rvm for fish
```curl --create-dirs -o ~/.config/fish/functions/rvm.fish https://raw.github.com/lunks/fish-nuggets/master/functions/rvm.fish```
4. 接下来就是iTerm2的字体设置，进入Preferences->Profiles->Text，将```Non-ASCII Font```改为任意一个已经patch过的字体。  
然后进入Terminal标签下，修改```Report terminal type```为```xterm-256color```，再勾选最下方的```Set local variables automaticly```。
5. 目前来看，fish对rvm的支持还是有点蛋疼，所以需要进入Preferences->Profiles，将```Send text at start```改成  
```rvm & clear```  
这样启动iTerm2就能将ruby版本设为rvm默认的了
