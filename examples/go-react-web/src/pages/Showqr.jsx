import React, { Component } from 'react';

class Showqr extends Component {
  componentWillMount() {
    console.log(this.props);
  }

  componentDidMount() {}

  componentWillReceiveProps(nextProps) {}

  shouldComponentUpdate(nextProps, nextState) {}

  componentWillUpdate(nextProps, nextState) {}

  componentDidUpdate(prevProps, prevState) {}

  componentWillUnmount() {}

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
