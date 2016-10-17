# Ruby bits 学习笔记  

记忆力是靠不住的，写下来！

1. Ruby里```or```运算符的含义是，前者为真，则无论后者真值如何，都返回前者；若前者为假，则无论后者真值如何，都返回后者。  
        
       nil or nil         # => nil
       false or true      # => true
       false or false     # => false
       false or nil       # => nil
       true or nil        # => true
       true or false      # => true

2. 而```and```运算符正好相反，前者为假时返回前者，前者为真时反悔后者，返回值与后者的真值无关。

3. Ruby里，Array的声明是放括号```[]```，Hash的声明是花括号```{}```。  

4. Ruby里使用简单的```float```存储小数是不精确的，因为二进制表示十进制存在一定的误差。可以使用```BigDecimal```类来声明十进制数。

        require 'bigdecimal'

        0.3 - 0.2 == 0.1    #//=> false

        num1 = BigDecimal('0.3')
        num2 = BigDecimal('0.2')
        num3 = BigDecimal('0.1')
        num1 - num2 == num3    #//=> true

5. 在Ruby中，哈希作为方法动态参数，应使用```{}```初始化，而数组应该用```*```表示，不能用```[]```初始化，用```[]```初始化表示这是一个数组参数。

        def test_hash(para1, para2 = {})   # right
        def test_arr(para1, para2 = [])   #wrong
        def test_arr(para1, *para2)   #right
       
6. 在Ruby中，方法中出现的异常可以用```raise```抛出，然后在执行方法时，使用```begin...end```语句，并且使用```rescue```捕获这个异常。

        def exce(num)
          if num == 1
            raise Exception.new
          end
          p num
        end

        begin 
          exce(1)
        rescue Exception
          warn 'Num1 Exception: It works!'
        end
        # => Num1 Exception: It works!
        # => It doesn't print the num

7. ```begin...end```在Ruby里被定义成一个```expression```，而```do...end```被定义成一个```block```，有一定区别。

8. ```private```完全没有权限，```public```拥有完全权限，```protected```方法：```hidden from outside but accessible from other instances of same class```。

9. 使用```include```包含一个```module```时，大致相当于继承父类，但是使用的```module```的话，通过多次```include```可以获得多个```module```的内部方法，但是使用继承的话只能获得一个父类的内部方法。

       module Shareable
         def share_on_facebook
           p "this is module 1"
         end
       end
 
       module Favoritable
         def add_to_delicious
           p "this is module 2"
         end
       end

       class Testclass
         def pclass
           p 'This is the class'
         end
       end
       
       class Post
         include Shareable
         include Favoritable
       end
        
       class Image < Testclass        
       end

       p = Post.new
       p.share_on_facebook   # => this is module 1
       p.add_to_delicious   # => this is module 2
       
       i = Image.new
       i.pclass   # => This is the class

10. 使用```extend```包含一个```module```，可以把```module```中方法当做类方法来使用，```include```是当做实例方法。  

        module Modtest
          def modmethod(para)
            p "this is the mothod: #{para}"
          end
        end

        class Clatest
          extend Modtest
        end

        Clatest.modmethod('Modtest')

11. 在类实例化后再包含```module```：
       
        module Modtest
          def modmethod(para)
            p "this is the mothod: #{para}"
          end
        end

        class Clatest
        end

        c = Clatest.new
        c.extend(Modtest)
        c.modmethod('Hehe')

12. ```module```内部```included```的用法：

        module Modtest
           def self.included(base)
             base.extend(Classmethod)
           end

           def modmethod(para)
             p "this is the mothod: #{para}"
           end

           module Classmethod
             def pclass
               p 'this is used as a class method'
             end
           end
         end 

         class Clatest
           include Modtest
         end

         c = Clatest.new
         c.modmethod("Hehe")
         Clatest.pclass
