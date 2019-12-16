<!DOCTYPE html>
<html>
<head>
	<title>签到结果</title>
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
		p {

		}
    </style>
</head>
<body>
<div class="container box">
	<p>{{.Info}}</p>
	<p>已签到：{{.Users}}</p>
</div>
</body>

</html>
