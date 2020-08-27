var data = {
    "title" : "title",
    "body" : "body",
}

var params = {
    headers:{
        "content-type" : "application/json"
    },
    body:data,
    method:"POST"
}

function makego() {
    var myHeaders = new Headers();
myHeaders.append("Content-Type", "application/json");

var raw = JSON.stringify({"date":"09/1293/01","body":"idk"});

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
}


