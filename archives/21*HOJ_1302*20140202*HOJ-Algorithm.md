# HOJ_1302  

今天因为要去江北那点有点事，下午回来又睡了一觉，所以就只刷了一道水题。  

昨天晚上九江起了很大的雾，今天中午还没散去，车子在大桥中间的时候，前方和后方的路都看不见了，这种感觉很有意思，好像我们就是行驶在一段凭空出现的道路上。  

还是说水题吧。  

这道题很简单，就是判断一个数是不是可以拆散成阶乘的和，本来想把可能性都列出来打表的，结果发现还不如用贪心呢，需要注意的是，0的阶乘是1，一定要带上，我就是这样```WA```了好几次。。。  

    #include <iostream>
    using namespace std;

    typedef struct Team{
        char name[100];
        int detail[4][2], correct, penalty;
    } Team;

    int main(){
        int n, winner = 0;
        Team t[100];
        cin >> n;
        for(int i = 0; i < n; i++){
            t[i].correct = 0;
            t[i].penalty = 0;  //如果把这两个值在结构体内确定的话，clang++是warning，但在HOJ上无法编译
            cin >> t[i].name >> t[i].detail[0][0] >> t[i].detail[0][1] >> t[i].detail[1][0] >> t[i].detail[1][1] >> t[i].detail[2][0] >> t[i].detail[2][1] >> t[i].detail[3][0] >> t[i].detail[3][1];
            for(int j = 0; j < 4; j++){
                if(t[i].detail[j][1] != 0){
                    t[i].correct += 1;
                    t[i].penalty = t[i].detail[j][0] > 1 ? t[i].penalty + t[i].detail[j][1] + (t[i].detail[j][0] - 1) * 20 : t[i].detail[j][1] + t[i].penalty;
                }
            }
        }
        for(int i = 0; i < n; i++){
            if(t[i].correct > t[winner].correct)
                winner = i;
            else if(t[i].correct == t[winner].correct){
                if(t[i].penalty < t[winner].penalty)
                    winner = i;
            }
        }
        cout << t[winner].name << ' ' << t[winner].correct << ' ' << t[winner].penalty << endl;
    }
