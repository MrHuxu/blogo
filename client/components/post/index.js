import React from 'react';
import { shape, string } from 'prop-types';
import { connect } from 'react-redux';
import { parse } from 'marked';

import { Container, Title, PostDivider, TagsInDivider, Content, Disqus } from './elements';

import { monthNames } from '../layout/constants';

const Post = ({ data }) => {
  const { title, time, tags, content } = data;

  return (
    <Container>
      <Title> { title } </Title>
      <PostDivider className="ui horizontal divider">
        { `${monthNames[parseInt(time.slice(5, 7)) - 1].slice(0, 3)} ${time.slice(8, 10)}, ${time.slice(0, 4)}` }
        { tags.map(tag => (
          <TagsInDivider>
            &nbsp;&nbsp;|&nbsp;&nbsp;
            <a href={ `/page/0?tag=${tag}` }>{ tag }</a>
          </TagsInDivider>
        )) }
      </PostDivider>
      <Content dangerouslySetInnerHTML={ { __html: parse(content) } } />

      <Disqus id="disqus_thread" />
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
