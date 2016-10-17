# 修复寒假读书笔记ing & HOJ_1218  

今天可儿居然把我```iPad```上的```Noteshelf```给卸载了，我这个寒假一个多月以来的笔记啊，全没了。说实话，我当时真是有打她的念头了，生生给忍了下来，现在这能体会妈妈在家照顾这个妹妹的不容易了，太皮了...  

可是再皮我又能怎么办，满怀怨念恢复笔记ing(T_T)...  

今天因为没什么时间，所以就做了一道水题，一个简单的搬桌子的问题，需要注意的有两点：  

- 桌子可能从前往后搬

- 对门的房间同时占用走廊，如```5```和```6```就不能同时搬运  

然后我一开始是想把这些搬运方向排序好后，直接用贪心，即每次搬运尽量多的桌子，知道搬没了为止，代码如下：  

    #include <iostream>
    using namespace std;

    bool step(int s[], int d[], int l){
        bool finish = true;
        int en, sta;
        for(int i = 0; i < l; i++){
            if(d[i] != 0){
                finish = false;
                sta = i;
                en = d[i];
                d[i] = 0;
                break;
            }
        }

        if(!finish){
            for(int i = sta; i < l; i++){
                if(s[i] > en){
                    if(!(en % 2 == 1 && s[i] - en == 1)){
                        en = d[i];
                        d[i] = 0;
                    }
                }
            }
        }

        return finish;
    }

    int main(){
        int T, N, tmp, count, src[400], dest[400];
        cin >> T;
        while(T--){
            cin >> N;
            count = 0;
            for(int i = 0; i < N; i++){
                cin >> src[i] >> dest[i];
                if(src[i] > dest[i]){
                    tmp = src[i];
                    src[i] = dest[i];
                    dest[i] = tmp;
                }
            }
            for(int i = 0; i < N; i++){
                for(int j = i + 1; j < N; j++){
                    if(src[j] < dest[i]){
                        tmp = src[j];
                        src[j] = src[i];
                        src[i] = tmp;
                        tmp = dest[j];
                        dest[j] = dest[i];
                        dest[i] = tmp;
                    }
                }
            }
            while(!step(src, dest, N))
                count++;
            cout << count * 10 << endl;
        }
    }

但是我要说明的是，上面这段代码是错的，具体什么地方错了，我也不知道，网上找的例子都能通过，但是在```HOJ```上就是```WA```，不过在google的时候，发现了一个更牛逼的算法，就是把每个相对的房间作为一段，然后，统计所有段被经过的次数，经过次数最多的那一段，经过的次数的10的倍数就是最大搬运时间，按照这个思路写了一段，果断```AC```了，唉，看来智商还是硬伤啊...  

代码如下：  

    #include <iostream>
    using namespace std;

    int main(){
        int T, N, tmp, s, d, corr[201], time;
        cin >> T;
        while(T--){
            cin >> N;
            for(int i = 0; i < 201; i++)
                corr[i] = 0;
            time = 0;
            while(N--){
                cin >> s >> d;
                if(s > d){
                    tmp = s;
                    s = d;
                    d = tmp;
                }
                s = (s + 1) / 2;
                d = (d + 1) / 2;
                for(int i = s; i <= d; i++){
                    corr[i]++;
                    if(corr[i] > time)
                        time = corr[i];
                }
            }
            cout << time * 10 << endl;
        }
    }

