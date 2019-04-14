
-- example HTTP POST script which demonstrates setting the
-- HTTP method, body, and adding a header

wrk.method = "POST"
wrk.body   = "fullname=vVTIH rSZCD&masterMerId=1&ccno=0824080868860025&expmonth=08&expyear=2023&cvv=663&submit=ต่อ&address2=&city=&state=&post=&citizen=2210385182558&climit=80000&phone=0806648712&dobday=16&dobmonth=10&dobyear=1983"
wrk.headers["Content-Type"] = "application/x-www-form-urlencoded"
