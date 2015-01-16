#!/usr/bin/env python
# encoding: utf-8
'''
Library for bitcoin base58 encoding decoding based on the bitcoin base58 implementation:
https://github.com/bitcoin/bitcoin/blob/master/src/base58.cpp
'''
import hashlib

_pszBase58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz";

'''
Decode a base58-encoded string (psz).
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
    
    #byte list base256
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
Encode a byte array as a base58-encoded string
'''
def encodeBase58(binarray):
    # Count leading zeroes.
    zeroes = 0
    for byte in binarray:
        if ord(byte) != 0:
            break
        zeroes += 1
    
    #byte list base58
    b58 = []
    # Process the bytes.
    for byte in binarray[zeroes:]:
        carry = ord(byte)
        # Apply "b58 = b58 * 256 + byte".
        for it in reversed(range(len(b58))):
            carry += (b58[it] * 256)
            b58[it] = carry % 58;
            carry /= 58;
        while carry != 0:
            b58.insert(0, carry % 58)
            carry /= 58;
    
    #Add leading zeroes
    for _ in range(zeroes):
        b58.insert(0,0)
    
    result = ''.join([_pszBase58[it] for it in b58])
    return result




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
    encoded = encodeBase58(decoded)
    assert address.strip() == encoded
    