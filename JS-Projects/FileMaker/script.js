function test() {
    fetch("http://localhost:8080/api/v1/webpack")
        .then((data) => {console.log(data)})
        .then((response) => {console.log(response)})
}

try {
    console.log("try");
} catch (err) {
    console.log(err)
}
