function scrollD() {
    window.scrollTo(0, document.body.scrollHeight);
}

window.setInterval(scrollD, 3000)
window.clearInterval()


const elementList = document.querySelectorAll('body > div.application-outlet > div.authentication-outlet > div > main > section > div > div.scaffold-finite-scroll__content > div > ul > li > div > div > div > div.entity-result__content-inner-container.entity-result__content-inner-container--right-padding > a')
const hrefValues = [];

elementList.forEach((element) => {
    hrefValues.push(element.href);
});

console.log(hrefValues);


function generateUUID() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
        return v.toString(16);
    });
}

// Create an array to hold the rows of the CSV
const rows = [];

// Iterate over the list of links
rows.push('name,link,category')
for (const link of hrefValues) {
    // Get the current date with miliseconds
    const date = generateUUID();
    // Create a row with the date, link, and category
    const row = `${date},${link},8`;
    // Add the row to the array
    rows.push(row);
}

// Join the rows with newline characters to create the CSV content
const csv = rows.join('\n');

// Use the CSV content to create a blob object
const blob = new Blob([csv], { type: 'text/csv' });

// Create a URL for the blob object
const url = URL.createObjectURL(blob);

// Create a link element and set its href to the URL of the blob
const link = document.createElement('a');
link.setAttribute('href', url);
link.setAttribute('download', 'file.csv');

// Append the link to the body of the document and click it
document.body.appendChild(link);
link.click();



