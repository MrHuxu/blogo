import React from 'react';
import { shape, arrayOf, string, objectOf, number, object } from 'prop-types';
import { connect } from 'react-redux';
import styled from 'styled-components';

const HomeContainer = styled.div`
  position: fixed;
  width: 100%;
  height: 100%;
  overflow: auto;
`;

const Home = ({ data, match }) => {
  const { titles, infos } = data;

  return (
    <HomeContainer>
      <p> { match.params.page } </p>
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
  }),
  match : object
};

const mapStateToProps = ({ page }) => ({ data: page });

export default connect(mapStateToProps)(Home);
