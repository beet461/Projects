let x = document.getElementById("message").value

document.addEventListener('astilectron-ready', function() {
    astilectron.sendMessage(x, function(message) {
        console.log("received " + message)
    });
})