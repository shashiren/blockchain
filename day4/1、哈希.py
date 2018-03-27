import hashlib
sha=hashlib.sha256()
sha.update("000000".encode("utf-8"))
print(sha.hexdigest())