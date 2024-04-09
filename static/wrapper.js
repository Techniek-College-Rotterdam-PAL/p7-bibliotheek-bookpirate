// function submitForm() {
//     const form = document.getElementById("LoginForm");
//     const formData = new FormData(form);
//     const xhr = new XMLHttpRequest();
//
//     xhr.open("POST", "http://127.0.0.1:8080/register", true);
//
//     xhr.setRequestHeader("Content-Type", "application/json");
//
//     xhr.onload = function () {
//         if (xhr.status === 200) {
//             console.log("Request successful:", xhr.responseText);
//         } else {
//             console.error("Request failed with status:", xhr.status);
//         }
//     };
//
//
//     xhr.send(formData);
// }

function sendRequest(elementId, endPoint) {
    const form = document.getElementById(elementId);
    const formData = new FormData(form);

    const jsonData = {};
    formData.forEach((value, key) => {
        jsonData[key] = value;
    });
    var resp = {}
    // Use the fetch API to send a POST request with JSON data
    fetch("http://127.0.0.1:8080/" + endPoint, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(jsonData),
    })
        .then(response => response.json())
        .then(data => {
            console.log(data)
            errorMessage(data)
            // if (!data === null) {
            //     console.log("fihrbefjuon")
            //     errorMessage(data)
            // }
            if (data.message === "Admin Needed") {
                window.location.replace("http://127.0.0.1:8080/contact-owner")
            }
            resp = response.json();
        })
        .then(data => {
            console.log("Request successful:", data);
        })
        .catch(error => {
            console.error("Request failed:", error);
        });
    return resp
}


function fetchMovies() {
    fetch("http://127.0.0.1:8080/api/v1/fetch-books", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({list: 30})
    })
        .then(response => response.json())
        .then(data => {
            renderMovies(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function renderMovies(response) {
    const booksMain = document.getElementById('books_main');
    booksMain.innerHTML = '';
    const data = response.data;

    if (Array.isArray(data)) {
        data.forEach(book => {
            const bookDiv = document.createElement('div');
            bookDiv.className = 'col-md-4';

            // language=HTML
            bookDiv.innerHTML = `
                <div class="card mb-4 shadow-sm" style="background: #3c0e0e; border-radius: 20px">
                    <a class="#" href="#">
                    <img class="card-img-top" src="static/images/no-image.svg" alt="Thumbnail"
                         style="height: 220px; width: 100%; display: block; border-radius: 20px">
                    <!-- text -->
                    <div class="card-body">
                        <h5 class="text-center pb-0" style="color: white">${book.title}</h5>
                    </div>
                    </a>
                </div>
            `;
            booksMain.appendChild(bookDiv);
        });
    } else {
        console.error('Data is not an array:', data);
    }
}

function createAccount() {
    d = sendRequest('registrationForm', 'api/v1/register')
    console.log(d)
    setAuthToken(d.data.token)
}

function setAuthToken(token) {
    localStorage.setItem('pb_token', token);
}

function getAuthToken() {
    return localStorage.getItem('pb_token');
}

function removeAuthToken() {
    localStorage.removeItem('bp_token');
}
