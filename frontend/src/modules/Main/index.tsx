import React, { Component } from 'react'
import moment from 'moment'
import { ConfigProvider, Layout, Divider } from 'antd'
import antZhCN from 'antd/lib/locale-provider/zh_CN'
import Workspace from '@/modules/WorkSpace'
import CodePreview from '@/modules/CodePreview'
import logo from '@/assets/images/teamlogo.png'
import './index.less'

const { Header, Content, Sider } = Layout

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
          <Header>
            <div className="title">
              <img className="logo" src={logo} />
              <Divider type="vertical" style={{ height: 45 }} />
              <h1>Billionaire Car Kit</h1>
            </div>
          </Header>
          <Layout style={{ padding: '10px 10px 0 10px' }}>
            <Content style={{ padding: '0 10px 0 0' }}>
              <Workspace />
            </Content>
            <Sider width={450}>
              <CodePreview />
            </Sider>
          </Layout>
        </Layout>
      </ConfigProvider>
    )
  }
}
