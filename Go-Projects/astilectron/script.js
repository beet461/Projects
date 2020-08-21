




function sendmsg() {
    document.addEventListener('astilectron-ready', function() {
        // This will send a message to GO
        astilectron.sendMessage("hello", function(message) {
            console.log("received " + message)
        });
    })
    document.addEventListener('astilectron-ready', function() {
        // This will listen to messages sent by GO
        astilectron.onMessage(function(message) {
            // Process message
            if (message === "hello") {
                return "world";
            }
        });
    })
}

