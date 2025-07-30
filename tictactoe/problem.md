# Tic-Tac-Toe — Board State Evaluator

## Objectives
- Choose appropriate data types.
- Use conditionals and for-loops.
- Use slices.
- Create and use functions.

## Problem
Diberikan sebuah string panjang **9** berisi karakter **'X'**, **'O'**, dan **'-'** yang merepresentasikan papan 3×3 dari kiri ke kanan, atas ke bawah. Tentukan status permainan dan cetak salah satu dari:

- `X wins!`
- `O wins!`
- `Its a draw!`
- `Game still in progress!`
- `Invalid game board`

### Representasi papan
Index 0..8 diisi ke papan 3×3:
```
0 1 2
3 4 5
6 7 8
```

Contoh input: `XOXXOOXXO` merepresentasikan:
```
X O X
X O O
X X O
```

## Aturan Validitas (Invalid game board)
- X selalu jalan **lebih dulu**.
- Hitungan langkah: `countX == countO` atau `countX == countO + 1`.
- Tidak boleh **kedua pemain** menang sekaligus.
- Jika **X menang** maka `countX == countO + 1`.
- Jika **O menang** maka `countX == countO`.

## Kriteria Penilaian
- Menemukan kemenangan jika ada 3 simbol sama di baris/kolom/diagonal.
- Jika tidak ada pemenang dan tidak ada sel kosong (`'-'`) → `Its a draw!`.
- Jika tidak ada pemenang dan masih ada sel kosong → `Game still in progress!`.

## Contoh
- `XOXXOOXXO` → `X wins!`
- `XOOXOXOXO` → `O wins!`
- `OXOXOXXOX` → `Its a draw!`
- `XOXX--O--` → `Game still in progress!`
- `XXXOOOXXO` → `Invalid game board`

## Saran Implementasi (Go)
- Representasikan papan sebagai `[][]rune` 3×3.
- Definisikan 8 kombinasi garis menang.
- Fungsi:
  - `validateInput(s string) (err string, blanks int)`
  - `validateWin(board [][]rune, p rune) bool`
  - `main()` untuk membaca input, validasi, cek menang, dan cetak status.

## Cara Menjalankan
```
go run .
```