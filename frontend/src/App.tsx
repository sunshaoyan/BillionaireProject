import React from 'react'
import { BrowserRouter as Router, Switch } from 'react-router-dom'
import { Provider } from 'react-redux'
import { hot } from 'react-hot-loader/root'
import { renderRouter, mainRouters } from '@/routers'
import { store } from './store'
import '@/styles/global.less'

const App = () => (
  <Provider store={store}>
    <Router>
      <Switch>{renderRouter(mainRouters)}</Switch>
    </Router>
  </Provider>
)

export default hot(App)
