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

axios.post('localhost:8080/api/v1/makego', {
    firstName: 'Fred',
    lastName: 'Flintstone'
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
