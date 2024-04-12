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
            "Authorization": getAuthToken()
        },
        body: JSON.stringify(jsonData),
    })
        .then(response => response.json())
        .then(data => {
            errorMessage(data)
            if (data.message === "Admin Needed") {
                window.location.replace("http://127.0.0.1:8080/contact-owner")
            }
            if (!data.data.token === null) {
                setAuthToken(data.data.token)
            }
            resp = data
        })
        .then(data => {
            console.log("Request successful:", data);
        })
        .catch(error => {
            console.error("Request failed:", error);
        });
    return resp
}


function fetchBooks() {
    fetch("http://127.0.0.1:8080/api/v1/fetch-books", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({list: 30})
    })
        .then(response => response.json())
        .then(data => {
            renderBooks(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

function fetchAdminBooks() {
    fetch("http://127.0.0.1:8080/api/v1/fetch-all-books", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": getAuthToken()
        },
        body: JSON.stringify({list: 30})
    })
        .then(response => response.json())
        .then(data => {
            generateTable(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}


function generateTable(response) {
    const Main = document.getElementById('adminbooklist');
    Main.innerHTML = '';
    const data = response.data;


    if (Array.isArray(data)) {
        const tbody = document.getElementById('adminbooklist'); // select the tbody element
        data.forEach(book => {
            if (book.stock === undefined) {
                book.stock = 1
            }
            const row = document.createElement('tr'); // create a table row
            row.innerHTML = `
            <td>${book.isbn}</td>
            <td>${book.title}</td>
            <td>${book.author}</td>
            <td>${book.language}</td>
            <td>${book.stock}</td>
            <td>
                <div class="btn-group" role="group" aria-label="Button for availability">
                    <button type="button" class="btn btn-success">✔</button>
                    <button type="button" class="btn btn-danger">✘</button>
                </div>
            </td>
        `;
            tbody.appendChild(row); // append the row to the tbody
        });
    } else {
        console.error('Data is not an array:', data);
    }

}


function renderBooks(response) {
    const booksMain = document.getElementById('books_main');
    booksMain.innerHTML = '';
    const data = response.data;

    if (Array.isArray(data)) {
        data.forEach(book => {
            const bookDiv = document.createElement('div');
            bookDiv.className = 'col-md-4';
            bookDiv.innerHTML = `
                <div class="card mb-4 shadow-sm" style="background: #3c0e0e; border-radius: 20px">
                    <a class="#" href="/book/${book.isbn}" id="bookLink"></a>
                    <a href="/book/${book.isbn}">
                        <img class="card-img-top" src="static/images/no-image.svg" alt="Thumbnail"
                             style="height: 220px; width: 100%; display: block; border-radius: 20px">
                    </a>

                    <!-- text -->
                    <div class="card-body">
                        <h5 class="text-center pb-0" style="color: white">${book.title}</h5>
                        <p class="card-text" style="color: white">Available: x</p>
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

function errorMessage(data) {
    const errorheader = document.getElementById('error_header');
    errorheader.innerHTML = '';
    const Div = document.createElement('div');
    Div.className = 'col-md-4';

    // language=HTML
    Div.innerHTML = `
        <div style="width: 1080px" class="alert alert-warning alert-dismissible fade show" role="alert">
            <strong>${data.code}</strong> ${data.message}.
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
            </button>
        </div>
    `;
    errorheader.appendChild(Div);
}

