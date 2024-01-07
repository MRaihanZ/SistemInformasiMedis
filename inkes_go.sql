-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 07, 2024 at 04:56 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `inkes_go`
--

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'poli umum', '2024-01-02 10:47:30', '2024-01-02 10:47:30'),
(2, 'poli gigi', '2024-01-02 10:47:30', '2024-01-02 10:47:30'),
(3, 'poli kandungan', '2024-01-02 10:47:55', '2024-01-02 10:48:03');

-- --------------------------------------------------------

--
-- Table structure for table `doctor`
--

CREATE TABLE `doctor` (
  `id` int(11) NOT NULL,
  `name` varchar(30) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `doctor`
--

INSERT INTO `doctor` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'RS dr. Sismadi', '2024-01-06 22:24:37', '2024-01-06 22:24:37'),
(2, 'RS dr. Doni', '2024-01-06 22:24:46', '2024-01-06 22:24:46'),
(3, 'RS dr. Sumardi', '2024-01-06 22:24:55', '2024-01-06 22:24:55');

-- --------------------------------------------------------

--
-- Table structure for table `medicine`
--

CREATE TABLE `medicine` (
  `id` int(11) NOT NULL,
  `name` varchar(30) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `medicine`
--

INSERT INTO `medicine` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Paracetamol', '2024-01-06 22:26:57', '2024-01-06 22:26:57'),
(2, 'Amoxicillin', '2024-01-06 22:27:29', '2024-01-06 22:27:29'),
(3, 'Diatabs', '2024-01-06 22:27:50', '2024-01-06 22:27:50'),
(4, 'Oralit', '2024-01-06 22:27:58', '2024-01-06 22:27:58'),
(5, 'Ambroxol', '2024-01-06 22:29:00', '2024-01-06 22:29:00'),
(6, 'Methylprednisolone', '2024-01-06 22:29:52', '2024-01-06 22:29:52');

-- --------------------------------------------------------

--
-- Table structure for table `pasien`
--

CREATE TABLE `pasien` (
  `id` int(11) NOT NULL,
  `name` varchar(30) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pasien`
--

INSERT INTO `pasien` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Raihan', '2024-01-07 08:48:49', '2024-01-07 08:48:49'),
(2, 'Rei', '2024-01-07 08:48:55', '2024-01-07 08:48:55'),
(3, 'Zaky', '2024-01-07 08:49:00', '2024-01-07 08:49:00'),
(4, 'Zhafran', '2024-01-07 08:49:43', '2024-01-07 08:49:43');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `category_id` int(11) NOT NULL,
  `stock` int(11) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `category_id`, `stock`, `description`, `created_at`, `updated_at`) VALUES
(7, 'kapal api', 7, 10, 'mantap', '2023-12-04 10:12:06', '2023-12-04 10:12:06'),
(8, 'asdfadsf', 1, 0, NULL, '2024-01-02 05:18:59', '2024-01-02 05:18:59'),
(9, 'test', 1, 0, NULL, '2024-01-02 05:19:09', '2024-01-02 05:19:09'),
(10, 'test', 1, 0, NULL, '2024-01-02 05:23:36', '2024-01-02 05:23:36');

-- --------------------------------------------------------

--
-- Table structure for table `record`
--

CREATE TABLE `record` (
  `id` int(11) NOT NULL,
  `id_pasien` int(11) NOT NULL,
  `id_categori` int(11) NOT NULL,
  `id_doctor` int(11) NOT NULL,
  `diagnose` text NOT NULL,
  `description` text NOT NULL,
  `id_medicine` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `record`
--

INSERT INTO `record` (`id`, `id_pasien`, `id_categori`, `id_doctor`, `diagnose`, `description`, `id_medicine`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 1, '', '', 1, '2024-01-02 10:49:56', '2024-01-07 08:57:49'),
(2, 2, 2, 2, '', '', 2, '2024-01-02 05:24:16', '2024-01-07 08:57:52'),
(3, 3, 1, 3, '', '', 3, '2024-01-02 05:43:46', '2024-01-07 14:56:26'),
(5, 4, 1, 1, 'batuk', 'batuk berdahak', 5, '2024-01-07 08:39:01', '2024-01-07 08:39:01'),
(6, 4, 1, 2, 'Demam', 'Demam', 1, '2024-01-07 08:39:58', '2024-01-07 08:39:58');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `doctor`
--
ALTER TABLE `doctor`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `medicine`
--
ALTER TABLE `medicine`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pasien`
--
ALTER TABLE `pasien`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `record`
--
ALTER TABLE `record`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `doctor`
--
ALTER TABLE `doctor`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `medicine`
--
ALTER TABLE `medicine`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `pasien`
--
ALTER TABLE `pasien`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `record`
--
ALTER TABLE `record`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
