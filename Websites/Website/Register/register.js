var eregex = /[a-zA-Z.]*[@][a-zA-Z]*[.][a-zA-Z.]*/;
var uregex = /[^ ][a-zA-Z_.]*/;

async function unblur() {
    window.setTimeout(function () {
        document.getElementById('everything').className =
            'form formatting unblur';
        console.log(document.getElementById('everything').className);
    }, 1000);
}
unblur();

function spremove(input) {
    if (input.trim().split(' ').length > 1) {
        return false;
    }
    return true;
}

function emailCheck(email, em, emerr) {
    if (!eregex.test(email)) {
        em.className = 'input is-danger is-rounded';
        emerr.innerHTML = 'Error! Value is incorrect!';
    } else {
        em.className = 'input is-primary is-rounded';
        document.getElementById('emlerr').className = 'display_none';
    }
}

function usernameCheck(usrnm, usr, usrerr) {
    if (usrnm === '' || spremove(usrnm)) {
        usr.className = 'input is-danger is-rounded';
        usrerr.innerHTML = 'Error! Value is incorrect!';

        if (!uregex.test(usrnm)) {
            usr.className = 'input is-danger is-rounded';
            usrerr.innerHTML = 'Error! Value is incorrect!';
        } else {
            usr.className = 'input is-primary is-rounded';
            usrerr.className = 'display_none';
        }
    }
}

function passwordCheck(pswrd, psw, psderr) {
    if (pswrd === '' || spremove(pswrd)) {
        psw.className = 'input is-danger is-rounded';
        psderr.innerHTML = 'Error! Value is incorrect!';

        if (!uregex.test(pswrd)) {
            psw.className = 'input is-danger is-rounded';
            psderr.innerHTML = 'Error! Value is incorrect!';
        } else {
            psw.className = 'input is-primary is-rounded';
            psderr.className = 'display_none';
        }
    }
}

async function register() {
    var em = document.getElementById('email');
    var usr = document.getElementById('usrnm');
    var psw = document.getElementById('pswrd');

    var emerr = document.getElementById('emlerr');
    var usrerr = document.getElementById('usrerr');
    var psderr = document.getElementById('psderr');

    var email = document.getElementById('email').value;
    var usrnm = document.getElementById('usrnm').value;
    var pswrd = document.getElementById('pswrd').value;

    emailCheck(email, em, emerr);

    usernameCheck(usrnm, usr, usrerr);

    passwordCheck(pswrd, psw, psderr);

    console.log(email);
    console.log(usrnm);
    console.log(pswrd);

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
}
