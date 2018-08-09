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
  height: 50px;
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
`;
