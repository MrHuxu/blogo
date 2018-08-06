import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';
import styled from 'styled-components';

const TagsContainer = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
  overflow: auto;
  background: linear-gradient(20deg, #566994, #9AFFFF);
`;

const Tags = ({ data }) => {
  const { tags, times } = data;

  return (
    <TagsContainer>
      <a href="/"> back to home </a>
      { tags.map(item => (
        <p style={ { fontSize: 16 + times[item] * 10 } }>
          { item.toString() }
        </p>
      )) }
    </TagsContainer>
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
