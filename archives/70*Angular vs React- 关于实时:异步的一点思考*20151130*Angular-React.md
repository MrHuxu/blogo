# Angular vs React: 关于实时/异步的一点思考

之所以写这篇文章，主要是因为看了[Avalon](http://avalonjs.github.io/)的作者司徒正美的[一条微博](http://weibo.com/1896474751/CkpppycVZ?type=comment)，这条微博里所描述的前端开发的难题，正是我自己在使用Angular重写Blog时所遇到的痛点之一，在解决问题的过程中，我也有了自己的一点感想和思考，就在这里写下来吧。

---

在编写这个blog的前端代码的时候，我遇到了一个很棘手的问题，也就是在显示单篇文章的时候，我需要在文章内容加载完成后，使用```highlightjs```来给代码着色，并且使用一段js代码来生成行号，最后把```Nprogress```的进度条走到100%，如果按照同步编程的思路，这个过程其实很简单：

![](https://raw.githubusercontent.com/MrHuxu/img-repo/master/blog/angular%20vs.%20react.jpg)

所以一开始，我的code是这样的:

1. 首先在页面上绑定$scope里content这个变量:

        <!-- in html -->
        <p ng-bind="content"></>

2. 按照上面的流程把相应函数写到controller中:

		// in controller
		$scope.content = data; // may this command refresh ui?
		highlightCode();
		showLineNum();
		Nprogress.done();

当然结果肯定是**WTF!!!!!!**

我们再回头看刚刚所绘制的流程图，可以明显看出我们设想的情况太过美好，因为在Angular的世界，填充数据 -> 渲染，这个过程并不是可控的，```$scope.content = data```这句代码给仅仅model绑定了了数据，但是这些数据什么时候被渲染到页面上我们并不知道，所谓的实时，其实是**在数据和界面之间建立了一个动态的平衡**，所以当我们在后来执行```highlightCode```和```showLineNum```这两个函数是，很容易出现错误，因为函数要操作的数据或者dom，并没有被渲染到页面上。

---

上面那一段是我很久之前写的，其实解决这个问题的关键就是编写一个```postRender```函数，能够明确地在dom渲染结束后执行一些操作，我只记得当时因为这个问题google了很久，但是很遗憾的，虽然这个问题是Angular开发者的一个普遍痛点，可是却很难有一个完全正确的解答，如果我没记错的话，很多方案都用到了directive相关的黑魔法，比如parent controller和directive的渲染顺序，我试了不少方式，并没有效果，而且找了官方文档，也并没有关于渲染顺序的详细定义，所以最后以失败告终。

当然我最后还是写出了我想要的效果，代码却不忍直视：

    blogModule.directive('postRender', function(){
      return {
        restrict: 'EA',
        scope: {
          content: '=',
          callback: '&'
        },
        template: '<p ng-bind-html="content"></p>',
        replace: true,
        link: function($scope, iElm, iAttrs, controller) {
          $scope.$watch('content', function (value) {
            if (value)
              postArticleRendered();
          });
        }
      };
    });
    
简单的说，这种方案就是把渲染的部分放到了一个directive里，当controller里的数据发生变化时，directive接收到数据并渲染，同时link函数里```$watch```函数监听到数据变化并执行我们指定的postRender操作。

当然我也不知道为什么在directive里监听到数据变化就已经渲染完成了，在文档里也没有找到一个完美的解释，不管怎么样，它就是这样工作的(摊手)。

在这里我们可以归纳一下遇到这种问题时的解决方案，其实就两种:

1. 使用```digest```或```$apply```，手动的去影响渲染进程
2. 使用```$watch```，通过一些方式来监听渲染进程

但是如果使用这两种方法，首先代码会非常丑陋，其次我们违背了Angular的初衷，双向绑定的本意就是让数据和界面统一起来，当有一方发生改变时，另一方就实时更新，如果在这种大前提下我们仍然要手动的干预渲染进程，代码的丑陋也就是必然的了。

---

当然我现在基本上已经不怎么写Angular了，在我看来，Angular这个框架，至少在1.x的版本，是有很大的问题的，对于一个前端框架来说，这个框架真的是显得太过厚重，推荐使用自带的```$http```来进行数据请求，使用自带的依赖注入来进行模块化，编写组件要使用参数多的一比的```directive```，更不用说傻傻分不清楚的```provider```/```service```/```factory```三兄弟了，这个框架在提供了一些新手看上去很美好的特性同时，却又在进阶的路上隐藏了太多的细节，所以像我这样的菜鸟，最后是义无反顾的投入了React的怀抱。

当然其实把Angular和React做对比其实是很不公平的，因为React本身也说了，就是一个View级别的框架，只是做一些渲染数据的事儿，渲染前要做什么 -> 渲染 -> 渲染后要做什么，React搞定这些就行了，至于数据从哪儿来，有一堆好的框架来解决(Redux/Flux甚至Backbone)，剩下的任务，就是让渲染快一点，更快一点。

而Angular除了View之外，还要额外处理数据的部分，其实**在我看来**，这些功能大大限制了Angular在另一些方面的改进，比如说为了保证双向绑定，dom其实是处在动态平衡的状态，dom树在不断刷新，当然平时肉眼可能无法发觉，但是当我们给页面添加上一些动画时，渲染进程和动画进程就会混合在一起，而且由于双向绑定的原因，渲染进程会不停地工作直到数据和界面达到稳定，最终的结果就是一点简单的动画都会导致页面出现非常明显的卡顿。

话说回来，作为一个前端框架，加一点点动画就卡的不行，那还要你何用，更别提写directive时的restrict以及scope里那几个用于限制权限的前缀，虽然在数据层面看上去很专业，但是弄一堆复杂~~而且好像并没什么卵用的~~概念真的大丈夫？

反观React，简洁易用的组件化，清晰的lifecycle，数据的更新和页面的渲染完全是异步可控的，虽然在初学的时候需要额外了解一个数据框架，但是一旦运用熟练，在稍微大一点的项目上，性能我不敢说，至少React项目的可控性和以及组件可复用性都比Angular高出了不少。

当然写完这篇文章的时候，Angular2已经整装待发了，最大的改动就是完全不兼容Angular1.x了，这个改动不可谓不大胆。就像当初提出双向绑定一样，希望Angular2的发布，能继续给我们带来新的惊喜吧。