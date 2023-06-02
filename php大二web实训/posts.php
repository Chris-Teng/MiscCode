<?php
header('content-type:text/html;charset=utf-8');
//连接数据库
require "./conn.php";
$pid = $_GET['id'];
$result = $conn->query("SELECT title,content,createTime FROM posts where pid=$pid");
while($row = $result->fetch_assoc()) {
$title = $row['title'];
$content = $row['content'];
$date = $row['createTime'];
}
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>首页</title>
    <link rel="stylesheet" href="123.css">
</head>
<body>
    <div id="top">
        <div class="top_left"><img src="023.png" alt=""></div>
        <div  class="top_right">信息技术学院</div>
    </div>
<div class="box">
        <input type="radio" name="button" id="button-1" checked="checked">
        <input type="radio" name="button" id="button-3">
        <input type="radio" name="button" id="button-2">
        <input type="radio" name="button" id="button-4">
        <ul>
            <li><img src="3.png" alt=""></li>
            <li><img src="10.jpg" alt=""></li>
            <li><img src="11.jpg" alt=""></li>
            <li><img src="v1.jpg" alt=""></li>
        </ul>
        <label for="button-1">1</label>
        <label for="button-2">2</label>
        <label for="button-3">3</label>
        <label for="button-4">4</label>
</div>
<div   class="poi">
<?php 
echo <<<EOF
<h1 class="plo">$title</h1>
<h3 class="poiu">$date</h3>
<p class="poki">$content</p>
EOF;
?>

</div>
<div id="zuihou">
    <div class="ti">Copyright 2021 All Rights Reserved 南京森林警察学院    地址：南京市仙林大学城文澜路28号 </div>
    <div class="ty">邮编：210023    电话：025-85878800</div>
</div>
</body>
</html>
