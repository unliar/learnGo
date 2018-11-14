import { GetPayInfo, GetUserInfo } from './action';

export default (state = { Id: 1, Brief: '' }, action) => {
  switch (action.type) {
    case GetPayInfo: {
      console.log(111);
      break;
    }
    case GetUserInfo: {
      console.log('GetUserInfo==>state', state);
      console.log('GetUserInfo==>action', action);
      return { ...state, ...action.payload };
    }
    // 默认一定要原封不动的返回初始状态
    default: {
      return state;
    }
  }
};
