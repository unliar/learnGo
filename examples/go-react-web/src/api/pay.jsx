import axios from './base';

/**
 * 获取用户支付信息
 * @param {number} uid 用户id
 */
export const getPayInfo = uid => axios.get(`/api/pay/${uid}`);

/**
 * 创建支付信息
 *
 * @param {number} uid
 * @param {string} alipay
 * @param {string} tenpay
 */
export const postPayInfo = (uid, alipay, tenpay) =>
  axios({
    method: 'POST',
    url: '/api/pay',
    data: {
      uid,
      alipay,
      tenpay
    }
  });

/**
 * 创建支付信息
 *
 * @param {number} uid
 * @param {string} alipay
 * @param {string} tenpay
 */
export const putPayInfo = (uid, alipay, tenpay) =>
  axios({
    method: 'PUT',
    url: '/api/pay',
    data: {
      uid,
      alipay,
      tenpay
    }
  });

export default {
  getPayInfo,
  postPayInfo,
  putPayInfo
};
