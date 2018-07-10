let content = document.querySelector('.friend-content');

function displayFriends() {
  var count;
  content.innerHTML = '';
  for (count = 0; count < 8; count++) {
    let friendPage = document.createElement('article');
    friendPage.innerHTML = '<img src="friend-icon.png" class="friend-img-center">';
    friendPage.classList.add("friend-article");
    content.appendChild(friendPage);
  }
}

window.onload = displayFriends();
