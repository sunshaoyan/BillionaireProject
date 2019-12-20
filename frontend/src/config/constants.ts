export const phoneRules = [
  { required: true, message: '请输入手机号码' },
  { pattern: /^1[3|4|5|7|8][0-9]\d{8}$/, message: '请输入正确的手机号码' }
]

export const emailRules = [
  {
    type: 'email',
    required: true,
    message: '请输入正确的邮箱'
  },
  { max: 65, message: '邮箱太长' }
]

export const passwordRules = [
  { required: true, message: '请输入密码' },
  { min: 8, message: '至少为8位' }
]

export const COLOR_POOL = [
  '#0066FF',
  '#AF593E',
  '#01A368',

  '#FF861F',
  '#ED0A3F',
  '#FF3F34',

  '#76D7EA',
  '#8359A3',
  '#03BB85',

  '#FFDF00',
  '#8B8680',
  '#0A6B0D',

  '#8FD8D8',
  '#A36F40',
  '#F653A6',

  '#CA3435',
  '#FFCBA4',
  '#FF99CC',

  '#FA9D5A',
  '#FFAE42',
  '#A78B00',

  '#788193',
  '#514E49',
  '#1164B4',

  '#F4FA9F',
  '#FED8B1',
  '#C32148',

  '#01796F',
  '#E90067',
  '#FF91A4',

  '#404E5A',
  '#6CDAE7',
  '#FFC1CC',

  '#006A93',
  '#867200',
  '#E2B631',

  '#6EEB6E',
  '#FFC800',
  '#CC99BA',

  '#FF007C',
  '#BC6CAC',
  '#DCCCD7',

  '#EBE1C2',
  '#A6AAAE',
  '#B99685',

  '#0086A7',
  '#5E4330',
  '#C8A2C8',

  '#708EB3',
  '#BC8777',
  '#B2592D',

  '#497E48',
  '#6A2963',
  '#E6335F',

  '#00755E',
  '#B5A895',
  '#0048ba',

  '#EED9C4',
  '#C88A65',
  '#FF6E4A',

  '#87421F',
  '#B2BEB5',
  '#926F5B',

  '#00B9FB',
  '#6456B7',
  '#DB5079',

  '#C62D42',
  '#FA9C44',
  '#DA8A67',

  '#FD7C6E',
  '#93CCEA',
  '#FCF686',

  '#503E32',
  '#FF5470',
  '#9DE093',

  '#FF7A00',
  '#4F69C6',
  '#A50B5E',

  '#F0E68C',
  '#FDFF00',
  '#F091A9',

  '#FFFF66',
  '#6F9940',
  '#FC74FD',

  '#652DC1',
  '#D6AEDD',
  '#EE34D2',

  '#BB3385',
  '#6B3FA0',
  '#33CC99',

  '#FFDB00',
  '#87FF2A',
  '#6EEB6E',

  '#FFC800',
  '#CC99BA',
  '#7A89B8',

  '#006A93',
  '#867200',
  '#E2B631'
]
