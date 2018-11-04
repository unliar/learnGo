import React, { Suspense } from 'react';

import { Route, Switch, BrowserRouter as Router } from 'react-router-dom';

import Reactloadable from 'react-loadable';
import noMatch from '../pages/noMatch';
// const home = lazy(() => import('../pages/home'));
// const payqr = lazy(() => import('../pages/Showqr'));
// const login = Reactloadable({
//   loader: () => import('../pages/login'),
//   loading: () => <div />
// });
const RouteMap = [
  {
    name: 'home',
    path: '/',
    exact: true,
    component: Reactloadable({
      loader: () => import('../pages/home'),
      loading: () => <div />
    })
  },
  {
    name: 'payqr',
    path: '/pay/:uid',
    exact: true,
    component: Reactloadable({
      loader: () => import('../pages/Showqr'),
      loading: () => <div />
    })
  },
  {
    name: 'login',
    path: '/login',
    exact: true,
    component: Reactloadable({
      loader: () => import('../pages/login'),
      loading: () => <div />
    })
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
                component={props => <item.component {...props} />}
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
