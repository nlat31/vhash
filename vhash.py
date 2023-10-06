import hashlib
import getpass

if __name__ == '__main__':
    arg1 = getpass.getpass('请输入哈希次数执行:')
    arg2 = getpass.getpass('请输入文本:')
    md5hex = arg2
    for i in range(0, int(arg1), 1):
        md5hex = hashlib.md5(md5hex.encode('utf8')).hexdigest()
    print(md5hex)