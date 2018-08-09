import React from 'react';
import { shape, arrayOf, string, objectOf, number, object } from 'prop-types';
import { connect } from 'react-redux';

import { Container, Year, Item, ItemDate, ItemLink } from './elements';

let monthNames = [
  'January', 'February', 'March', 'April', 'May', 'June',
  'July', 'August', 'September', 'October', 'November', 'December'
];

const Home = ({ data, match }) => {
  const { titles, infos } = data;

  let year = infos[titles[0]].time.slice(0, 4);

  return (
    <Container>
      <a href="/tags"> to tags </a>
      <Year> /* { year } */ </Year>
      { titles.map(title => {
        let eles = [
          <Item>
            <ItemDate> {
              `${monthNames[parseInt(infos[title].time.slice(6))].slice(0, 3)} ${infos[title].time.slice(5, 7)}`
            } </ItemDate>
            <ItemLink href={ `/post/${title}` }>{ title }</ItemLink>
          </Item>
        ];

        if (infos[title].time.slice(0, 4) !== year) {
          year = infos[title].time.slice(0, 4);
          eles.unshift(<Year> /* { year } */ </Year>);
        }

        return eles;
      }) }
    </Container>
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
