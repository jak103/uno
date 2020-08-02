export default {
  get(key, def) {
    const val = localStorage.getItem(key) || def;
    try {
      return JSON.parse(val);
    } catch {
      // do nothing
    }
  },
  set(key, val) {
    localStorage.setItem(key, JSON.stringify(val));
  }
}