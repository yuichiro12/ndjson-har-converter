package main

import (
	"bytes"
	"testing"
)


const HAR = `{
  "log": {
    "version": "1.2",
    "creator": {
      "name": "WebInspector",
      "version": "537.36"
    },
    "pages": [
      {
        "startedDateTime": "2016-10-14T17:28:53.866Z",
        "id": "page_1",
        "title": "https://golang.org/",
        "pageTimings": {
          "onContentLoad": 518,
          "onLoad": 1523
        }
      }
    ],
    "entries": [
      {
        "startedDateTime": "2016-10-14T17:28:53.866Z",
        "time": 74,
        "request": {
          "method": "GET",
          "url": "https://golang.org/",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Pragma",
              "value": "no-cache"
            },
            {
              "name": "Accept-Encoding",
              "value": "gzip, deflate, sdch, br"
            },
            {
              "name": "Host",
              "value": "golang.org"
            },
            {
              "name": "Accept-Language",
              "value": "en-US,en;q=0.8,zh-CN;q=0.6"
            },
            {
              "name": "Upgrade-Insecure-Requests",
              "value": "1"
            },
            {
              "name": "User-Agent",
              "value": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"
            },
            {
              "name": "Accept",
              "value": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
            },
            {
              "name": "Cache-Control",
              "value": "no-cache"
            },
            {
              "name": "Connection",
              "value": "keep-alive"
            }
          ],
          "queryString": [],
          "cookies": [],
          "headersSize": 427,
          "bodySize": 0
        },
        "response": {
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Strict-Transport-Security",
              "value": "max-age=31536000; preload"
            },
            {
              "name": "Content-Encoding",
              "value": "gzip"
            },
            {
              "name": "Server",
              "value": "Google Frontend"
            },
            {
              "name": "Date",
              "value": "Fri, 14 Oct 2016 17:28:54 GMT"
            },
            {
              "name": "Vary",
              "value": "Accept-Encoding"
            },
            {
              "name": "Content-Type",
              "value": "text/html; charset=utf-8"
            },
            {
              "name": "X-Cloud-Trace-Context",
              "value": "dfba54934a6d4feaeac42e060a739c6d"
            },
            {
              "name": "Cache-Control",
              "value": "private"
            },
            {
              "name": "Transfer-Encoding",
              "value": "chunked"
            },
            {
              "name": "Alt-Svc",
              "value": "quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""
            }
          ],
          "cookies": [],
          "content": {
            "size": 7902,
            "mimeType": "text/html",
            "compression": 4593
          },
          "redirectURL": "",
          "headersSize": 385,
          "bodySize": 3309,
          "_transferSize": 3694
        },
        "cache": {},
        "timings": {
          "blocked": 1,
          "dns": -1,
          "connect": -1,
          "send": 0,
          "wait": 69,
          "receive": 2,
          "ssl": -1
        },
        "serverIPAddress": "216.58.219.177",
        "connection": "69617",
        "pageref": "page_1"
      },
      {
        "startedDateTime": "2016-10-14T17:28:53.972Z",
        "time": 74,
        "request": {
          "method": "GET",
          "url": "https://golang.org/lib/godoc/style.css",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Pragma",
              "value": "no-cache"
            },
            {
              "name": "Accept-Encoding",
              "value": "gzip, deflate, sdch, br"
            },
            {
              "name": "Host",
              "value": "golang.org"
            },
            {
              "name": "Accept-Language",
              "value": "en-US,en;q=0.8,zh-CN;q=0.6"
            },
            {
              "name": "User-Agent",
              "value": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"
            },
            {
              "name": "Accept",
              "value": "text/css,*/*;q=0.1"
            },
            {
              "name": "Referer",
              "value": "https://golang.org/"
            },
            {
              "name": "Connection",
              "value": "keep-alive"
            },
            {
              "name": "Cache-Control",
              "value": "no-cache"
            }
          ],
          "queryString": [],
          "cookies": [],
          "headersSize": 390,
          "bodySize": 0
        },
        "response": {
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Strict-Transport-Security",
              "value": "max-age=31536000; preload"
            },
            {
              "name": "Content-Encoding",
              "value": "gzip"
            },
            {
              "name": "Server",
              "value": "Google Frontend"
            },
            {
              "name": "Date",
              "value": "Fri, 14 Oct 2016 17:28:54 GMT"
            },
            {
              "name": "Vary",
              "value": "Accept-Encoding"
            },
            {
              "name": "Content-Type",
              "value": "text/css; charset=utf-8"
            },
            {
              "name": "X-Cloud-Trace-Context",
              "value": "81465de5d9dfde5c424f59c5e92c2afe"
            },
            {
              "name": "Cache-Control",
              "value": "private"
            },
            {
              "name": "Transfer-Encoding",
              "value": "chunked"
            },
            {
              "name": "Accept-Ranges",
              "value": "bytes"
            },
            {
              "name": "Alt-Svc",
              "value": "quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""
            }
          ],
          "cookies": [],
          "content": {
            "size": 12116,
            "mimeType": "text/css",
            "compression": 9067
          },
          "redirectURL": "",
          "headersSize": 406,
          "bodySize": 3049,
          "_transferSize": 3455
        },
        "cache": {},
        "timings": {
          "blocked": 3,
          "dns": -1,
          "connect": -1,
          "send": 0,
          "wait": 68,
          "receive": 1,
          "ssl": -1
        },
        "serverIPAddress": "216.58.219.177",
        "connection": "69617",
        "pageref": "page_1"
      },
      {
        "startedDateTime": "2016-10-14T17:28:53.972Z",
        "time": 78,
        "request": {
          "method": "GET",
          "url": "https://golang.org/lib/godoc/jquery.treeview.css",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Pragma",
              "value": "no-cache"
            },
            {
              "name": "Accept-Encoding",
              "value": "gzip, deflate, sdch, br"
            },
            {
              "name": "Host",
              "value": "golang.org"
            },
            {
              "name": "Accept-Language",
              "value": "en-US,en;q=0.8,zh-CN;q=0.6"
            },
            {
              "name": "User-Agent",
              "value": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"
            },
            {
              "name": "Accept",
              "value": "text/css,*/*;q=0.1"
            },
            {
              "name": "Referer",
              "value": "https://golang.org/"
            },
            {
              "name": "Connection",
              "value": "keep-alive"
            },
            {
              "name": "Cache-Control",
              "value": "no-cache"
            }
          ],
          "queryString": [],
          "cookies": [],
          "headersSize": 400,
          "bodySize": 0
        },
        "response": {
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Strict-Transport-Security",
              "value": "max-age=31536000; preload"
            },
            {
              "name": "Content-Encoding",
              "value": "gzip"
            },
            {
              "name": "Server",
              "value": "Google Frontend"
            },
            {
              "name": "Date",
              "value": "Fri, 14 Oct 2016 17:28:54 GMT"
            },
            {
              "name": "Vary",
              "value": "Accept-Encoding"
            },
            {
              "name": "Content-Type",
              "value": "text/css; charset=utf-8"
            },
            {
              "name": "X-Cloud-Trace-Context",
              "value": "e207f5602faa76b1ed575b7d0653d2b6"
            },
            {
              "name": "Cache-Control",
              "value": "private"
            },
            {
              "name": "Transfer-Encoding",
              "value": "chunked"
            },
            {
              "name": "Accept-Ranges",
              "value": "bytes"
            },
            {
              "name": "Alt-Svc",
              "value": "quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""
            }
          ],
          "cookies": [],
          "content": {
            "size": 2755,
            "mimeType": "text/css",
            "compression": 1927
          },
          "redirectURL": "",
          "headersSize": 406,
          "bodySize": 828,
          "_transferSize": 1234
        },
        "cache": {},
        "timings": {
          "blocked": 7,
          "dns": -1,
          "connect": -1,
          "send": 0,
          "wait": 70,
          "receive": 0,
          "ssl": -1
        },
        "serverIPAddress": "216.58.219.177",
        "connection": "69532",
        "pageref": "page_1"
      },
      {
        "startedDateTime": "2016-10-14T17:28:53.972Z",
        "time": 141,
        "request": {
          "method": "GET",
          "url": "https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js",
          "httpVersion": "unknown",
          "headers": [
            {
              "name": "path",
              "value": "/ajax/libs/jquery/1.8.2/jquery.min.js"
            },
            {
              "name": "pragma",
              "value": "no-cache"
            },
            {
              "name": "accept-encoding",
              "value": "gzip, deflate, sdch, br"
            },
            {
              "name": "accept-language",
              "value": "en-US,en;q=0.8,zh-CN;q=0.6"
            },
            {
              "name": "user-agent",
              "value": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"
            },
            {
              "name": "x-chrome-uma-enabled",
              "value": "1"
            },
            {
              "name": "accept",
              "value": "*/*"
            },
            {
              "name": "cache-control",
              "value": "no-cache"
            },
            {
              "name": "authority",
              "value": "ajax.googleapis.com"
            },
            {
              "name": "referer",
              "value": "https://golang.org/"
            },
            {
              "name": "scheme",
              "value": "https"
            },
            {
              "name": "x-client-data",
              "value": "CIq2yQEIo7bJAQjEtskBCKmZygE="
            },
            {
              "name": "method",
              "value": "GET"
            }
          ],
          "queryString": [],
          "cookies": [],
          "headersSize": -1,
          "bodySize": 0
        },
        "response": {
          "status": 200,
          "statusText": "",
          "httpVersion": "unknown",
          "headers": [
            {
              "name": "date",
              "value": "Wed, 12 Oct 2016 19:59:10 GMT"
            },
            {
              "name": "content-encoding",
              "value": "gzip"
            },
            {
              "name": "x-content-type-options",
              "value": "nosniff"
            },
            {
              "name": "age",
              "value": "163784"
            },
            {
              "name": "status",
              "value": "200"
            },
            {
              "name": "alt-svc",
              "value": "quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""
            },
            {
              "name": "content-length",
              "value": "33397"
            },
            {
              "name": "x-xss-protection",
              "value": "1; mode=block"
            },
            {
              "name": "last-modified",
              "value": "Fri, 16 Oct 2015 18:27:31 GMT"
            },
            {
              "name": "server",
              "value": "sffe"
            },
            {
              "name": "vary",
              "value": "Accept-Encoding"
            },
            {
              "name": "content-type",
              "value": "text/javascript; charset=UTF-8"
            },
            {
              "name": "access-control-allow-origin",
              "value": "*"
            },
            {
              "name": "cache-control",
              "value": "public, max-age=31536000, stale-while-revalidate=2592000"
            },
            {
              "name": "timing-allow-origin",
              "value": "*"
            },
            {
              "name": "expires",
              "value": "Thu, 12 Oct 2017 19:59:10 GMT"
            }
          ],
          "cookies": [],
          "content": {
            "size": 93435,
            "mimeType": "text/javascript"
          },
          "redirectURL": "",
          "headersSize": -1,
          "bodySize": -1,
          "_transferSize": 33724
        },
        "cache": {},
        "timings": {
          "blocked": 7,
          "dns": 0,
          "connect": 8,
          "send": 0,
          "wait": 32,
          "receive": 11,
          "ssl": 34
        },
        "serverIPAddress": "173.194.209.95",
        "connection": "69914",
        "pageref": "page_1"
      },
      {
        "startedDateTime": "2016-10-14T17:28:53.974Z",
        "time": 121,
        "request": {
          "method": "GET",
          "url": "https://golang.org/lib/godoc/jquery.treeview.js",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Pragma",
              "value": "no-cache"
            },
            {
              "name": "Accept-Encoding",
              "value": "gzip, deflate, sdch, br"
            },
            {
              "name": "Host",
              "value": "golang.org"
            },
            {
              "name": "Accept-Language",
              "value": "en-US,en;q=0.8,zh-CN;q=0.6"
            },
            {
              "name": "User-Agent",
              "value": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"
            },
            {
              "name": "Accept",
              "value": "*/*"
            },
            {
              "name": "Referer",
              "value": "https://golang.org/"
            },
            {
              "name": "Connection",
              "value": "keep-alive"
            },
            {
              "name": "Cache-Control",
              "value": "no-cache"
            }
          ],
          "queryString": [],
          "cookies": [],
          "headersSize": 384,
          "bodySize": 0
        },
        "response": {
          "status": 200,
          "statusText": "OK",
          "httpVersion": "HTTP/1.1",
          "headers": [
            {
              "name": "Strict-Transport-Security",
              "value": "max-age=31536000; preload"
            },
            {
              "name": "Content-Encoding",
              "value": "gzip"
            },
            {
              "name": "Server",
              "value": "Google Frontend"
            },
            {
              "name": "Date",
              "value": "Fri, 14 Oct 2016 17:28:54 GMT"
            },
            {
              "name": "Vary",
              "value": "Accept-Encoding"
            },
            {
              "name": "Content-Type",
              "value": "application/x-javascript"
            },
            {
              "name": "X-Cloud-Trace-Context",
              "value": "0f3abeb58802e643e1fb46121e3ec9a9"
            },
            {
              "name": "Cache-Control",
              "value": "private"
            },
            {
              "name": "Transfer-Encoding",
              "value": "chunked"
            },
            {
              "name": "Accept-Ranges",
              "value": "bytes"
            },
            {
              "name": "Alt-Svc",
              "value": "quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""
            }
          ],
          "cookies": [],
          "content": {
            "size": 8264,
            "mimeType": "application/x-javascript",
            "compression": 5662
          },
          "redirectURL": "",
          "headersSize": 407,
          "bodySize": 2602,
          "_transferSize": 3009
        },
        "cache": {},
        "timings": {
          "blocked": 9,
          "dns": 0,
          "connect": 45,
          "send": 0,
          "wait": 63,
          "receive": 1,
          "ssl": 7
        },
        "serverIPAddress": "216.58.219.177",
        "connection": "69925",
        "pageref": "page_1"
      }
    ]
  }
}
`

const ENTRIES = `{"pageref":"page_1","startedDateTime":"2016-10-14T17:28:53.866Z","time":74,"request":{"method":"GET","url":"https://golang.org/","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Pragma","value":"no-cache"},{"name":"Accept-Encoding","value":"gzip, deflate, sdch, br"},{"name":"Host","value":"golang.org"},{"name":"Accept-Language","value":"en-US,en;q=0.8,zh-CN;q=0.6"},{"name":"Upgrade-Insecure-Requests","value":"1"},{"name":"User-Agent","value":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"},{"name":"Accept","value":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},{"name":"Cache-Control","value":"no-cache"},{"name":"Connection","value":"keep-alive"}],"queryString":[],"headersSize":427,"bodySize":0},"response":{"status":200,"statusText":"OK","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Strict-Transport-Security","value":"max-age=31536000; preload"},{"name":"Content-Encoding","value":"gzip"},{"name":"Server","value":"Google Frontend"},{"name":"Date","value":"Fri, 14 Oct 2016 17:28:54 GMT"},{"name":"Vary","value":"Accept-Encoding"},{"name":"Content-Type","value":"text/html; charset=utf-8"},{"name":"X-Cloud-Trace-Context","value":"dfba54934a6d4feaeac42e060a739c6d"},{"name":"Cache-Control","value":"private"},{"name":"Transfer-Encoding","value":"chunked"},{"name":"Alt-Svc","value":"quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""}],"content":{"size":7902,"compression":4593,"mimeType":"text/html"},"redirectURL":"","headersSize":385,"bodySize":3309},"cache":{},"timings":{"blocked":1,"dns":-1,"connect":-1,"send":0,"wait":69,"receive":2,"ssl":-1},"serverIPAddress":"216.58.219.177","connection":"69617"}
{"pageref":"page_1","startedDateTime":"2016-10-14T17:28:53.972Z","time":74,"request":{"method":"GET","url":"https://golang.org/lib/godoc/style.css","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Pragma","value":"no-cache"},{"name":"Accept-Encoding","value":"gzip, deflate, sdch, br"},{"name":"Host","value":"golang.org"},{"name":"Accept-Language","value":"en-US,en;q=0.8,zh-CN;q=0.6"},{"name":"User-Agent","value":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"},{"name":"Accept","value":"text/css,*/*;q=0.1"},{"name":"Referer","value":"https://golang.org/"},{"name":"Connection","value":"keep-alive"},{"name":"Cache-Control","value":"no-cache"}],"queryString":[],"headersSize":390,"bodySize":0},"response":{"status":200,"statusText":"OK","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Strict-Transport-Security","value":"max-age=31536000; preload"},{"name":"Content-Encoding","value":"gzip"},{"name":"Server","value":"Google Frontend"},{"name":"Date","value":"Fri, 14 Oct 2016 17:28:54 GMT"},{"name":"Vary","value":"Accept-Encoding"},{"name":"Content-Type","value":"text/css; charset=utf-8"},{"name":"X-Cloud-Trace-Context","value":"81465de5d9dfde5c424f59c5e92c2afe"},{"name":"Cache-Control","value":"private"},{"name":"Transfer-Encoding","value":"chunked"},{"name":"Accept-Ranges","value":"bytes"},{"name":"Alt-Svc","value":"quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""}],"content":{"size":12116,"compression":9067,"mimeType":"text/css"},"redirectURL":"","headersSize":406,"bodySize":3049},"cache":{},"timings":{"blocked":3,"dns":-1,"connect":-1,"send":0,"wait":68,"receive":1,"ssl":-1},"serverIPAddress":"216.58.219.177","connection":"69617"}
{"pageref":"page_1","startedDateTime":"2016-10-14T17:28:53.972Z","time":78,"request":{"method":"GET","url":"https://golang.org/lib/godoc/jquery.treeview.css","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Pragma","value":"no-cache"},{"name":"Accept-Encoding","value":"gzip, deflate, sdch, br"},{"name":"Host","value":"golang.org"},{"name":"Accept-Language","value":"en-US,en;q=0.8,zh-CN;q=0.6"},{"name":"User-Agent","value":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"},{"name":"Accept","value":"text/css,*/*;q=0.1"},{"name":"Referer","value":"https://golang.org/"},{"name":"Connection","value":"keep-alive"},{"name":"Cache-Control","value":"no-cache"}],"queryString":[],"headersSize":400,"bodySize":0},"response":{"status":200,"statusText":"OK","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Strict-Transport-Security","value":"max-age=31536000; preload"},{"name":"Content-Encoding","value":"gzip"},{"name":"Server","value":"Google Frontend"},{"name":"Date","value":"Fri, 14 Oct 2016 17:28:54 GMT"},{"name":"Vary","value":"Accept-Encoding"},{"name":"Content-Type","value":"text/css; charset=utf-8"},{"name":"X-Cloud-Trace-Context","value":"e207f5602faa76b1ed575b7d0653d2b6"},{"name":"Cache-Control","value":"private"},{"name":"Transfer-Encoding","value":"chunked"},{"name":"Accept-Ranges","value":"bytes"},{"name":"Alt-Svc","value":"quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""}],"content":{"size":2755,"compression":1927,"mimeType":"text/css"},"redirectURL":"","headersSize":406,"bodySize":828},"cache":{},"timings":{"blocked":7,"dns":-1,"connect":-1,"send":0,"wait":70,"receive":0,"ssl":-1},"serverIPAddress":"216.58.219.177","connection":"69532"}
{"pageref":"page_1","startedDateTime":"2016-10-14T17:28:53.972Z","time":141,"request":{"method":"GET","url":"https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js","httpVersion":"unknown","cookies":[],"headers":[{"name":"path","value":"/ajax/libs/jquery/1.8.2/jquery.min.js"},{"name":"pragma","value":"no-cache"},{"name":"accept-encoding","value":"gzip, deflate, sdch, br"},{"name":"accept-language","value":"en-US,en;q=0.8,zh-CN;q=0.6"},{"name":"user-agent","value":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"},{"name":"x-chrome-uma-enabled","value":"1"},{"name":"accept","value":"*/*"},{"name":"cache-control","value":"no-cache"},{"name":"authority","value":"ajax.googleapis.com"},{"name":"referer","value":"https://golang.org/"},{"name":"scheme","value":"https"},{"name":"x-client-data","value":"CIq2yQEIo7bJAQjEtskBCKmZygE="},{"name":"method","value":"GET"}],"queryString":[],"headersSize":-1,"bodySize":0},"response":{"status":200,"statusText":"","httpVersion":"unknown","cookies":[],"headers":[{"name":"date","value":"Wed, 12 Oct 2016 19:59:10 GMT"},{"name":"content-encoding","value":"gzip"},{"name":"x-content-type-options","value":"nosniff"},{"name":"age","value":"163784"},{"name":"status","value":"200"},{"name":"alt-svc","value":"quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""},{"name":"content-length","value":"33397"},{"name":"x-xss-protection","value":"1; mode=block"},{"name":"last-modified","value":"Fri, 16 Oct 2015 18:27:31 GMT"},{"name":"server","value":"sffe"},{"name":"vary","value":"Accept-Encoding"},{"name":"content-type","value":"text/javascript; charset=UTF-8"},{"name":"access-control-allow-origin","value":"*"},{"name":"cache-control","value":"public, max-age=31536000, stale-while-revalidate=2592000"},{"name":"timing-allow-origin","value":"*"},{"name":"expires","value":"Thu, 12 Oct 2017 19:59:10 GMT"}],"content":{"size":93435,"compression":0,"mimeType":"text/javascript"},"redirectURL":"","headersSize":-1,"bodySize":-1},"cache":{},"timings":{"blocked":7,"connect":8,"send":0,"wait":32,"receive":11,"ssl":34},"serverIPAddress":"173.194.209.95","connection":"69914"}
{"pageref":"page_1","startedDateTime":"2016-10-14T17:28:53.974Z","time":121,"request":{"method":"GET","url":"https://golang.org/lib/godoc/jquery.treeview.js","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Pragma","value":"no-cache"},{"name":"Accept-Encoding","value":"gzip, deflate, sdch, br"},{"name":"Host","value":"golang.org"},{"name":"Accept-Language","value":"en-US,en;q=0.8,zh-CN;q=0.6"},{"name":"User-Agent","value":"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36"},{"name":"Accept","value":"*/*"},{"name":"Referer","value":"https://golang.org/"},{"name":"Connection","value":"keep-alive"},{"name":"Cache-Control","value":"no-cache"}],"queryString":[],"headersSize":384,"bodySize":0},"response":{"status":200,"statusText":"OK","httpVersion":"HTTP/1.1","cookies":[],"headers":[{"name":"Strict-Transport-Security","value":"max-age=31536000; preload"},{"name":"Content-Encoding","value":"gzip"},{"name":"Server","value":"Google Frontend"},{"name":"Date","value":"Fri, 14 Oct 2016 17:28:54 GMT"},{"name":"Vary","value":"Accept-Encoding"},{"name":"Content-Type","value":"application/x-javascript"},{"name":"X-Cloud-Trace-Context","value":"0f3abeb58802e643e1fb46121e3ec9a9"},{"name":"Cache-Control","value":"private"},{"name":"Transfer-Encoding","value":"chunked"},{"name":"Accept-Ranges","value":"bytes"},{"name":"Alt-Svc","value":"quic=\":443\"; ma=2592000; v=\"36,35,34,33,32\""}],"content":{"size":8264,"compression":5662,"mimeType":"application/x-javascript"},"redirectURL":"","headersSize":407,"bodySize":2602},"cache":{},"timings":{"blocked":9,"connect":45,"send":0,"wait":63,"receive":1,"ssl":7},"serverIPAddress":"216.58.219.177","connection":"69925"}
`

func Test_har2entries(t *testing.T) {
	br := bytes.NewBuffer([]byte(HAR))
	bw := new(bytes.Buffer)
	if err := Har2Entries(br, bw); err != nil {
		t.Fatal(err)
	}
	if bw.String() != ENTRIES {
		t.Errorf("input and output has not match")
	}
}

func Test_entries2har(t *testing.T) {
	br := bytes.NewBuffer([]byte(ENTRIES))
	bw := new(bytes.Buffer)
	if err := Entries2Har(br, bw); err != nil {
		t.Fatal(err)
	}
	if bw.String() != ENTRIES {
		t.Errorf("input and output has not match")
	}
}
