import styled from 'styled-components';

export const Container = styled.div`
  height: 100%;
  overflow: auto;
`;

export const Year = styled.p`
  color: #BBB;
  margin: 10px 0 20px 0;
`;

export const Item = styled.div`
  position: relative;
  padding: 10px 0 10px 0;
  width: calc(100% - 10px);
  height: 50px;
  transition: transform 0.6s;

  &:hover {
    background-color: #EEE;
    transform: translate(10px, 0);
  }
`;

export const ItemDate = styled.span`
  position: absolute;
  color: #888;
`;

export const ItemLink = styled.a`
  position: absolute;
  top: 7px;
  left: 60px;
  color: #666;
  font-size: 18px;
  line-height: 24px;
  text-decoration: underline;

  &:hover {
    color: #666;
    text-decoration: underline;
  }
`;

export const PrevNext = styled.div`
  display: inline-block;
  text-align: center;
  width: 110px;
  font-size: 14px;
  letter-spacing: 2px;
  color: #666;
  border: 1px solid #DADADA;
  background: #FFF;
  padding: 7px 8px 7px 10px;
  margin: 30px 20px 50px 0;

  &:hover {
    background: #EEE;
  }
`;
