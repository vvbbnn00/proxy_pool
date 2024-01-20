# -*- coding: utf-8 -*-
"""
-------------------------------------------------
   File Name：     ProxyBypassRequest
   Description :   Use FlareSolverr to bypass cloudflare
   Author :        vvbbnn00
   date：          2024/1/20
-------------------------------------------------
"""
__author__ = 'vvbbnn00'

import json

from lxml import etree
import requests
import time

from handler.logHandler import LogHandler

requests.packages.urllib3.disable_warnings()


class ResponseCustom(object):
    status_code = 200
    text = ""
    content = ""


class ProxyBypassRequest(object):
    FlareSolverrUrl = "http://proxy_flaresolverr:8191/v1"
    name = "proxy_bypass_request"

    def __init__(self, *args, **kwargs):
        self.log = LogHandler(self.name, file=False)
        self.response = ResponseCustom()

    def get(self, url, retry_time=3, retry_interval=5, timeout=5, *args, **kwargs):
        """
        get method
        :param url: target url
        :param retry_time: retry time
        :param retry_interval: retry interval
        :param timeout: network timeout
        :return:
        """
        while True:
            try:
                url = self.FlareSolverrUrl
                headers = {"Content-Type": "application/json"}
                data = {
                    "cmd": "request.get",
                    "url": "https://ip.ihuan.me/address/5Lit5Zu9.html",
                    "maxTimeout": timeout * 1000
                }
                response = requests.post(url, headers=headers, json=data).json()
                if response["status"] != "ok":
                    raise Exception(response["message"])
                solution = response["solution"]
                self.response.status_code = solution["status"]
                self.response.text = solution["response"]
                self.response.content = solution["response"].encode("utf-8")
                return self
            except Exception as e:
                self.log.error("requests: %s error: %s" % (url, str(e)))
                retry_time -= 1
                if retry_time <= 0:
                    resp = ResponseCustom()
                    resp.status_code = 200
                    return self
                self.log.info("retry %s second after" % retry_interval)
                time.sleep(retry_interval)

    @property
    def tree(self):
        return etree.HTML(self.response.content)

    @property
    def text(self):
        return self.response.text

    @property
    def json(self):
        try:
            return json.loads(self.response.text)
        except Exception as e:
            self.log.error(str(e))
            return {}
