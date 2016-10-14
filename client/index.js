import React, { Component } from 'react';
import { render } from 'react-dom';

class Blogo extends Component {
  render () {
    return (
      <h1> This is Blogo </h1>
    );
  }
}

render(<Blogo />, document.getElementById('blogo'));