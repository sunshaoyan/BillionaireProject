/**
 * author: shuisheng.zhang
 * createTime: 2018/9/28 下午2:38
 * description: 将 zh-CN.json 新增的字段同步到其他语言文件
 */
const fs = require('fs')
const path = require('path')

const basePath = path.join(__dirname, '../src/locales')

const sourceFile = JSON.parse(fs.readFileSync(`${basePath}/zh-CN.json`, 'utf8'))

const targetFiles = fs
  .readdirSync(basePath)
  .filter(item => item !== 'zh-CN.json')

const addItem = (source, target) => {
  Object.entries(source).forEach(item => {
    const key = item[0]
    const value = item[1]
    if (!target[key]) {
      target[key] = value
    } else if (value instanceof Object) {
      addItem(value, target[key])
    }
  })
}

const removeItem = (source, target) => {
  Object.entries(target).forEach(item => {
    const key = item[0]
    const value = item[1]
    if (!source[key]) {
      delete target[key]
    } else if (value instanceof Object) {
      removeItem(source[key], value)
    }
  })
}

targetFiles.forEach(item => {
  const targetFile = JSON.parse(fs.readFileSync(`${basePath}/${item}`, 'utf8'))
  addItem(sourceFile, targetFile)
  removeItem(sourceFile, targetFile)

  fs.writeFileSync(
    `${basePath}/${item}`,
    JSON.stringify(targetFile, null, 2),
    'utf8'
  )
})

fs.writeFileSync(
  `${basePath}/zh-CN.json`,
  JSON.stringify(sourceFile, null, 2),
  'utf8'
)
