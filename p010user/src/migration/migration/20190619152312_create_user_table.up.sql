CREATE TABLE  IF NOT EXISTS user (
  id int(11) unsigned NOT NULL PRIMARY KEY AUTO_INCREMENT ,
  username VARCHAR(255) NOT NULL UNIQUE ,
  password VARCHAR(255),
  email VARCHAR(255) ,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  oauth ENUM ('google', 'facebook'),
  native_lang VARCHAR(255),
  is_validated TINYINT(3)
);
