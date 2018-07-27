export const createStaticStore = data => ({
  getState : data ? () => JSON.parse(data) : {}
});
