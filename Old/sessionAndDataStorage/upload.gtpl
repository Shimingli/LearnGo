<html>
<head>
	<title>cookie的练习</title>
</head>
<body>
<form  action="/cookieDemo" method="post">
用户名:<input type="text" name="username">
密码:<input type="password" name="password">
<input type="hidden" name="token" value="{{.}}">
<input type="hidden" name="cookie" value="{{.}}">
<input type="submit" value="登陆">
</form>
</body>
</html>