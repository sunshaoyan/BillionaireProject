import axios from 'axios'
import pathToRegexp from 'path-to-regexp'
import { message } from 'antd'

message.config({ maxCount: 1 })

// 请求返回成功后，对返回数据进行处理
const processResponse = response => {
  if (response && response.data) {
    if (response.data.error === 2) {
      message.error('未登录，请先登录')
      window.location.href = '/data-label/login'
    } else {
      return response.data
    }
  }
}

// 请求返回错误后，对错误进行处理
const processResponseError = error => {
  Promise.reject(error)
}

// 对返回进行拦截
axios.interceptors.response.use(
  response => {
    return processResponse(response)
  },
  error => {
    return processResponseError(error)
  }
)

/**
 * 发送请求
 * /api/v1/foo/bar/:id => request(api.xxx, data, {id: 1})
 * /api/v1/foo/bar?id=1 => request(api.xxx, {id: 1})
 * @param api 对象 {url, method}，见 api.ts
 * @param data 需要发送的数据，可以是 querystring(get) 也可以是 body(post, put, ...)
 * @param urlParams 组成 url 的变量
 * @param options 额外的配置选项，同 axios
 */
const request: any = (
  api: { url: string; method: string },
  data?: object,
  urlParams?: object,
  options?: object
) => {
  if (urlParams) {
    api = {
      ...api,
      url: pathToRegexp.compile(api.url)(urlParams)
    }
  }

  let config: any = {
    ...options,
    ...api
  }

  if (api.method === 'get') {
    config = {
      ...config,
      params: data
    }
  } else {
    config = {
      ...config,
      data
    }
  }
  return axios(config)
}

export { request }
