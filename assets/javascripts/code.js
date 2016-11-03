$(() => {
  $('pre code').each((i, block) => {
    CodeMirror(block.parentElement, {
      lineNumbers: true,
      value: block.innerText,
      mode: 'javascript',
      tabSize: 2,
      readOnly: true
    });
    block.remove();
  });
});