LINE=$1
# ./amath 2**2 && ./amath 2+2
echo 'print(eval("'$LINE'"))' > out.out.out.py && python3 out.out.out.py
rm out.out.out.py
