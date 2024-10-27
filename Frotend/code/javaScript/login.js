function login(event) {
    event.preventDefault();
    
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            username: username,
            password: password,
        }),
    })
        .then((response) => {
            if (!response.ok) {
                throw new Error(`Network response was not ok: ${response.statusText}`);
            }
            return response.json();
        })
        .then((data) => {
            console.log('Response Data:', data);

            if (data.status === 'ok') {
                alert("รหัสถูกต้อง"); // Correct password
                window.location.href = "http://127.0.0.1:5500/Frotend/index.html"; // Redirect to the home page on successful login
            } else {
                alert("รหัสไม่ถูกต้อง"); // Incorrect password
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert("An error occurred: " + error.message);
        });
}
