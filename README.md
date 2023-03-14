<h3>1.1</h3>
<p>1.使用原生语言编写</p>
<p>2.只实现了两个接口</p>
<p>3.抽象出了server结构体来启动服务</p>

<h3>1.2</h3>
<p>1.抽象出mycontext类用于解析和生成json</p>
<p>2.进一步封装暴露给用户的路由方法Router参数ctx</p>
<p>3.利用http.Handle方法，利用map传入Handle接口来实现GET/PUT方法啊访问</p>
<p>4.抽象出BaseHandle和RouteBle接口，并为BaseHandleOnMap添加Router方法以此来隐藏自己的map</p>
