Name Matcher
------------

### Algorithm draft

To calculate similarity score of Football Team name, we will use couple of
strategies:

- We clean both input strings

- We assume that order of words in alternative name is the same as in original.
  For example alternative to West Bromwich will not be Bromwich West

- Next we expand alternative string by capital letter as if it was separate word
  For example: ABeC becomes A Be C

- We iterate over original words and checking much of alternative word is
  in its original counterpart
    B - Bromwich, that is 1/8
    If we have partial match, we can assume it is abbreviation of the word
    We fit the weight to the range of 1.0>=weight>=0.5
    If there is no match then we put 0 weight

- We calculate the similarity as sum of weights divided by number
   of match attempts performed

- If resulting similarity is >=0.5 we can assume we have a likely match
    If similarity is 1.0 then it is a perfect match
    If similarity is <0.5 then match is unlikely
    If similarity is 0.0 then for sure there is no match

### Usage

#### Installation:

I assume you have Go installed (tested with 1.7 version). Run:

```
go get github.com/elcct/name-matcher
```

Name matcher should be in your $GOPATH/bin. You can try it by issuing:

```
./name-matcher -original="West Bromwich Albion Football Club" -alternative="WBA FC"
```

Should give you:
```
0.5916666666666667
```

Which is a likely match.
