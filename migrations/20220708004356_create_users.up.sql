CREATE  TABLE users (
  id serial not null primary key,
  email varchar not null unique,
  encrypted_password varchar not null,
  user_name varchar not null,
  seccond_name varchar not null,
  user_role varchar ,
  education_department boolean DEFAULT false,
  source_tracking_department boolean DEFAULT false,
  periodic_reporting_department boolean DEFAULT false,
  international_department boolean DEFAULT false,
  documentation_department boolean DEFAULT false,
  nr_department boolean DEFAULT false,
  db_department boolean DEFAULT false, 
  monitoring_specialist boolean DEFAULT false,
  monitoring_responsible int 
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
