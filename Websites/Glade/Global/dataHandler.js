//Global data for use in navbar and elsewhere
var gData = {
    email: localStorage.getItem("glade.email"),
    username: localStorage.getItem("glade.username"),
    password: localStorage.getItem("glade.password")
}

async function apiRequest(raw, tags) {
    //The below method is called a 'try and catch' method. You try some code and if there are errors the 'catch' past catches it in the variable 'err' 
    try {
        //These set the headers
        var myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/json');

        //Here you can set options for the request, like what method it is and what data to send
        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow',
        };

        //This part is where the request is actually sent and the response is stored it in the variable 'response'
        var response = await fetch(
            tags.requestLocation, requestOptions
        );
    } catch (err) {
        //If there is an error, then it logged to the console to see what exactly went wrong
        console.log(err)
    }
    //The switch statement runs functions based on the response from the api - the four functions below are defined in handleData.js, not here 
    //The '^' at the start means the response is from the register part of the api and the '%' at the start means the response is from the login part of the api
    //'^++' means that the api did not find any existing values in the database that match the input values and has succesfully inserted the new values into the database 
    //'^--' means that the api found existing values in the database that were the same as the input values and so did not insert new values into the database
    //'%++' means that the api found matches for the input values and then the javascript here will log the user into the website
    //'%--' means that the api hasn't found any matches for the input values, which means the user either needs to create a new account or the values were inputted incorrectly and need to be re-typed
    switch (response) {
        case "^++":
            displayCorrect(tags)
            break;
        case "^--":
            displayErr(tags)
            break;
        case "^+-":
            break;
        case "%++":
            displayCorrect(tags)
            break;
        case "%--":
            displayErr(tags)
    }
}

function displayErr(tags) {
    var wrong = "input is-danger is-rounded"

    tags.emerr.classname = "red-text"
    tags.usrerr.classname = "red-text"
    tags.psderr.classname = "red-text"

    if (tags.type === "reg") {
        var errMessage = "Error! This username already exists!"
    } else if (tags.type === "log") {
        errMessage = `Error! The ${tags.inptType} or password do not match any existing accounts!`
    }

    if (tags.email !== "nil") {
        tags.email.classname = wrong
    } else if (tags.username !== "nil") {
        tags.username.classname = wrong
    } else if (tags.password !== "nil") {
        tags.password.classname = wrong
        tags.psderr.innerHTML = errMessage
    }

}

function displayWarn(tags) {
    var warn = "input is-warning is-rounded"

    if (tags.type === "reg") {
        var warnMessage = "Warning this password already exists! You might want to change it!"
    }

}

function displayCorrect(tags) {
    var correct = "input is-primary is-rounded"

    if (tags.email !== "nil") {
        tags.email.classname = correct
        tags.emerr.className = "invisible"
    } else if (tags.username !== "nil") {
        tags.username.classname = correct
        tags.usrerr.classname = "invisible"
    } else if (tags.password !== "nil") {
        tags.password.classname = correct
        tags.psderr.classname = "invisible"
    }

}

function inserted(input, tags) {
    storedata(input)
    displayCorrect(tags)

}

function ntInserted(input, tags) {
    displayErr(tags)
}

function match(input, tags) {

}

function ntMatch(input, tags) {

}

function storedata(tbStored) {
    localStorage.setItem("glade.email", tbStored.email)
    localStorage.setItem("glade.username", tbStored.username)
    localStorage.setItem("glade.password", tbStored.password)
}
