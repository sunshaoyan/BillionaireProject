import requests


class HttpClient(object):
    def __init__(self, config=None):
        config = config or {}
        _config = {
            "connection_retries": 3, "connection_timeout": 300,
            "host": "localhost:8080"}
        for k, v in config.items():
            _config.update({k: v})
        self.config = _config
        self.host = self.config['host']
        self.connection_timeout = self.config['connection_timeout']
        self._session = requests.Session()
        adapter = requests.adapters.HTTPAdapter(
            max_retries=self.config['connection_retries'])
        self._session.mount('http://', adapter)

    def _get_url(self, uri):
        url = 'http://{0}{1}'.format(self.host, uri)
        return url

    def get(self, uri, params=None, headers=None):
        params = params or {}
        headers = headers or {}
        url =  self._get_url("/hackathon/{0}".format(uri))
        print(url)
        try:
            r = self._session.get(
                url,
                params=params,
                headers=headers,
                timeout=self.connection_timeout)
        except Exception as e:
            return ResponseInfo(None, e)
        return ResponseInfo(r)


class ResponseInfo(object):
    code = -1
    message = None
    request_id = None
    data = None
    status_code = -1
    exception = None
    raw_data = None

    def __init__(self, response, exception=None):
        if exception is not None:
            self.exception = exception
            return
        if response is not None:
            self.status_code = response.status_code
            self.raw_data = response.text
            if response.status_code in [200, 401]:
                ret = response.json(
                    encoding='utf-8') if response.text != '' else {}
                self.request_id = ret["request_id"]
                self.code = ret["code"]
                self.message = ret["message"]
                self.data = ret["data"]

    def ok(self):
        return self.status_code == 200

    def connect_failed(self):
        return self.request_id is None

    def __str__(self):
        return ', '.join(['%s:%s' % item for item in self.__dict__.items()])

    def __repr__(self):
        return self.__str__()
