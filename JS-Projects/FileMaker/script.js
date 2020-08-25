function test() {
    try {
        fetch("http://localhost:8080/api/v1/webpack")
            .then((data) => {console.log(data)})
            .then((response) => {console.log(response.message)})
    } catch (err) {
        console.log(err)
    }
}


