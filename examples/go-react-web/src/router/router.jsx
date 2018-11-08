import React from 'react';


import Reactloadable from 'react-loadable';

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


export default RouteMap;
