var eregex = /[a-zA-Z.]*[@][a-zA-Z]*[.][a-zA-Z.]*/;
var uregex = /[^ ][a-zA-Z_.]*/;
var validation = 0

//This part is css animations
async function unblur() {
    window.setTimeout(function () {
        document.getElementById('everything').className =
            'form formatting unblur';
    }, 1000);
}
unblur();

//Api requests
async function apiRequest(input) {
    try {
        var myHeaders = new Headers();
        myHeaders.append('Content-Type', 'application/json');

        var raw = JSON.stringify(input);

        var requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow',
        };

        var request = await fetch(
            'http://localhost:8081/api/v1/register',
            requestOptions
        );
    } catch (err) {
        console.log(err)
    }

    console.log("Request has been sent")
    return request
}

//The next few functions are to do with validation
function spremove(tbtrmd) {
    if (tbtrmd.trim().split(' ').length > 1) {
        return false;
    }
    return true;
}

function emailCheck(email, em, emerr) {
    if (!eregex.test(email)) {
        em.className = 'input is-danger is-rounded';
        emerr.innerHTML = 'Error! Something is incorrect!';
    } else {
        em.className = 'input is-primary is-rounded';
        emerr.className = 'display_none';
        validation++
    }
}

function usernameCheck(usrnm, usr, usrerr) {
    //The first thing this function does, is run the below 'if' statement
    //The below 'if' statement checks if the username value is empty OR it consists of more than one word(s). And if either condition is true, it means the value has an error so an error message is displayed
    if (usrnm === '' || spremove(usrnm)) {
        usr.className = 'input is-danger is-rounded';
        usrerr.innerHTML = 'Error! Something is incorrect!';
    } else {
        //After the above 'if' statement is run and there are no errors so far, another 'if' statement is run
        //This time the value is parsed by a regular expression (regex). The regex checks the value for any special characters or spaces and displays an error if there are any
        if (!uregex.test(usrnm)) {
            usr.className = 'input is-danger is-rounded';
            usrerr.innerHTML = 'Error! Something is incorrect!';
        } else {
            //If there are no errors in the values so far, that means the values are acceptable and can be used
            usr.className = 'input is-primary is-rounded';
            usrerr.className = 'display_none';
            //Lastly the 'validation' variable has 1 added to it for the reason explained below in the register() function
            validation++
        }
    }

}

function passwordCheck(pswrd, psw, psderr) {
    //The first thing this function does, is run the below 'if' statement
    //The below 'if' statement checks if the password value is empty OR it consists of more than one word(s). And if either condition is true, it means the value has an error so an error message is displayed
    if (pswrd === '' || spremove(pswrd)) {
        psw.className = 'input is-danger is-rounded';
        psderr.innerHTML = 'Error! Something is incorrect!';
    } else {
        //After the above 'if' statement is run and there are no errors so far, another 'if' statement is run
        //This time the value is parsed by a regular expression (regex). The regex checks the value for any special characters or spaces and displays an error if there are any
        if (!uregex.test(pswrd)) {
            psw.className = 'input is-danger is-rounded';
            psderr.innerHTML = 'Error! Something is incorrect!';
        } else {
            //If there are no errors in the values so far, that means the values are acceptable and can be used
            psw.className = 'input is-primary is-rounded';
            psderr.className = 'display_none';
            //Lastly the 'validation' variable has 1 added to it for the reason explained below in the register() function
            validation++
        }
    }
}

//This is the main function, which is run when the 'Register' button is pressed on the website
async function register() {
    //These variables are linked to the html input tags in 'register.html', you can add 'attributes', such as ones that change the look or values of the input tags
    var em = document.getElementById('email');
    var usr = document.getElementById('usrnm');
    var psw = document.getElementById('pswrd');

    //These variables are linked to the normally hidden 'p' tags in the html, but are used to display errors in the values, when needed. As it says above you can add attributes to these as well
    var emerr = document.getElementById('emlerr');
    var usrerr = document.getElementById('usrerr');
    var psderr = document.getElementById('psderr');

    //The input values are stored in the 'input' object
    //To access the individual values (like email or username) you can just type in 'input.email' for example
    var input = {
        email: document.getElementById('email').value,
        username: document.getElementById('usrnm').value,
        password: document.getElementById('pswrd').value
    }

    //These functions parse the input values for any inacceptable characters and tell the user if there are 
    emailCheck(input.email, em, emerr);
    //The functions for parsing the username and password are basically the same but they display the error in different places so there needs to be different functions
    usernameCheck(input.username, usr, usrerr);
    passwordCheck(input.password, psw, psderr);

    //The 'if' statement checks if the 'validation' variable's value is equal to 3
    //The 'validation' variable is to check if the three inputs from the user (email, username, password) have been checked and verified to be acceptable
    if (validation === 3) {
        var response = apiRequest(input)
    }
    //Then the variable is set to 0 to not interfere with any future attempts at inputting the right value
    //e.g. If the user has two acceptable value (let's say the email and the username), but the other value is inacceptable, the 'validation' variable would only equal to 2
    //But if the second attempt has three acceptable values then the 'validation' variable's value would be 5, which means a request wouldn't be sent to the api even if all the correct values had been inputted
    validation = 0;

    //The switch statement functions based on the response from the api
    //The '^' at the start means the response is from the register part of the api and the '%' at the start means the response is from the login part of the api
    //'^++' means that the api did not find any existing values in the database that match the input values and has succesfully inserted the new values into the database 
    //'^--' means that the api found existing values in the database that were the same as the input values and so did not insert new values into the database
    //'%++' means that the api found matches for the input values and then the javascript here will log the user into the website
    //'%--' means that the api hasn't found any matches for the input values, which means the user either needs to create a new account or the values were inputted incorrectly and need to be re-typed
    switch (response) {
        case "^++":
            inserted(input)
            break;
        case "^--":
            ntInserted(input)
            break;
        case "%++":
            match(input)
            break;
        case "%--":
            ntMatch(input)
    }
}
