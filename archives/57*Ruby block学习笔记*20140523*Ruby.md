# Ruby block学习笔记  

今天中午上了会儿QQ，本来是想看一下年级群又没有什么新的通知，这时Ruby爱好者这个群的图标闪了起来，进去看，发现有个哥们儿问了一个关于```block```和```yield```的问题。  

翻了翻镐头书，发现大致的语法就是方法名后面接```block```，将会把代码块放到方法中```yield```关键字所在的地方执行。  

比如以下的代码：  

    def say_hello
      yield
      yield
      yield
    end
    
    say_hello { puts 'hello world' }  # print "hello world" three tiems  
	
当然，在```yield```后面接变量的话，```block```也是可以接受带参数的代码块的。于是我把我的理解发了过去，结果群里面有个哥们儿提到了另一种用法，据他的说法，是```不带yield```的用法，代码如下：  

    def add&block
	
这个语法从来没见过啊有木有，果断Google之，大致的解释是，Ruby中的```&```符号，在修饰方法的参数时，表示方法中引用的block。然后又在StackOverflow上发现了一个[问答](http://stackoverflow.com/questions/814739/whats-this-block-in-ruby-and-how-does-it-get-passed-in-a-method-here)，其中的代码如下：  

    def meth_captures(arg, &block)
      puts block.call(arg, 0) + block.call(arg.reverse, 1)
    end

    meth_captures('pony') do |word, num|
      puts "in callback! word = #{word.inspect}, num = #{num.inspect}"
      word + num.to_s
    end  
	
输出结果是：  

    in callback! word = "pony", num = 0
    in callback! word = "ynop", num = 1
    pony0ynop1  
	
这样的程序真是把我的看醉了，下面的那个输出语句居然被执行了两次，而且那两个参数完全就不知道从哪里过来的，简直就像魔术一样有木有~  

突然我想明白了，我又犯了把```代码块```简单地看作```过程```这个错误。  

在初学C++的时候，为了便于理解，老师都会说，代码代表的就是程序的执行过程。比如```while```之后就是一个过程，```if```之后就是一个过程，当然，函数后面接一个大括号也是一个过程。这在C/C++这样的静态语言中，大多时候的确是正确的。在学习了Ruby后，很多时候，我还是在使用C++的方式编程，所以，也理所当然的把这个想法带到了Ruby里。比如在这段代码里，我一开始以为```do...end```里的代码段是方法执行后的一个过程，然后就开始到处找这两个参数的来源了Orz。。。  

这显然是错的，因为当遇到block这个概念后，就不能再这样看，在Ruby的世界里，block就是一个整体，block里的语句在内部是一个过程，而在外部看来，应该是这一整体的一部分。  

正因为如此，在C++里，用大括号包含的代码块，一般只能作为```过程```来执行（C++11支持lambda表达式了，巨大的进步有木有），但在Ruby里，这样的代码块组成了block，可以作为参数传递给一个方法，在运行时改变一个方法的内容，这就是动态编程的魅力所在！  

比如这段代码，具体的顺序应该是这样理解的：  

1. 在声明方法的时候，带了两个参数，一个是普通参数，另一个是```block```。这个```block```带有两个参数，在方法中被执行了两次，方法最后输出两次执行结果的相加值。  

2. 执行方法，并给其中的普通变量赋值为```pony```，后面的这个```do...end```才是关键所在，这一段代码，就是参数中```block```！也就是说，原来方法的两个参数，一个是字符串，另一个就是后面这个代码块。  

3. ```block```使用字符串```pony```和一个数字作为参数，执行两次，前两行输出就是这样来的。  

4. 最终得到的字符串相加，于是就有了第三行输出。  

不得不说，这段代码让我对Ruby的```block```有了更深的理解，还有```Proc```，```lambda```表达式这些概念，以后有机会，我也会通过博客写出学习感想的。  

上面那个链接还给出了一个不使用```&```而使用```yield```的方法，代码如下：  

    def meth_yields(arg)
      puts yield(arg, 0) + yield(arg.upcase, 1)
    end
    
    meth_yields('frog') do |word, num|
      puts "in callback! word = #{word.inspect}, num = #{num.inspect}"
      word + num.to_s
    end  
	
输出内容是：  

    in callback! word = "frog", num = 0
    in callback! word = "FROG", num = 1
    frog0FROG1  
	
这段代码的执行过程和上面那段类似，不过```block```不是通过作为参数被方法调用，而是通过```yield```关键字被调用的。