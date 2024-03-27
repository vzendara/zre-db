Implement RESP decoding. 

* Read simple string
* Read error
* Read integer
* Read bulk string
* Read array

To start, write test cases for each decode operation. 

* TestDecodeSimpleString
* TestDecodeError
* TestDecodeInt64
* TestDecodeBulkString
* TestDecodeArray

Some questions to help warm up. In RESP, the type of some data is indicated by the first byte:
* What is the first byte for a Simple String?
* For Errors?
* For Integers?
* For Bulk Strings?
* For Arrays?

Additionally, each part of the RESP protocol is terminated by what? 
=> "\r\n" (CRLF)

It would help to read the spec:
[Redis Protocol Specification](https://redis-doc-test.readthedocs.io/en/latest/topics/protocol/)

The following will likely be helpful as well:
* https://pkg.go.dev/testing
