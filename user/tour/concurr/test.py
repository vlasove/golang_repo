from math import sqrt
a=float(input("Enter a: "))
b=float(input("Enter b: "))
c=float(input("Enter c: "))

if a==0 and b!=0 and c==0:
    x=-c/b
    print("x= %0.1f"%(x))
elif a==0 and b==0:
    print('Пустое множество')
elif a!=0 and b==0:
    x1=c**0.5
    x2=-(c**0.5)
    print("x1= %0.1f"%(x1))
    print("x2= %0.1f"%(x2))
elif a!=0 and b!=0:
    d1=(b*b)-(4*a*c)
    if d1>0:
         d=d1**0.5
         x1=(-b+d)/(2*a)
         x2=(-b-d)/(2*a)
         print("x1= %0.1f"%(x1))
         print("x2= %0.1f"%(x2))
    else:  
        print('В действительном множестве корней нет')