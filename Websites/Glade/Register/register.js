//These variables are used to validate the input values
//This regular expression is for the email, and when used to parse the input values, it will only return a match if the input looks something like a regular email 
var eregex = /[^ ][a-zA-Z.]*[@][a-zA-Z]*[.][a-zA-Z.]*/
var uregex = /[^ ][a-zA-Z_.*]*/
var validation = 0

//This part is css animations
async function unblur() {
    window.setTimeout(function () {
        document.getElementById('everything').classList = 'form formatting unblur';
    }, 1000);
}
unblur();


//The next few functions are to do with validation
//This function is used to check if the values it is given are empty or there are spaces
function spremove(input) {
    //The '.trim().split(' ').length' removes the spaces at the start and end of the value and splits it into seperate words every time there is a space
    //Then the '.length' measures how many individual words there are and, if there are more than 1 it returns false  
    if (input === "" || input.trim().split(' ').length > 1) {
        return false;
    }
    //If there is nothing wrong then this function returns true
    return true;
}

function check(input, inputTag, psderr, regex) {
    psderr.className = "red-text";

    if (!spremove(input)) {
        inputTag.className = "input is-danger is-rounded";
        psderr.innerHTML = "Error! There seem to be spaces, or an empty field!";
    } else if (!regex.test(input)) {
        inputTag.className = "input is-danger is-rounded";
        psderr.innerHTML = "Error! There seem to be special characters (remember you can only use _ , . or *!)! Or a value was entered incorrectly";
    } else {
        inputTag.className = "input is-info is-rounded"
        psderr.className = "invisible"
        validation++;
    }
}

//This is the main function, which is run when the 'Register' button is pressed on the website
async function register() {
    //These variables are linked to the html input tags in 'register.html', you can add 'attributes', such as ones that change the look or values of the input tags
    var em = document.getElementById('email');
    var usr = document.getElementById('usrnm');
    var psw = document.getElementById('pswrd');

    //These variables are linked to the normally hidden 'p' tags in the html, but are used to display errors in the values, when needed. As it says above you can add attributes to these as well
    var emerr = document.getElementById('emlerr');
    var usrerr = document.getElementById('usrerr');
    var psderr = document.getElementById('psderr');

    //The input values are stored in the 'input' object
    //To access the individual values (like email or username) you can just type in 'input.email' for example
    var input = {
        email: em.value.trim(),
        username: usr.value.trim(),
        password: psw.value.trim(),
        type: "register"
    }

    var tags = {
        email: em,
        username: usr,
        password: psw,
        emerr: emerr,
        usrerr: usrerr,
        psderr: psderr,
        type: "register"
    }

    check(input.email, em, psderr, eregex)
    check(input.username, usr, psderr, uregex)
    check(input.password, psw, psderr, uregex)

    //The 'if' statement checks if the 'validation' variable's value is equal to 3
    //The 'validation' variable is to check if the three inputs from the user (email, username, password) have been checked and verified to be acceptable
    //If everuthing has been checked then a request to the api, along with the data, is sent
    //Then the 'validation' variable is set to 0 to not interfere with any future attempts at inputting the right value
    //e.g. If the user has two acceptable value (let's say the email and the username), but the other value is inacceptable, the 'validation' variable would only equal to 2
    //But if the second attempt has three acceptable values then the 'validation' variable's value would be 5, which means a request wouldn't be sent to the api even if all the correct values had been input
    if (validation === 3) {
        apiRequest(input, tags)
        validation = 0;
    }
}
