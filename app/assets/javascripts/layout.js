$(() => {
  const re = /[a-zA-Z0-9_#: .\-/\\[\]]+/g;
  const getWordsFromText = text => (text.match(re) || []).filter(t => (/[a-zA-Z0-9]+/g).test(t));

  const replaceWordsInText = (text, words) => {
    return words.length ? words.reduce((prev, word, index) => {
      var wordEndIdx = text.indexOf(word) + word.length;
      if (index === words.length - 1) {
        prev += text.replace(word, ` ${word} `);
      } else {
        prev += (text.slice(0, wordEndIdx)).replace(word, ` ${word} `);
        text = text.slice(wordEndIdx);
      }
      return prev;
    }, '') : text;
  };

  var containers = $('p, li, strong, a');
  containers.each((idx, container) => {
    container.childNodes.forEach(child => {
      if (0 === child.childNodes.length) {
        let words = getWordsFromText(child.textContent);
        child.textContent = replaceWordsInText(child.textContent, words);
      }
    });
  });

  $('pre code').each((i, block) => {
    hljs.highlightBlock(block);
  });

  $('.logo').mouseenter(function () {
    $(this).animateCss('swing');
  });

  var $toTop = $('#back-to-top');
  $toTop.click(() => {
    $('body').animate({ scrollTop: 0 });
  });

  $(window).scroll(() => {
    if ($('body').scrollTop() > 10) {
      $toTop.show(400);
    } else {
      $toTop.hide(400);
    }
  });
});