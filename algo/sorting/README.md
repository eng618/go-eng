# Sorting Algorithms

## Merge Sort

The mere sort algorithm was inspired by [this](https://www.golangprograms.com/golang-program-for-implementation-of-mergesort.html) example.

### pseudo code

```python
 split each element into partitions of size 1
 recursively merge adjacent partitions
   for i = leftPartIdx to rightPartIdx
     if leftPartHeadValue <= rightPartHeadValue
       copy leftPartHeadValue
     else: copy rightPartHeadValue; Increase InvIdx
 copy elements back to original array
```
