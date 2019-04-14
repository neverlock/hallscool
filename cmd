extract proxylist from http://free-proxy-list.net

cat proxy.txt|awk -F" " '{print "\"http://"$1":"$2"\","}' > proxy.go





http://free-proxy.cz/en/proxylist/country/TH/http/ping/all

cat proxy.txt |awk -F" " '{print "\"http://"$1"\","}' > proxy.go
