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
    var resp= {}
    // Use the fetch API to send a POST request with JSON data
    fetch("http://127.0.0.1:8080/" + endPoint, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(jsonData),
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
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
        body: JSON.stringify({ list: 30 })
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
            bookDiv.className = 'row row-cols-6';

            const bookContent = `
                <div class="col">
                     <div class="book_main">
                        <h1 class="title">${book.title}</h1>
                        <p class="author">Author: ${book.author}</p>
                        <p class="isbn">ISBN: ${book.isbn}</p>
                        <p class="language">Language: ${book.language}</p>
                    </div>
                </div>
                    `;
            bookDiv.innerHTML = bookContent;
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