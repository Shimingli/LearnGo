<html>
<head>
<title>仕明测试登录</title>
</head>
<body>
<form action="/shiming" method="post">
	用户名:<input type="text" name="username">
	密码:<input type="password" name="password">
	age:<input type="text" name="age">
	请输入英文:<input type="text" name="eng">
	email:<input type="text" name="email">
	mobile:<input type="text" name="mobile">
	usercard:<input type="text" name="usercard">
	<input type="submit" value="登录">
</form>

<form  action="/shiming" method="get">
<select name="fruit">
<option value="apple">apple</option>
<option value="pear">pear</option>
<option value="banana">banana</option>
</select>
<input type="submit" value="登录">
</form>



<form  action="/shiming" method="get">
<input type="radio" name="gender" value="1">男
<input type="radio" name="gender" value="2">女
<input type="submit" value="提交">
</form>

<form  action="/shiming" method="get">
<input type="checkbox" name="interest" value="football">足球
<input type="checkbox" name="interest" value="basketball">篮球
<input type="checkbox" name="interest" value="tennis">网球
<input type="submit" value="提交">
</form>




</body>
</html>