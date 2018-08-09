import styled from 'styled-components';

export const Container = styled.div`
  width: 100%;
  height: 100%;
`;

export const App = styled.div`
  padding: 60px 0 0 0;
  min-height: calc(100% - 133px);
  width: 720px;
  margin-left: calc((100% - 720px) / 2);
  margin-right: calc((100% - 720px) / 2);

  @media screen and (max-width: 720px) {
    margin-left: 0;
    margin-right: 0;
    padding: 60px 15px 0 15px;
  }
`;

export const Title = styled.div`
  margin: 0 0 60px 0;
  font-size: 18px;
`;

export const Footer = styled.div`
  background-color: #FFFEEC;
  height: 60px;
  padding: 15px 0 0 0;
  text-align: center;
  color: #444444;

  @media screen and (max-width: 720px) {
    height: 102px;
    padding: 22px 0 22px 0;
  }
`;
