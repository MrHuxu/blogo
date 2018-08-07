import React from 'react';
import { shape, arrayOf, string, objectOf, number } from 'prop-types';
import { connect } from 'react-redux';
import styled from 'styled-components';

const HomeContainer = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
  overflow: auto;
  background: linear-gradient(20deg, rgb(219, 112, 147), #daa357);
`;

const Home = ({ data }) => {
  const { titles, infos } = data;

  return (
    <HomeContainer>
      <a href="/tags"> to tags </a>
      { titles.map(title => (
        <div>
          <a href={ '/post/' + title }> { title } </a>
          <p> { infos[title].seq } </p>
          <p> { infos[title].time } </p>
          { infos[title].tags.map(tag => (
            <p> { tag } </p>
          )) }
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

const mapStateToProps = ({ page }) => ({ data: page });

export default connect(mapStateToProps)(Home);
