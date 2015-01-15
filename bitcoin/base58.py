#!/usr/bin/env python
# encoding: utf-8
'''
Library for bitcoin base58 encoding decoding based on the bitcoin base58 implementation:
https://github.com/bitcoin/bitcoin/blob/master/src/base58.cpp
'''


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
    
    return b256

def hexEncodeByteList(b256):
    return ''.join(''.join('{:02x}'.format(x) for x in b256))

if __name__ == '__main__':
    address = ' 1DTjvhLV6S72NQrSDrCX1GTCb9B3D5pmCB '
    decoded = hexEncodeByteList(decodeBase58(address))
    print decoded
    assert '0088b028348642ad1bbaa8fcc054273070eda045fe238fa750' == decoded