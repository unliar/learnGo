import React, { Component } from 'react';

import qrcode from 'qrcode';
import api from '../api';
class Showqr extends Component {
  state = {
    UID: null,
    Alipay: null,
    TenPay: null,
    AliPayData: null,
    TenPayData: null
  };

  async componentDidMount() {
    const { data } = await api.pay.getPayInfo(this.props.match.params.uid);
    if (data.statusCode !== 200) return;
    const AliPayData = await qrcode.toDataURL(data.result.PayInfo.Alipay);
    const TenPayData = await qrcode.toDataURL(data.result.PayInfo.TenPay);

    this.setState({
      ...data.result.PayInfo,
      AliPayData,
      TenPayData
    });
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

        {this.state.AliPayData ? (
          <img src={this.state.AliPayData} alt="ali" />
        ) : null}
        {this.state.TenPayData ? (
          <img src={this.state.TenPayData} alt="tenpay" />
        ) : null}
      </div>
    );
  }
}

export default Showqr;
