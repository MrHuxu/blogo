# Vultr VPS入门 & 编译安装Nginx

首先必须要推荐Vultr的日本节点，在大陆访问的话ping一般是在80到100左右，速度相当感人，每月$5的价格，能享受到768内存，15G SSD以及1000G的流量，不论是价格还是性能都是相当不错的。

这次购买Vultr之后，我倒是没那么急着搭建ShadowSocks，原因有两点，首先是我基本上不玩Facebook、Twitter之类，并没有那种随时随地翻墙的需求，其次是我还想在这个vps上搭建一些自己的网站和blog之类，还不想这么快就被封掉，所以目前只是搭建了一个Google的反代，能够随时使用Google搜索我就知足了，当然在反代Google的过程中也遇到了一些磕磕绊绊，这些我在之后的文章里会写到。

接下来就是怎么折腾这个刚入手还热乎的vps了。

### VPS入门操作

1. 系统的选择

    既然是vps那就果断上CentOS了，毕竟是RedHat系的，服务器的正统。直接在Vultr的Deploy页面就可以选择CentOS了。
    
2. 包管理器入门

    既然选择了CentOS，就必须要适应一些RedHat系的系统工具，在命令行模式下首当其冲的就是包管理工具了，CentOS上使用的```yum```，和Debian系的```apt-get```以及OS X上的```homebrew```类似，这个工具也提供了软件包的安装卸载以及依赖管理，下面是一些常用命令：

        yum check-update   # == brew update
        yum update   # == brew upgrade
        yum search   # == brew search
        yum install   # == brew install
        yum erase   # == brew uninstall
        
    详细的命令可以查看这篇[博客][1]，另外，yum缺省的安装路径为```/usr/bin```
    
        yum install mongodb
        /usr/bin/mongo
        
3. 使用密钥进行ssh登录

    安装好系统后，就可以使用ssh登录了，比如登录我的vps
    
        ssh root@45.32.254.138
        
    默认情况下使用用户名和密码登录，Vultr的初始密码着实反人类，所以最好在vps里添加ssh公钥，然后就可以直接使用密钥登录了
    
    - 首先在本地生成ssh私钥和公钥，可以查看GayHub(大雾)上的[这篇文档][2]
    - 然后使用scp把本地的公钥上传到vps上

            scp ~/.ssh/id_rsa.pub root@45.32.254.138:/root/xxx_dir/
    - 将公钥里面的内容作为验证的依据
    
            ssh root@45.32.254.138
            cat /root/xxx_dir/id_rsa.pub > ~/.ssh/authorized_keys
    
    - 多台电脑登录的话，也可以通过```scp```和```cat```把多个公钥内容添加到```authorized_keys```文件里来登录
    - 编辑一下ssh的config文件，给vps设置一个别名，就不用每次都输IP了

            # ~/.ssh/config
            Host vps
              HostName 45.32.254.138
              User     root
              
            # usage
            ssh vps
            
4. 买了vps一般不外乎翻墙和建站这些用途，像我这种想把自己的blog放上去的话，还要做一件事，就是打开vps的80和443端口，前者用于http请求，后者用于https请求。

    一般的vps可能默认并没有开这些端口，比如Vultr就只开了ssh使用的22端口，所以需要手动打开这些端口，[这里][3]有在CentOS下开端口的详细教程。

### 手动编译安装Nginx

现在建站基本上都是用nginx来作为反代的工具了，灵活的语法以及强大的功能让nginx基本上就是VPS的标配，特别是在VPS里跑着多个应用的情况下，nginx可以灵活的对多个域名进行相应的端口转发，甚至可以非常轻松的的反代Google。

安装nginx，最简单的方式就是通过yum安装了：

    yum install nginx
    
但是这样安装的nginx缺点不少，首先是版本可能会比较旧，其次就是不支持一些常用模块，所以一般还是会手动安装：

使用```wget```下载```pcre```，这个正则支持很多nginx模块都需要

    wget ftp://ftp.csx.cam.ac.uk/pub/software/programming/pcre/pcre-8.37.tar.gz ~/workspace/
    cd ~/workspace && tar -zxvf pcre-8.37.tar.gz
    
接着下载nginx，并解压

    wget http://nginx.org/download/nginx-1.8.0.tar.gz ~/workspace/
    tar -zxvf nginx-1.8.0.tar.gz
    
configure, make一气呵成：

    ./configure --user=www --group=www --prefix=/usr/local/nginx --with-http_stub_status_module --with-http_ssl_module --with-http_sub_module --with-md5=/usr/lib --with-sha1=/usr/lib --with-http_gzip_static_module --with-http_stub_status_module --with-http_ssl_module --with-http_sub_module --with-pcre=/root/workspace/pcre-8.37 --with-md5=/usr/lib          
    make
    make install

这样安装的nginx在```/usr/local/nginx```目录，可以通过```/usr/local/nginx/sbin/nginx -v```来查看版本

当然这样的nginx只是一个应用，我们可以通过[这个脚本][4]把它注册为一个服务，把这个脚本命名为```nginx```并放在```/etc/init.d```目录下，就可以通过```service```命令来控制nginx了：

    service nginx start
    service nginx stop
    service nginx restart


  [1]: https://www.centos.bz/2011/07/yum-all-command-explanation/
  [2]: https://help.github.com/articles/generating-ssh-keys/
  [3]: https://www.vultr.com/docs/setup-iptables-firewall-on-centos-6
  [4]: https://gist.github.com/MrHuxu/bfc4731694e84185c93a