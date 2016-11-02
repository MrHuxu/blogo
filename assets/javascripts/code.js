$(() => {
  $('pre code').each((i, block) => {
    CodeMirror(block, {
      lineNumbers: true,
      value: block.innerText,
      mode: 'javascript',
      tabSize: 2,
      autofocus: true
    });
  });
});