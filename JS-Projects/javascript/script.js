var time = 0;
var time2 = 0;
var af = 0;
var af2 =  0;
var chec = false;


function time10() {
    chec = true;
    console.log(chec);
}

function reset() {
    chec = false;
    console.log(chec);
}

function seven() {
    var h = document.getElementById("7").nodeValue;
    console.log(h)
    if (h !== undefined) {
    if (chec === true) {
        time += 7;
        af++;
        document.getElementById("7").innerHTML = h * 10;
    } else {
        time += 7;
        af++;
        document.getElementById("7").innerHTML ="7 x " + af + " = " + time;
    }      
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
