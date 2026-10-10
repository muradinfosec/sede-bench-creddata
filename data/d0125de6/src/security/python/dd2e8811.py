#!/usr/bin/env python
# -*- coding: UTF-8 -*-

"""
@Desc   ：timing attack Against Hash
"""
import hmac
import hashlib

key = "y375003d-17f7-2805-3m57-a93eg7p9"
msg = "Test"

def sign(pre_key, imsg, alg):
    return hmac.new(pre_key, imsg, alg).digest()

def verify(msg, sig):
    return sig == sign(key, msg, hashlib.sha256) #bad
