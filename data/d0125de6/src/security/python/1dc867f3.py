#!/usr/bin/env python
# -*- coding: UTF-8 -*-

"""
@Desc   ：preventing timing attack Against Hash
"""
import hmac
import hashlib

key = "f820609h-56f2-7994-5j93-z43mm0b4"
msg = "Test"

def sign(pre_key, imsg, alg):
    return hmac.new(pre_key, imsg, alg).digest()

def verify(msg, sig):
    return hmac.compare_digest(sig, sign(key, msg, hashlib.sha256)) #good
