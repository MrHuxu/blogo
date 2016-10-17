# C/C++标准库学习 & HOJ_1280 & HOJ_1284  

1. ```C++```中的哈希：```C++```中其实已经提供了哈希表(散列表)的实现，即标准库```map```，语法如下：
	
		#include <map>
		
		map<int, string> test;
	
	这样就建立了一个```key```为整数，```value```为字符串的哈希表，添加一个```key-value```对和获得一个```key```的```value```也很简单，使用与数组类似的方法即可
		
		test[1] = "test";
		cout << test[1] << endl;     // => test
		
	当然，这个标准库也有自己的迭代器，在查找元素和删除元素时都要用到，使用```find```函数创造迭代器后，可以进行相关的操作，代码如下  

        //查找操作，在找不到键值对时，迭代器最终会指向表尾	    
        search_test = test.find(3);
        if(search_test == test.end())
          cout << "there is not a key-value pair which key is 3" << endl;
        	
        //删除操作，如果迭代器找到该键值对，会停止迭代，这时可以调用删除的函数
        delete_test = test.find(2);
        test.erase(delete_test);   //即删除该迭代器指向的键值对
        
2. ```atoi```函数：即将字符串转为整数的函数，包含在```stdlib.h```中，用法很简单：

		int test = atoi("123");
		
然后，照理是两道```HOJ```水题，话说我今天把我之前做的一些网页的```CSS```也加上了```overflow: auto```，好看了好多啊有木有~~

```HOJ_1280```: 简单的字符串处理，理解了题意基本上就能做出来，代码如下：  

    #include <iostream>
    #include <cstring>
    #include <string>
    using namespace std;

    int main(){
        int n, order_len, message_len;
        char order[10], message[30], tmp;
        cin >> n;
        while(n){
            cin >> order;
            order_len = strlen(order);
            cin >> message;
            message_len = strlen(message);
            //cout << order << ' ' << message << endl;
            for(int i = order_len - 1; i >= 0; i--){
                switch(order[i]){
                    case 'J':
                        tmp = message[message_len - 1];
                        for(int j = message_len - 1; j > 0; j--)
                            message[j] = message[j - 1];
                        message[0] = tmp;
                        break;
                    case 'C':
                        tmp = message[0];
                        for(int j = 0; j < message_len - 1; j++)
                            message[j] = message[j + 1];
                        message[message_len - 1] = tmp;
                        break;
                    case 'E':
                        for(int j = message_len / 2 - 1, k = message_len - 1; j >= 0; j--, k--){
                            tmp = message[k];
                            message[k] = message[j];
                            message[j] = tmp;
                        }
                        break;
                    case 'A':
                        for(int j = 0, k = message_len - 1; j < k; j++, k--){
                            tmp = message[k];
                            message[k] = message[j];
                            message[j] = tmp;
                        }
                        break;
                    case 'P':
                        for(int j = 0; j < message_len; j++){
                            if(message[j] >= '1' && message[j] <= '9')
                                message[j] = (char)((int)message[j] - 1);
                            else if(message[j] == '0')
                                message[j] = '9';
                        }
                        break;
                    case 'M':
                        for(int j = 0; j < message_len; j++){
                            if(message[j] >= '0' && message[j] <= '8')
                                message[j] = (char)((int)message[j] + 1);
                            else if(message[j] == '9')
                                message[j] = '0';
                        }
                        break;
                    default:
                        break;
                }
            }
            cout << message << endl;
            n--;
        }
    }  
	
```HOJ_1284```: 这道题本来我一看到就想用后缀表达式来做，但是想了一会儿，觉得用后缀表达式有点难，再想，我去，这题用栈不就行了么，大致思想就是，在遍历开始和左括号开始，都独立出一个栈，而遇到右括号，则将目前栈里的值反悔并作为上一个栈的栈顶元素，遇到数字则与当前栈顶相加，最后将剩下的一个栈里所有元素相加就可以得到结果。  

好吧，我觉得我这样说是没人能懂的，但是这道题的确不错，我觉得做这样的栈模拟题很有意思啊，要是能有人给我一个HOJ上的栈模拟题列表就好了。。。  

话不多说，代码如下：

    #include <iostream>
    #include <map>
    #include <string>
    #include <cstring>
    #include <cstdlib>
    using namespace std;
    #define N 100

    typedef struct {
        int data[N], top;
    } STACK;

    map<string, int> atlist;

    int isupper(char ch) {
        return ch >= 'A' && ch <= 'Z';
    }

    int islower(char ch) {
        return ch >= 'a' && ch <= 'z';
    }

    int isdigit(char ch) {
        return ch >= '0' && ch <= '9';
    }

    char cmd[N];

    int solve(int &index) {
        STACK s;
        string tmp;
        char ch;
        int ans = 0, i, t;
        s.top = -1;
        while (cmd[index] && cmd[index] != ')') {
            ch = cmd[index];
            if (isdigit(ch)) {
                char num[N];
                i = 0;
                while (isdigit(cmd[index]))
                    num[i++] = cmd[index++];
                num[i] = 0;
                t = atoi(num);
                s.data[s.top] *= t;
            } else if (isupper(ch)) {
                index++;
                tmp.clear();
                tmp += ch;
                while (islower(cmd[index]))
                    tmp += cmd[index++];
                if (atlist[tmp] == 0)
                    return -1;
                s.data[++s.top] = atlist[tmp];
            } else if (ch == '(') {
                t = solve(++index);
                if (t == -1)
                    return -1;
                else
                    s.data[++s.top] = t;
            }
        }
        index++;
        for (i = 0; i <= s.top; i++)
            ans += s.data[i];
        return ans;
    }

    int main() {
        while (cin >> cmd) {
            int num;
            if (!strcmp(cmd, "END_OF_FIRST_PART"))
                break;
            cin >> num;
            atlist[string(cmd)] = num;
        }
        while (cin >> cmd) {
            int ans, index = 0;
            if (cmd[0] == '0')
                break;
            ans = solve(index);
            if (ans == -1)
                cout << "UNKNOWN" << endl;
            else
                cout << ans << endl;
        }
        return 0;
    }

