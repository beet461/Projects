//These variables are used to validate the input values
//This regular expression is for the email, and when used to parse the input values, it will only return a match if the input looks something like a regular email 
var eregex = /[^ ][a-zA-Z.]*[@][a-zA-Z]*[.][a-zA-Z.]*/;
var uregex = /[^ ][a-zA-Z_.]*/;
var validation = 0

//This part is css animations
async function unblur() {
    window.setTimeout(function () {
        document.getElementById('everything').className =
            'form formatting unblur';
    }, 1000);
}
unblur();


//The next few functions are to do with validation
//This function is used to check if the values it is given are empty or there are spaces
function spremove(valin) {
    //The '.trim().split(' ').length' removes the spaces at the start and end of the value and splits it into seperate words every time there is a space
    //Then the '.length' measures how many individual words there are and, if there are more than 1 it returns false  
    if (valin === "" || valin.trim().split(' ').length > 1) {
        return false;
    }
    //If there is nothing wrong then this function returns true
    return true;
}

function emailCheck(email, em, emerr) {
    //The first thing this function does is make sure that the error message can be displayed, because if the values are correct then the 'p' tag where the error message is displayed is forced to be invisible
    emerr.className = "red-text"
    //The next thing this function does, is run the below 'if' statement
    //The below 'if' statement checks if the email value is empty OR it consists of more than one word(s) by running the spremove function. And if either condition is true, it means the value has an error so an error message is displayed
    if (!spremove(email)) {
        em.className = 'input is-danger is-rounded'
        emerr.innerHTML = "Error! Something is incorrect"
    } else {
        //After the above 'if' statement is run and there are no errors so far, another 'if' statement is run
        //This time the value is parsed by a regular expression (regex). The regex checks the value for any special characters or spaces and displays an error if there are any
        if (!eregex.test(email)) {
            em.className = 'input is-danger is-rounded';
            emerr.innerHTML = 'Error! Something is incorrect!';
        } else {
            //Lastly the 'validation' variable has 1 added to it for the reason explained below in the register() function
            validation++
        }
    }
}

function usernameCheck(usrnm, usr, usrerr) {
    //The first thing this function does is make sure that the error message can be displayed, because if the values are correct then the 'p' tag where the error message is displayed is forced to be invisible
    usrerr.className = "red-text"
    //The next thing this function does, is run the below 'if' statement
    //The below 'if' statement checks if the username value is empty OR it consists of more than one word(s) by running the spremove function. And if either condition is true, it means the value has an error so an error message is displayed
    if (!spremove(usrnm)) {
        usr.className = 'input is-danger is-rounded';
        usrerr.innerHTML = 'Error! Something is incorrect!';
    } else {
        //After the above 'if' statement is run and there are no errors so far, another 'if' statement is run
        //This time the value is parsed by a regular expression (regex). The regex checks the value for any special characters or spaces and displays an error if there are any
        if (!uregex.test(usrnm)) {
            usr.className = 'input is-danger is-rounded';
            usrerr.innerHTML = 'Error! Something is incorrect!';
        } else {
            //Lastly the 'validation' variable has 1 added to it for the reason explained below in the register() function
            validation++
        }
    }

}

function passwordCheck(pswrd, psw, psderr) {
    //The first thing this function does is make sure that the error message can be displayed, because if the values are correct then the 'p' tag where the error message is displayed is forced to be invisible
    psderr.className = "red-text"
    //The next thing this function does, is run the below 'if' statement
    //The below 'if' statement checks if the password value is empty OR it consists of more than one word(s) by running the spremove function. And if either condition is true, it means the value has an error so an error message is displayed
    if (!spremove(pswrd)) {
        psw.className = 'input is-danger is-rounded';
        psderr.innerHTML = 'Error! Something is incorrect!';
    } else {
        //After the above 'if' statement is run and there are no errors so far, another 'if' statement is run
        //This time the value is parsed by a regular expression (regex). The regex checks the value for any special characters or spaces and displays an error if there are any
        if (!uregex.test(pswrd)) {
            psw.className = 'input is-danger is-rounded';
            psderr.innerHTML = 'Error! Something is incorrect!';
        } else {
            //Lastly the 'validation' variable has 1 added to it for the reason explained below in the register() function
            validation++
        }
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
        email: document.getElementById('email').value,
        username: document.getElementById('usrnm').value,
        password: document.getElementById('pswrd').value
    }

    var tags = {
        email: em,
        username: usr,
        password: psw,
        emerr: emerr,
        usrerr: usrerr,
        psderr: psderr,
        type: "reg"
    }

    //These functions parse the input values for any inacceptable characters and tell the user if there are any 
    emailCheck(input.email, em, emerr);
    //The functions for parsing the username and password are basically the same but they display the error in different places so there needs to be different functions
    usernameCheck(input.username, usr, usrerr);
    passwordCheck(input.password, psw, psderr);

    //The 'if' statement checks if the 'validation' variable's value is equal to 3
    //The 'validation' variable is to check if the three inputs from the user (email, username, password) have been checked and verified to be acceptable
    //If everuthing has been checked then a request to the api, along with the data, is sent
    if (validation === 3) {
        apiRequest(input, tags)
    }
    //Then the variable is set to 0 to not interfere with any future attempts at inputting the right value
    //e.g. If the user has two acceptable value (let's say the email and the username), but the other value is inacceptable, the 'validation' variable would only equal to 2
    //But if the second attempt has three acceptable values then the 'validation' variable's value would be 5, which means a request wouldn't be sent to the api even if all the correct values had been inputted
    validation = 0;



}
//Function for tellign user values already exist
//Figure out how to find out which value already exists
