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