function test() {
    fetch("http://localhost:8000/api/v1/webpack")
        .then((data) => {console.log(data)})
        .then((response) => {console.log(response)})
}