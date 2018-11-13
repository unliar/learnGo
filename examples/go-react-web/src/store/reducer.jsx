import { getUserInfo } from '../api';
import { GetPayInfo, GetUserInfo, GetUserInfoAction } from './action';

export const GetUserInfoFromRemote = id => async dispatch => {
  try {
    const { data } = await getUserInfo(id);
    await dispatch(GetUserInfoAction(data.result));
  } catch (error) {
    console.log(error);
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