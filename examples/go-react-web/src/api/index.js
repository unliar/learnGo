import axios from 'axios';

axios.defaults.baseURL = 'http://192.168.31.236:8088';

axios.defaults.withCredentials = true;

export const getPayinfo = uid => {
  return axios.get(`/api/pay/${uid}`);
};

export const postPayinfo = ({ uid, alipay, tenpay }) => {
  return axios.post('/api/pay', { uid, alipay, tenpay });
};

export const getUserInfo = uid => axios.get(`/api/account/users/${uid}`);
