extract proxylist from http://free-proxy-list.net

cat proxy.txt|awk -F" " '{print "\"http://"$1":"$2"\","}' > proxy.go
