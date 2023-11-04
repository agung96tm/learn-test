Manipulate Records
=================================================


#### Question
```
first input: length of records
second input: input records by length from first_input process
last input:
    if "edit <key> <value>" then
        edit: update value by key from records
        print current records

    if "exit" then
        close process

    else
        print "invalid action"
```

#### Answer
```
# Input:
3
name John
lastName Bliss
city Florida
edit city Seattle

# Output:
city Seattle
lastName Bliss
name John
```
