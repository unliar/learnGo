import axios from './base';

/**
 * 获取用户信息
 *
 * @param {number} uid - 用户id
 */
export const getUserInfo = uid => axios.get(`/api/account/users/${uid}`);

/**
 * 检测用户手机号邮箱登录是否存在
 *
 * @param {string} type
 * @param {string} value
 */
export const checkAccountUnique = (type, value) =>
  axios({
    url: `/api/account/unique`,
    params: {
      type,
      value
    }
  });

/**
 * 注册用户
 *
 * @param {string} type 类型
 * @param {string} value 值
 * @param {string} password 密码
 */
export const postUser = (type, value, password) =>
  axios({
    url: '/api/account/users',
    method: 'POST',
    data: {
      type,
      value,
      password
    }
  });

/**
 * 登录或者刷新token
 *
 * @param {string} type 类型
 * @param {string} value 值
 * @param {string} password 密码
 * @param {string} opType 类型
 */
export const getToken = (type, value, password, opType = 'login') =>
  axios({
    url: '/api/account/tokens',
    method: 'POST',
    params: {
      type: opType
    },
    data: {
      type,
      value,
      password
    }
  });

export default {
  getToken,
  postUser,
  getUserInfo,
  checkAccountUnique
};
