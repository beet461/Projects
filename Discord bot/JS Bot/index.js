const Discord = require("discord.js");
const client = new Discord.Client();
let splitmessage;


client.on('ready', () => {
    console.log("Bot has **beep** connected!");
    
});

client.on("message", message => {
    let prblmnum = 1;
    splitmessage = message.content.trim().split(' ')
    try {
        if (splitmessage[0] === "!calc") {
            message.channel.send(eval(splitmessage[1]))
            console.log(`Problem No.${prblmnum}'s answer is ` + eval(splitmessage[1]) + ".")
            prblmnum++
        }
        if (splitmessage[0] === "!help") {
            message.channel.send("Hello *Beep Boop* I am Le Bot. I have many functions *Beep* \n 1)Calculations - I offer *Boop* a calculator! Just type in !calc then your question, e.g. !calc 8+8. Remember that there should be 0 spaces in your equation. *Beep Boop* \n 2)Help - If yoou want help on what I can *Beep* do, just type !help")
        }
        if (splitmessage[0] === "hello") {
            message.channel.send("Hello!")
        } else if (splitmessage[0] === "hi") {
            message.channel.send("Hi!")
        }
    } catch(err) {
        console.log("Oops. *Beep* There seems to be a typo! Aborting *Beep* the current process!");
    } 
})

client.login("NzMxNTI0NDgzNDQxMDk4ODI0.XwnTQw.qhEHak_6tYfogaGCi2CL94iL4jo")