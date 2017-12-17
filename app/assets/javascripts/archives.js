$(() => {
  var tagContainer = $('.tag-list');
  var postContainer = $('.post-list');
  if (tagContainer.height() > postContainer.height()) {
    postContainer.height(tagContainer.height());
  }

  $('.tag-link').hover(function () {
    $(this).animateCss('tada');
  }, () => {});
});
