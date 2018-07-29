import React from 'react';
import { shape, string } from 'prop-types';
import { connect } from 'react-redux';
import styled from 'styled-components';

const PostContainer = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
  overflow: auto;
  background: linear-gradient(20deg, rgb(219, 112, 147), #daa357);
`;

const Post = ({ data }) => {
  const { title, time, content } = data;

  return (
    <PostContainer>
      <a href="/"> back to home </a>
      <p> { title }</p>
      <p> { time } </p>
      <p> { decodeURI(content) } </p>
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
