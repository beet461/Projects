function register() {
  console.log(document.getElementById("email").value)
  console.log(document.getElementById("usrnm").value)
  console.log(document.getElementById("pswrd").value)
}

var email = document.getElementById("email").value
var usrnm = document.getElementById("usrnm").value
var pswrd = document.getElementById("pswrd").value

var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

var input = {
  Email : email,
  Username : usrnm,
  Password : pswrd
}

var raw = JSON.stringify({input});

var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: raw,
  redirect: 'follow'
};

fetch("http://localhost:8080/api/v1/makego", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));
