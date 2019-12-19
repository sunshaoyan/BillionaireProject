/**
 * 配置导航
 */
export default [
  {
    path: '/',
    icon: 'dashboard',
    title: 'Dashboard'
  },
  {
    path: '/error',
    icon: 'warning',
    title: '错误页面',
    children: [
      {
        path: '/403',
        title: '403'
      },
      {
        path: '/404',
        title: '404'
      },
      {
        path: '/500',
        title: '500'
      }
    ]
  },
  {
    path: '/table',
    icon: 'unordered-list',
    title: '表格示例',
    children: [
      {
        path: '/table/base',
        title: '基础表格'
      },
      {
        path: '/table/advance',
        title: '高级表格'
      },
      {
        path: '/table/search',
        title: '搜索表格'
      }
    ]
  },
  {
    path: '/form',
    icon: 'form',
    title: '表单示例'
  },
  {
    path: '/device',
    icon: 'tool',
    title: '设备管理',
    children: [
      {
        path: '/device/list',
        title: '设备列表'
      },
      {
        path: '/device/ota',
        title: '升级包'
      },
      {
        path: '/device/task',
        title: '任务'
      }
    ]
  },
  {
    path: '/drive',
    icon: 'car',
    title: '驾驶行为分析'
  }
]
