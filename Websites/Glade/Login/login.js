var eregex = /[^ ][a-zA-Z.]*[@][a-zA-Z]*[.][a-zA-Z.]*/;
var pregex = /[^ ][a-zA-Z_.-*]*/;
var inptType
var validation = 0

async function unblur() {
    window.setTimeout(function () {
        document.getElementById('everything').className =
            'form formatting unblur';
    }, 1000);
}
unblur();

function spremove(valin) {
    if (valin === "" || valin.trim.split(' ').length > 1) {
        return false;
    }
    return true;
}

function usernameCheck(input, emusr, usrerr) {
    if (!spremove(input.username)) {
        emusr.className = "input is-danger is-rounded transparent"
        usrerr.innerHtml = "Error! You seemed to have spaces present!"
    } else if (!eregex.test(input.username) || !pregex.test(input.username)) {
        emusr.className = "input is-danger is-rounded transparent"
        usrerr.innerHtml = "Error! There seem to be special characters!"
    } else {
        if (eregex.text(input.username)) {
            inptType = "email"
        } else if (pregex.test(input.username)) {
            inptType = "username"
            validation++
        }
    }
}

function passwordCheck(input, psd, psderr) {
    if (!spremove(input.password)) {
        psd.className = "input is-danger is-rounded transparent"
        psderr.innerHtml = "Error! There seems to be spaces present!"
    } else if (!pregex.test(input.password)) {
        emusr.className = "input is-danger is-rounded transparent"
        psderr.innerHtml = "Error! There seem to be special characters! (Remember you can only use _ , . , - or *!)"
    } else {
        validation++
    }
}

function login() {
    var emusr = document.getElementById("username/email")
    var psd = document.getElementById("password")

    var usrerr = document.getElementById("usrerr")
    var psderr = document.getElementById("psderr")

    var input = {
        email: "nil",
        username: emusr.value.trim(),
        password: psd.value.trim(),
        usrerr: usrerr,
        psderr: psderr
    }

    var tags = {
        emusr: emusr,
        psd: psd,
        usrerr: usrerr,
        psderr: psderr,
        inptType: inptType,
        type: "log"
    }

    usernameCheck(input, emusr, usrerr)
    passwordCheck(inputs, psd, psderr)

    if (validation === 2) {
        apiRequest(inputs, tags)
    }
} 