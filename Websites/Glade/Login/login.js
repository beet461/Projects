var eregex = /[^ ][a-zA-Z.]*[@][a-zA-Z]*[.][a-zA-Z.]*/;
var pregex = /[^ ][\da-zA-Z_.*]*/;
var inptType;
var validation = 0;

async function unblur() {
    window.setTimeout(function () {
        document.getElementById('everything').classList = 'form formatting unblur';
    }, 1000);
}
unblur();

function spremove(valin) {
    if (valin === "" || valin.trim().split(' ').length > 1) {
        return false;
    }
    return true;
}

function check(input, inputTag, psderr) {
    psderr.className = "red-text";

    if (!spremove(input)) {
        inputTag.className = "input is-danger is-rounded";
        psderr.innerHTML = "Error! There seem to be spaces, or an empty field!";
    } else if (!pregex.test(input)) {
        inputTag.className = "input is-danger is-rounded";
        psderr.innerHTML = "Error! There seem to be special characters (remember you can only use _ , . or *!)! Or a value was entered incorrectly";
    } else {
        inputTag.className = "input is-info is-rounded"
        psderr.className = "invisible"
        validation++;
    }
}

function login() {
    var usr = document.getElementById("username");
    var psd = document.getElementById("password");

    var usrerr = document.getElementById("usrerr");
    var psderr = document.getElementById("psderr");

    var input = {
        username: usr.value.trim(),
        password: psd.value.trim(),
    };

    var tags = {
        email: "nil",
        username: usr,
        password: psd,
        usrerr: usrerr,
        psderr: psderr,
        type: "login",
        requestLocation: "localhost:8081/api/v1/login"
    };

    check(input.username, usr, usrerr);
    check(input.password, psd, psderr);

    var raw = JSON.stringify(input)

    if (validation === 2) {
        apiRequest(raw, tags);
    }
} 