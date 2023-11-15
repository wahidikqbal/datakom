CREATE TABLE `categories` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `pangkats` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `kesatuans` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `stocks` (
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
  `serialnumber` text,
  `stock_id` integer NOT NULL,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

ALTER TABLE `categories` ADD FOREIGN KEY (`id`) REFERENCES `products` (`category_id`);

ALTER TABLE `pangkats` ADD FOREIGN KEY (`id`) REFERENCES `products` (`pangkat_id`);

ALTER TABLE `kesatuans` ADD FOREIGN KEY (`id`) REFERENCES `products` (`kesatuan_id`);

ALTER TABLE `stocks` ADD FOREIGN KEY (`id`) REFERENCES `products` (`stock_id`);
