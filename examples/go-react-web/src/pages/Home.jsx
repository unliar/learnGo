import React, { Component } from 'react';
import { connect } from 'react-redux';

class Home extends Component {
  componentDidMount() {}
  render() {
    return <div onClick={this.props.hi}>嗨！首页君 {this.props.User} </div>;
  }
}
const mapStateToProps = (state, prop) => {
  console.log(state, prop);
  return { User: state.home.User };
};
const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    hi: () => {
      console.log('ownProps', ownProps);
      dispatch({
        type: 'GetUserInfo',
        payload: { User: 1 }
      });
    }
  };
};
export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Home);
