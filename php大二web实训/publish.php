<?php
header('content-type:text/html;charset=utf-8');
//连接数据库
require "./conn.php";
//检查登录状态
$status = 0;
//首先判断登录状态
//如果已经设置了cookie，将cookie值和数据库比对，如果存在该cookie，视为已经登录
//如果没有设置cookie，检查$_POST中是否有用户名和密码
//如果设置了，说明用户尝试登录，相应的返回登陆成功和登录失败
//如果没有设置，则展示默认首页
if(isset($_COOKIE['USERID'])&&isset($_COOKIE['USER'])){
    $usr = $_COOKIE['USER'];
    $pwd = $_COOKIE['USERID'];
    $query = "SELECT password from accounts where username='$usr'";
    $result = $conn->query($query);
    $row = $result->fetch_assoc();
    $password = $row["password"];
    if($password==strval($pwd)){
        $status = 1;
    }
    }
elseif(isset($_POST['pwd'])&&isset($_POST['usr'])){
    //获取用户提交的用户名和密码
    $pwd = md5($_POST['pwd']);
    $usr = $_POST['usr'];
    $query = "SELECT password from accounts where username='$usr'";
    $result = $conn->query($query);
    $row = $result->fetch_assoc();
    $password = $row["password"];
    if($password==strval($pwd)){
        //设置用户为登录状态
        $status = 1;
        //为用户保留登录状态
        setcookie('USERID',$pwd,time()+7200);
        setcookie('USER',$usr,time()+7200);
    }else echo "<script>alert('登录失败，请检查用户名或密码！');window.location.href='./index.php'</script>";
    }
?>


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>文章发布</title>
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
<div id="header">
    <div class="navigation">
            <ul>
                <li ><a href="" class="index">2020信息学院大事件</a></li>
                <li><a href="">今日必看</a></li>
                <li><a href="">推荐</a></li>
                <li><a href="">小小网安</a></li>
                <li><a href="">情报那点事</a></li>
                <li><a href="">老师之间的较量</a></li>
            </ul>
    </div>
</div>
<div id="middle">
    <div class="middle_left1">
    <?php
if($status === 1){
    echo <<<EOF
    <p>欢迎用户 $usr !</p>
    <div><a href="./publish.php" class="one12" >发布文章</a></div> 
    <div><a href="./exit.php" class="one123">退出</a></div> 
    EOF;
}else
    echo <<<EOF
    <div>
            <form action="" method="POST">   
                用户名:<br>
                <input type="text" name="usr"><br>       
                密码:<br>         
                <input type="password" name="pwd"><br>
                <input type="submit" value="登录" class="delu">
            </form>
        </div>
    EOF;
?>
        <div class="middle_left2">
        <h2>最新文章</h2>

            <?php
//文章展示
$articles = $conn->query("SELECT pid,title FROM posts order by pid desc");
if($articles->num_rows>0){
    while($row = $articles->fetch_assoc()) {
        echo "<a href='".$url."posts.php?id=".$row["pid"]."'>".$row["title"]."</a><br>";
    }
}else{
    echo '还没有文章哦';
}
?>
        </div>
    </div>
    <div class="middle_right2">
        <div class="middle_right21">
            <div class="zhaung8">发布文章</div>
            <div class="middle_right211">
                <form action="./pub.php" method="POST">
                <div class="0p">标题:<input type="text" class="op" name="title"></div>                   
                <div class="0d">内容:<input type="text" class="od" name="content"></div>
                <div class="09">
                        <input type="submit" value="提交" class="lkl">
                        <input type="reset" value="重置" class="lkl2">
                </div>
                </form>
            </div>
        </div>
    </div>
</div>
<div id="zuihou">
    <div class="ti">Copyright 2021 All Rights Reserved 南京森林警察学院    地址：南京市仙林大学城文澜路28号 </div>
    <div class="ty">邮编：210023    电话：025-85878800</div>
</div>
</body>
</html>