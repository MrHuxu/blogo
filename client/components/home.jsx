import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';
import styled from 'styled-components';

const HomeContainer = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
  background: linear-gradient(20deg, rgb(219, 112, 147), #daa357);
`;

const Home = ({ data }) => {
  const { titles, infos } = data;

  return (
    <HomeContainer>
      { titles.map(title => (
        <div>
          <p> { title } </p>
          <p> { infos[title] } </p>
          <a href="/test"> to test </a>
        </div>
      )) }
    </HomeContainer>
  );
};

Home.propTypes = {
  data : shape({
    list  : arrayOf(string),
    infos : objectOf(number)
  })
};

const mapStateToProps = ({ home }) => ({ data: home });

export default connect(mapStateToProps)(Home);
