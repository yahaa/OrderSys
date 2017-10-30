-- 表设计语句
CREATE TABLE USERS (
  U_id    NUMBER      NOT NULL,
  Name    VARCHAR(20) NOT NULL,
  Sex     VARCHAR(20),
  Phone   VARCHAR(15),
  Role_id NUMBER,
  PRIMARY KEY (U_id),
  FOREIGN KEY (Role_id) REFERENCES ROLE (Role_id)
);


CREATE TABLE ROLE (
  Role_id   NUMBER      NOT NULL,
  Role_name VARCHAR(20) NOT NULL,
  PRIMARY KEY (Role_id)
);


CREATE TABLE PERMIT (
  Per_id   NUMBER      NOT NULL,
  Per_name VARCHAR(50) NOT NULL,
  PRIMARY KEY (Per_id)
);

CREATE TABLE PERMIT_ROLE (
  Role_id NUMBER NOT NULL,
  Per_id  NUMBER NOT NULL,
  PRIMARY KEY (Role_id, Per_id),
  FOREIGN KEY (Role_id) REFERENCES ROLE (Role_id),
  FOREIGN KEY (Per_id) REFERENCES PERMIT (Per_id)
);


CREATE TABLE PRODUCT (
  Pro_id     NUMBER       NOT NULL,
  Pro_name   VARCHAR(100) NOT NULL,
  Pro_sm     VARCHAR(100),
  Pro_com    VARCHAR(100),
  Pro_price  NUMBER,
  Pro_counts NUMBER,
  Pro_style  VARCHAR(100),
  Order_time TIMESTAMP,
  Marks      VARCHAR(100),
  PRIMARY KEY (Pro_id)
);

CREATE TABLE ORDERS (
  Ord_id     NUMBER NOT NULL,
  U_id       NUMBER NOT NULL,
  Saler_id   NUMBER NOT NULL,
  Pro_id     NUMBER NOT NULL,
  Nums       NUMBER,
  Total      NUMBER,
  Order_time TIMESTAMP,
  Marks      VARCHAR(100),
  Ord_com    VARCHAR(100),
  Play_time  TIMESTAMP,
  PRIMARY KEY (Ord_id),
  FOREIGN KEY (U_id) REFERENCES USERS (U_id),
  FOREIGN KEY (Pro_id) REFERENCES PRODUCT (Pro_id),
  FOREIGN KEY (Saler_id) REFERENCES USERS (U_id)
);

CREATE TABLE USER_SALER (
  U_id         NUMBER NOT NULL,
  Pro_id       NUMBER NOT NULL,
  Sale_type    VARCHAR(100),
  Deliver_type VARCHAR(100),
  Marks        VARCHAR(100),
  PRIMARY KEY (U_id, Pro_id),
  FOREIGN KEY (U_id) REFERENCES USERS (U_id),
  FOREIGN KEY (Pro_id) REFERENCES PRODUCT (pro_id)
);


-- 数据语句

INSERT INTO USERS(USERID,NAME,PASSWORD,SEX,PHONE) VALUES (14122260,'Admin','123456','男','18777859598');

INSERT INTO ROLE(ROLEID,ROLENAME) VALUES (1,'Admin');

