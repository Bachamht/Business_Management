CREATE TABLE IF NOT EXISTS `t_user_info`
(
    `profile_photo` VARCHAR(100)                 COMMENT '工作照',
    `name` 			VARCHAR(30)         NOT NULL COMMENT '姓名',
    `phone_number`  VARCHAR(20)			NOT NULL COMMENT '联系电话',
    `password`		VARCHAR(20)         NOT NULL COMMENT '密码',
    `company`  		VARCHAR(100)  		NOT NULL COMMENT '公司信息',
    `permitted`     BOOLEAN                NOT NULL COMMENT '是否已经得到批准 true:是 false:否',
    `created_at` 	TIMESTAMP    		NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    		NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`phone_number`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `t_client_info`
(
    `id`                  BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `salesman_name`       VARCHAR(30)         NOT NULL COMMENT '负责人姓名',
    `salesman_number`     VARCHAR(20)		  NOT NULL COMMENT '负责人电话',
    `salesman_company`     VARCHAR(20)		  NOT NULL COMMENT '负责人公司',
    `client_name`         VARCHAR(30)         NOT NULL COMMENT ' 客户姓名',
    `client_number`       VARCHAR(30)         NOT NULL COMMENT '客户电话',
    `client_company`             VARCHAR(100)  	  NOT NULL COMMENT '公司信息',
    `detail`              TEXT		          NOT NULL COMMENT '业务内容',
    `progress`            VARCHAR(100)                 COMMENT '进展情况',
    `finished`            BOOLEAN                NOT NULL COMMENT '是否完成 true:是 false:否',
    `created_at` 	      TIMESTAMP    		  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`          TIMESTAMP    		  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `t_admin_info`
(
    `account`       VARCHAR(20)			NOT NULL COMMENT '账号',
    `password`		VARCHAR(20)         NOT NULL COMMENT '密码',
    `created_at` 	TIMESTAMP    		NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    		NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`account`)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `t_conflict_info`
(
    `id`                    BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name_b`                VARCHAR(20)			    NOT NULL COMMENT '另一冲突业务员姓名',
    `phone_b`		        VARCHAR(20)             NOT NULL COMMENT '另一冲突业务员电话',
    `company_b`             VARCHAR(20)             NOT NULL COMMENT '另一冲突业务员公司',
    `phone`		            VARCHAR(20)             NOT NULL COMMENT     '本人电话',
    `name`		            VARCHAR(20)             NOT NULL COMMENT     '本人姓名',
    `company`		        VARCHAR(20)             NOT NULL COMMENT     '本人公司',
    `client_name`           VARCHAR(20)             NOT NULL COMMENT     '客户姓名',
    `client_number`         VARCHAR(20)             NOT NULL COMMENT     '客户电话',
    `client_company`        VARCHAR(100)  	        NOT NULL COMMENT     '公司信息',
    `detail`                TEXT		            NOT NULL COMMENT     '业务内容',
    `conflict_content`		VARCHAR(80)             NOT NULL COMMENT     '冲突内容',
    `created_at` 	TIMESTAMP    		NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    		NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;
