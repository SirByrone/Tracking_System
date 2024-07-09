-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jul 02, 2024 at 11:37 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `attachee_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` longtext DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT NULL,
  `reg_number` longtext DEFAULT NULL,
  `course` longtext DEFAULT NULL,
  `year_of_study` longtext DEFAULT NULL,
  `department` longtext DEFAULT NULL,
  `college` longtext DEFAULT NULL,
  `address` longtext DEFAULT NULL,
  `mobile_number` longtext DEFAULT NULL,
  `attachment_from` longtext DEFAULT NULL,
  `attachment_to` longtext DEFAULT NULL,
  `lecturer_in_charge_name` longtext DEFAULT NULL,
  `lecturer_contact` longtext DEFAULT NULL,
  `s_name` longtext DEFAULT NULL,
  `s_address` longtext DEFAULT NULL,
  `sphone` longtext DEFAULT NULL,
  `s_county` longtext DEFAULT NULL,
  `s_constituency` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `username`, `password`, `is_admin`, `reg_number`, `course`, `year_of_study`, `department`, `college`, `address`, `mobile_number`, `attachment_from`, `attachment_to`, `lecturer_in_charge_name`, `lecturer_contact`, `s_name`, `s_address`, `sphone`, `s_county`, `s_constituency`) VALUES
(1, 'SCIADMIN', 'SCIADMIN', 'SCI123', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(2, 'Byrone ', 'Opande', 'incorrect', 0, '21/05174', 'Software Development', '3rd Year', 'Department of Software Development', 'KCA University', 'Nairobi', '0745130473', '2024-05-07', '2024-07-07', 'Ernest Madara', '0722892206', 'Joseph', 'Nairobi', '0721348049', 'Nairobi Municipality', 'Nairobi');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
