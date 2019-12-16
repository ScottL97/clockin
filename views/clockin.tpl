<!DOCTYPE html>
<html>
<head>
	<title>签到</title>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<style>
		body {
			height: 100%;
			margin: 0px;
			background: #8a5c5c;
		}
		.box {
			text-align: center;
			background: #e48b8b;
			padding: 50px;
			box-shadow: 0px 5px 16px 7px black;
		}
		#img {
			width: 70%;	
		}
		.loading {
			position: fixed;
			width: 100%;
			height: 100%;
			background-color: #eee;
			text-align: center;
			font-size: 30px;
			top: 0px;
			left: 0px;
		}
		.progress {
			margin-top: 50px;
		}
		.footer {
			text-align: center;
			font-size: 30px;
			padding: 30px;
		}
	</style>
</head>
<body>
	<div class="box">
		<img src="{{.Imgurl}}" id="qrcode">
	</div>
	<div class="loading">
		<p class="progress">0%</p>
	</div>
	<div class="footer">
		<p>新员工晨会签到二维码</p>
	</div>
	<script src="http://code.jquery.com/jquery-3.4.1.min.js"></script>
	<script src="/static/js/preload.js"></script>
	<script>
		var imgs = ['/static/img/qrcode.png'];
		var len = imgs.length;
		$.preload(imgs, {
			each: function(count) {
				$('.progress').html(Math.round((count + 1) / len * 100) + '%');
			},
			all: function() {
				$('.loading').hide();
			}
		});
	</script>
</body>
</html>
