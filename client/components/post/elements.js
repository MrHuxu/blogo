import styled from 'styled-components';

export const Container = styled.div`
  height: 100%;
  overflow: auto;
`;

export const Title = styled.h1`
  margin: 40px 0;
  text-align: center;
  font-weight: 500;
  color: #646464;
  font-family: "Lato", sans-serif;
`;

export const PostDivider = styled.div`
  margin: 20px 0 40px 0 !important;
`;

export const TagsInDivider = styled.span`
  & a {
    color: #A3717F;
  }
`;

export const Content = styled.div`
  letter-spacing: .2px;
  font-size: 13px;
  color: #555;

  & h1, h2, h3, h4, h5, h6 {
    margin: 20px 0 15px;
    font-weight: 500;
    color: #646464;
  }

  & p, & li {
    line-height: 1.9;
  }

  & blockquote {
    padding: 15px 0 15px 15px;
    margin: 0 0 18px;
    border-left: 5px solid #D1D0CE;
    line-height: 28px;
    font-weight: normal;
    font-size: 15px;
    font-style: italic;
    color: #696969;
  }

  & img {
    max-width: 100%;
  }

  & a {
    color: #4183c4;
    text-decoration: none;
  }

  & hr {
    border: 0;
    color: #ddd;
    background-color: #ddd;
    height: 2px;
    margin: 5px 0 19px 0;
  }

  & code {
    display: inline;
    word-wrap: break-word;
    font-size: 10.6px;
    color: rgb(85, 85, 85);
    background: rgb(255, 255, 255);
    border-width: 1px;
    border-style: solid;
    border-color: rgb(221, 221, 221);
    border-image: initial;
    border-radius: 4px;
    padding: 1px 3px;
    margin: -1px 1px 0px;
  }

  & pre code {
    display: block;
    font-size: 10.8px;
    line-height: 18px;
    font-weight: 12px;
    letter-spacing: .5px;
    margin: 0 0 20px 0;
    padding: 15px !important;
    background-color: #f7f7f7 !important;
    border-width: 0;
  }
`;

export const Disqus = styled.div`
  padding: 18px 0 54px 0;
`;
