$(() => {
  var tagContainer = $('.tag-container');
  var postContainer = $('.post-container');
  if (tagContainer.height() > postContainer.height()) {
    postContainer.height(tagContainer.height());
  }
});