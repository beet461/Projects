//Global data for use in navbar and elsewhere
var gData = {
    email: localStorage.getItem("glade.email"),
    username: localStorage.getItem("glade.username"),
    password: localStorage.getItem("glade.password")
}

function inserted(input) {
    storedata(input)
}

function ntInserted(input) {

}

function match() {

}

function ntMatch() {

}

function storedata(tbStored) {
    localStorage.setItem("glade.email", tbStored.email)
    localStorage.setItem("glade.username", tbStored.username)
    localStorage.setItem("glade.password", tbStored.password)
}