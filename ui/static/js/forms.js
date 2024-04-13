const urlParams = new URLSearchParams(window.location.search);
const queryContainer = urlParams.get('query');
document.querySelector('#query').value = queryContainer;

function handleClientSubmit() {
    const value = document.querySelector('#query').value
    urlParams.set('query', value)
    window.location.search = urlParams.toString();
}
