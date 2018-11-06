import React, { Suspense } from 'react';

import { Route, Switch, BrowserRouter as Router, Link } from 'react-router-dom';

import Reactloadable from 'react-loadable';
import PageNotFound from '../pages/PageNotFound';
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
      loader: () => import('../pages/Home'),
      loading: () => <div />
    })
  },
  {
    name: 'GetPayInfo',
    path: '/get-pay-info/:uid',
    exact: true,
    component: Reactloadable({
      loader: () => import('../pages/GetPayInfo'),
      loading: () => <div />
    })
  },
  {
    name: 'PostPayInfo',
    path: '/post-pay-info',
    exact: true,
    component: Reactloadable({
      loader: () => import('../pages/PostPayInfo'),
      loading: () => <div />
    })
  },
  {
    name: 'login',
    path: '/login',
    exact: true,
    component: Reactloadable({
      loader: () => import('../pages/Login'),
      loading: () => <div />
    })
  },
  {
    name: 'userInfo',
    path: '/users/:uid',
    exact: true,
    component: Reactloadable({
      loader: () => import('../pages/GetUserInfo'),
      loading: () => <div>...</div>
    })
  }
];

const routes = () => {
  return (
    <Router>
      <Suspense fallback={<div>Loading...</div>}>
        <header>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/login">Login</Link>
            </li>
            <li>
              <Link to="/users/1">GetUserInfo</Link>
            </li>
            <li>
              <Link to="/get-pay-info/1">GetPayInfo</Link>
            </li>
            <li>
              <Link to="/post-pay-info">PostPayInfo</Link>
            </li>
          </ul>
        </header>
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
          <Route component={PageNotFound}> </Route>
        </Switch>
        <footer>这是你的脚气</footer>
      </Suspense>
    </Router>
  );
};

export default routes;
