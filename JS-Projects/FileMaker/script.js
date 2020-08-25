function test() {
    try {
        fetch("192.168.0.28:8080")
            .then((data) => {console.log(data)})
            .then((response) => {console.log(response.message)})
    } catch (err) {
        console.log(err)
    }
}


