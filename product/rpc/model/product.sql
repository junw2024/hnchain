-- Active: 1708569287902@@127.0.0.1@5432@haidb

-- goctl model pg datasource -url="postgres://hai:hai@127.0.0.1:5432/haidb?sslmode=disable"   -t product -dir="./" -c


drop table if exists product;
CREATE TABLE product (
    id bigint  NOT NULL primary key,
    categoryid bigint  NULL DEFAULT 0,
    name varchar(100) NOT NULL DEFAULT '',
    subtitle varchar(200) NOT NULL DEFAULT '',
    imageurl varchar(256) NOT NULL DEFAULT '', 
    images varchar(2048) NOT NULL DEFAULT '',
    detail varchar(2048) NOT NULL DEFAULT '',
    price  decimal(20,2) NOT NULL DEFAULT 0.0,
    stock  int NOT NULL DEFAULT 0,
    status int NOT NULL DEFAULT 1,
    createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

comment on table  product is '商品表';
comment on column product.categoryid is '商品类别';
comment on column product.name is '商品名称';
comment on column product.subtitle is '商品副标题';
comment on column product.images is '图片地址,逗号分隔';
comment on column product.imageurl is '商品主图';

comment on column product.detail is '商品详情';
comment on column product.price is '价格,单位-元保留两位小数';
comment on column product.stock is '库存数量';
comment on column product.status is '商品状态.1-在售 2-下架 3-删除';

create  index idx_product_cateid on product(categoryid);
create  index idx_product_status on product(status);
create  index idx_product_createtime on product(createtime);
create  index idx_product_updatetime on product(updatetime);

drop table if exists product_category;
CREATE TABLE product_category (
    id bigint  NOT NULL PRIMARY key,
    parentid bigint NOT NULL DEFAULT 0,
    name varchar(200) NOT NULL DEFAULT '',
    status int NOT NULL DEFAULT 1,
    createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

comment on table  product_category is '商品类别表';
comment on column product_category.parentid is '父类别id当id=0时说明是根节点,一级类别';
comment on column product_category.name is '类别名称';
comment on column product_category.status is '类别状态1-正常,2-已废弃';

create  index idx_product_category_status on product_category(status);
create  index idx_product_category_createtime on product_category(createtime);
create  index idx_product_category_updatetime on product_category(updatetime);

drop table if exists product_operation;
CREATE TABLE product_operation (
    id bigint  NOT NULL primary key,
    productid bigint NOT NULL DEFAULT 0,
    status int NOT NULL DEFAULT 1,
    createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
comment on table  product_operation is '商品运营表';
comment on column product_operation.productid is '运营商品状态 0-下线 1-上线';
comment on column product_operation.status is '商品ID';
create  index idx_product_operation_productid on product_operation(productid);
create  index idx_product_operation_createtime on product_operation(createtime);
create  index idx_product_operation_updatetime on product_operation(updatetime);


drop table if exists product_recommend;
CREATE TABLE product_recommend (
    id bigint  NOT NULL primary key,
    productid bigint NOT NULL DEFAULT 0,
    status int NOT NULL DEFAULT 1,
    heat int not null DEFAULT 0,
    imageurl varchar(500) not null DEFAULT '',
    createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
comment on table  product_recommend is '推荐商品';
comment on column product_recommend.status is '状态1:在推荐 0:不推荐';
comment on column product_recommend.imageurl is '推荐主图';

comment on column product_recommend.heat is '推荐指数，越大越靠前';
create  index idx_product_recommend_productid on product_recommend(productid);
create  index idx_product_recommend_heat on product_recommend(heat);
create  index idx_product_recommend_createtime on product_recommend(createtime);
create  index idx_product_recommend_updatetime on product_recommend(updatetime);



//---select * from product_recommend where status=1 order by heat desc limit 10;
//---


















