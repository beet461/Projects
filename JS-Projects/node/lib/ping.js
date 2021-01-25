export function ping(thing) {
    if (thing === "pinglog") {
        pinglog()
    } else {
        pingtwo()
    }
}

function pinglog() {
    console.log("ping")
}

function pingtwo() {
    console.log("ping ping")
}