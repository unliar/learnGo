import React, { Suspense } from 'react';

import { Route, Switch, BrowserRouter as Router, Link } from 'react-router-dom';

import Reactloadable from 'react-loadable';
import PageNotFound from '../pages/PageNotFound';

import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
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
const styles = {
  footer: {
    backgroundColor: '#fff',
    bottom: '0px',
    padding: '40px 0',
    textAlign: 'center',
    width: '100%',
    position: 'fixed'
  }
};
class r extends React.Component {
  state = {
    value: 'index'
  };
  ChangeTabValue = (e, v) => {
    this.setState({ value: v });
    console.log(v);
  };
  render() {
    return (
      <Router>
        <Suspense fallback={<div>Loading...</div>}>
          <AppBar position="static">
            <Tabs value={this.state.value} onChange={this.ChangeTabValue}>
              <Tab label="首页" value="index" component={Link} to="/" />
              <Tab
                label="用户信息"
                value="info"
                component={Link}
                to="/users/1"
              />
              <Tab
                label="二维码"
                value="code"
                component={Link}
                to="/get-pay-info/1"
              />
            </Tabs>
          </AppBar>
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
          <footer style={styles.footer}>曾有容颜惑少年</footer>
        </Suspense>
      </Router>
    );
  }
}

export default r;
