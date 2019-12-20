// import { request } from './request'

const api = {
  toCar: {
    url: '/datasys/rbac/user/info/',
    method: 'get'
  }
}

export const workspace = {
  async sendToCar(list) {
    console.log(list, api)
  }
}
