var time = 0;
var time2 = 0;
var af = 0;
var af2 =  0;
var chec = false;
var a = 1;
var b = 0;


function time10() {
    chec = true;
    console.log(chec);
}

function reset() {
    chec = false;
    console.log(chec); 
}

function sev() {
    var sev = "7 x " + f10 + " = " + t10;
    document.getElementById("7").innerHTML = sev;
}

function seven() {
    b = af + a
    time += 7
    af ++
    var h = document.getElementById("7").nodeValue;
    console.log(h)
    if (chec === true) {
        f10 = af * 10
        t10 = time * 10
        document.getElementById("7").innerHTML ="7 x " + af * 10 + " = " + time * 10;
    } else {
        sev()
    }      
 
}

function eleven() {
    if (chec === true) {
        time2 += 11;
        af2++
        document.getElementById("11").innerHTML ="11 x " + af2 * 10+ " = " + time2 * 10;
    } else {
        time2 += 11;
        af2++
        document.getElementById("11").innerHTML ="11 x " + af2 + " = " + time2;
    }
}
