import React, { Component } from 'react';
import { connect } from 'react-redux';

class Home extends Component {
  render() {
    return <div>嗨！首页君 </div>;
  }
}

export default connect()(Home);
