const Discord = require("discord.js");
const client = new Discord.Client();
let splitmessage;
let sDate;
let date;
date = Date();
sDate = date.trim().split(' ');
var eDate = "The *Beep* date : " + sDate[0] + " " + sDate[1] + " " + sDate[2] + " " + sDate[3] + "!";
var eTime = "And the *Boop* time : " + sDate[4] + "!";

client.on('ready', () => {
    console.log("Bot has **beep** connected!");

});

client.on("message", message => {
    let prblmnum = 1;
    splitmessage = message.content.trim().split(' ');
    try {
        if (message.author.bot) {
            var bot = true
        }
        if (splitmessage[0] === "!calc" && !bot) {
            message.channel.send(eval(splitmessage[1]))
            console.log(`Problem No.${prblmnum}'s answer is ` + eval(splitmessage[1]) + ".")
            prblmnum++
        }

        if (splitmessage[0] === "!help" && !bot) {
            message.channel.send("Hello *Beep Boop* I am Le Bot. I have many functions *Beep* \n")
            message.channel.send("1)Calculator - I offer *Boop* a calculator! Just type in !calc then your question, e.g. !calc 8+8. Remember that there should be 0 spaces in your equation. *Beep* \n")
            message.channel.send("2)If you want to know the *Boop* time and date, just type !date and you will get the time!")
            message.channel.send("And remember, if you ever need help just type '!help' for a list of all my functions! Bye! *Beep Boop Beep*")
        }

        if (splitmessage[0] === "hello" || splitmessage[0] == "Hello" && !bot) {
            message.channel.send("Hello!")
        } else if (splitmessage[0] === "hi" || splitmessage[0] === "Hi" && !bot) {
            message.channel.send("Hi!")
        }

        if (splitmessage[0] === "!date" && !bot) {
            message.channel.send(eDate);
            message.channel.send(eTime);
        }

        if (splitmessage[0] === "hello" && splitmessage[1] === "there") {
            message.channel.send("General Kenobi")
        }

        if (splitmessage[0] === "!newrank") {
            message.channel.send("Ok! A new rank has been created for you!")

        }
    } catch (err) {
        console.log("Oops. *Beep* There seems to be a typo! Aborting *Boop* the current process! Error : " + err);
    }
})

client.login("NzMxNTI0NDgzNDQxMDk4ODI0.XwnTQw.67kYj3jw2wWoMjGNV9ntzd0vYBI")