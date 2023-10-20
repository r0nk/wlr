word list replace

usage:
wlr "http://www.example.com/FUZZ" wordlist.txt | httpx
wlr "http://DOM.example.com/PATH" "path/domains.txt:DOM" "path/paths.txt:PATH"
cat passwords.txt | wlr "username:FUZZ"
