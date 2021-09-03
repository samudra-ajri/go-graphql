--name: check-table-existence
SELECT count(*) as exist 
FROM information_schema.tables 
WHERE table_name = ? 
LIMIT 1;

--name: author
INSERT INTO `author` 
VALUES 
  (1,'Iman Tumorang','2017-05-18 13:50:19','2017-05-18 13:50:19'),
  (2,'Samudra Ajri','2021-09-02 13:50:19','2017-05-18 13:50:19');

--name: article
INSERT INTO `article` 
VALUES 
  (1,'Article 1','This is content 1',2,'2021-09-01 15:49:25','2021-09-01 15:49:25'),
  (2,'Article 2','This is content 2',2,'2021-09-01 15:49:25','2021-09-01 15:49:25'),
  (3,'Article 3','This is content 3',1,'2021-09-01 15:49:25','2021-09-01 15:49:25');

--name: category
INSERT INTO `category` 
VALUES 
  (1,'Makanan','food','2021-09-01 15:49:25','2021-09-01 15:49:25'),
  (2,'Kehidupan','life','2021-09-01 15:49:25','2021-09-01 15:49:25'),
  (3,'Cinta','love','2021-09-01 15:49:25','2021-09-01 15:49:25');

--name: article_category
INSERT INTO `article_category` 
VALUES (1,1,1), (2,1,2), (3,1,3), (4,2,1), (5,2,2), (6,3,3);