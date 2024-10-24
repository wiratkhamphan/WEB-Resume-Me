<?php
session_start();
session_unset();
session_destroy();

// หลังจาก logout จะเปลี่ยนหน้าไปที่ login.html
header("Location: login.php");
exit();
