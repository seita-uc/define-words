#!/usr/bin/env python
# -*- coding: utf-8 -*- 

import re

text = '‘fathomless'.decode('cp932')
print text
highlighted_word = re.sub(r'[!\'‘"#$%&()\*\+\-\.,\/:;<=>?@\[\\\]^_`{|}~]', "", text)
print highlighted_word
