bilangan = int(input('masukan bilangan'))
# menggunakan modulus 2 jika bialangan nya habis di bagi 2 berati genap
genap = bilangan%2
print(genap)
if genap == 0:
    print("bilangan genap")
else:
    print("bilangan ganjil")