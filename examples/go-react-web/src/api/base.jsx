import axios from 'axios';

const baseURL =
  process.env.NODE_ENV === 'production'
    ? 'https://hipoor.com:8088'
    : 'http://127.0.0.1:8088';

axios.defaults.baseURL = baseURL;
axios.defaults.withCredentials = true;
export default axios;
