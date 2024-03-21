

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

    console.log(formData)

    const jsonData = {};
    formData.forEach((value, key) => {
        jsonData[key] = value;
    });

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
            return response.json();
        })
        .then(data => {
            console.log("Request successful:", data);
        })
        .catch(error => {
            console.error("Request failed:", error);
        });
}


function fetchMovies() {
    fetch("http://127.0.0.1:8080/api/v1/home/books")
        .then(response => response.json())
        .then(data => {
            // Call a function to render the movies data on the page
            renderMovies(data);
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

// {
//     "movies": [
//         {
//             "image": "images/img-3.png",
//             "title": "CADE Prlor",
//             "content": "There are many variations",
//             "stars": 5
//         }
//     ]
// }
function renderMovies(data) {
    // Get the movies_main element
    const moviesMain = document.getElementById('books_main');

    // Clear any existing movie elements
    moviesMain.innerHTML = '';

    // Loop through the movies data and create HTML elements for each movie
    data.forEach(movie => {
        const movieDiv = document.createElement('div');
        movieDiv.className = 'iamge_movies_main';

        const movieContent = `
                    <div class="iamge_movies">
                        <div class="image_3">
                            <img src="${movie.image}" class="image" style="width:100%">
                            <div class="middle">
                                <div class="playnow_bt">Play Now</div>
                            </div>
                        </div>
                        <h1 class="code_text">${movie.title}</h1>
                        <p class="there_text">${movie.content}</p>
                        <div class="star_icon">
                            <ul>
                                ${'*'.repeat(movie.stars)}
                            </ul>
                        </div>
                    </div>
                `;
        movieDiv.innerHTML = movieContent;

        // Append the movie element to movies_main
        moviesMain.appendChild(movieDiv);
    });
}
