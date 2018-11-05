import React, { Component } from 'react';

import { getPayinfo } from '../api/index';
import { Object } from 'core-js';
class Showqr extends Component {
  state = {
    UID: null,
    Alipay: null,
    TenPay: null
  };

  async componentDidMount() {
    const { data } = await getPayinfo(this.props.match.params.uid);

    this.setState(Object.assign(this.state, { ...data.PayInfo }));
  }

  render() {
    return (
      <div>
        <h1>
          UID:
          {this.state.UID ? this.state.UID : '无用户id'}
        </h1>
        <p>
          Alipay:
          {this.state.Alipay ? this.state.Alipay : <span>无</span>}
        </p>
        <p>
          TenPay:
          {this.state.TenPay ? this.state.TenPay : <span>无</span>}
        </p>
      </div>
    );
  }
}

export default Showqr;
