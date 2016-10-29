-- phpMyAdmin SQL Dump
-- version 4.1.13
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2016-10-28 10:23:30
-- 服务器版本： 5.1.73
-- PHP Version: 5.6.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `timer`
--

-- --------------------------------------------------------

--
-- 表的结构 `run_logs`
--

CREATE TABLE IF NOT EXISTS `run_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tp_id` int(10) unsigned NOT NULL,
  `start_runtime` int(11) NOT NULL,
  `end_runtime` int(11) NOT NULL COMMENT '结束运行时间',
  `url` varchar(1000) NOT NULL,
  `result` varchar(1000) NOT NULL,
  `status` tinyint(4) NOT NULL COMMENT '1=运行成功/0=运行失败',
  PRIMARY KEY (`id`),
  KEY `tp_id` (`tp_id`),
  KEY `status` (`status`),
  KEY `start_runtime` (`start_runtime`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 AUTO_INCREMENT=12437 ;

-- --------------------------------------------------------

--
-- 表的结构 `timer_process`
--

CREATE TABLE IF NOT EXISTS `timer_process` (
  `tp_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1=正常/0=禁用',
  `base_time` int(11) NOT NULL COMMENT '基准时间',
  `interval_minute` int(11) NOT NULL COMMENT '间隔分钟',
  `timeout` int(11) NOT NULL COMMENT '超时时间-秒',
  `url` varchar(1000) NOT NULL,
  `uid` int(11) NOT NULL,
  `memos` varchar(100) NOT NULL,
  `last_run_status` tinyint(4) NOT NULL COMMENT '1=成功/0=失败',
  `last_run_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`tp_id`),
  KEY `uid` (`uid`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 COMMENT='定时程序配置表' AUTO_INCREMENT=10 ;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
