




function sendmsg() {
    document.addEventListener('astilectron-ready', function() {
        // This will send a message to GO
        astilectron.sendMessage("hello", function(message) {
            console.log("received " + message)
        });
    })
}

