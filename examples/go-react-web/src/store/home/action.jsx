export const GetUserInfo = 'GetUserInfo';
export const GetPayInfo = 'GetPayInfo';

/**
 * 获取用户信息
 * @param {number} UID 用户id
 */
export const GetUserInfoAction = UserInfo => ({
  type: GetUserInfo,
  payload: UserInfo
});

/**
 * 获取用户支付信息
 * @param {number} UID 用户id
 */
export const GetPayInfoAction = PayInfo => ({
  type: GetPayInfo,
  payload: PayInfo
});
