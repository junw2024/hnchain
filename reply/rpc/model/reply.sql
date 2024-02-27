-- Active: 1708569287902@@127.0.0.1@5432@haidb
-- goctl model pg datasource -url="postgres://hai:hai@127.0.0.1:5432/haidb?sslmode=disable"   -t reply -dir="./" -c


DROP TABLE IF EXISTS reply;
CREATE TABLE reply(
    id bigint NOT NULL PRIMARY key,
    business varchar(64) NOT NULL DEFAULT '',
    targetid bigint NOT NULL DEFAULT 0,
    reply_userid bigint NOT NULL DEFAULT 0,
    be_reply_userid bigint NOT NULL DEFAULT 0,
    parentid bigint NOT NULL DEFAULT 0,
    content varchar(500) NOT NULL DEFAULT '',
    imageurl varchar(255) NOT NULL DEFAULT '',
    createtime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatetime timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

comment on table  reply is '评论列表';
comment on COLUMN  reply.business is '评论业务类型';
comment on COLUMN  reply.targetid is '评论目标id';
comment on COLUMN  reply.reply_userid is '回复用户id';
comment on COLUMN  reply.be_reply_userid is '被回复用户id';
comment on COLUMN  reply.parentid is '父评论id';
comment on COLUMN  reply.content is '评论内容';
comment on COLUMN  reply.imageurl is '评论内容';

create INDEX idx_reply_targetid on reply(targetid);
create INDEX idx_reply_reply_userid on reply(reply_userid);
create INDEX idx_reply_reply_be_reply_userid on reply(be_reply_userid);
create INDEX idx_reply_reply_parentid on reply(parentid);
create INDEX idx_reply_createtime on reply(createtime);
create INDEX idx_reply_updatetime on reply(updatetime);






