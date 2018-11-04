import React from 'react';

import { BrowserRouter as Router, Route, Link } from 'react-router-dom';

import home from '../pages/home';
const routes = () => {
  return (
    <Router>
      <h1>111</h1>
      <Route exact path="/" component={home} />
      <Route exact path="/pay/:uid" component={home} />
      <Route exact path="/login" component={home} />
    </Router>
  );
};

export default routes;
