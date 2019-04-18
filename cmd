extract proxylist from http://free-proxy-list.net

cat proxy.txt|awk -F" " '{print "\"http://"$1":"$2"\","}' > proxy.go





http://free-proxy.cz/en/proxylist/country/TH/http/ping/all

echo "package main" > proxy2.go
echo "var proxy []string = []string{" >> proxy2.go

cat proxy2.txt |awk -F" " '{print "\"http://"$1"\","}' >> proxy.go

echo "}" >> proxy2.go
