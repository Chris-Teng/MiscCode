<?php
header('content-type:text/html;charset=utf-8');
require "./conn.php";//连接数据库

//检查登录状态
if(isset($_COOKIE['USERID'])&&isset($_COOKIE['USER'])){
    $title = $_POST['title'];
    $content = $_POST['content'];
    if($title==''||$content=='') echo "<script>alert('标题和内容不能为空！');window.location.href='./publish.php';</script>";
    $sql = "INSERT into posts(title,content) values('$title','$content')";
    $result = $conn->query($sql);
    if($result){
    $result = $conn->query("SELECT * FROM posts order by pid desc limit 1");
    $row = $result->fetch_assoc();
    echo "<script>window.location.href='".$url."posts.php?id=".$row["pid"]."'</script>";
    }
}else{
    echo "<script>alert('你还没登陆');window.location.href='./index.php';</script>";
}
?>
