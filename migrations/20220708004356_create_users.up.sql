CREATE  TABLE users (
  id serial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null,
  userName varchar not null,
  seccondName varchar not null,
  user_role varchar ,
  educationdepartment boolean DEFAULT false,
  sourceTrackingdepartment boolean DEFAULT false,
  periodicreportingdepartment boolean DEFAULT false,
  internationaldepartment boolean DEFAULT false,
  documentationdepartment boolean DEFAULT false,
  nrdepartment boolean DEFAULT false,
  dbdepartment boolean DEFAULT false, 
  monitoringspecialist boolean DEFAULT false,
  monitoringresponsible int 
);

CREATE TABLE apilogs(
  request_id serial not null primary key,
  remote_ip varchar,
  method varchar,
  request_url varchar not null ,
  resp_status varchar not null ,
  duration varchar not null ,
  request_time varchar not null ,
  request_message varchar not null, 
)
