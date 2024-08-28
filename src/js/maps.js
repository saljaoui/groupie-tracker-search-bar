
async function fetchLocation(index) {
    const response = await fetch(`/geoMap?index=${index}`);

    if (!response.ok) {
        throw new Error('Network response was not ok');
    }

    const result = await response.json();
    const listElement = document.getElementById('location-info');
    listElement.innerHTML = '';

    // Create an <a> element
    const link = document.createElement('a');
    link.href = result; 
    link.textContent = 'View Location >';
    link.style.display = 'block'; 
    link.style.textAlign = 'center'; 
    link.classList.add("href-location")

    listElement.appendChild(link);
    console.log(link);
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