import React, { Component } from 'react';
import { connect } from 'react-redux';

import api from '../api';
class Home extends Component {
  // 页面请求数据调用reducer
  async componentWillMount() {
    const { data } = await api.account.getUserInfo(1);
    console.log('send statusCode===>', data);
    if (data.statusCode === 200) {
      this.props.hi(data.result);
    }
  }
  render() {
    return (
      <div>
        嗨！首页君 !
        {Object.keys(this.props.UserInfo).map(item => {
          return (
            <div key={item}>
              {item}={this.props.UserInfo[item]}
            </div>
          );
        })}
      </div>
    );
  }
}
const mapStateToProps = (state, prop) => {
  console.log(state, prop);
  return {
    UserInfo: {
      ...state.home
    }
  };
};
const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    hi: payload => {
      console.log('ownProps', ownProps);
      dispatch({
        type: 'GetUserInfo',
        payload: payload
      });
    }
  };
};
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Home);
