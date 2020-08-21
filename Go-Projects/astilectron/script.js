

astilectron.sendMessage("hello", function(message) {
    console.log("received " + message)
});
astilectron.onMessage(function(message) {
    // Process message
    if (message === "world") {
        return "world";
    }
});
