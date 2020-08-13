var time = 0;
var time2 = 0;
var af = 0;
var af2 =  0;
var chec = false;

function time10() {
    chec = false;
    console.log(chec);
}

function clear() {
    chec = true;
    console.log(chec);
}

function seven() {
    if (chec === true) {
        time += 7;
        af++;
        document.getElementById("7").innerHTML ="7 x " + af * 10 + " = " + time * 10;
    } else {
        time += 7;
        af++;
        document.getElementById("7").innerHTML ="7 x " + af + " = " + time;
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
