CREATE TABLE IF NOT EXISTS functions (
    `id` VARCHAR(50) NOT NULL,
    `tenant_id` VARCHAR(50) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `container_image` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
)  ENGINE=INNODB;