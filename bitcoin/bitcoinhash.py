#!/usr/bin/env python
# encoding: utf-8

import hashlib

'''
A hasher for Bitcoin's 256-bit hash (double SHA-256).
'''
def hash256(data):
    return hashlib.sha256(hashlib.sha256(data).digest()).digest()

'''
A hasher for Bitcoin's 160-bit hash(SHA-256 + RIPEMD-160)).
'''
def hash160(data):
    sha256hash = hashlib.sha256(data).digest()
    return hashlib.new('ripemd160', sha256hash).digest()


if __name__ == '__main__':
    pass