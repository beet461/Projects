
function sendmsg() {
    let x = document.getElementById("message").value
    document.addEventListener('astilectron-ready', function() {
        astilectron.sendMessage(x, function(message) {
            console.log("received " + message)
        });
        if (message === "world")  {
            console.log(420)
        }
    })
    console.log(69)
}
