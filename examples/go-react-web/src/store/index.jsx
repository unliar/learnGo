import { createStore, combineReducers } from 'redux';

import UserInfo from './reducer';

export default createStore(combineReducers({ UserInfo }));
