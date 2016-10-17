# Event#stopPropagation & Unbind Submit of a Button in Form

Javascript里，一个元素上的事件会经理事件捕获和事件冒泡的阶段，一个元素被整体监听的情况下，在其内部的事件会被传递到被监听的元素。但是有时候我们希望内部的元素不受其父元素事件的影响，这时可以使用```stopPropagation```方法:

    $('#XXX').click(function(event){ e.stopPropagation(); });

这样xxx元素的click事件将不会被父元素上绑定的click事件捕获。

比如有这么一个需求，在一个div里面有一个input，我们需要点击div的时候弹出一个窗口，点击input的时候打印一条信息，一开始代码是这么写的:

    <html>
    <body>
      <div id="parent">
        <input id="child" />
      </div>
    </body>
    <script type="text/javascript" src="./jquery.min.js"></script>
    <script type="text/javascript">
      $('#parent').on('click', function (e) {
        alert('this is parent');
      });

      $('#child').on('click', function (e) {
        console.log('this is child');
      });
    </script>
    </html>

这时点击div的时候的确弹出了窗口，但是点击input的时候，在打印信息的同时，也弹出了窗口，就是因为div对click事件的监听也影响到了子元素，这显然不是期望的行为。

这时，我们只需要阻止子元素事件的冒泡，就可以防止在点击input的时候弹出窗口了，改动如下:

    $('#child').on('click', function (e) {
      e.stopPropagation();
      console.log('this is child');
    });

---

When you want to bind a event to a button in form, at first you should diable the submit event, because any button in one form will be binded submit event by default.

    <form>
      <button onClick='xxx_method'>Btn</button>
      <!-- 
        * when click this button, the xxx_method will be invoked and then the form will be fucking submitted
        * and this is same as the code below
      -->
      <button type='submit' onClick='xxx_method'>Btn</button>
    </form>
    
    <form>
      <button type='button' onClick='xxx_method'>Btn</button>
      <!-- the form will never be fucking submitted -->
    </form>
