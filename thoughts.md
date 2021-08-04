


### jump game 2
```
algo(A):
    help = []
    help[len(A)-1] = 0

    for i=len(A)-2; i>=0; i--:
        minJump = MAX_INT
        for j=A[i]; j>0; j--:
            if minJump > A[i+j]+1:
                minJump=A[i+j]+1
        A[i]=minJump
    return A[0]
```
