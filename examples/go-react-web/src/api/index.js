import axios from 'axios';

axios.defaults.baseURL =
  process.env.NODE_ENV === 'production'
    ? 'https://hipoor.com:8088'
    : 'http://127.0.0.1:8088';

axios.defaults.withCredentials = true;
/**
 *
 * @param {number} uid 用户id
 */
export const getPayinfo = uid => {
  return axios.get(`/api/pay/${uid}`);
};

/**
 *
 * @param {number} uid
 *
 */
export const postPayinfo = ({ uid, alipay, tenpay }) => {
  return axios.post('/api/pay', { uid, alipay, tenpay });
};

export const getUserInfo = uid => axios.get(`/api/account/users/${uid}`);
