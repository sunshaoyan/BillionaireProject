import React, { Component } from 'react'
import moment from 'moment'
import { ConfigProvider, Layout } from 'antd'
import antZhCN from 'antd/lib/locale-provider/zh_CN'
import Workspace from '@/modules/WorkSpace'
import CodePreview from '@/modules/CodePreview'
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
            <div className="logo" />
          </Header>
          <Layout style={{ padding: 20 }}>
            <Content style={{ padding: '0 10px 0 0' }}>
              <Workspace />
            </Content>
            <Sider width={600}>
              <CodePreview />
            </Sider>
          </Layout>
        </Layout>
      </ConfigProvider>
    )
  }
}
