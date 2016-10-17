# 个人Ubuntu/LinuxMint笔记  

用了```Ubuntu/Linux Mint```这么久了，慢慢的也有了自己的一点心得，别的不说，光每次重装系统后要装的一堆软件就够麻烦的，为了避免下次重装系统后再到处百度，这里把那些重装系统后的必装软件都记录下来，顺便也向大家推荐一下经过我试用被确定为经典的软件：      

#### 首先是一堆必须要添加和安装的PPA源：

	sudo add-apt-repository ppa:albert748/ppa    #为知笔记，跨平台笔记软件，在linux上打表现明显好于evernote
	sudo add-apt-repository ppa:synapse-core/ppa    #synapse，一个超级强大快速启动工具
	sudo add-apt-repository ppa:kokoto-java/omgubuntu-stuff    #nitrux图标，超级好看（install nitrux-umd)
	sudo add-apt-repository ppa:marlin-devs/marlin-daily     #marlin,非常适合程序员试用的资源管理器，可全键盘操作
	sudo add-apt-repository ppa:webupd8team/sublime-text-2     #sublime text2，一个美观的傻瓜式文本编辑器 （install sublime-text-2-beta or install sublime-text-2-dev ）

#### 然后是不需要添加源的必要软件：

	sudo apt-get install ia32-libs*    #64位系统下的32位依赖包，WineQQ，WPS均需要
	sudo apt-get install im-switch libapt-pkg-perl fcitx fcitx-table-wbpy    #fcitx输入法，目前感觉最好用的汉语输入法
	sudo apt-get install libqt4-dev libqt4-dbg libqt4-gui libqt4-sql qt4-dev-tools qt4-doc qt4-designer qt4-qtconfig qtcreator     #qt开发环境

#### 然后是一些常用软件的安装方法：

```deb```文件安装方式为  
	
	sudo dpkg -i ***.deb

```sh```文件安装方式为 

	sudo sh ***.sh

解决依赖问题 
		
	sudo apt-get -f install

#### 一些针对特定软件的TIPS：
##### 添加鼠标主题：

方法，将下载下来的包解压，将含有	```cursors```文件夹和```.theme```文件的文件夹复制到```/usr/share/icons```文件夹下，打开主题设置即可看见.  
PS：```wiz```这版本居然无法现实我换的鼠标主题，还是系统最开始的白色的，真难看啊，怨念ing。。。  

##### 安装```JDK```：   
 话说系统自带的OPENJDK真废柴啊，果断安装oracle jdk替换之，首先下载64位JDK的tar.gz文件

	sudo tar zxvf ./jdk-7-linux-i586.tar.gz -C /usr/lib/jvm
	cd /usr/lib/jvm`

以下内容均视为已将```tar.gz```压缩包内所有内容都解压缩到了```/usr/lib/jvm/```下的```jdk1.7.0_17```文件夹里  
  其实就是解压这个文件夹然后放到```/usr/lib/jvm下```，如果不记得解压命令，完全可以手动解压然后```sudo cp```到目标文件夹，反正我就是这么做的。。

	sudo gedit ~/.bashrc
打开配置文件，然后在里面加上
	
	export JAVA_HOME=/usr/lib/jvm/jdk1.7.0_17
	export JRE_HOME=${JAVA_HOME}/jre
	export CLASSPATH=.:${JAVA_HOME}/lib:${JRE_HOME}/lib
	export PATH=${JAVA_HOME}/bin:$PATH

这就是环境变量的配置了，然后输入
	
	source ~/.bashrc
使修改立即生效，然后在终端里输入

	java -version
	
如果粗线了```oracle java```的信息，就证明安装成功了。
PS：以上```JDK```安装方法中的目录和文件名有些视安装的```JDK```版本而定，请自行判断。
 在```linux mint```下配置```fcitx```输入法：
好吧，偷个懒，直接看这吧
PS：在```Linux Mint15```上```fcitx```已经不用任何配置了，所以这里就不写了～

##### ```Netbeans```安装```Android```和```Ruby```插件：

	http://nbandroid.org/release72/updates/updates.xml     #Android插件
	https://blogs.oracle.com/geertjan/resource/nb-72-community-ruby.xml      #Ruby插件