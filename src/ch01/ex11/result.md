演習問題11
=========
* より長い引数リストでfetchallを実行した結果
```
./fetchall http://google.com http://facebook.co
m http://youtube.com http://yahoo.com http://amazon.com http://wikipedia
.org http://google.co.in http://twitter.com http://live.com http://notresponse.site
Get http://notresponse.site: dial tcp: lookup notresponse.site: no such host
0.32s   21280 http://google.co.in
0.33s   19412 http://google.com
1.47s  251390 http://twitter.com
1.76s  393937 http://youtube.com
1.96s   54537 http://wikipedia.org
2.10s  411081 http://yahoo.com
2.48s  380514 http://amazon.com
2.62s    9656 http://live.com
2.87s   72796 http://facebook.com
2.87s elapsed
```
