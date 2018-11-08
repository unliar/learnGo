import React, { Component, Suspense } from 'react';

import { Route, Switch, BrowserRouter as Router, Link } from 'react-router-dom';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import AppBar from '@material-ui/core/AppBar';
import PageNotFound from './pages/PageNotFound';
import RouteMap from "./router/router"
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

class App extends Component {
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

export default App;