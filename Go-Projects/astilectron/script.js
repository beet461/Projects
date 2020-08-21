function sendmsg() {
    astilectron.sendMessage("hello", function(message) {
        console.log("received " + message)
    });
    document.addEventListener('astilectron-ready', function() {
        // This will send a message to GO
        
        astilectron.onMessage(function(message) {
            // Process message
            if (message === "world") {
                return "world";
            }
        });
    })
}

sendmsg()