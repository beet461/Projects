class NavBar extends HTMLElement {
    constructor() {
        super();
        this.container = document.createElement('div');
        this.logo = document.createElement('img');
        this.reg = document.createElement('button');
        this.login = document.createElement('button');

        this.reg.innerHTML = 'register';
        this.login.innerHTML = 'login';

        this.logo.className = 'nav_logo';
        this.login.className = 'nav_button nav_login';
        this.reg.className = 'nav_button nav_register';
    }
    connectedCallback() {
        document.body.append(this.container);
        this.container.append(this.logo);
        this.container.append(this.reg);
    }
}

customElements.define('nav-bar', NavBar);
