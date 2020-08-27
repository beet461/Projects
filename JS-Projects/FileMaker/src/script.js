var data = {
    "title" : "title",
    "body" : "body",
}

var params = {
    headers:{
        "content-type" : "application/json"
    },
    body:JSON.stringify(data),
    method:"POST"
}

function makego() {
    try {
        fetch("192.168.0.28:8080/api/v1/makego", params)
            .then((data) => {console.log(data)})
            .then((response) => {console.log("this is the response")})
    } catch (err) {
        console.log(err)
    }
}


