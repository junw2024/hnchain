-- Active: 1708569287902@@127.0.0.1@5432@haidb

-- goctl model pg datasource -url="postgres://hai:hai@127.0.0.1:5432/haidb?sslmode=disable"   -t orders -dir="./" -c

DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
     id BIGINT NOT NULL  PRIMARY KEY,
     ordernum   VARCHAR(64) NOT NULL DEFAULT '',
     userid bigint NOT NULL DEFAULT 0,
     shoppingid bigint NOT NULL DEFAULT 0,
     payment    decimal(20,2) NOT NULL DEFAULT 0,
     paymenttype INT NOT NULL DEFAULT 1,
     postage  decimal(20,2)  NOT NULL DEFAULT 0,
     status INT NOT NULL DEFAULT 10,
     createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
comment on table  orders is '订单表';
comment on COLUMN  orders.ordernum is '订单号';
comment on COLUMN  orders.userid is '用户id';
COMMENT ON COLUMN orders.shoppingid IS '收货信息表id';
COMMENT ON COLUMN orders.payment IS '实际付款金额,单位是元,保留两位小数';
COMMENT ON COLUMN orders.paymenttype IS '支付类型,1-在线支付';
COMMENT ON COLUMN orders.postage IS '运费,单位是元';
COMMENT ON COLUMN orders.status IS '订单状态:0-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭';

CREATE UNIQUE INDEX  idx_union_orders_ordernum on orders(ordernum);
CREATE  INDEX idx_orders_userid on orders(userid);
CREATE  INDEX idx_orders_shoppingid on orders(shoppingid);
CREATE  INDEX idx_orders_status on orders(status);
CREATE  INDEX idx_orders_createtime on orders(createtime);
CREATE  INDEX idx_orders_updatetime on orders(updatetime);

DROP TABLE IF EXISTS orderitem;
CREATE TABLE orderitem (
      id  bigint NOT NULL PRIMARY KEY,
      ordernum VARCHAR(64) NOT NULL DEFAULT '',
      userid   bigint NOT NULL DEFAULT 0,
      productid bigint NOT NULL DEFAULT 0,
      productname varchar(200) NOT NULL DEFAULT '',
      productimage varchar(500) NOT NULL DEFAULT '',
      currentprice decimal(20,2) NOT NULL DEFAULT 0,
      quantity int NOT NULL DEFAULT 0,
      totalprice decimal(20,2) NOT NULL DEFAULT 0,
      createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

comment on table  orderitem  is '订单明细表';
comment on COLUMN  orderitem.ordernum  is '订单号';
comment on COLUMN  orderitem.productid  is '商品id';
comment on COLUMN  orderitem.userid  is '用户id';
comment on COLUMN  orderitem.productname  is '商品名称';
comment on COLUMN  orderitem.productimage  is '商品图片地址';
comment on COLUMN  orderitem.currentprice  is '生成订单时的商品单价，单位是元,保留两位小数';
comment on COLUMN  orderitem.quantity  is '商品数量';
comment on COLUMN  orderitem.quantity  is '商品总价,单位是元,保留两位小数';

CREATE INDEX  idx_orderitem_ordernum on orderitem(ordernum);  
CREATE INDEX  idx_orderitem_userid on orderitem(userid);  
CREATE INDEX  idx_orderitem_productid on orderitem(productid);  
CREATE INDEX  idx_orderitem_createtime on orderitem(createtime);  
CREATE INDEX  idx_orderitem_updatetime on orderitem(updatetime);  

DROP TABLE IF EXISTS shipping;
CREATE TABLE shipping (
     id bigint NOT NULL PRIMARY KEY,
     ordernum varchar(64) NOT NULL DEFAULT '',
     userid  bigint NOT NULL DEFAULT 0,
     receiver_name varchar(40) NOT NULL DEFAULT '',
     receiver_phone varchar(20) NOT NULL DEFAULT '',
     receiver_mobile varchar(20) NOT NULL DEFAULT '',
     receiver_province varchar(20) NOT NULL DEFAULT '',
     receiver_city varchar(20) NOT NULL DEFAULT '',
     receiver_district varchar(20) NOT NULL DEFAULT '',
     receiver_address varchar(200) NOT NULL DEFAULT '',
     createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
     updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);
comment on table   shipping  is '收货信息表';
comment on column  shipping.userid  is '用户id';
comment on column  shipping.ordernum  is '订单号';
comment on column  shipping.receiver_name  is '收货姓名';
comment on column  shipping.receiver_phone  is '收货固定电话';
comment on column  shipping.receiver_mobile  is '收货移动电话';
comment on column  shipping.receiver_province  is '省份';
comment on column  shipping.receiver_city  is '城市';
comment on column  shipping.receiver_district  is '区/县';
comment on column  shipping.receiver_address  is '详细地址';

CREATE INDEX idx_shipping_ordernum on shipping(ordernum);
CREATE INDEX idx_shipping_userid on shipping(userid);
CREATE INDEX idx_shipping_createtime on shipping(createtime);
CREATE INDEX idx_shipping_updatetime on shipping(updatetime);






















