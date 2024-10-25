function login(event) {
    event.preventDefault();
    
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    fetch("http://localhost:3000/login", {
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
            if (!response.ok) { // Check if the request was successful
                throw new Error(`Network response was not ok: ${response.statusText}`);
            }
            return response.json();
        })
        .then((data) => {
            console.log('Response Data:', data);

            if (data.status === 'ok') {
                alert("รหัสถูกต้อง");
            } else {
                alert("รหัสไม่ถูกต้อง");
            }
        })
        .catch((error) => {
            console.error('Error:', error);
            alert("An error occurred: " + error.message);
        });
}
