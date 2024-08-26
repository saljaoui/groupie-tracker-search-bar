let showInfodate = document.getElementById("showInfodate");
let showInfoRelation = document.getElementById("showInfoRelation");
let showInfolocation = document.getElementById("showInfolocation");

let date = document.getElementById("date");
let relation = document.getElementById("relation");
let locations = document.getElementById("location");
let color = document.querySelector(".color");

// Default state: date is visible, others are hidden
toggleVisibility(date, [relation, locations]);

function toggleVisibility(showElement, hideElements) {
    showElement.classList.add('visible');
    showElement.classList.remove('hidden');
    hideElements.forEach(element => {
        element.classList.add('hidden');
        element.classList.remove('visible');
    });
   
}

showInfodate.onclick = function() {
    toggleVisibility(date, [relation, locations]);
}

showInfoRelation.onclick = function() {
    toggleVisibility(relation, [date, locations]);
     color.classList.remove('color')
}

showInfolocation.onclick = function() {
    toggleVisibility(locations, [date, relation]);
    color.classList.remove('color')
}

const navbarLinks = document.querySelectorAll('.navbar-link');

navbarLinks.forEach(link => {
    link.addEventListener('click', function() {
        navbarLinks.forEach(nav => nav.style.color = '');
        // Change the color of the clicked link
        this.style.color = '#d39c51';
    });
});

async function fetchLocation(index) {
        const response = await fetch(`/geoMap?index=${index}`);

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const result = await response.json();


        const listElement = document.getElementById('location-info');
        listElement.innerHTML = '';
        
        const iframe = document.createElement('iframe');
        iframe.src = result; 
        iframe.width = '97%';
        iframe.height = '450';
        iframe.style.border = '0';
        listElement.appendChild(iframe);
}

function handleSelectChange() {
    const selectElement = document.getElementById('locationSelect');
    const selectedIndex = selectElement.value;
    const listElement = document.getElementById('location-info');
    
    if (selectedIndex !== '') {
        fetchLocation(selectedIndex);
        listElement.style.display = 'block'; 
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const selectElement = document.getElementById('locationSelect');
    selectElement.addEventListener('change', handleSelectChange);

    handleSelectChange();
});