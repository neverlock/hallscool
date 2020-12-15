extract proxylist from http://free-proxy-list.net
http://www.freeproxylists.net/?c=TH&pt=&pr=&a%5B%5D=0&a%5B%5D=1&a%5B%5D=2&u=0
http://www.freeproxylists.net/
https://github.com/ShiftyTR/Proxy-List

cat proxy.txt|awk -F" " '{print "\"http://"$1":"$2"\","}' > proxy.go

cat proxy.txt |awk -F" " '{print $1":"$2}' >> proxy2.txt



http://free-proxy.cz/en/proxylist/country/TH/http/ping/all

echo "package main" > proxy1.go
echo "var listProxy []string = []string{" >> proxy1.go

//for http proxy
cat proxy.txt |awk -F" " '{print "\"http://"$1"\","}' >> proxy1.go

echo "}" >> proxy1.go







echo "package main" > proxy2.go
echo "var listProxy []string = []string{" >> proxy2.go

//for socks proxy
cat proxy.txt |awk -F" " '{print "\""$1"\","}' >> proxy2.go

echo "}" >> proxy2.go
