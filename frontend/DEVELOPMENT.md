### 警告&错误解决方案

1、LocaleProvider 使用不当

```
Warning: [antd: LocaleProvider] `LocaleProvider` is deprecated. Please use `locale` with `ConfigProvider` instead: http://u.ant.design/locale
```

解决方案：将 `LocaleProvider` 替换为 `ConfigProvider`

- [参考资料](https://ant.design/components/locale-provider/)

2、热更新

```
React-Hot-Loader: react-🔥-dom patch is not detected. React 16.6+ features may not work.
```

解决方案

```js
// package.json
{
  "dependencies": {
    "react-hot-loader": "^4.8.3",
    "@hot-loader/react-dom": "^16.8.6"
    // ...
  }
  // ...
}
// webpack.config.js
{
  // ...
  resolve: {
    alias: {
      'react-dom': '@hot-loader/react-dom'
    }
  }
}
```

- [参考资料](https://github.com/gaearon/react-hot-loader/issues/1227)

3、请求

若请求超时则 `axios.interceptors.response` 中间件会报，其中 15000ms 表示在 `webpack.DefinePlugin` 中配置的超时时间 `process.env.REQUEST_TIMEOUT`。

通过 [mocky](https://www.mocky.io/) 来模拟返回超时情况

```
Error: timeout of 15000ms exceeded
  at createError (createError.js:16)
  at XMLHttpRequest.handleTimeout (xhr.js:89)
```

解决方案：

1. 调整超时时间配置
2. 让接口提供方解决
