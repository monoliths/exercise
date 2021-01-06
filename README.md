# How to run
After building: 
```
./main path-to-urls dir-to-download-files target-ttfb target-total-time
```

# Example run 
- `/tmp/urls.txt` file looks like
```
http://localhost:41337/sm/20KB-0
http://localhost:41337/sm/20KB-1
http://localhost:41337/sm/20KB-2
http://localhost:41337/sm/20KB-3
http://localhost:41337/sm/20KB-4
http://localhost:41337/sm/20KB-5
http://localhost:41337/sm/20KB-6
http://localhost:41337/sm/20KB-7
http://localhost:41337/sm/20KB-8
http://localhost:41337/sm/20KB-9
```
- command: 
```
./main /tmp/urls.txt /tmp/downloads 200 10000 
```