/*
Navicat PGSQL Data Transfer

Source Server         : localhost
Source Server Version : 90408
Source Host           : localhost:5432
Source Database       : word
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 90408
File Encoding         : 65001

Date: 2016-11-02 18:05:51
*/


-- ----------------------------
-- Table structure for dictionary
-- ----------------------------
DROP TABLE IF EXISTS "public"."dictionary";
CREATE TABLE "public"."dictionary" (
"id" int4 NOT NULL,
"spell" varchar(255) COLLATE "default" NOT NULL,
"interpretation" varchar(255) COLLATE "default" NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Records of dictionary
-- ----------------------------
INSERT INTO "public"."dictionary" VALUES ('1', 'leaks', 'n. 泄漏，渗漏（leak的复数）
v. 泄漏；渗（leak的单三形式）');
INSERT INTO "public"."dictionary" VALUES ('2', 'alternatively', 'adv. 非此即彼；二者择一地；作为一种选择');
INSERT INTO "public"."dictionary" VALUES ('3', 'compromise', 'vt. 妥协；危害
vi. 妥协；让步
n. 妥协，和解；折衷');
INSERT INTO "public"."dictionary" VALUES ('4', 'bulk', 'n. 体积，容量；大多数，大部分；大块
vt. 使扩大，使形成大量；使显得重要
n. (Bulk)人名；(土)布尔克');
INSERT INTO "public"."dictionary" VALUES ('5', 'allocator', 'n. 分配算符');
INSERT INTO "public"."dictionary" VALUES ('6', 'determine', 'v. （使）下决心，（使）做出决定
vt. 决定，确定；判定，判决；限定
vi. 确定；决定；判决，终止；[主用于法律]了结，终止，结束');
INSERT INTO "public"."dictionary" VALUES ('7', 'iteration', 'n. [数] 迭代；反复；重复');
INSERT INTO "public"."dictionary" VALUES ('8', 'various', 'adj. 各种各样的；多方面的');
INSERT INTO "public"."dictionary" VALUES ('9', 'maintain', 'vt. 维持；继续；维修；主张；供养');
INSERT INTO "public"."dictionary" VALUES ('10', 'proficient', 'adj. 熟练的，精通的
n. 精通；专家，能手');
INSERT INTO "public"."dictionary" VALUES ('11', 'effort', 'n. 努力；成就');
INSERT INTO "public"."dictionary" VALUES ('12', 'policy', 'n. 政策，方针；保险单');
INSERT INTO "public"."dictionary" VALUES ('13', 'further', 'adv. 进一步地；而且；更远地
adj. 更远的；深一层的
vt. 促进，助长；增进');

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------

-- ----------------------------
-- Indexes structure for table dictionary
-- ----------------------------
CREATE UNIQUE INDEX "spell" ON "public"."dictionary" USING btree ("spell");

-- ----------------------------
-- Primary Key structure for table dictionary
-- ----------------------------
ALTER TABLE "public"."dictionary" ADD PRIMARY KEY ("id");
