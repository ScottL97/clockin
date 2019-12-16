<!DOCTYPE html>
<html>
<head>
	<title>签到</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">	
	<link rel="stylesheet" href="/static/css/bootstrap.min.css" />
	<style>
		body {
			background: #8a5c5c;
		}
		.box {
			padding: 50px;
			background: #e48b8b;
			box-shadow: 0px 5px 16px 7px black;
			margin-top: 50px;
		}
		.title {
			text-align: center;
			margin: 50px;	
		}
	</style>
</head>
<body>
	<div class="container box">
	<h2 class="title">新员工晨会签到</h2>
	<form action="/clockin/result" method="post">
		<div class="form-group">
			<label for="name">姓名：</label>
			<input id="name" type="text" name="Name" class="form-control"/>
			<small class="form-text text-muted">请输入中文姓名</small>
		</div>
		<div style="text-align: right">
			<button type="submit" value="提交" class="btn btn-primary">提交</button>
		</div>
	</form>
	</div>
	<script src="http://code.jquery.com/jquery-3.4.1.slim.min.js"></script>
	<script src="http://cdnjs.cloudflare.com/ajax/libs/popper.js/1.15.0/umd/popper.min.js"></script>
	<script src="/static/js/bootstrap.min.js"></script>
</body>
</html>
