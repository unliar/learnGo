import React, { Component } from 'react';

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

export default Hello;
