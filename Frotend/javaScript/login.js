function login() {
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  const xhttp = new XMLHttpRequest();
  xhttp.open("POST", "http://localhost:8080/login");
  xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

  xhttp.onerror = function () {
    Swal.fire({
      text: "Unable to connect to the server. Please check your connection and try again.",
      icon: 'error',
      confirmButtonText: 'OK'
    });
  };

  xhttp.onreadystatechange = function () {
    if (this.readyState === 4) {
      if (this.status === 200) {
        try {
          const objects = JSON.parse(this.responseText);
          if (objects['status'] === 'ok') {
            localStorage.setItem("jwt", objects['accessToken']);
            Swal.fire({
              text: objects['message'],
              icon: 'success',
              confirmButtonText: 'OK'
            }).then((result) => {
              if (result.isConfirmed) {
                window.location.href = './index.html';
              }
            });
          } else {
            Swal.fire({
              text: objects['message'],
              icon: 'error',
              confirmButtonText: 'OK'
            });
          }
        } catch (error) {
          Swal.fire({
            text: "Unexpected error. Please try again later.",
            icon: 'error',
            confirmButtonText: 'OK'
          });
        }
      } else {
        Swal.fire({
          text: "Server error. Please try again later.",
          icon: 'error',
          confirmButtonText: 'OK'
        });
      }
    }
  };

  xhttp.send(JSON.stringify({
    "username": username,
    "password": password
  }));
  return false;
}