async function unblur() {
    window.setTimeout(function() {
    document.getElementById("everything").className = "form formatting unblur"
    console.log(document.getElementById("everything").className)
    }, 2000)
    
}

unblur()

function register() {
  var email = document.getElementById("email").value
  var usrnm = document.getElementById("usrnm").value
  var pswrd = document.getElementById("pswrd").value
  if (email === "") {
    document.getElementById("emlerr").className = "red-text"
    document.getElementById("email").className = "input is-danger is-rounded"
    document.getElementById("emlerr").innerHTML = "Error! Value is incorrect!"
  } else {
    document.getElementById("email").className = "input is-primary is-rounded"
    document.getElementById("emlerr").innerHTML = "✓"
    document.getElementById("emlerr").className = "green-text"
  }
  
  if (usrnm === "") {
    document.getElementById("usrerr").className = "red-text"
    document.getElementById("usrnm").className = "input is-danger is-rounded"
    document.getElementById("usrerr").innerHTML = "Error! Value is incorrect!"
  } else {
    document.getElementById("usrnm").className = "input is-primary is-rounded"
    document.getElementById("usrerr").innerHTML = "✓"
    document.getElementById("usrerr").className = "green-text"
  }
  
  if (pswrd === "") {
    document.getElementById("psderr").className = "red-text"
    document.getElementById("pswrd").className = "input is-danger is-rounded"
    document.getElementById("psderr").innerHTML = "Error! Value is incorrect!"
  } else {
    document.getElementById("pswrd").className = "input is-primary is-rounded"
    document.getElementById("psderr").innerHTML = "✓"
    document.getElementById("psderr").className = "green-text"
  } 
  console.log(document.getElementById("email").value)
  console.log(document.getElementById("usrnm").value)
  console.log(document.getElementById("pswrd").value)

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

  var raw = JSON.stringify({});

  var requestOptions = {
    method: 'POST',
    headers: myHeaders,
    body: raw,
    redirect: 'follow'
  };

  fetch("http://localhost:8081", requestOptions)
    .then(response => response.text())
    .then(result => console.log(result))
    .catch(error => console.log('error', error));

    window.location.assign("../OtpCheck/index.html")

}

