var jwt = localStorage.getItem("jwt");
if (jwt != null) {
  window.location.href = './index.html'
}

function login() {
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:3000/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log('Response Data:', data);  // Debugging the response
        
        if (data.status === 'ok') {
            console.log('Access Token:', data.accessToken);  // Log the token
            
            localStorage.setItem("jwt", data.accessToken);  // Store JWT token
            
            Swal.fire({
                title: data.message,
                icon: 'success'
            }).then((result) => {
                if (result.isConfirmed) {
                    window.location.href = './index.html';  // Redirect to index.html after login
                }
            });
        } else {
            Swal.fire(data.status, data.message, 'error');
        }
    })
    .catch(error => {
        console.error('Error:', error);
        Swal.fire('Error', 'Something went wrong, please try again.', 'error');
    });

    return false; // Prevent form submission
}
