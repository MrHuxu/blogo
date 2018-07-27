import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';
import styled from 'styled-components';

const TestContainer = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
  background: linear-gradient(20deg, #566994, #9AFFFF);
`;

const Test = ({ data }) => {
  const { list, infos } = data;

  return (
    <TestContainer>
      { list.map(item => (
        <p style={ { fontSize: 16 + infos[item] * 10 } }>
          { item }
        </p>
      )) }
      <a href="/"> back to home </a>
    </TestContainer>
  );
};

Test.propTypes = {
  data : shape({
    list  : arrayOf(string),
    infos : objectOf(number)
  })
};

const mapStateToProps = ({ test }) => ({ data: test });

export default connect(mapStateToProps)(Test);
