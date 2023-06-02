<?php
$servername = "localhost";
$username = "news";
$password = "123456";
$database = "news";
$url = 'http://192.168.158.32/';
// 创建连接
$conn = new mysqli($servername, $username, $password,$database);
// 检测连接
if ($conn->connect_error) {
    die("连接失败: " . $conn->connect_error);
}
$conn->query('set names utf8');
    //MySQL初始化，运行一次后注释掉下面的内容
    //创建accounts表
//     $sql = "CREATE TABLE accounts(uid int primary key auto_increment,username varchar(10) not null unique,password varchar(32))";
//     if ($conn->query($sql) === TRUE) {
//         die ("Table accounts successfully created.\n");
//     }else {
//         echo "cant create CREATE TABLE accounts\n";
//     }
//     //创建posts表
//     $sql = "CREATE TABLE posts(pid int primary key auto_increment,title varchar(100) not null,content text not null,createTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP)";
//     if ($conn->query($sql) === TRUE) {
//         echo "数据表创建成功";
//     } else {
//         echo "Error creating table: " . $conn->error;
// }
?>