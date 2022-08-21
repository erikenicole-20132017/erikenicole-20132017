const memoize = (fn: Function): Function => {
  const cache: { [key: string]: any } = {};
  return (...args: any[]): any => {
    const key = JSON.stringify(args);
    if (cache[key]) {
      return cache[key];
    }
    const result = fn(...args);
    cache[key] = result;
    return result;
  };
}
