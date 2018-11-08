import { getUserInfo } from '../api';
import { GetUserInfoAction, GetPayInfo, GetUserInfo } from './action';

// export const GetUserInfoFromRemote = id => async dispatch => {
//   try {
//     const { data } = await getUserInfo(id);
//     await dispatch(GetUserInfoAction(data.result));
//   } catch (error) {
//     console.log(error);
//   }
// };

export default async (state = { user: 1 }, action) => {
  switch (action.type) {
    case GetPayInfo: {
      console.log(111);
      break;
    }
    case GetUserInfo: {
      console.log(222);
      break;
    }
    default: {
      console.log(3333);
      getUserInfo(1);
    }
  }
};
