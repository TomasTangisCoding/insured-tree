CREATE TABLE your_schema.users (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '保戶編號',
  `name` varchar(50) NOT NULL COMMENT '保戶名稱',
  `referrer_id` int NOT NULL COMMENT '推薦人編號',
  `email` varchar(50) DEFAULT NULL COMMENT '保戶email',
  `parent_id` int DEFAULT NULL COMMENT '父節點',
  `left_child` int DEFAULT NULL COMMENT '左子節點',
  `right_child` int DEFAULT NULL COMMENT '右子節點',
  `is_delete` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否刪除 0:未刪除 1:刪除',
  `created_at` datetime NOT NULL COMMENT '建立時間',
  `updated_at` datetime DEFAULT NULL COMMENT '更新時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

insert into users (id, name, referrer_id, email, parent_id, left_child, right_child, is_delete, created_at, updated_at)
values  (1, 'a', 0, null, null, 2, 3, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (2, 'b', 1, null, 1, 4, 5, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (3, 'c', 1, null, 1, 6, 7, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (4, 'd', 1, null, 1, 8, 9, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (5, 'e', 3, null, 1, 10, null, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (6, 'f', 3, null, 1, 11, null, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (7, 'g', 3, null, 11, null, null, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (8, 'h', 2, null, 1, null, null, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (9, 'i', 4, null, 1, null, null, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (10, 'j', 9, null, 11, null, null, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55'),
        (11, 'k', 1, null, 1, null, null, 0, '2023-04-30 08:49:49', '2023-04-30 16:27:55');
