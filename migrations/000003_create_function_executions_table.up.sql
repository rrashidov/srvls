CREATE TABLE IF NOT EXISTS function_executions (
    `id` VARCHAR(50) NOT NULL,
    `function_id` VARCHAR(50) NOT NULL,
    `triggered` DATE,
    `started` DATE,
    `finished` DATE,
    `status` INT,
    PRIMARY KEY (`id`)
)  ENGINE=INNODB;