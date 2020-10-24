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

async function apiRequest(email, usrnm, pswrd) {
    try {
        var input = {
            Email: email.trim(),
            Username: usrnm.trim(),
            Password: pswrd.trim(),
        };
    
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
    
    
}

//The next few functions are to do with validation
function spremove(input) {
    if (input.trim().split(' ').length > 1) {
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
    if (usrnm === '' || spremove(usrnm)) {
        usr.className = 'input is-danger is-rounded';
        usrerr.innerHTML = 'Error! Something is incorrect!';
        if (!uregex.test(usrnm)) {
            usr.className = 'input is-danger is-rounded';
            usrerr.innerHTML = 'Error! Something is incorrect!';
        } else {
            usr.className = 'input is-primary is-rounded';
            usrerr.className = 'display_none';
            validation++
        }
    }
}

function passwordCheck(pswrd, psw, psderr) {
    if (pswrd === '' || spremove(pswrd)) {
        psw.className = 'input is-danger is-rounded';
        psderr.innerHTML = 'Error! Something is incorrect!';

        if (!uregex.test(pswrd)) {
            psw.className = 'input is-danger is-rounded';
            psderr.innerHTML = 'Error! Something is incorrect!';
        } else {
            psw.className = 'input is-primary is-rounded';
            psderr.className = 'display_none';
            validation++
        }
    }
}

//The main function, which is run when the 'Register' button is pressed on the website
async function register() {
    //Variables from which the data is taken and used to display any errors
    var em = document.getElementById('email');
    var usr = document.getElementById('usrnm');
    var psw = document.getElementById('pswrd');

    var emerr = document.getElementById('emlerr');
    var usrerr = document.getElementById('usrerr');
    var psderr = document.getElementById('psderr');

    var email = document.getElementById('email').value;
    var usrnm = document.getElementById('usrnm').value;
    var pswrd = document.getElementById('pswrd').value;

    //Validating input
    emailCheck(email, em, emerr);
    usernameCheck(usrnm, usr, usrerr);
    passwordCheck(pswrd, psw, psderr);

    //Logging the data
    console.log("Email : ", email);
    console.log("Username : ", usrnm);
    console.log("Password", pswrd);

    //Checking if data has been checked and is fine then sending a request to the api
    if (validation === 3) {
        apiRequest(email, usrnm, pswrd)
    }
}
