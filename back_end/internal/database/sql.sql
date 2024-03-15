CREATE TABLE UDO_ADDITIONAL_CONDITION (
      ID NUMBER(17, 0) NOT NULL,
      NORM_MUNITION_ID NUMBER(17, 0) NOT NULL,
      MUNITION_ID NUMBER(17, 0) NOT NULL,
      DESCRIPTION VARCHAR2(250) NOT NULL,
      SEX NUMBER(1, 0), --пол
      DIVISION_ID NUMBER(17, 0), --подразделение
      POSITION_ID NUMBER(17, 0), --должность
      RANK_ID NUMBER(17, 0), --звание
      CLIMATE VARCHAR(100), --климат
      REPLACE_MUNITION_ID NUMBER(17, 0), --имущество (если требуется замена)
      CONSTRAINT ADDITIONAL_CONDITION_ID_PK PRIMARY KEY (id)
)

CREATE TABLE UDO_DIVISION (
      ID NUMBER(17, 0) NOT NULL,
      NAME VARCHAR2(50) NOT NULL,
      DESCRIPTION VARCHAR2(100),
      CONSTRAINT DIVISION_ID_PK PRIMARY KEY (id),
      CONSTRAINT DIVISION_UNIQUE UNIQUE (name)
)

CREATE TABLE UDO_RANK (
      ID NUMBER(17, 0) NOT NULL,
      NAME VARCHAR2(50) NOT NULL,
      DESCRIPTION VARCHAR2(100),
      CONSTRAINT RANK_ID_PK PRIMARY KEY (id),
      CONSTRAINT RANK_UNIQUE UNIQUE (name)
)

CREATE TABLE UDO_POSITION (
      ID NUMBER(17, 0) NOT NULL,
      NAME VARCHAR2(50) NOT NULL,
      DESCRIPTION VARCHAR2(100),
      CONSTRAINT POSITION_ID_PK PRIMARY KEY (id),
      CONSTRAINT POSITION_UNIQUE UNIQUE (name)
)