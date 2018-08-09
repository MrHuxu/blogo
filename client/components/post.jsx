import React from 'react';
import { shape, string } from 'prop-types';
import { connect } from 'react-redux';
import { parse } from 'marked';
import styled from 'styled-components';

const PostContainer = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
  overflow: auto;
`;

const Post = ({ data }) => {
  const { title, time, content } = data;

  return (
    <PostContainer>
      <a href="/"> back to home </a>
      <p> { title }</p>
      <p> { time } </p>
      <div dangerouslySetInnerHTML={ { __html: parse(content) } } />
    </PostContainer>
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
