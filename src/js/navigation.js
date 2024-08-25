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


// document.getElementById('locationSelect').addEventListener('change', function() {
//     var iframes = document.getElementsByTagName('iframe');
//     for (var i = 0; i < iframes.length; i++) {
//       iframes[i].style.display = 'none';
//     }
//     var selectedIndex = this.value;
//     document.getElementById('map' + selectedIndex).style.display = 'block';
//   });
async function fetchLocation(index) {
    try {
        const response = await fetch(`/geoMap?index=${index}`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const result = await response.json();
        const listElement = document.getElementById('location-info');
        listElement.innerHTML = ''; // Clear any existing content
        
        const iframe = document.createElement('iframe');
        iframe.src = result; 
        iframe.width = '97%';
        iframe.height = '450';
        iframe.style.border = '0'; // Changed from '5px' to '0'
        iframe.allowFullscreen = true;
        iframe.loading = 'lazy';
        
        listElement.appendChild(iframe);
        console.log(listElement);
    } catch (error) {
        console.error('Fetch error:', error);
        const listElement = document.getElementById('location-info');
        listElement.innerHTML = '<p>Error fetching location</p>';
    }
}

function handleSelectChange() {
    const selectElement = document.getElementById('locationSelect');
    const selectedIndex = selectElement.value;
    const listElement = document.getElementById('location-info');
    
    if (selectedIndex !== '') {
        fetchLocation(selectedIndex);
        listElement.style.display = 'block'; // Show the container
    } else {
        listElement.innerHTML = '<p>Select a location...</p>';
        listElement.style.display = 'none'; // Hide the container when no location is selected
    }
}

document.addEventListener('DOMContentLoaded', () => {
    const selectElement = document.getElementById('locationSelect');
    selectElement.addEventListener('change', handleSelectChange);
    
    // Initialize the view
    handleSelectChange();
});