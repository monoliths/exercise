# Test Runs
- 20KB files
- Client and nginx server running on same machine
- TTFB Target = 200ms
- Total Target = 1000ms (1 second)
- NOTE: Indivual request results omited for shorter length

# Result Analysis
- When does concurrency performance starts breaking down? 
    ```
    According to my test targets, performance starts to break down when concurrent request reach 2k. The primary culprit is TTFB which I set to a realistic number. Total time was still doing great at this point. 
    ```
- What is the bottleneck? 
    ```
    Its hard for me to say, but my educated guess is that we are just hitting the normal scaling of what my VM can perform. Opening up so many connections (2k) takes time causing a delay in that inital first byte. If we take a look at total time we can see that it is great across all tests, this allows me to rule out the file size and disk as a bottle neck. 
    ```

# Raw Results
- 10 files/requests
```
./main urls-localhost/20KB/10.txt downloads 200 1000
...
===================== Average Results ====================
 TTFB PASS @ 1ms
 TOTAL PASS @ 2ms
```

- 50 files/requests
```
./main urls-localhost/20KB/50.txt downloads 200 1000
...
===================== Average Results ====================
 TTFB PASS @ 5ms
 TOTAL PASS @ 6ms
 ```

 - 100 files/requests
 ```
 ./main urls-localhost/20KB/100.txt downloads 200 1000
 ...
 ===================== Average Results ====================
 TTFB PASS @ 42ms
 TOTAL PASS @ 46ms
 ```

 - 200 files/requests
 ```
./main urls-localhost/20KB/200.txt downloads 200 1000
...
===================== Average Results ====================
 TTFB PASS @ 70ms
 TOTAL PASS @ 97ms
 ```

 - 300 files/requests
 ```
./main urls-localhost/20KB/300.txt downloads 200 1000
...
 ===================== Average Results ====================
 TTFB PASS @ 64ms
 TOTAL PASS @ 85ms
 ```
- 500 files/requests
 ```
 ./main urls-localhost/20KB/300.txt downloads 200 1000
 ...
 ======================== Average Results ====================
 TTFB PASS @ 149ms
 TOTAL PASS @ 177ms
 ```

 - 700 files/requests
 ```
 ./main urls-localhost/20KB/300.txt downloads 200 1000
 ...
 ===================== Average Results ====================
 TTFB PASS @ 165ms
 TOTAL PASS @ 188ms
 ```

- 1000 files/requests
```
./main urls-localhost/20KB/1000.txt downloads 200 1000
...
===================== Average Results ====================
 TTFB PASS @ 190ms
 TOTAL PASS @ 301ms
```

- 2000 files/requests (lets try this for fun)
```
./main urls-localhost/20KB/2000.txt downloads 200 1000
===================== Average Results ====================
 TTFB FAIL @ 348ms
 TOTAL PASS @ 405ms
```