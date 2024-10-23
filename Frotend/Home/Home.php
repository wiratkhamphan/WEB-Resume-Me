<?php
session_start();

// ตรวจสอบว่า session มีการบันทึกข้อมูลการเข้าสู่ระบบหรือไม่
if (!isset($_SESSION['username'])) {
    // ถ้าไม่มีการ login ให้เปลี่ยนหน้าไปที่ login.html
    header("Location: http://localhost/WEBResumeMe/Frotend/login/login.php");
    exit();
}
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Home</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
</head>
<body>
    <div class="container">
        <h2 class="mt-4">Welcome to the Home Page</h2>
        <p>Hello, <?php echo htmlspecialchars($_SESSION['username']); ?>! You are logged in.</p>

        <!-- ปุ่มออกจากระบบ -->
        <form action="logout.php" method="POST">
            <button type="submit" class="btn btn-danger">Logout</button>
        </form>
    </div>
</body>
</html>
