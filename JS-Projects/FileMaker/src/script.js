function register() {
  console.log(document.getElementById("email").value)
  console.log(document.getElementById("usrnm").value)
  console.log(document.getElementById("pswrd").value)
}

var myHeaders = new Headers();
myHeaders.append("Access-Control-Allow-Origin", "*")
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({"username":"heeeeeeellllllllo","password":"heeeeeeelllllllloo"});

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
