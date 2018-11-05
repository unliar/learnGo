import React, { Component } from 'react';

import { getUserInfo } from '../api';
class userInfo extends Component {
  state = {
    Id: null,
    LoginName: null,
    Nickname: null,
    Age: null,
    Gender: null,
    Avatar: null,
    Location: null,
    Profession: null,
    Status: null,
    Phone: null,
    Brief: null,
    NationCode: null
  };
  async componentDidMount() {
    const uid = this.props.match.params.uid;
    const { data } = await getUserInfo(uid);
    this.setState(Object.assign(this.state, { ...data.result }));
  }
  render() {
    return (
      <div>
        {Object.keys(this.state).map(item => {
          const key = item;
          const value = this.state[key];
          return (
            <div>
              {key}:{value}
            </div>
          );
        })}
      </div>
    );
  }
}

export default userInfo;
