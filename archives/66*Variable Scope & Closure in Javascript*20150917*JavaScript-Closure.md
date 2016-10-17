# Variable Scope & Closure in Javascript

About Javascript closure:

> 闭包可以避免全局变量的污染；有利于对变量的控制； 
其实闭包原理很简单的，就是开辟一个小的栈内存；也可以理解为，函数执行就是生成一个闭包(作用域)；
但是闭包也有缺点：如果闭包外面有变量接收闭包内的引用类型返回值；那么这个作用域不销毁，浪费浏览器性能；（正常的闭包应该是运行过后，浏览器在空闲的时候销毁）；

Why Javascript use closure:
    
> JS没有class的概念,所以变量是全局的,很容易造成变量的污染,闭包的作用在于,闭包环境内,可以访问全局变量,在闭包环境里定义的变量,类似于class内定义的 private 变量,外部无法访问.

---

In Javascript, a immediately-execute function is not a property of the object ```window```, we can charge that it as a property of an extra environment which named ```env```, and when we use a variable in Javascript, the sequence to find the variable is as follow:

    local -> env -> window
        
Here is a snippet to explain the theory:
    
    !function A() {
      A = 1;
      console.log(A);   // => function (...)[...]
    }
    !function A() {
      var A = 1;
      console.log(A);   // => 1
    }
        
Another intereting sample:
 
    function B() {
      B = 2;
      console.log(B);
    }
    
    B();   // => 2, function B is a property of object window
    B();   // error, the first time execute function B, the property B of window has been set 2, not a function
    