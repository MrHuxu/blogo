import React from 'react';
import { shape, string } from 'prop-types';
import { connect } from 'react-redux';
import { parse } from 'marked';

import { Container } from './elements';

const Post = ({ data }) => {
  const { title, time, content } = data;

  return (
    <Container>
      <a href="/"> back to home </a>
      <p> { title }</p>
      <p> { time } </p>
      <div dangerouslySetInnerHTML={ { __html: parse(content) } } />
    </Container>
  );
};

Post.propTypes = {
  data : shape({
    title   : string,
    content : string
  })
};

const mapStateToProps = ({ post }) => ({ data: post });

export default connect(mapStateToProps)(Post);
