# coding test GITS Indonesia
source code answer of the test from gits indonesia

Name : Aria Shabry
Email: aria.shabry4@gmail.com

Fulstack

1. Weighted Strings
2. Highest Palindrome
3. Balanced Bracket

Kompleksitas Code Soal No 3

Kompleksitas Waktu:
- Iterasi melalui string input: O(n), di mana n adalah panjang string input.
- Pengecekan tanda kurung buka dan tutup menggunakan isOpeningBracket dan isMatchingBracket: O(1) (kompleksitas konstan).
- Pemanggilan rekursif pada isBalancedBracketRecursive: Terjadi sebanyak karakter dalam string input, yaitu O(n).
- Jadi, kompleksitas waktu keseluruhan adalah O(n), di mana n adalah panjang string input.

Kompleksitas Ruang:
- Stack digunakan untuk menyimpan karakter buka yang belum ditutup.
- Jumlah elemen pada stack pada titik tertentu adalah sebanding dengan jumlah karakter buka yang belum ditutup.
- Pemanggilan rekursif pada isBalancedBracketRecursive: Terjadi sebanyak kedalaman rekursi, yang pada kasus terburuk adalah panjang string input, yaitu O(n).
- Jadi, kompleksitas ruang keseluruhan adalah O(n), di mana n adalah panjang string input.

Sehingga, kodingan di soal no 3 memiliki kompleksitas waktu O(n) dan kompleksitas ruang O(n), di mana n adalah panjang string input.
