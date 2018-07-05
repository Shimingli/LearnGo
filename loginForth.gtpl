<html>
<head>
<title>仕明测试登录</title>
</head>
<body>

<form  action="/shiming" method="post">
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球
用户名:<input type="text" name="username">
密码:<input type="password" name="password">
<input type="hidden" name="token" value="{{.}}">
<input type="submit" value="登陆">
</form>
<div class="separator"></div>
<div class="wiki-common-headTabBar">
<a href="https://www.baidu.com/" nslog="normal" nslog-type="10600112" data-href="https://www.baidu.com/s?ie=utf-8&fr=bks0000&wd=">网页</a>
<a href="http://news.baidu.com/" nslog="normal" nslog-type="10600112" data-href="http://news.baidu.com/ns?tn=news&cl=2&rn=20&ct=1&fr=bks0000&ie=utf-8&word=">新闻</a>
<a href="https://tieba.baidu.com/" nslog="normal" nslog-type="10600112" data-href="https://tieba.baidu.com/f?ie=utf-8&fr=bks0000&kw=">贴吧</a>
<a href="https://zhidao.baidu.com/" nslog="normal" nslog-type="10600112" data-href="https://zhidao.baidu.com/search?pn=0&&rn=10&lm=0&fr=bks0000&word=">知道</a>
<a href="http://music.baidu.com/" nslog="normal" nslog-type="10600112" data-href="http://music.baidu.com/search?f=ms&ct=134217728&ie=utf-8&rn=&lm=-1&pn=30&fr=bks0000&key=">音乐</a>
<a href="http://image.baidu.com/" nslog="normal" nslog-type="10600112" data-href="http://image.baidu.com/search/index?tn=baiduimage&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=">图片</a>
<a href="http://v.baidu.com/" nslog="normal" nslog-type="10600112" data-href="http://v.baidu.com/v?ct=301989888&rn=20&pn=0&db=0&s=22&ie=utf-8&fr=bks0000&word=">视频</a>
<a href="http://map.baidu.com/" nslog="normal" nslog-type="10600112" data-href="http://map.baidu.com/m?ie=utf-8&fr=bks0000&word=">地图</a>
<a href="https://wenku.baidu.com/" nslog="normal" nslog-type="10600112" data-href="https://wenku.baidu.com/search?lm=0&od=0&ie=utf-8&fr=bks0000&word=">文库</a>
<b class="baike">百科</b>
</div>

</body>
</html>