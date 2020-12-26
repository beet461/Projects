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

function check(input, inputTag, psderr, tag) {
    psderr.className = "red-text";

    if (!spremove(input)) {
        var err = true;
        inputTag.className = "input is-danger is-rounded";
        psderr.innerHTML = "Error! There seem to be spaces, or an empty field!";
    } else if (!eregex.test(input) || !pregex.test(input)) {
        var match = false
        var err = true
        inputTag.className = "input is-danger is-rounded";
        psderr.innerHTML = "Error! There seem to be special characters (remember you can only use _ , . or *!)! Or a value was entered incorrectly";
    } else {
        if (tag === "email/user" && match) {
            if (eregex.test(input)) {
                inptType = "email";
            } else if (pregex.test(input)) {
                inptType = "username"
            }
        } else if (tag === "password" && !err) {
            inputTag.className = "input is-info is-rounded"
            psderr.className = "invisible"
            validation++;
        }
    }
}

function login() {
    var usr = document.getElementById("username/email");
    var psd = document.getElementById("password");

    var usrerr = document.getElementById("usrerr");
    var psderr = document.getElementById("psderr");

    var input = {
        username: usr.value.trim(),
        password: psd.value.trim(),
        logType: inptType
    };

    var tags = {
        email: "nil",
        username: usr,
        password: psd,
        usrerr: usrerr,
        psderr: psderr,
        inptType: inptType,
        type: "login",
        requestLocation: "localhost:8081/api/v1/login"
    };

    check(input.username, usr, usrerr, "email/user");
    check(input.password, psd, psderr, "password");

    var raw = JSON.stringify(input)

    if (validation === 2) {
        apiRequest(raw, tags);
    }
} 