use orm_test;

INSERT INTO profile (age)
VALUES (20),(20),(20),(20),(20),(20),(20),(20),(20),(20),(20),(20),(20),(20),(20),
       (20),(20),(20),(20),(20),(20),(20),(20),(20),(20);

INSERT INTO post (title, user_id)
VALUES ("title1",1),("title2",1),("title3",1),("title4",1),("title5",1);

INSERT INTO tag (name)
VALUES ("tag1"),("tag2"),("tag3"),("tag4"),("tag5");

INSERT INTO user (name,username,password,profile_id)
VALUES("admin","admin","admin",1),
		("123456","123456","123456",2),
		("234567","234567","234567",3),
		("000001","000001","000001",4),
		("000002","000002","000002",5),
		("000003","000003","000001",6),
		("000004","000004","000001",7),
		("000005","000005","000001",8),
		("000006","000006","000001",9),
		("000007","000007","000001",10),
		("000008","000008","000001",11),
		("000009","000009","000001",12),
		("0000010","0000010","000001",13),
		("0000011","0000011","000001",14),
		("0000012","0000012","000001",15),
		("0000013","0000013","000001",16),
		("0000014","0000014","000001",17),
		("0000015","0000015","000001",18),
		("0000016","0000016","000001",19),
		("0000017","0000017","000001",20),
		("0000018","0000018","000001",21);
