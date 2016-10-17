# Nodejs + Redis Blog缓存实践

写在前面，我们先来看一下Redis的官方介绍:

> Redis is an open source (BSD licensed), in-memory data structure store, used as database, cache and message broker.

Redis是一个运行在内存中的数据库，支持一些常用的数据结构并且在内存中以key-value键值对的形式存在，正因为Redis是运行在内存中的，所以它的速度非常快，所以很多网站都会选择使用Redis作为一些常用数据的缓存，来提高访问速度。

当然这个Blog一开始是没有用缓存的，因为压根就没有查询数据库的操作，之前我在编写个人项目的时候也没有想过会有需要缓存的场景。一开始的访问速度也是可以接受的，但是在加入一些新的功能之后，访问速度却出现了明显的瓶颈，主要有以下两个: 

1. 因为使用前端分页，在Home页面一次性请求全部数据的时候，server端解析markdown耗时过高，文章数量一多就会有明显的延迟，最好在部署的时候就解析好并且缓存起来;
2. 在Projects页面使用的是GitHub的数据，对于没有认证的请求GitHub API有频率限制，而且Repo的更新时间也不需要每次都从GitHub API获取，所以必须使用缓存。

这就是这个Blog使用缓存的原因了，使用数据库显得太重，而Redis轻量并且非常易于上手，所以果断选择用Redis进行缓存了。

---

### 第三方库 3rd Part Library

Node的第三方Redis库主要有

1. [node_redis](https://github.com/NodeRedis/node_redis), 曾经的官方Redis库，但是作者已经基本不更新了，而且对Redis的特性也支持的不是很完善，操作基本上都是基于回调，需要借助bluebird等来实现Promise。
2. [ioredis](https://github.com/luin/ioredis), 当前的官方推荐Redis库，对```Sentinel```和```Cluster```等功能支持良好，更新速度快，并且原生支持Promise。

支持Promise还有一个好处，就是需要执行多个查询操作的时候，回调的方式很容易陷入回调地狱，但是用Promise的话，就可以避免这个问题(参考[这篇文章](http://blog.xhu.me/#/archives/67*ES6: 回调将死, Promise永生*20151018*JavaScript-Promise.md?_k=mhw2ck))。

所以我在这里就毫不犹豫选择ioredis了:)

---

### 基本操作 Basic Operation

Redis的基本操作可以看这个页面 [Redis: Commands](http://redis.io/commands), 每一个基本操作都对应一个输出，在ioredis里，每一个基本操作都被封装成了一个方法并返回一个promise，函数的输出会作为resolve的参数。

    var Redis = require('ioredis');
    var redis = new Redis();

    redis.ping().then((result) => {
      console.log(result);
      redis.end();
    });
    // => 'PONG', connected to the redis client successfully
    
在这段代码中，我们先使用Redis初始化了一个redis对象，然后调用了```ping```这个方法，这个方法对应Redis里的[PING](http://redis.io/commands/ping)命令，这个命令在Redis中被用来检测客户端是否成功连接，然后对返回的promise调用then方法，打印出结果```PONG```。

当我们在上面的代码中初始化redis对象的时候，其实可以看作是在终端中执行```redis-cli```命令，输入命令并且获得输出，所以当操作完成的时候，需要手动调用```end```方法来结束这个进程，否则程序将会hang住。

---

### 存 Save

熟悉了基本操作之后，就是要把数据给存到Redis里了，首先我们通过GitHub的API获得repo的信息，得到的应该是这样的一个JS对象:

    var repo = {
      name        : 'blog',
      fullName    : 'MrHuxu/blog',
      url         : 'https://github.com/MrHuxu/blog',
      star        : 0,
      homepage    : 'http://blog.xhu.me',
      description : 'My Blog',
      updatedAt   : '2015-12-19T15:37:46Z'
    }
    
存储这样的对象，最适合的就是Redis中的散列类型了，这个数据类型在Redis中专门用来存储带字段的键，[HSET](http://redis.io/commands/hset)用来给一个字段赋值，而[HMSET](http://redis.io/commands/hmset)用来一次给多个字段赋值，并且ioredis的hmset方法和redis的HMSET命令一样，同样是接受数组作为参数，所以我写了一个[obj2arr](https://github.com/MrHuxu/blog/blob/cf6439bc6da67212c8d78672799f0db65d853397/lib%2Fcommon.js#L48)函数来将对象转换成数组。

    redis.hmset(`repo:${repo.name}`, obj2arr(repo)).then((result) => {
      redis.end();
      console.log(result);   // => 'OK'
    });
    
这样我们就把一个repo的信息存到了Redis里。

---

### 取 Read

取数据话使用[HGETALL](http://redis.io/commands/hgetall)指令就可以一次性获得一个散列中的所有字段，并且redis的```hgetall```方法直接返回的就是一个JS对象了:

    redis.hgetall('repo:blog').then((result) => {
      redis.end();
      console.log(result);   // => an object contains all infos of a repo
    });
    
并且由于ioredis返回的都是promise，所以在一次性从cache中获得多个repo的信息的时候，可以用ES6的```Promise#all```很优雅的完成这个任务。

    var promiseSet = repoNames.map(name => redis.hgetall(`repo:${name}`));
    
    Promise.all(promiseSet).then((repos) => {
      ...
    });
    
---

### 过期时间 Expiration

在文章开头提到的两个使用缓存的场景，第一个生成缓存后就不需要改变了，但是对于从GitHub上获取的repo信息，是需要更新的，所以我们需要对缓存做过期时间设置，在这里我并不是对每个repo的信息都做过期设置，而是存储一个名为```repo:count```的键并让它兼职作为过期标志:

    redis.set('repo:count', repoNames.length).then((flag) => {
      if ('OK' === flag)
        return redis.expire('repo:count', 1800);
    });
    
[EXPIRE](http://redis.io/commands/expire)命令的第一个参数是键名，第二个参数是过期时间，单位是秒，这段代码就是给```repo:count```这个键设为半个小时之后过期，半个小时候这个键就会被删除。

这时从缓存中取数据的地方也要有所改动，也就是加一个判断这个标志键是否存在的过程:

    redis.exists('repo:count').then((flag) => {
      if (flag) {
        ...   // flag === 1, get repos from cache
      } else {
        ...   // flag === 0, re-cache repos and then get them from cache
      }
    });
    
Redis还有一个[TTL](http://redis.io/commands/ttl)命令来差看一个键还有多久过期:

    redis.ttl('repo:count').then(val => console.log(val));   // => 1718

---

### 总结

这个Blog对于Redis的使用还是比较简单的，基本上就只用到了上面所述的操作，不过在使用的过程中我深深感受到，简单和强大这两个词在Redis身上得到了统一，而且配合ioredis和Promise，多个键并行存取也是易如反掌，Blog的访问速度也有了很大的提高（目前瓶颈应该就是网络问题了，万恶的GFW），缓存的使用也为我今后的编程开发提供了一个良好的优化思路，我也会努力在今后的学习中去探索Redis的更多高级特性。