
function sendmsg() {
    let x = document.getElementById("message").value
    document.addEventListener('astilectron-ready', function() {
        astilectron.sendMessage(x, function(x) {
            console.log("received " + x)
            if (message === "world")  {
                console.log(420)
            }
        });
        
    })
    console.log(69)
}
