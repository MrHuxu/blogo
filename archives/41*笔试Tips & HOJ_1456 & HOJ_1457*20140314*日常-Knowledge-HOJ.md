# 笔试Tips & HOJ_1456(未完成) & HOJ_1457  

1、面向对象的三个基本特征：  

   - ```封装```，客观实物封装成抽象类，是```对象```和```类```概念的主要特性。
   - ```继承```，使用现有类的所有功能，并在不重写该类的情况下进行扩展，通过继承生成的类被称为```子类```或```派生类```，被继承的类被称为```父类```、```基类```或```超类```。
   - ```多态```，允许将子类类型的指针赋值给父类类型的指针。可以通过```覆盖```或```重载```实现多态。  

这里有一个对面向对象的[详细讲解](http://blog.csdn.net/ztj111/article/details/1854611)。   
	
2、Linux系统的系统分区有```Ext2```，```Ext3```，```Ext4```。其中Ext3比Ext2多了```日志```，而Ext4与Ext3兼容，可以不格式化直接升级，支持```更大的文件系统```和```更大的文件```，可以有```无限的子目录数量```，这里也有一个详细的Ext4和Ext3的区别[详解](http://zxboom.iteye.com/blog/383986)。  

3、```TCP/IP```的```三次握手```过程（已经在笔试里看到了两次了），具体过程如下：

  - 第一次握手：建立连接时，客户端发送```syn```包(```syn=j```)到服务器，并进入```SYN_SEND```状态，等待服务器确认； 
  - 第二次握手：服务器收到```syn```包，必须确认客户的```SYN```(```ack=j+1```)，同时自己也发送一个```SYN```包（```syn=k```），即```SYN+ACK```包，此时服务器进入```SYN_RECV```状态； 
  - 第三次握手：客户端收到服务器的```SYN＋ACK```包，向服务器发送确认包```ACK```(```ack=k+1```)，此包发送完毕，客户端和服务器进入```ESTABLISHED```状态，完成三次握手。 
	
4、```do...while```和```while```区别在于，前者会在不论判断条件成不成立的情况下都执行一次循环体里面的内容，在某些情况需要这样用。（但是我个人认为这个语句写起来很难看，还不如先手写一遍循环体再使用while）  

5、C/C++的按位运算符是```&```/```|```/```^(异或)```，使用方法如下：

    int num1 = 8, num2 = 1;
    int num3 = num1 | num2;
    cout << num3 << endl;     // => 9，必须把这个结果赋给一个变量，否则编译器会有warning并且无法输出正确结果

6、```.```运算符是不能被重载的。  

7、如果一个变量一开始就声明成指针型，那么它是可以使用```自增```运算符的，而如果一开始被声明成数组，那么就不能使用自增运算符，虽然这两者内在是一种东西。测试代码：

    int b[100];
    int *pb = b;
    *(pb + 1) = 1;
    *(b + 1) = 1;
    *(++pb) = 1;     // => 上面三个结果都是b[1] = 1
    *(++b) = 1;     // 无法正确赋值，而且编译器汇报错，b++也是一样  
		
##### ~~~~~~~~~~~~我是萌萌的昏割线~~~~~~~~~~~~~  

```1456(未完成)```：这道题就是模拟一个```队列```，然后向里面存数，但是入队的数也是有```Team```区别的，如果一个Team已经有数在队列里，那么这个Team接下来入队的数直接放在这个数的后面，否则放到队尾。  

题意是很简单的，本来我是用数组存Team，但是判断起来太麻烦，后来想到用```map```的方法，即建立一个容量够大的数组，下脚标是Team的数，对应的值就是Team序号，这样就好判断两个数是不是同一个Team里的了。这样跑例子以及POJ上同学给的数据都是对的，但是在HOJ上却是```TLE```。  

仔细看了一下，感觉代码里面没有什么卡循环的啊，算了，不纠结了，放代码吧：  

    #include <iostream>
    #include <string>
    #include <memory>
    using namespace std;
    
    int main(){
        int t, elem_num, tmp, input_num, front, rear, location, visit, k = 0;
        int elem[1000000], queue[1000000];
        string input_str;
        while(cin >> t && t != 0){
            memset(elem, 0, sizeof(elem));
            memset(queue, -1, sizeof(queue));
            for(int i = 1; i <= t; i++){
                cin >> elem_num;
                for(int j = 1; j <= elem_num; j++){
                    cin >> tmp;
                    elem[tmp] = i;
                }
            }
            front = rear = 0;
            cout << "Scenario #" << ++k << endl;
            while(cin >> input_str && input_str != "STOP"){
                if(input_str == "ENQUEUE"){
                    cin >> input_num;
                    visit = 0;
                    for(location = rear; location >= front; location--){
                        if(elem[queue[location]] == elem[input_num]){
                            visit = 1;
                            break;
                        }
                    }
                    if(visit){
                        for(int i = rear; i > location; i--)
                            queue[i] = queue[i - 1];
                        queue[location + 1] = input_num;
                        rear++;
                    }else{
                        queue[rear] = input_num;
                        rear++;
                    }
                }else if(input_str == "DEQUEUE"){
                    cout << queue[front++] << endl;
                }
            }
            cout << endl;
        }
    }


```1457```：这道题很简单，就是看矩阵里行和列有没有和为奇数的，如果没有，那么肯定就是```OK```，如果各有一行一列不是，那么改变```交叉点```的值就可以了，如果一个有一个没有，或者数量大于1，那就肯定没救，输出```Corrupt```就行了。  

代码如下：

    #include <iostream>
    #include <memory.h>
    using namespace std;
    
    int main(){
        int n, data[110][110], row[110], col[110], odd_row, odd_col, tmp1, tmp2;
        while(cin >> n && n != 0){
            memset(row, 0, sizeof(row));
            memset(col, 0, sizeof(col));
            for(int i = 0; i < n; i++){
                for(int j = 0; j < n; j++){
                    cin >> data[i][j];
                    row[i] += data[i][j];
                    col[j] += data[i][j];
                }
            }
            tmp1 = tmp2 = 0;
            for(int i = 0; i < n; i++){
                if(row[i] % 2 != 0){
                    tmp1++;
                    odd_row = i;
                }
                if(col[i] % 2 != 0){
                    tmp2++;
                    odd_col = i;
                }
            }
            if(!tmp1 && !tmp2)
                cout << "OK" << endl;
            else if(tmp1 == 1 && tmp2 == 1)
                cout << "Change bit (" << odd_row + 1 << ',' << odd_col + 1 << ')' << endl;
            else
                cout << "Corrupt" << endl;
        }
    }
