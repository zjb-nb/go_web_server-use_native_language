<h3>1.1</h3>
<p>1.使用原生语言编写</p>
<p>2.只实现了两个接口</p>
<p>3.抽象出了server结构体来启动服务</p>

<h3>1.2</h3>
<p>1.抽象出mycontext类用于解析和生成json</p>
<p>2.进一步封装暴露给用户的路由方法Router参数ctx</p>
<p>3.利用http.Handle方法，利用map传入Handle接口来实现GET/PUT方法啊访问</p>
<p>4.抽象出BaseHandle和RouteBle接口，并为BaseHandleOnMap添加Router方法以此来隐藏自己的map</p>
<p>5.利用闭包特性实现filter来像洋葱一样包装自身的请求</p>
<p>6.新增panic</p>

<p>1.3</p>
<p>1.3.1</p>
<p>利用树实现简易路由，即完全匹配查找</p>


<p>1.4</p>
<p>1.增加监听函数waitshutdown来监听用户信号强制退出</p>
<p>2.增加hook函数</p>
<p>3.新增GracefulShutDown结构体来实现请求的拒绝和优雅的退出，实际是hook函数</p>