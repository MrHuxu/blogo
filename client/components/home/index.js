import React from 'react';
import { shape, arrayOf, string, objectOf, number, object } from 'prop-types';
import { connect } from 'react-redux';

import { Container, Year, Item, ItemDate, ItemLink, PrevNext } from './elements';

import { monthNames } from '../layout/constants';

const Home = ({ data, match }) => {
  const { titles, infos, maxPage } = data;
  const page = parseInt(match.params.page);

  return (
    <Container>
      {titles.reduce((pre, title) => {
        if (infos[title].time.slice(0, 4) !== pre.year) {
          pre.year = infos[title].time.slice(0, 4);
          pre.eles.push(<Year> /* {pre.year} */ </Year>);
        }

        pre.eles.push(
          <Item>
            <ItemDate> {
              `${monthNames[parseInt(infos[title].time.slice(5, 7)) - 1].slice(0, 3)} ${infos[title].time.slice(8, 10)}`
            } </ItemDate>
            <ItemLink href={`/post/${title}`}>{title}</ItemLink>
          </Item>
        );

        return pre;
      }, { year: null, eles: [] }).eles}

      {page > 0 ? (
        <a href={`/page/${page - 1}`}>
          <PrevNext> <i className="icon-left-arrow link" />NEWER </PrevNext>
        </a>
      ) : null}

      {page < maxPage - 1 ? (
        <a href={`/page/${page + 1}`}>
          <PrevNext> OLDER<i className="icon-right-arrow link" /> </PrevNext>
        </a>
      ) : null}

    </Container>
  );
};

Home.propTypes = {
  data: shape({
    list: arrayOf(string),
    infos: objectOf(number),
    maxPage: number
  }),
  match: object
};

const mapStateToProps = ({ page }) => ({ data: page });

export default connect(mapStateToProps)(Home);
