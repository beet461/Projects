class NavBar extends HTMLElement {
    constructor() {
        super();
        this.container = document.createElement('div');
        this.logo = document.createElement('img');
        this.reg = document.createElement('button');
        this.login = document.createElement('button');

        this.logo.className = 'nav_logo';
        this.reg.className = 'nav_register';
    }
    connectedCallback() {
        if (window.location !== 'http://localhost:8080/home.html') {
            //this.logo.onclick = window.location.assign('../home.html');
            console.log('sdf');
        }
        document.body.append(this.container);
        this.container.append(this.logo);
        this.container.append(this.reg);
    }
}

customElements.define('nav-bar', NavBar);
