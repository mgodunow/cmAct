const modalWrapper = document.querySelector('.modal-wrapper');
const wrapper = document.querySelector('.wrapper');
const loginLink = document.querySelector('.login-link');
const registerLink = document.querySelector('.register-link');
const btnPopup = document.querySelector('.login-button');
const iconClose = document.querySelector('.icon-close');
const overlay = document.querySelector('.overlay');

registerLink.addEventListener('click', ()=> {
    wrapper.classList.add('active');
});

loginLink.addEventListener('click', ()=> {
    wrapper.classList.remove('active');
});

btnPopup.addEventListener('click', ()=> {
    wrapper.classList.add('active-popup');
    modalWrapper.classList.add('active-popup')
})

iconClose.addEventListener('click', ()=> {
    wrapper.classList.remove('active-popup');
    modalWrapper.classList.remove('active-popup')
})

overlay.addEventListener('click', ()=> {
    wrapper.classList.remove('active-popup')
    modalWrapper.classList.remove('active-popup')
})

var formLogin = document.getElementById('login-form');
var formRegister = document.getElementById('register-form');

formLogin.onsubmit = sendRequest(formLogin, '/login') 
formRegister.onsubmit = sendRequest(formRegister, '/registration') 

function sendRequest(form, root) {
            var xhr = new XMLHttpRequest();
            var formData = new FormData(form);
            //open the request
            xhr.open('POST','http://localhost:8080' + root)
            xhr.setRequestHeader("Content-Type", "application/json");

            //send the form data
            xhr.send(JSON.stringify(Object.fromEntries(formData)));

            xhr.onreadystatechange = function() {
                if (xhr.readyState == XMLHttpRequest.DONE) {
                    form.reset(); //reset form after AJAX success or do something else
                }
            }
            //Fail the onsubmit to avoid page refresh.
            return false; 
    }