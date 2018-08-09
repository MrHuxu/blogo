import styled, { keyframes } from 'styled-components';

export const Container = styled.div`
  width: 100%;
  height: 100%;
`;

export const App = styled.div`
  padding: 60px 0 0 0;
  min-height: calc(100% - 60px);
  width: 720px;
  margin-left: calc((100% - 720px) / 2);
  margin-right: calc((100% - 720px) / 2);

  @media screen and (max-width: 720px) {
    width: 100%;
    margin-left: 0;
    margin-right: 0;
    padding: 60px 15px 0 15px;
  }
`;

export const Title = styled.div`
  margin: 0 0 45px 0;
  font-size: 18px;
`;

const blink = keyframes`
  0% {
    opacity: 1;
  }
  50% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
`;

export const TitleArrow = styled.a`
  margin: 0 0 0 6px;
  animation: ${blink} 1.2s infinite linear;
`;

export const Footer = styled.div`
  background-color: #FFFEEC;
  height: 60px;
  padding: 18px 0 0 0;
  text-align: center;
  color: #444444;
  opacity: .8;
  letter-spacing: .8px;
  font-family: Lato,sans-serif;

  @media screen and (max-width: 720px) {
    height: 60px;
    padding: 10px 0 18px 0;
  }
`;
