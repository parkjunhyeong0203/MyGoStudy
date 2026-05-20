import math
def get_lcm(x, y):
    return (x * y) // math.gcd(x, y)
signals = [] #세 개 이상의 객체(신호등)를 다룰려면 주체를 신호등으로 보면 안되고 시간으로 봐서 그 시간을 신호등에 대입
for i in range(3):
    signals.append(list(map(int, input().split())))
ans, com, lsn, status = 0, 0, 1, True
a, b =  [], []
num = len(signals)
for i in range(num):
    rsum = sum(signals[i])
    lsn = get_lcm(lsn, rsum)
    a.append(sum(signals[i]) - signals[i][2] - signals[i][1])
    b.append(sum(signals[i]) - signals[i][2])
for n in range(1,lsn + 1):
    status = True
    for m in range(num):
        ans = n % sum(signals[m])
        if (a[m] < ans) and (ans <= b[m]):
            ans = n
        else :
            status = False
            break
    if (status == True):
        break
    else :
        ans = -1

print(ans)