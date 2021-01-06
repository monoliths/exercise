# Test Run
- 200MB files 
- Client and nginx server running on same machine
- TTFB Target = 200ms
- Total Target = 10000ms (10 seconds)
- NOTE: Indivual request results omited for shorter length
- NOTE: Request are limited due to disk space on my end

# Result Analysis
- When does concurrency performance starts breaking down? 
    ```
    According to my test targets, performance starts to break down when concurrent request reach 40. The primary culprit is total time average going over our target total time. It can be noted that TTFB is still performing well under our target TTFB.

    However, once we reach 200 concurrent request both total and TTFB targets are failing. This is the point where our performance is "bad".
    ```
- What is the bottleneck? 
    ```
    During these test runs I had "Netdata" running which I could see performance and health of my VM. The bottleneck was easily the Disk I/O performance. I could see I was reaching ~550MiB/s which was capping the limits of what my VM could do. This bottleneck would slow down the total time.
    ```

# Raw Results
- 10 files/requests
    ```
    ./main urls-localhost/200MB/10.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB PASS @ 21ms
    TOTAL PASS @ 2213ms
    ```

- 20 files/requests
    ```
    ./main urls-localhost/200MB/20.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB PASS @ 34ms
    TOTAL PASS @ 6316ms
    ```

- 30 files/requests
    ```
    ./main urls-localhost/200MB/30.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB PASS @ 54ms
    TOTAL PASS @ 9165ms
    ```

- 40 files/requests
    ```
    ./main urls-localhost/200MB/40.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB PASS @ 87ms
    TOTAL FAIL @ 13110ms
    ```

- 50 files/requests
    ```
    ./main urls-localhost/200MB/50.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB PASS @ 84ms
    TOTAL FAIL @ 15310ms
    ```

 - 100 files/requests
    ```
    ./main urls-localhost/200MB/100.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB PASS @ 177ms
    TOTAL FAIL @ 31667ms
    ```

 - 200 files/requests
    ```
    ./main urls-localhost/200MB/200.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB FAIL @ 336ms
    TOTAL FAIL @ 61067ms 
    ```


 - 300 files/requests
    ```
    ./main urls-localhost/200MB/300.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB FAIL @ 560ms
    TOTAL FAIL @ 121816ms
    ```

- 500 files/requests
    ```
    ./main urls-localhost/200MB/500.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB FAIL @ 522ms
    TOTAL FAIL @ 249130ms
    ```

 - 700 files/requests
    ```
    ./main urls-localhost/200MB/700.txt downloads 200 10000
    ...
    ===================== Average Results ====================
    TTFB FAIL @ 1269ms
    TOTAL FAIL @ 382581ms
    
    ```
