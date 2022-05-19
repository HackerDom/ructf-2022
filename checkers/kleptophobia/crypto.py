# func performXor(data, result, key []byte) {
# 	for i := 0; i < len(data); i++ {
# 		result[i] = data[i] ^ key[i%len(key)]
# 	}
# }
#
# func Encrypt(data, key []byte) []byte {
# 	resultLen := (len(data)/BlockLength)*BlockLength + BlockLength
# 	result := make([]byte, resultLen)
#
# 	performXor(data, result, key)
#
# 	ln := make([]byte, 4)
# 	binary.LittleEndian.PutUint32(ln, uint32(len(data)))
#
# 	return append(result, ln...)
# }
#
# func Decrypt(data, key []byte) []byte {
# 	ln := binary.LittleEndian.Uint32(data[len(data)-4:])
#
# 	result := make([]byte, len(data)-4)
#
# 	performXor(data[:len(data)-4], result, key)
#
# 	return result[:ln]
# }
import sys
from hashlib import md5


def perform_xor(data, result, key):
    for i in range(len(data)):
        result[i] = data[i] ^ key[i % len(key)]


def decrypt(data, key):
    ln = int.from_bytes(data[len(data) - 4:], "little")

    result = [0] * (len(data) - 4)
    perform_xor(data[:len(data) - 4], result, key)
    return bytes(result[:ln])


def get_hash(data):
    hasher = md5()
    hasher.update(data)
    return hasher.digest()
