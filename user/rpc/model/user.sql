-- Active: 1708569287902@@127.0.0.1@5432@haidb



-- Active: 1702970850204@@127.0.0.1@5432@haidb

-- goctl model pg datasource -url="postgres://hai:hai@127.0.0.1:5432/haidb?sslmode=disable"   -t hnuser -dir="./" -c


drop table if exists hnuser;

create table hnuser
(
    id  bigint not null primary key,
	identitytype char(1) not null default '0',
	identity varchar(100) not null default '',
    name varchar(256) not null  default  '',
    nick varchar(256) not null  default '',
    sex char(1) default '0',
    phone varchar(15) not null,
    username varchar(200) not null,
    question varchar(100) not null default '',
    answer   varchar(100) not null default '',
    password varchar(50) not null default '',
    address varchar(500) not null default '',
    email varchar(100) not null default '',
    wxopenid varchar(200) not null default '',
    wxunionid varchar(200) not null default '',
	loginaddr varchar(200) not null default '',
	createtime  timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);



comment on column hnuser.name is '名称';
comment on column hnuser.nick is '昵称';
comment on column hnuser.identitytype is '证件类型 0:身份证 1:护照 2:驾照';
comment on column hnuser.identity is '证件号码';
comment on column hnuser.username is '用户名称';
comment on column hnuser.question is '找回密码问题';
comment on column hnuser.answer is '问题答案';
comment on column hnuser.sex is '性别 0:未知 1:男 2:女';
comment on column hnuser.wxopenid is  '微信openid';
comment on column hnuser.wxunionid is '微信unionid';
comment on column hnuser.loginaddr is '最后登录ip';
comment on column hnuser.createtime is '创建时间';
comment on column hnuser.updatetime is '更新时间';



create unique index idx_hnuser_fusername on hnuser  (username);
create unique index idx_hnuser_fphone on hnuser (phone);
create  index idx_hnuser_fpassword on hnuser  (password);
create  index idx_hnuser_fcreatetime on hnuser (createtime);
create  index idx_hnuser_fupdatetime on hnuser  (updatetime);
create  index idx_hnuser_fidentity on hnuser  (identity);


drop table if exists hnuser_collection;
create table hnuser_collection
(
     id bigint not null primary key,
     uid bigint NOT NULL DEFAULT 0,
     productid bigint NOT NULL DEFAULT 0,
     isdelete  boolean NOT NULL DEFAULT false,
     createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

comment on table hnuser_collection is '收藏';
comment on column hnuser_collection.uid is '用户id';
comment on column hnuser_collection.productid is '商品id';
comment on column hnuser_collection.isdelete is '是否删除';

create  index idx_hnuser_collection_uid on hnuser_collection  (uid);
create  index idx_hnuser_collection_productid on hnuser_collection  (productid);
create  index idx_hnuser_collection_createtime on hnuser_collection  (createtime);
create  index idx_hnuser_collection_updatetime on hnuser_collection  (updatetime);

drop table if exists hnuser_rev_addr;

CREATE TABLE hnuser_rev_addr (
   id bigint NOT NULL PRIMARY KEY,
   uid bigint NOT NULL DEFAULT 0,
   name varchar(100) NOT NULL DEFAULT '',
   phone varchar(20) NOT NULL DEFAULT '' ,
   isdefault boolean NOT NULL DEFAULT false,
   postcode varchar(100) NOT NULL DEFAULT '',
   province varchar(100) NOT NULL DEFAULT '',
   city  varchar(100) NOT NULL DEFAULT '',
   region varchar(100) NOT NULL DEFAULT '',
   detail_address varchar(128) NOT NULL DEFAULT '',
   isdelete boolean NOT NULL DEFAULT false,
   createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
comment on table hnuser_rev_addr is '用户收货地址表';
comment on column hnuser_rev_addr.name is '收货人名称';
comment on column hnuser_rev_addr.phone is '手机号';
comment on column hnuser_rev_addr.isdefault is '是否为默认地址';
comment on column hnuser_rev_addr.postcode is '邮政编码';
comment on column hnuser_rev_addr.city is '省份/直辖市';
comment on column hnuser_rev_addr.region is '区';
comment on column hnuser_rev_addr.detail_address is '详细地址(街道)';
comment on column hnuser_rev_addr.isdelete is '是否删除';

create  index idx_hnuser_rev_addr_uid on hnuser_rev_addr(uid);
create  index idx_hnuser_rev_addr_phone on hnuser_rev_addr(phone);

create  index idx_hnuser_rev_addr_createtime on hnuser_rev_addr(createtime);

create  index idx_hnuser_rev_addr_updatetime on hnuser_rev_addr(updatetime);




select * from "public"."hnuser" where id =7149513293248335872;
select * from hnuser_collection;


select id,uid,productid,isdelete,createtime,updatetime from "public"."hnuser_collection" where uid=7149513293248335872 
