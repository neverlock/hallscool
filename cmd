extract proxylist from http://free-proxy-list.net
http://www.freeproxylists.net/?c=TH&pt=&pr=&a%5B%5D=0&a%5B%5D=1&a%5B%5D=2&u=0
http://www.freeproxylists.net/

cat proxy.txt|awk -F" " '{print "\"http://"$1":"$2"\","}' > proxy.go

cat proxy.txt |awk -F" " '{print $1":"$2}' >> proxy2.txt



http://free-proxy.cz/en/proxylist/country/TH/http/ping/all

echo "package main" > proxy2.go
echo "var proxy []string = []string{" >> proxy2.go

cat proxy2.txt |awk -F" " '{print "\"http://"$1"\","}' >> proxy2.go

echo "}" >> proxy2.go
