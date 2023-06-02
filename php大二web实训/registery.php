<?php
require "./conn.php";//连接数据库
//检查登录状态
if(isset($_COOKIE['USERID'])&&isset($_COOKIE['USER'])){
    echo "<script>alert('你已经登录了，请勿重复注册！');window.location.href='./index.php';</script>";
}
if(isset($_POST['pwd'])&&isset($_POST['usr'])){
$pwd = md5($_POST['pwd']);
$usr = $_POST['usr'];
// //检查是否已有同账户名用户
// $check = "SELECT * from accounts where username=$usr";
// if($check->num_rows>0){
//     echo "<script>alert('已存在该用户名，请勿重复注册！');window.location.href='./index.php';</script>";
// }else{
$sql = "INSERT into accounts(username,password) values('$usr','$pwd')";
$conn->query($sql);
//设置用户为登录状态
setcookie('USERID',$pwd,time()+7200);
setcookie('USER',$usr,time()+7200);
echo "<script>window.location.href='./index.php';</script>";
}
?>