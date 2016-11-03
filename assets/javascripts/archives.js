$(() => {
  var tagContainer = $('.tag-container');
  var postContainer = $('.post-container');
  if (tagContainer.height() > postContainer.height()) {
    postContainer.height(tagContainer.height());
  }

  $('.tag-link').hover(function () {
    $(this).addClass('animated tada');
  }, function () {
    $(this).removeClass('animated tada');
  });

  var selectedTag = location.search.split('=')[1];
  if (selectedTag) {
    $(`.tag-${selectedTag} a`).css({ color: '#EC354C' });
    $('.show-all-link').addClass('animated flipInX');
  } else {
    $('.show-all-link').hide();
  }

  $(window).scroll(() => {
    if ($('body').scrollTop() < 120) {
      $('.tag-container').css({ top: 100 - $('body').scrollTop() });
    }
  });
});