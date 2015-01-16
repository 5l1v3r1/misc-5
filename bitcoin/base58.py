#!/usr/bin/env python
# encoding: utf-8
'''
Library for bitcoin base58 encoding decoding based on the bitcoin base58 implementation:
https://github.com/bitcoin/bitcoin/blob/master/src/base58.cpp
'''
import hashlib

_pszBase58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz";

'''
Decode a base58-encoded string (psz) into a byte list.
psz cannot be None.
'''
def decodeBase58(psz):
    # Strip leading and trailing spaces
    psz = psz.strip()
    
    # Count leading '1's.
    zeroes = 0
    for psz_char in psz:
        if psz_char != '1':
            break
        zeroes += 1
    
    #bytearray base256
    b256 = []
    # Process the characters.
    for psz_char in psz[zeroes:]:
        # Decode base58 character
        carry = _pszBase58.find(psz_char)
        if carry == -1:
            return None

        # Apply "b256 = b256 * 58 + carry".
        for it in reversed(range(len(b256))):
            carry += (b256[it] * 58)
            b256[it] = carry % 256;
            carry /= 256;
        if carry != 0:
            b256.insert(0, carry)
            
    #Add leading zeroes
    for _ in range(zeroes):
        b256.insert(0,0)
    
    binarystring = b''.join([chr(byte) for byte in b256])
    return binarystring


'''
Decode a base58-encoded string (psz) that includes a checksum into a byte list.
Also returns a boolean with the checksum match
psz cannot be None.
'''
def decodeBase58Check(psz):
    ret = decodeBase58(psz)
    if len(ret) < 4:
        return ret, False
    #re-calculate the checksum, ensure it matches the included 4-byte checksum
    hash = hashlib.sha256(hashlib.sha256(ret[:-4]).digest()).digest()
    hashmatch = hash[:4] == ret[-4:]
    return ret, hashmatch


if __name__ == '__main__':
    address = ' 1DTjvhLV6S72NQrSDrCX1GTCb9B3D5pmCB '
    import binascii
    decoded, hashmatch = decodeBase58Check(address)
    assert '0088b028348642ad1bbaa8fcc054273070eda045fe238fa750' == binascii.hexlify(decoded)
    assert hashmatch
    