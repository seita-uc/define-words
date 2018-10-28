#!/usr/bin/env python
# -*- coding: utf-8 -*- 

import time
import os
from selenium import webdriver
from selenium.webdriver.chrome.options import Options

email = os.environ['AMAZON_EMAIL']
password = os.environ['AMAZON_PASSWORD']
#options = Options()
#options.add_argument('--headless')
#options.add_argument('--disable-gpu')
#driver = webdriver.Chrome(chrome_options=options)

driver = webdriver.Chrome("/usr/local/bin/chromedriver")

driver.get("https://www.amazon.co.jp/ap/signin?clientContext=357-6962270-7432721&openid.return_to=https%3A%2F%2Fread.amazon.co.jp%2Fkp%2Fnotebook%3Fpurpose%3DNOTEBOOK%26amazonDeviceType%3DA2CLFWBIMVSE9N%26appName%3Dnotebook&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.assoc_handle=amzn_kp_jp&openid.mode=checkid_setup&marketPlaceId=A1VC38T7YXB528&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&pageId=amzn_kp_notebook_us&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0&openid.pape.max_auth_age=1209600&siteState=clientContext%3D355-5895142-4373139%2CsourceUrl%3Dhttps%253A%252F%252Fread.amazon.co.jp%252Fkp%252Fnotebook%253Fpurpose%253DNOTEBOOK%2526amazonDeviceType%253DA2CLFWBIMVSE9N%2526appName%253Dnotebook%2Csignature%3DAKjePsRJ1oCrL9r4ZcbuxhfKZokj3D&language=ja_JP&auth=Customer+is+not+authenticated")
#driver.find_element_by_xpath("//a[@id='nav-link-accountList']/span").click()

driver.find_element_by_id("ap_email").send_keys(email)
#driver.find_element_by_xpath("//input[@id='continue']").click()
driver.find_element_by_id("ap_password").send_keys(password)
driver.find_element_by_id("signInSubmit").click()
driver.find_element_by_id("ap_password").send_keys(password)
driver.find_element_by_id("signInSubmit").click()
#content = driver.find_element_by_xpath("//div[@id='B01ERWBOBU']/span/a/h2").click()
contents = driver.find_elements_by_xpath("//div/span/a/h2[contains(@class, 'kp-notebook-searchable')]")
#contents = driver.find_element_by_class_name("kp-notebook-searchable")
for c in contents:
    #print(c.text)
    #highlights = driver.find_elements_by_classname("//div[contains(@class, 'kp-notebook-highlight-yellow')]/span[@id='highlight']")
    c.click()
    time.sleep(2)
    print(c.text)
    highlights = driver.find_elements_by_id("highlight")
    for h in highlights:
        print(h.text)

#print(driver.page_source)
