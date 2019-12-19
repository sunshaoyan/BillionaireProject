import React, { Component } from 'react'
import moment from 'moment'
import { ConfigProvider, Layout } from 'antd'
import antZhCN from 'antd/lib/locale-provider/zh_CN'

export default class Main extends Component<null, null> {
  antLocale: any

  constructor(props) {
    super(props)
    this.initLocale()
  }

  initLocale = () => {
    this.antLocale = antZhCN
    moment.locale('zh-cn')
  }

  render() {
    return (
      <ConfigProvider locale={this.antLocale}>
        <Layout style={{ height: '100%' }}>
          <div>三万块</div>
        </Layout>
      </ConfigProvider>
    )
  }
}
