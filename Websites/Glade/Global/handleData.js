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

function storedata(incData) {
    localStorage.setItem("glade.email", incData.email)
    localStorage.setItem("glade.username", incData.username)
    localStorage.setItem("glade.password", incData.password)
}