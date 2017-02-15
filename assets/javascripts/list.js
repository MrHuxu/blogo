$(() => {
  const hasCondition = ($dom, condition) => (
    $dom.find('.item-title a')[0].innerText.toLowerCase().indexOf(condition) !== -1 ||
    $dom.find('.item-tags')[0].innerText.toLowerCase().indexOf(condition) !== -1
  );

  $('.search-container input').keyup((e) => {
    let condition = e.target.value.toLowerCase();

    $('.item-container').each((index, item) => {
      let $item = $(item);

      if (hasCondition($item, condition)) {
        $item.show();
        $item.next().show();
      } else {
        $item.hide();
        $item.next().hide();
      }
    });
  });
});