import { getUserInfo } from '../../api';
import { GetPayInfo, GetUserInfo, GetUserInfoAction } from './action';

export const GetUserInfoFromRemote = async id => {
  try {
    const { data } = await getUserInfo(id);
    console.log('request data', data);
    return data;
  } catch (error) {
    console.log(error);
    return {};
  }
};

export default (state = { User: 1 }, action) => {
  switch (action.type) {
    case GetPayInfo: {
      console.log(111);
      break;
    }
    case GetUserInfo: {
      console.log('GetUserInfo==>state', state);
      console.log('GetUserInfo==>action', action);
      return { User: action.payload.User + state.User || 0 };
    }
    // 默认一定要原封不动的返回初始状态
    default: {
      return state;
    }
  }
};
