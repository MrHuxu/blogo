import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';

import { Container } from './elements';

const Tags = ({ data }) => {
  const { tags, times } = data;

  return (
    <Container>
      <a href="/"> back to home </a>
      { tags.map(item => (
        <span style={ { fontSize: 16 + times[item] * 10 } }>
          { item }
        </span>
      )) }
    </Container>
  );
};

Tags.propTypes = {
  data : shape({
    list  : arrayOf(string),
    infos : objectOf(number)
  })
};

const mapStateToProps = ({ tags }) => ({ data: tags });

export default connect(mapStateToProps)(Tags);
