# Troubleshooting about Version

## linux에서 버전 관리

### update-alternatives 사용

```
$ ll /usr/bin/go
lrwxrwxrwx 1 root root 21 Apr 16  2020 /usr/bin/go -> ../lib/go-1.13/bin/go*

$ sudo update-alternatives --install /usr/bin/go go /usr/lib/go-1.14/bin/go 1014
```

### 결과
```
lrwxrwxrwx 1 root root 20 Nov  6 01:07 /usr/bin/go -> /etc/alternatives/go*
```

- [alternatives 명령어 알아보기](https://skyoo2003.github.io/post/2017/03/17/what-is-alternatives-command)