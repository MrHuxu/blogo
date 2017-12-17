$(() => {
  $('p>img').each((index, img) => {
    $(img.parentElement).css({ textAlign: 'center' });
  });
});
