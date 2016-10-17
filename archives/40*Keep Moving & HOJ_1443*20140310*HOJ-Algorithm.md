# Keep Moving! & HOJ_1443  

今天豆豆告诉我河南省设计院已经给她回复了，让她过几天就回去面试，不愧是我的女朋友，一出手就有了，我也要加油了~  

已经笔试过的两家企业，感觉都还行，在结果出来之前我就不评论了。事在人为，就算这次没能通过，我相信这两次笔试的经历也将是我今后找工作道路上宝贵的经验。  

今天下午把报告交给金野老师，开题算是正式完成，然后就是看代码和敲代码的节奏。  

最近还一直在看各种程序员笔试宝典之类的东西，在看到好的解题思路时，真是有一种深深的智商被碾压感，的确是要更加努力了。  

标题是一句我非常喜欢的广告语，话说在国产企业里，这文案我觉得算是相当不错了。  

Keep Moving，永不止步！  

##### ~~~~~~~~~~~~我是萌萌的昏割线~~~~~~~~~~~~~  

1443：很简单的一道```栈模拟```，题意是计算矩阵运算过程中参与运算的数据对有多少，如果参加运算的矩阵不符合矩阵乘法，即左边的列数不等于右边的行数，就报错。  

关于括号处理，这题的思路就是模拟一个栈，扫描输入字符串，如果是左括号不处理，是字母就将其对应的行数和列数压栈，如果是右括号，则将栈顶的两个矩阵的行列数弹出进行计算，如果报错，就跳出扫描，如果有结果，则将数据对数目累加到结果中，并将乘积矩阵压栈，继续扫描。最后输出结果即可。  

在这题里我还复习了一下C++里```map```的用法，果然是方便啊，还有一点要注意的是，```(cin >> str) != EOF```在HOJ上会报错，还是得用scanf的形式。  

话不多说，代码如下：  

    #include <iostream>
    #include <cstdio>
    #include <string>
    #include <cstring>
    #include <map>
    using namespace std;
    
    typedef struct Stack{
        int r[100];
        int c[100];
        int top;
    }Stack;

    void pop(Stack &s){
        s.top--;
    }

    void push(Stack &s, int row, int column){
        s.r[++s.top] = row;
        s.c[s.top] = column;
    }
    
    int main(){
        Stack s;
        map<char, int> rows, columns;
        int n, row, column, result, err;
        char name, expression[100];
        cin >> n;
        while(n--){
            cin >> name >> row >> column;
            rows[name] = row;
            columns[name] = column;
        }
        while((scanf("%s", expression)) != EOF){
            err = result = 0;
            s.top = -1;
            for(int i = 0; i < strlen(expression); i++){
                if(expression[i] >= 'A' && expression[i] <= 'Z')
                    push(s, rows[expression[i]], columns[expression[i]]);
                else if(expression[i] == ')'){
                    if(s.c[s.top - 1] != s.r[s.top]){
                        err = 1;
                        break;
                    }else{
                        result += s.r[s.top - 1] * s.c[s.top] * s.c[s.top - 1];
                        s.c[s.top - 1] = s.c[s.top];
                        pop(s);
                    }
                }
            }
            if(err)
                cout << "error" << endl;
            else
                cout << result << endl;
        }
    }
