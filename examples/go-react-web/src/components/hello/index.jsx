import React, { Component } from 'react';

import PropTypes from 'prop-types';
import stypes from './index.module.css';
class Hello extends Component {
  render() {
    return (
      <div className={stypes.gloabal}>
        <h1> Hello, World</h1>
        <p> {this.props.text}</p>
      </div>
    );
  }
}

Hello.propTypes = {
  text: PropTypes.string
};
Hello.defaultProps = {
  text: 'Hello-Component-default-props'
};
export default Hello;
