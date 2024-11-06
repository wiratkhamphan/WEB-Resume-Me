// var jwt = localStorage.getItem("jwt");
// // console.log("JWT Token:", jwt); 
// if (jwt == null) {
//     window.location.href = './login.html';
// }


function loadUser() {
  const jwt = localStorage.getItem("jwt"); // Ensure you're getting the JWT from localStorage
  if (jwt == null) {
      window.location.href = './login.html';
      return; // Exit if JWT is not found
  }

  const xhttp = new XMLHttpRequest();
  xhttp.open("GET", "http://localhost:8080/api/user", true);
  xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
  xhttp.setRequestHeader("Authorization", "Bearer " + jwt);
  xhttp.send();
  xhttp.onreadystatechange = function() {
      if (this.readyState == 4) {
          const objects = JSON.parse(this.responseText);
          if (objects["status"] == "ok") {
              const user = objects["user"];
              document.getElementById("username").innerHTML = user;
              // document.getElementById("username").innerHTML = user["username"];
              document.getElementById("fname").innerHTML = user
              // document.getElementById("fname").innerHTML = user["username"];
          } else {
              console.error(objects["message"]); // Log error message
          }
      }
  };
}

loadUser();
function logout() {
  localStorage.removeItem("jwt");
  window.location.href = './login.html';
}
