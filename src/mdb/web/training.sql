-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 17 Okt 2019 pada 13.00
-- Versi server: 10.1.37-MariaDB
-- Versi PHP: 7.3.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `training`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `databarang`
--

CREATE TABLE `databarang` (
  `ID` varchar(11) NOT NULL,
  `NamaBarang` varchar(255) NOT NULL,
  `Harga` int(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `databarang`
--

INSERT INTO `databarang` (`ID`, `NamaBarang`, `Harga`) VALUES
('b001', 'odol', 10000),
('b002', 'sabun', 5000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `login`
--

CREATE TABLE `login` (
  `ID` int(255) NOT NULL,
  `uname` varchar(255) NOT NULL,
  `passw` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `login`
--

INSERT INTO `login` (`ID`, `uname`, `passw`) VALUES
(1, 'kasir', 'kasir');

-- --------------------------------------------------------

--
-- Struktur dari tabel `simpanbarang`
--

CREATE TABLE `simpanbarang` (
  `ID` int(11) NOT NULL,
  `IDBarang` varchar(7) NOT NULL,
  `NamaBarang` varchar(255) NOT NULL,
  `Harga` int(11) NOT NULL,
  `Jumlah` int(11) NOT NULL,
  `Total` int(11) NOT NULL,
  `Tanggal` int(11) NOT NULL,
  `Bayar` int(11) NOT NULL,
  `Kembalian` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `simpanbarang`
--

INSERT INTO `simpanbarang` (`ID`, `IDBarang`, `NamaBarang`, `Harga`, `Jumlah`, `Total`, `Tanggal`, `Bayar`, `Kembalian`) VALUES
(18, 'b002', 'sabun', 5000, 2, 0, 0, 0, 0),
(19, 'b002', 'sabun', 5000, 2, 0, 0, 0, 0),
(20, 'b002', 'sabun', 5000, 3, 0, 0, 0, 0),
(21, 'b002', 'sabun', 5000, 5, 25000, 2019, 25000, 25000),
(22, 'b002', 'sabun', 5000, 4, 20000, 2019, 20000, 0);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `databarang`
--
ALTER TABLE `databarang`
  ADD PRIMARY KEY (`ID`);

--
-- Indeks untuk tabel `login`
--
ALTER TABLE `login`
  ADD PRIMARY KEY (`ID`);

--
-- Indeks untuk tabel `simpanbarang`
--
ALTER TABLE `simpanbarang`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `login`
--
ALTER TABLE `login`
  MODIFY `ID` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `simpanbarang`
--
ALTER TABLE `simpanbarang`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
