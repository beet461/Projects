var eregex = /[^ ][a-zA-Z.]*[@][a-zA-Z]*[.][a-zA-Z.]*/;
var pregex = /[^ ][a-zA-Z_.-*]*/;
var email
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

function emailusernameCheck(input, emusr, lusrerr) {
    if (!spremove(input.emailusername)) {
        emusr.className = "input is-danger is-rounded transparent"
        lusrerr.innerHtml = "Error! You seemed to have spaces present!"
    } else if (!eregex.test(input.emailusername) || !pregex.test(input.emailusername)) {
        emusr.className = "input is-danger is-rounded transparent"
        lusrerr.innerHtml = "Error! There seem to be special characters!"
    } else {
        if (eregex.text(input.emailusername)) {
            email = true
        } else if (pregex.test(input.emailusername)) {
            email = false
            validation++
        }
    }
}

function passwordCheck(input, psd, lpsderr) {
    if (!spremove(input.password)) {
        psd.className = "input is-danger is-rounded transparent"
        lpsderr.innerHtml = "Error! There seems to be spaces present!"
    } else if (!pregex.test(input.password)) {
        emusr.className = "input is-danger is-rounded transparent"
        lpsderr.innerHtml = "Error! There seem to be special characters! (Remember you can only use _ , . , - or *!)"
    } else {
        validation++
    }
}

function login() {
    var emusr = document.getElementById("username/email")
    var psd = document.getElementById("password")

    var lusrerr = document.getElementById("lusrerr")
    var lpsderr = document.getElementById("lpsderr")

    var input = {
        emailusername: emusr.value.trim(),
        password: psd.value.trim(),
        lusrerr: lusrerr,
        lpsderr: lpsderr
    }

    emailusernameCheck(input, emusr, lusrerr)
}