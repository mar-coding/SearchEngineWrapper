const searchInput = document.getElementById('searchInput');
const searchResults = document.getElementById('searchResults');
const pagination = document.getElementById('pagination');

let currentPage = 1;
const itemsPerPage = 10;

function handleSearch() {
  const searchQuery = searchInput.value.toLowerCase();
  const options = {
    method: 'GET',
    // headers: {
    //   'ngrok-skip-browser-warning': '1'
    // }
  };

  fetch(`http://127.0.0.1:8000/search?q=${encodeURIComponent(searchQuery)}&pageSize=${itemsPerPage}&pageNo=${currentPage}`, options)
    .then(response => response.json())
    .then(data => {
      displaySearchResults(data);
      renderPagination(data.totalItemsCount);
    })
    .catch(error => {
      console.error('Error fetching search results:', error);
    });
}

function displaySearchResults(results) {
  const searchResultsContainer = document.getElementById('searchResults');
  searchResultsContainer.innerHTML = '';

  results.items.forEach(item => {
    const listItem = document.createElement('li');

    const titleLink = document.createElement('a');
    titleLink.href = item.url;
    titleLink.textContent = item.title;

    const description = document.createElement('p');
    description.textContent = item.description;

    listItem.appendChild(titleLink);
    listItem.appendChild(description);

    searchResultsContainer.appendChild(listItem);
  });
}

function handlePageChange(pageNumber) {
  currentPage = pageNumber;
  handleSearch();
}

function renderPagination(totalItemsCount) {
  pagination.innerHTML = '';

  const totalPages = Math.ceil(totalItemsCount / itemsPerPage);

  for (let i = 1; i <= totalPages; i++) {
    const button = document.createElement('button');
    button.textContent = i;
    button.addEventListener('click', () => handlePageChange(i));
    if (i === currentPage) {
      button.disabled = true;
    }
    pagination.appendChild(button);
  }
}

// Initial rendering
handleSearch();