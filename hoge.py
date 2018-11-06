#!/usr/bin/env python
# -*- coding: utf-8 -*- 

import requests
from bs4 import BeautifulSoup

url = 'https://ejje.weblio.jp/content/'
page = requests.get(url + "hoksks")
try:
    soup = BeautifulSoup(page.content, 'html.parser')
    res = soup.find(class_='content-explanation ej').get_text()
    print(res)
except:
    print "errrorororo"




