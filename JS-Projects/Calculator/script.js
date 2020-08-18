function Prompt() {
    var x = prompt("Calculation")
    var y = eval(x)
    document.getElementById("eval").innerHTML = y;
}