const loginTab = document.getElementById('login-tab');
const registerTab = document.getElementById('register-tab');
const loginForm = document.getElementById('login-form');
const registerForm = document.getElementById('register-form');

console.log("PENISI");

registerTab.addEventListener('click',()=>{
    registerForm.style.display = 'block';
    loginForm.style.display = 'none';
    loginTab.className = "tab";
    registerTab.className = "tab-active";
})