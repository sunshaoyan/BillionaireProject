import { request } from './request'

const api = {
  toCar: {
    url: '/hackathon/set',
    method: 'post'
  }
}

export const workspace = {
  async sendToCar(list) {
    const response: any = await request(api.toCar, {
      data: JSON.stringify(list)
    })
    return response
  }
}
