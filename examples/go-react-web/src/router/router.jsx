import React, { lazy, Suspense } from 'react';

import { Route, Switch, BrowserRouter as Router } from 'react-router-dom';

import noMatch from '../pages/noMatch';
const home = lazy(() => import('../pages/home'));
const payqr = lazy(() => import('../pages/Showqr'));
const login = lazy(() => import('../pages/login'));
const RouteMap = [
  {
    name: 'home',
    path: '/',
    exact: true,
    component: home
  },
  {
    name: 'payqr',
    path: '/pay/:uid',
    exact: true,
    component: payqr
  },
  {
    name: 'login',
    path: '/login',
    exact: true,
    component: login
  }
];

const routes = () => {
  return (
    <Router>
      <Suspense fallback={<div>Loading...</div>}>
        <Switch>
          {RouteMap.map(item => {
            return (
              <Route
                exact={!!item.exact}
                key={item.name}
                path={item.path}
                // 这里的路由有bug....
                render={props => <item.component {...props} />}
              />
            );
          })}
          <Route component={noMatch}> </Route>
        </Switch>
      </Suspense>
    </Router>
  );
};

export default routes;
