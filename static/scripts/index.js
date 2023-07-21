const modalWrapper = document.querySelector('.modal-wrapper');
const wrapper = document.querySelector('.wrapper');
const loginLink = document.querySelector('.login-link');
const registerLink = document.querySelector('.register-link');
const btnPopup = document.querySelector('.login-button');
const iconClose = document.querySelector('.icon-close');
const overlay = document.querySelector('.overlay');
var formLogin = document.getElementById('login-form');
var formRegister = document.getElementById('register-form');

registerLink.addEventListener('click', ()=> {
    wrapper.classList.add('active');
});

loginLink.addEventListener('click', ()=> {
    wrapper.classList.remove('active');
});

btnPopup.addEventListener('click', ()=> {
    wrapper.classList.add('active-popup');
    modalWrapper.classList.add('active-popup');
})

iconClose.addEventListener('click', ()=> {
    wrapper.classList.remove('active-popup');
    modalWrapper.classList.remove('active-popup');
})

overlay.addEventListener('click', ()=> {
    wrapper.classList.remove('active-popup');
    modalWrapper.classList.remove('active-popup');
})

formRegister.onsubmit = function(event){
        event.preventDefault();
        var xhr = new XMLHttpRequest();
        var formData = new FormData(formRegister);
        var values = formData.values();

        username = values.next().value;
        email = values.next().value;
        password = values.next().value;
        
        if (!isValidUsename(username)) {
            alert('The username can contain only numbers and Latin letters and have a length of 5 to 30 characters');
            return;
          }
        
        if (!isValidEmail(email)) {
            alert('Email is incorrect. Try again or use another email');
            return;
        }

        if (!isValidPassword(password)) {
            alert('The password must contain at least one uppercase letter, one lowercase letter and one digit and be between 8 and 30 characters long');
            return;
        }
        //open the request
        xhr.open('POST','http://localhost:8080/registration')
        xhr.setRequestHeader("Content-Type", "application/json");

        //send the form data
        xhr.send(JSON.stringify(Object.fromEntries(formData)));

        xhr.onreadystatechange = function() {
            if (xhr.readyState == XMLHttpRequest.DONE) {
                if (xhr.status != 201) {
                    alert(xhr.responseText);
                    return;
                }
                wrapper.classList.remove('active-popup');
                modalWrapper.classList.remove('active-popup');
                formRegister.reset();
                alert(xhr.responseText);
            }
        }
        //Fail the onsubmit to avoid page refresh.
        return false; 
}

function isValidUsename(username) {
    const pattern = /^[a-zA-Z0-9]{5,30}$/;
    if (username.length >= 5 || username.length <= 30) {
        return pattern.test(username);
    }
    return;
    }
      
function isValidPassword(password) {
    const pattern = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{8,30}$/;
    return pattern.test(password);
    }

function isValidEmail(email) {
    const pattern = /^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$/;
    return pattern.test(email);
}

formLogin.onsubmit = function(event) {
    event.preventDefault();
    var xhr = new XMLHttpRequest();
    var formData = new FormData(formLogin);
    var values = formData.values();

    Email = values.next().value;

    //In future will be able to login with username, but now only with email.
    // if (!usernameOrEmail.contains("@")) {
    //     if (!isValidUsename(usernameOrEmail)) {
    //         alert('The user name can contain only numbers and Latin letters and have a length of 5 to 30 characters');
    //         return;
    //     }}
    if (!isValidEmail(Email)) {
        alert('Email is incorrect. Try again or use another email');
        return;
        }

    xhr.open('POST','http://localhost:8080/login')
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.send(JSON.stringify(Object.fromEntries(formData)));

    xhr.onreadystatechange = function() {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            if (xhr.status != 200) {
                alert(xhr.responseText);
                return;
            }
            token = jwtParse(xhr.responseText);
            localStorage.setItem("jwt", token);
            formLogin.reset();
            wrapper.classList.remove('active-popup');
            modalWrapper.classList.remove('active-popup');
            formLogin.reset();
        }
    return false;
    }
}

function jwtParse(json) {
    const obj = JSON.parse(json);
    return obj.token;
}