import { AuthorizationUtils } from '@horizon/dawn-ui'

const ak = 'lWjiEcWyRDconCmZbp6Du3j7'
const sk = 'KkMC5k7BVtjqThj6Ab9m4q72U4k2wTVc'

const akskUrls = ['/dms/v1/driver_behavior']

const getSign = config => {
  let host = window.location.host
  host = host.split(':')[0]

  const url = config.url
  const options = {
    ak,
    sk,
    method: config.method.toUpperCase(),
    headers: {
      Host: host,
      'Content-Type': 'application/json;charset=UTF-8'
    },
    timestamp: Math.round(Number(new Date()) / 1000)
  }

  return AuthorizationUtils.sign(url, options)
}

export { getSign, akskUrls }
