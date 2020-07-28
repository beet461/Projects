var calculation = "";
document.getElementById("text-element").innerHTML = "Calculation: " + calculation;
document.getElementById("result").innerHTML = "Result: ";

function append(value) {
    calculation = calculation + value;
    document.getElementById("text-element").innerHTML = "Calculation: " + calculation;
}

function delall() {
    calculation = "";
    document.getElementById("text-element").innerHTML = "Calculation: " + calculation;
    document.getElementById("result").innerHTML = "Result: ";
}

function dellast() {
    var arrayLetters = calculation.trim().split('');
    arrayLetters.pop();
    var result = "";

    for (var i = 0; i < arrayLetters.length; i++) {
        result = result + arrayLetters[i];
    }

    calculation = result;
    document.getElementById("text-element").innerHTML = "Calculation: " + calculation;
}

function calculate() {
    var isErr = false; 
    var shouldICont = true;
    try {
        var endResult = eval(calculation);
    } catch (er) {
        isErr = true;
        console.log(er);
    } finally {
        if (isErr) {
            shouldICont = false;
        }
    }
    if (shouldICont) {
        document.getElementById("result").innerHTML = "Result: " + eval(calculation);
    }
}