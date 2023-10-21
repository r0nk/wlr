# wlr
wlr (Word List Replace) is a tool to replace parts of strings with lines from files.

If you've used ffuf and httpx before, wlr = ffuf - httpx

## Usage Example:
```
wlr "http://www.example.com/FUZZ" wordlist.txt | httpx
wlr "http://SUB.example.com/PATH" "path/domains.txt:SUB" "path/paths.txt:PATH"
cat passwords.txt | wlr "username:FUZZ"
cat ips.txt | wlr "ping FUZZ" | parallel
```

## Flags:
```
  -clusterbomb
    	Enable clusterbomb mode (1 1, 1 2, 1 3...) (default true)
  -pitchfork
    	Enable pitchfork mode (1 1, 2 2, 3 3...)
```

## Install:
```
go install -v github.com/r0nk/wlr@latest
```
