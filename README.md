# Sample Sort
Sample sort is a parallel version of quicksort.

# How Does It Work?
You can read more on Wikipedia [here](https://en.wikipedia.org/wiki/Samplesort)

## Input
- An array of N items
- An integer P for number of processes to use.

## Result:
- The input array has been sorted.

## Steps:
- Sample
    - Randomly select 3*(P-1) items from the array.
    - Sort those items.
    - Take the median of each group of three in the sorted array,
    producing an array (samples) of (P-1) items.
    - Add 0 at the start and +inf at the end (or the min and max
    values of the type being sorted) of the samples array so it
    has (P+1) items numbered (0 to P).
- Partition
    - Spawn P processes, numbered p in (0 to P-1).
    - Each process builds a local array of items to be sorted by
    scanning the full input and taking items between samples[p]
    and samples[p+1].
    - Write the number of items (n) taken to a shared array sizes
    at slot p.
    - Sort locally
    - Each process uses quicksort to sort the local array.
    - Copy local arrays to input.
    - Each process calculates where to put its result array as follows:
        - start = sum(sizes[0 to p - 1]) # that’s sum(zero items) = 0 for p = 0
        - end = sum(sizes[0 to p]) # that’s sum(1 item) for p = 0
        - Warning: Data race if you don’t synchronize here.
        Each process copies its sorted array to input[start..end]
- Cleanup
    - Terminate the P subprocesses. Array has been sorted “in place”.
