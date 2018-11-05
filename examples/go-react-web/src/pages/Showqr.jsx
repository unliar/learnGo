import React, { Component } from 'react';

import { getPayinfo } from "../api/index"
class Showqr extends Component {
  componentWillMount() {
    console.log(getPayinfo(1))
  }

  componentDidMount() { }

  componentWillReceiveProps(nextProps) { }

  shouldComponentUpdate(nextProps, nextState) { }

  componentWillUpdate(nextProps, nextState) { }

  componentDidUpdate(prevProps, prevState) { }

  componentWillUnmount() { }

  render() {
    const uid = this.props.match.params.uid;
    return (
      <div>
        hi qr===>
        {uid}
      </div>
    );
  }
}

export default Showqr;
