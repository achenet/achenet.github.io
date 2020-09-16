
let heading = document.querySelector('h1');
let button = document.querySelector('button');


function setUsername() {
	let name = prompt('Please enter your name');
	localStorage.setItem('name', name);
	heading.textContent = 'Welcome ' + name +' !';

}

if (!localStorage.getItem('name')) {
	setUsername();

}
else {
	let storedName = localStorage.getItem('name');
	heading.textContent = 'Welcome ' + name +' !';
}

button.onclick = function () {
	setUsername();
}
