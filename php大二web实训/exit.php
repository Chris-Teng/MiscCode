<?php
setcookie("USERID",'',time()-1);
setcookie('USER','',time()-1);
echo "<script>alert('退出成功');window.location.href='./index.php'</script>";
?>