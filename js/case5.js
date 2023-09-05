let bil1 = 10
let bil2  = 20
let bil3 = 30

let arr = [bil1,bil2,bil3]
const min = Math.min(...arr)
const max = Math.max(...arr)
const avg = bil1+bil2+bil3/arr.length
console.log("bilangan tersbesar adalah",max)
console.log("bilangan terkecil adalah",min)
console.log("rata rata bilangan  adalah",avg)