function Register() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;
    const password_pcf = document.getElementById("password_pcf").value;

   
    if (!username || !password || !password_pcf) {
        Swal.fire("Error", "All fields are required!", "error");
        return false;
    }
    if (password !== password_pcf) {
        Swal.fire("Error", "Passwords do not match!", "error");
        return false;
    }

    // Data to be sent to the server
    const data = JSON.stringify({
        username: username,
        password: password
    });

    console.log(data);
    
    // Create an XMLHttpRequest object
    const xhttp = new XMLHttpRequest();
    xhttp.open("POST", "http://localhost:8080/api/Register", true);
    xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

    // Define a callback function to handle the response
    xhttp.onreadystatechange = function() {
        if (this.readyState === 4) {
            if (this.status === 201) {
                Swal.fire("Success", "Registration successful!", "success").then(() => {
                    window.location.href = "login.html"; // Redirect to login page
                });
            } else {
                Swal.fire("Error", "Registration failed!", "error");
            }
        }
    };

    // Send the request
    xhttp.send(data);
    return false; // Prevent form submission
}
