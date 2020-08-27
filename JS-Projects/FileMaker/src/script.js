/*var data = {
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
*/

function makego() {
   
}

var data = JSON.stringify({"date":"09/1293/01","body":"idk"});

var config = {
  method: 'post',
  url: 'http://localhost:8080/api/v1/makego',
  headers: { 
    'Content-Type': 'application/json'
  },
  data : data
};

axios(config)
.then(function (response) {
  console.log(JSON.stringify(response.data));
})
.catch(function (error) {
  console.log(error);
});


