# 用ocra打包Ruby项目 & HOJ_1272  

今天突然心血来潮，想试试把```Ruby```项目打包成exe文件，原来我用```shoes```写```GUI```程序的时候，曾经打包成功过一次，不过后来再也没好使过，平心而论，```shoes```写```GUI```程度绝对是最舒服的，但是可惜_why先生不再维护这个项目了，现在```shoes```项目在```Windows```下几乎没法打包成功了，所以我还是得求助于别的方法。  

```exerb```，```ruby2exe```两个也是很久都没人打理了，还是得靠```ocra```，不过在学校的时候，用过这个```gem```，打包纯```tk```项目还行，但是如果我包含别的库的话，比如我做的图片下载器，要使用```net/http```库，就会打包失败，用什么选项都不好使，今天再次在```Windows```下装上了```ocra```，试了一下，居然打包成功了，看来```ocra```也对```Win8.1```不感冒啊。  

不过即使打包成功，也不是一句简单的```ocra xxx.rb```就能成功了，还要加上参数，经过多方归纳，以下这条命令基本能完成我的所有打包任务了~  

	ocra --no-autoload --add-all-core --windows --icon c:/Users/Huxu/Desktop/cc.ico cc_gui.rb c:\ruby200\lib\tcltk  
	
其中```--icon```后面跟的是图标位置，这是我第一次在```Windows```下编写```Ruby```，有一个小发现，就是，其实是```Windows```下，使用```/```也是可以进行目录跳转的，因为我用```tk```自带的```chooseDirectory```方法，打印处选择的目录后，发现是用```/```表示的，而且在上面的那条命令中，图标的路径是用```/```表示的，也可以正确使用。  

总体来说，算得上是一次美妙的编程经历，而且我再也不用为了下个图片切到```Mac```下了~  

然后仍然是一道水题，这题的题目真是不能更简单了，就不说什么了，直接放代码吧：  

    //HOJ_1272
    #include <iostream>
    using namespace std;

    void step(int data[], int length){
        int tmp[100];
        for(int i = 0; i < length; i++)
            tmp[i] = data[i] / 2;
        data[0] = (tmp[length - 1] + tmp[0]) % 2 == 0 ? (tmp[length - 1] + tmp[0]) : (tmp[length - 1] + tmp[0] + 1);
        for(int i = 1; i < length; i++)
            data[i] = (tmp[i - 1] + tmp[i]) % 2 == 0 ? (tmp[i - 1] + tmp[i]) : (tmp[i - 1] + tmp[i] + 1);
    }

    int main(){
        int N, counts[100] = {0}, steps;
        bool unfinished;
        while(cin >> N){
            if(N == 0)
                break;
            else{
                steps = 0;
                unfinished = true;
                for(int i = 0; i < N; i++)
                    cin >> counts[i];
                while(unfinished){
                    unfinished = false;
                    for(int i = 0; i < N - 1; i++){
                        if(counts[i] != counts[i + 1]){
                            unfinished = true;
                            step(counts, N);
                            steps += 1;
                            break;
                        }
                    }
                }
                cout << steps << ' ' << counts[0] << endl;
            }
        }
    }
