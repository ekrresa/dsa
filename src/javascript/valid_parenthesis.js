function isValidParenthesis(s) {
  if (s.length % 2 === 1) return false;

  let store = [];
  let tokens = { "(": ")", "{": "}", "[": "]" };

  for (const item of s.split("")) {
    if (tokens[item]) {
      store.push(item);
    } else {
      let lastItem = store.pop();
      if (tokens[lastItem] !== item) {
        return false;
      }
    }
  }

  return store.length === 0;
}
