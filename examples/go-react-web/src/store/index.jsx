import { createStore, combineReducers } from 'redux';

import home from './home/reducer';

export default createStore(combineReducers({ home }));
