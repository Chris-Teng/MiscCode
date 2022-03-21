-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2022-02-14 20:22:46
-- 服务器版本： 8.0.12
-- PHP 版本： 7.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `game`
--

-- --------------------------------------------------------

--
-- 表的结构 `city`
--

CREATE TABLE `city` (
  `id` int(11) NOT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
  `population` int(11) DEFAULT NULL,
  `economy` int(11) DEFAULT NULL,
  `army` int(11) DEFAULT NULL,
  `averageCombatPower` float NOT NULL DEFAULT '1',
  `food` int(11) DEFAULT NULL,
  `growthOfPopulation` float DEFAULT NULL,
  `growthOfEconomy` float DEFAULT NULL,
  `growthOfArmy` float DEFAULT NULL,
  `belongToPower` int(11) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `city`
--

INSERT INTO `city` (`id`, `name`, `population`, `economy`, `army`, `averageCombatPower`, `food`, `growthOfPopulation`, `growthOfEconomy`, `growthOfArmy`, `belongToPower`) VALUES
(1, '洛阳', 350051, 17502, 3615, 1.1, 65484, 0.08, 0.05, 0.01, 1),
(2, '长安', 851131, 56412, 4634, 1.1, 120354, 0.09, 0.1, 0.01, NULL),
(3, '徐州', 451502, 180021, 4777, 1.3, 130124, 0.06, 0.04, 0.04, 3),
(4, '成都', 670000, 65400, 3200, 1, 78000, 0.1, 0.6, 0.2, 2),
(5, '建业', 450000, 20000, 4000, 1, 60000, 0.08, 0.06, 0.01, NULL),
(6, '巴中', 250000, 12000, 3500, 1, 50000, 0.07, 0.04, 0.01, NULL),
(7, '荆州', 260000, 12000, 6615, 1, 50000, 0.07, 0.04, 0.01, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `coordinates`
--

CREATE TABLE `coordinates` (
  `id` int(11) NOT NULL,
  `longitude` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `attitude` timestamp NULL DEFAULT NULL,
  `cityId` int(11) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `power`
--

CREATE TABLE `power` (
  `id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- 转存表中的数据 `power`
--

INSERT INTO `power` (`id`, `name`) VALUES
(1, '魏'),
(2, '蜀'),
(3, '吴');

-- --------------------------------------------------------

--
-- 替换视图以便查看 `statistics`
-- （参见下面的实际视图）
--
CREATE TABLE `statistics` (
`id` int(11)
,`name_of_power` varchar(255)
,`num_of_cities` bigint(21)
,`total_army` decimal(32,0)
,`total_population` decimal(32,0)
);

-- --------------------------------------------------------

--
-- 视图结构 `statistics`
--
DROP TABLE IF EXISTS `statistics`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `statistics` (`id`, `name_of_power`, `num_of_cities`, `total_population`, `total_army`) AS   select `power`.`id` AS `id`,`power`.`name` AS `name`,count(`city`.`id`) AS `COUNT(city.id)`,sum(`city`.`population`) AS `SUM(city.population)`,sum(`city`.`army`) AS `SUM(city.army)` from (`power` join `city`) where (`power`.`id` = `city`.`belongToPower`) group by `power`.`id`  ;

--
-- 转储表的索引
--

--
-- 表的索引 `city`
--
ALTER TABLE `city`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `coordinates`
--
ALTER TABLE `coordinates`
  ADD PRIMARY KEY (`id`),
  ADD KEY `cityId` (`cityId`);

--
-- 表的索引 `power`
--
ALTER TABLE `power`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `city`
--
ALTER TABLE `city`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- 使用表AUTO_INCREMENT `coordinates`
--
ALTER TABLE `coordinates`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `power`
--
ALTER TABLE `power`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
