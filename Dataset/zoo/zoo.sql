CREATE TABLE type (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  description CLOB NULL,
  PRIMARY KEY (id))
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE type_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER type_seq_tr
 BEFORE INSERT ON type FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT type_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/


-- -----------------------------------------------------
-- Table `mydb`.`class`
-- -----------------------------------------------------
CREATE TABLE class (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  description CLOB NULL,
  type_id NUMBER(10) NOT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_class_type1
    FOREIGN KEY (type_id)
    REFERENCES type (id)
   )
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE class_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER class_seq_tr
 BEFORE INSERT ON class FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT class_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/

CREATE INDEX fk_class_type1_idx ON class (type_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`order_a`
-- -----------------------------------------------------
CREATE TABLE order_a (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  description CLOB NULL,
  class_id NUMBER(10) NOT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_order_a_class1
    FOREIGN KEY (class_id)
    REFERENCES class (id)
   )
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE order_a_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER order_a_seq_tr
 BEFORE INSERT ON order_a FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT order_a_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/

CREATE INDEX fk_order_a_class1_idx ON order_a (class_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`family`
-- -----------------------------------------------------
CREATE TABLE family (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  description CLOB NULL,
  order_a_id NUMBER(10) NOT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_family_order_a1
    FOREIGN KEY (order_a_id)
    REFERENCES order_a (id)
   )
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE family_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER family_seq_tr
 BEFORE INSERT ON family FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT family_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/

CREATE INDEX fk_family_order_a1_idx ON family (order_a_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`genus`
-- -----------------------------------------------------
CREATE TABLE genus (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  description CLOB NULL,
  family_id NUMBER(10) NOT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_genus_family1
    FOREIGN KEY (family_id)
    REFERENCES family (id)
   )
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE genus_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER genus_seq_tr
 BEFORE INSERT ON genus FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT genus_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/

CREATE INDEX fk_genus_family1_idx ON genus (family_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`species`
-- -----------------------------------------------------
CREATE TABLE species (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  conservation_status VARCHAR2(2) NULL,
  description CLOB NULL,
  genus_id NUMBER(10) NOT NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_species_genus1
    FOREIGN KEY (genus_id)
    REFERENCES genus (id)
   )
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE species_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER species_seq_tr
 BEFORE INSERT ON species FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT species_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/

CREATE INDEX fk_species_genus1_idx ON species (genus_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`animals`
-- -----------------------------------------------------
CREATE TABLE animals (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  species_id NUMBER(10) NOT NULL,
  birth DATE NULL,
  death DATE NULL,
  weight NUMBER(8,3) NULL,
  length NUMBER(7,2) NULL,
  height NUMBER(7,2) NULL,
  sex CHAR(2) NULL,
  PRIMARY KEY (id)
 ,
  CONSTRAINT fk_animals_species1
    FOREIGN KEY (species_id)
    REFERENCES species (id)
   )
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE animals_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER animals_seq_tr
 BEFORE INSERT ON animals FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT animals_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/

CREATE INDEX fk_animals_species1_idx ON animals (species_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`childs`
-- -----------------------------------------------------
CREATE TABLE childs (
  child_id NUMBER(10) NOT NULL,
  parent_id NUMBER(10) NOT NULL
 ,
  PRIMARY KEY (child_id, parent_id),
  CONSTRAINT fk_childs_animals1
    FOREIGN KEY (child_id)
    REFERENCES animals (id)
   ,
  CONSTRAINT fk_childs_animals2
    FOREIGN KEY (parent_id)
    REFERENCES animals (id)
   )
;

CREATE INDEX fk_childs_animals1_idx ON childs (child_id ASC);
CREATE INDEX fk_childs_animals2_idx ON childs (parent_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`keepers`
-- -----------------------------------------------------
CREATE TABLE keepers (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  surname VARCHAR2(255) NULL,
  patronymic VARCHAR2(255) NULL,
  birth DATE NULL,
  experience NUMBER(10) NULL,
  PRIMARY KEY (id))
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE keepers_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER keepers_seq_tr
 BEFORE INSERT ON keepers FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT keepers_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/


-- -----------------------------------------------------
-- Table `mydb`.`k_a`
-- -----------------------------------------------------
CREATE TABLE k_a (
  k_id NUMBER(10) NOT NULL,
  a_id NUMBER(10) NOT NULL,
  start_d DATE NULL,
  finish_d DATE NULL
 ,
  PRIMARY KEY (k_id, a_id),
  CONSTRAINT fk_watchers_keepers1
    FOREIGN KEY (k_id)
    REFERENCES keepers (id)
   ,
  CONSTRAINT fk_watchers_animals1
    FOREIGN KEY (a_id)
    REFERENCES animals (id)
   )
;

CREATE INDEX fk_watchers_keepers1_idx ON k_a (k_id ASC);
CREATE INDEX fk_watchers_animals1_idx ON k_a (a_id ASC);


-- -----------------------------------------------------
-- Table `mydb`.`guardians`
-- -----------------------------------------------------
CREATE TABLE guardians (
  id NUMBER(10) NOT NULL,
  name VARCHAR2(255) NULL,
  PRIMARY KEY (id))
;

-- Generate ID using sequence and trigger
CREATE SEQUENCE guardians_seq START WITH 1 INCREMENT BY 1;

CREATE OR REPLACE TRIGGER guardians_seq_tr
 BEFORE INSERT ON guardians FOR EACH ROW
 WHEN (NEW.id IS NULL)
BEGIN
 SELECT guardians_seq.NEXTVAL INTO :NEW.id FROM DUAL;
END;
/


-- -----------------------------------------------------
-- Table `mydb`.`g_a`
-- -----------------------------------------------------
CREATE TABLE g_a (
  g_id NUMBER(10) NOT NULL,
  a_id NUMBER(10) NOT NULL
 ,
  PRIMARY KEY (g_id, a_id),
  CONSTRAINT fk_guardian_animal_guardians1
    FOREIGN KEY (g_id)
    REFERENCES guardians (id)
   ,
  CONSTRAINT fk_guardian_animal_animals1
    FOREIGN KEY (a_id)
    REFERENCES animals (id)
   )
;

CREATE INDEX fk_animal_guardians1_idx ON g_a (g_id ASC);
CREATE INDEX fk_animal_animals1_idx ON g_a (a_id ASC);


Insert into type (name, description) values ('Хордовые', 'тип вторичноротых животных, для которых характерно наличие энтодермального осевого скелета в виде хорды, которая у высших форм заменяется позвоночником. По строению и функции нервной системы тип хордовых занимает высшее место среди всех животных. В мире известно более 60 000 видов хордовых.');

Insert into class (name, description, type_id) values ('Млекопитающие', 'класс позвоночных животных, основной отличительной особенностью которых является вскармливание детёнышей молоком[1]. Класс входит в кладу Synapsidomorpha надкласса четвероногих. По оценкам биологов, известно до 5500 современных видов млекопитающих[2], среди которых — и вид Человек разумный (лат. Homo sapiens); общее же число видов млекопитающих, включая вымершие, превышает 20 000[3].', 1);

Insert into order_a (name, description, class_id) values ('Грызуны','самый многочисленный отряд млекопитающих. По данным справочника Mammal Species of the World (2005) описано 2277 видов грызунов[1]. Отличительным признаком представителей отряда является наличие диастемы и одной пары крупных резцов в верхней и нижней челюстях.',1);

Insert into family (name, description, order_a_id) values ('Беличьи','К беличьим относятся разнообразные по внешнему виду и образу жизни животные, объединяемые единством происхождения и сходством анатомического строения. Размеры мелкие и средние. Длина тела от 6 (мышиная белка) до 60 см (сурки). По внешнему виду беличьих можно разделить на 3 типа: на древесных беличьих, наземных беличьих и летяг, причём считается, что общие предки были наиболее близки к древесным беличьим. ',1);

Insert into genus (name, description, family_id) values ('Сурки','Сурки образуют отчётливо выраженную группу из 14 или 15 видов (статус лесостепного сурка как отдельного вида является предметом обсуждения), в рамках семейства беличьих. Это относительно крупные, весом в несколько килограммов, животные, обитающие в открытых ландшафтах, в сооружаемых самостоятельно норах. Прародина сурков — Северная Америка, откуда они распространились через Берингию в Азию, и дальше — в Европу.',1);

Insert into species (name, conservation_status, description, genus_id) values ('Степной сурок','LC','Байбак является одним из самых крупных беличьих: длина его тела 50—70 см, масса нажировавшихся самцов достигает 10 кг. Тело у байбака толстое, на коротких, сильных лапах, вооружённых крупными когтями. Голова большая, уплощённая, шея короткая.',1);

Insert into animals (name, species_id, birth, death, weight, length, height, sex) values ('Сурок1',1,'10.11.2011',null,5.1,52,null,'м');
Insert into animals (name, species_id, birth, death, weight, length, height, sex) values ('Сурок2',1,'08.12.2011',null,4.8,47,null,'ж');
Insert into animals (name, species_id, birth, death, weight, length, height, sex) values ('Сурок3',1,'09.12.2017',null,1.5,15,null,'м');

Insert into childs values (3,1);
Insert into childs values (3,2);