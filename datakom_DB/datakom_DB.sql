CREATE TABLE `categories` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `pangkat` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `kesatuan` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `stock` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `products` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `pangkat_id` integer NOT NULL,
  `nrp` varchar(255) NOT NULL,
  `kesatuan_id` integer NOT NULL,
  `category_id` integer NOT NULL,
  `serial_number` text,
  `stock_id` integer NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

ALTER TABLE `categories` ADD FOREIGN KEY (`id`) REFERENCES `products` (`category_id`);

ALTER TABLE `pangkat` ADD FOREIGN KEY (`id`) REFERENCES `products` (`pangkat_id`);

ALTER TABLE `kesatuan` ADD FOREIGN KEY (`id`) REFERENCES `products` (`kesatuan_id`);

ALTER TABLE `stock` ADD FOREIGN KEY (`id`) REFERENCES `products` (`stock_id`);
